// Package batch splits work into batches and runs it.
package batch

import "sync"

// Chunk splits ids into batches of at most size ids, in order.
// The last batch holds the remainder and may be shorter.
func Chunk(ids []string, size int) [][]string {
	var chunks [][]string
	for i := 0; i+size <= len(ids); i += size {
		chunks = append(chunks, ids[i:i+size])
	}
	return chunks
}

// FetchAll calls fetch for every id in parallel.
// It returns the results in the same order as ids, or the first error.
func FetchAll(ids []string, fetch func(string) (string, error)) ([]string, error) {
	out := make([]string, len(ids))
	errs := make([]error, len(ids))
	var wg sync.WaitGroup
	for i, id := range ids {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out[i], errs[i] = fetch(id)
		}()
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
