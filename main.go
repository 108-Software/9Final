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

	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = rand.Int() 
	}

	return slice
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := 0
	for _, value := range data {
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

	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)
	var waitGroup sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		waitGroup.Add(1)

		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize
		if i == CHUNKS-1 {
			endIndex = len(data)
		}

		chunk := data[startIndex:endIndex]
		
		go func(chunkIndex int, chunkData []int) {
			defer waitGroup.Done()
			
			maxValues[chunkIndex] = maximum(chunkData)
		}(i, chunk)
	}

	waitGroup.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n\n", SIZE)
	
	startGeneration := time.Now()
	randomData := generateRandomElements(SIZE)
	generationTime := time.Since(startGeneration).Microseconds()
	
	if len(randomData) == 0 {
		fmt.Println("Ошибка: не удалось сгенерировать данные")
		return
	}
	
	fmt.Printf("Генерация данных заняла: %d микросекунд\n", generationTime)


	fmt.Println("Ищем максимальное значение в один поток")
	startSingleThread := time.Now()
	maxSingle := maximum(randomData)
	elapsedSingle := time.Since(startSingleThread).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\n", maxSingle)
	fmt.Printf("Время поиска: %d микросекунд\n\n", elapsedSingle)


	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startMultiThread := time.Now()
	maxMulti := maxChunks(randomData)
	elapsedMulti := time.Since(startMultiThread).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\n", maxMulti)
	fmt.Printf("Время поиска: %d микросекунд\n\n", elapsedMulti)

}
