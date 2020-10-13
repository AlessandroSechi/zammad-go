package zammad

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func NewClient(client *Client) (*Client, error) {

	if client.Url == "" {
		return nil, errors.New("APIBase is required to create a Client")
	}

	if (client.Username == "" || client.Password == "") && client.Token == "" && client.OAuth == "" {
		return nil, errors.New("Username and password, token or OAuth2 are required to create a Client")
	}

	if (client.Username != "" || client.Password != "") && (client.Token != "" || client.OAuth != "") {
		return nil, errors.New("Only one authentication type between username and password, token and OAuth is supported")
	}

	if client.Token != "" && client.OAuth != "" {
		return nil, errors.New("Only one between token and oauth must be specified")
	}

	client.Client = &http.Client{}

	return client, nil
}

// NewRequest constructs a request
// Convert payload to a JSON
func (c *Client) NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequest(method, url, buf)
}

// Send makes a request to the API, the response body will be
// unmarshaled into v, or if v is an io.Writer, the response will
// be written to it without decoding
func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// Set default headers
	req.Header.Set("Accept", "application/json")

	// Default values for headers
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	resp, err = c.Client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &ErrorResponse{}
		data, err = ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}

		return errResp
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		io.Copy(w, resp.Body)
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

// SendWithAuth makes a request to the API and apply proper authentication header automatically.
func (c *Client) SendWithAuth(req *http.Request, v interface{}) error {

	//Detect Authentication Type
	if c.Username != "" && c.Password != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}
	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Token token=%s", c.Token))
	}
	if c.OAuth != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.OAuth))
	}

	return c.Send(req, v)
}
