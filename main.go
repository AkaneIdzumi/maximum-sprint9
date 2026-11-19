package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for index := range data {
		data[index] = rand.Intn(size) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxima := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var waitGroup sync.WaitGroup

	for index := 0; index < CHUNKS; index++ {
		start := index * chunkSize
		end := start + chunkSize
		if index == CHUNKS-1 {
			end = len(data)
		}

		waitGroup.Add(1)
		go func(chunkIndex int, chunk []int) {
			defer waitGroup.Done()
			if len(chunk) == 0 {
				maxima[chunkIndex] = data[0]
				return
			}
			maxima[chunkIndex] = maximum(chunk)
		}(index, data[start:end])
	}

	waitGroup.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}
