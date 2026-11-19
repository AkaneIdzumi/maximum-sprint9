package main

import (
	"math"
	"reflect"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "negative size", size: -10, want: 0},
		{name: "zero size", size: 0, want: 0},
		{name: "single element", size: 1, want: 1},
		{name: "multiple elements", size: 1000, want: 1000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := generateRandomElements(test.size)
			if len(data) != test.want {
				t.Fatalf("length = %d, want %d", len(data), test.want)
			}
			for index, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive integer", index, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{7}, want: 7},
		{name: "maximum at start", data: []int{9, 3, 1}, want: 9},
		{name: "maximum in middle", data: []int{1, 9, 3}, want: 9},
		{name: "maximum at end", data: []int{1, 3, 9}, want: 9},
		{name: "equal elements", data: []int{5, 5, 5}, want: 5},
		{name: "repeated maximum", data: []int{9, 1, 9}, want: 9},
		{name: "zero values", data: []int{0, 0, 0}, want: 0},
		{name: "negative elements", data: []int{-9, -2, -5}, want: -2},
		{name: "mixed signs", data: []int{-5, 0, 8, -1}, want: 8},
		{name: "integer limits", data: []int{math.MinInt, math.MaxInt}, want: math.MaxInt},
		{name: "minimum integer", data: []int{math.MinInt}, want: math.MinInt},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			original := append([]int{}, test.data...)
			if got := maximum(test.data); got != test.want {
				t.Errorf("maximum(%v) = %d, want %d", test.data, got, test.want)
			}
			if len(test.data) > 0 && !reflect.DeepEqual(test.data, original) {
				t.Error("maximum changed the input slice")
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	for _, data := range [][]int{nil, {}, {math.MinInt}, {-9, -2, -5}, {0, 0, 0}, {math.MinInt, math.MaxInt}} {
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Errorf("maxChunks(%v) = %d, want %d", data, got, want)
		}
	}

	for _, size := range []int{1, 2, 7, 8, 9, 15, 16, 17, 63, 64, 65, 1001} {
		for position := 0; position < size; position++ {
			data := make([]int, size)
			for index := range data {
				data[index] = -5
			}
			data[position] = 100
			original := append([]int{}, data...)
			if got := maxChunks(data); got != 100 {
				t.Fatalf("size %d, maximum at %d: got %d, want 100", size, position, got)
			}
			if !reflect.DeepEqual(data, original) {
				t.Fatalf("maxChunks changed the input slice of size %d", size)
			}
		}
	}
}
