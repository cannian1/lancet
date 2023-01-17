// Copyright 2023 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package stream implements a sequence of elements supporting sequential and parallel aggregate operations.
// this package is an experiment to explore if stream in go can work as the way java does. it's complete, but not
// powerful like other libs
package stream

import (
	"bytes"
	"encoding/gob"

	"golang.org/x/exp/constraints"
)

// A stream should implements methods:
// type StreamI[T any] interface {

// 	// part methods of Java Stream Specification.
// 	Distinct() StreamI[T]
// 	Filter(predicate func(item T) bool) StreamI[T]
// 	FlatMap(mapper func(item T) StreamI[T]) StreamI[T]
// 	Map(mapper func(item T) T) StreamI[T]
// 	Peek(consumer func(item T)) StreamI[T]

// 	Sort(less func(a, b T) bool) StreamI[T]
// 	Max(less func(a, b T) bool) (T, bool)
// 	Min(less func(a, b T) bool) (T, bool)

// 	Limit(maxSize int) StreamI[T]
// 	Skip(n int64) StreamI[T]

// 	AllMatch(predicate func(item T) bool) bool
// 	AnyMatch(predicate func(item T) bool) bool
// 	NoneMatch(predicate func(item T) bool) bool
// 	ForEach(consumer func(item T))
// 	Reduce(accumulator func(a, b T) T) (T, bool)
// 	Count() int

// 	FindFirst() (T, bool)

// 	ToSlice() []T

// 	// part of methods custom extension
// 	Reverse() StreamI[T]
// 	Range(start, end int64) StreamI[T]
// 	Concat(streams ...StreamI[T]) StreamI[T]
// }

type stream[T any] struct {
	source []T
}

// Of creates a stream stream whose elements are the specified values.
func Of[T any](elems ...T) stream[T] {
	return FromSlice(elems)
}

// Generate stream where each element is generated by the provided generater function
// generater function: func() func() (item T, ok bool) {}
func Generate[T any](generator func() func() (item T, ok bool)) stream[T] {
	source := make([]T, 0)

	var zeroValue T
	for next, item, ok := generator(), zeroValue, true; ok; {
		item, ok = next()
		if ok {
			source = append(source, item)
		}
	}

	return FromSlice(source)
}

// FromSlice create stream from slice.
func FromSlice[T any](source []T) stream[T] {
	return stream[T]{source: source}
}

// FromChannel create stream from channel.
func FromChannel[T any](source <-chan T) stream[T] {
	s := make([]T, 0)

	for v := range source {
		s = append(s, v)
	}

	return FromSlice(s)
}

// FromRange create a number stream from start to end. both start and end are included. [start, end]
func FromRange[T constraints.Integer | constraints.Float](start, end, step T) stream[T] {
	if end < start {
		panic("stream.FromRange: param start should be before param end")
	} else if step <= 0 {
		panic("stream.FromRange: param step should be positive")
	}

	l := int((end-start)/step) + 1
	source := make([]T, l, l)

	for i := 0; i < l; i++ {
		source[i] = start + (T(i) * step)
	}

	return FromSlice(source)
}

// Distinct returns a stream that removes the duplicated items.
func (s stream[T]) Distinct() stream[T] {
	source := make([]T, 0)

	distinct := map[string]bool{}

	for _, v := range s.source {
		// todo: performance issue
		k := hashKey(v)
		if _, ok := distinct[k]; !ok {
			distinct[k] = true
			source = append(source, v)
		}
	}

	return FromSlice(source)
}

func hashKey(data any) string {
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic("stream.hashKey: get hashkey failed")
	}
	return buffer.String()
}

// Filter returns a stream consisting of the elements of this stream that match the given predicate.
func (s stream[T]) Filter(predicate func(item T) bool) stream[T] {
	source := make([]T, 0)

	for _, v := range s.source {
		if predicate(v) {
			source = append(source, v)
		}
	}

	return FromSlice(source)
}

// Map returns a stream consisting of the elements of this stream that apply the given function to elements of stream.
func (s stream[T]) Map(mapper func(item T) T) stream[T] {
	source := make([]T, s.Count(), s.Count())

	for i, v := range s.source {
		source[i] = mapper(v)
	}

	return FromSlice(source)
}

// Count returns the count of elements in the stream.
func (s stream[T]) Count() int {
	return len(s.source)
}

// ToSlice return the elements in the stream.
func (s stream[T]) ToSlice() []T {
	return s.source
}
