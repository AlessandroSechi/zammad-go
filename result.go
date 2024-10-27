package zammad

// Result is a generic type that holds a slice of results and provides methods
// to paginate through them.
type Result[T any] struct {
	res     []T                                       // Slice of results of type T
	resFunc func(options RequestOptions) ([]T, error) // Function to fetch results based on RequestOptions
	opts    RequestOptions                            // Options for the request
	lastErr error                                     // Last error encountered during fetching
}

// Next fetches the next page of results. It returns false if there are no more
// results to fetch. If an error occurs during fetching, it sets the lastErr
// field and returns true.
func (tr *Result[T]) Next() bool {
	res, err := tr.resFunc(tr.opts)
	if err != nil {
		tr.res = nil
		tr.lastErr = err
		return true
	}

	if tr.opts.page == 0 {
		tr.opts.page = 1
	}

	tr.opts.page++
	tr.res = res

	if len(res) == 0 {
		return false
	}

	return true
}

// Fetch returns the current slice of results and the last error encountered.
func (tr *Result[T]) Fetch() ([]T, error) {
	return tr.res, tr.lastErr
}

// FetchAll fetches all pages of results and returns them as a single slice.
// If an error occurs during fetching, it returns the results fetched so far
// along with the error.
func (tr *Result[T]) FetchAll() ([]T, error) {
	resAll := make([]T, 0)

	for tr.Next() {
		res, err := tr.Fetch()
		if err != nil {
			return resAll, err
		}

		resAll = append(resAll, res...)
	}

	return resAll, nil
}
