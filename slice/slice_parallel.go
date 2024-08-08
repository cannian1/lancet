// Copyright 2024 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

package slice

import (
	"runtime"
	"sync"
)

// UniqueByParallel removes duplicate elements from the slice by parallel
// The comparator function is used to compare the elements
// The numOfThreads parameter specifies the number of threads to use
// If numOfThreads is less than or equal to 0, it will be set to 1
// The comparator function should return true if the two elements are equal
// Play: todo
func UniqueByParallel[T comparable](slice []T, numOfThreads int, comparator func(item T, other T) bool) []T {
	if numOfThreads <= 0 {
		numOfThreads = 1
	} else if numOfThreads > len(slice) {
		numOfThreads = len(slice)
	}

	maxThreads := runtime.NumCPU()
	if numOfThreads > maxThreads {
		numOfThreads = maxThreads
	}

	removeDuplicate := func(items []T, comparator func(item T, other T) bool) []T {
		var result []T
		for _, item := range items {
			seen := false
			for _, r := range result {
				if comparator(item, r) {
					seen = true
					break
				}
			}
			if !seen {
				result = append(result, item)
			}
		}
		return result
	}

	chunkSize := (len(slice) + numOfThreads - 1) / numOfThreads

	chunks := make([][]T, 0, numOfThreads)
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	type resultChunk struct {
		index int
		data  []T
	}
	resultCh := make(chan resultChunk, numOfThreads)
	var wg sync.WaitGroup

	for i, chunk := range chunks {
		wg.Add(1)
		go func(index int, chunk []T) {
			defer wg.Done()
			resultCh <- resultChunk{index, removeDuplicate(chunk, comparator)}
		}(i, chunk)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	results := make([][]T, len(chunks))
	for r := range resultCh {
		results[r.index] = r.data
	}

	result := []T{}
	seen := make(map[T]bool)

	for _, chunk := range results {
		for _, item := range chunk {
			if !seen[item] {
				seen[item] = true
				result = append(result, item)
			}
		}

	}

	return result
}