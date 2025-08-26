package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"Zero size", 0, 0},
		{"Positive size", 100, 100},
		{"Large size", 10000, 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			require.Len(t, result, tt.expected)
		})
	}
}

func TestGenerateRandomElements_Soglasovannost(t *testing.T) {
	t.Run("Consistency between runs", func(t *testing.T) {
	
		results := make([][]int, 3)
		for i := 0; i < 3; i++ {
			results[i] = generateRandomElements(100)
			require.Len(t, results[i], 100)
		}

		
		assert.NotEqual(t, results[0], results[1], "Результаты должны быть разными между запусками")
		assert.NotEqual(t, results[1], results[2], "Результаты должны быть разными между запусками")
		assert.NotEqual(t, results[0], results[2], "Результаты должны быть разными между запусками")
	})

	t.Run("Empty input consistency", func(t *testing.T) {
		result1 := generateRandomElements(0)
		result2 := generateRandomElements(0)
		
		assert.Empty(t, result1)
		assert.Empty(t, result2)
		assert.Equal(t, result1, result2)
	})
}

func TestGenerateRandomElements_KraynieSluchai(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"Negative size", -1},
		{"Very large size", 1000000},
		{"Size 1", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			
			if tt.size <= 0 {
				assert.Empty(t, result, "Для неположительного размера должен возвращаться пустой слайс")
			} else {
				assert.Len(t, result, tt.size, "Для положительного размера должен возвращаться слайс нужной длины")
				
				if tt.size == 1 {
					assert.True(t, result[0] >= 0, "Единственный элемент должен быть неотрицательным")
				}
			}
		})
	}
}

func BenchmarkGenerateRandomElements(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run("Size_%d", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				generateRandomElements(size)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{42}, 42},
		{"Multiple elements", []int{1, 5, 3, 9, 2}, 9},
		{"All zeros", []int{0, 0, 0}, 0},
		{"Negative numbers", []int{-5, -1, -3}, 0}, 
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Smaller than chunks", []int{1, 5, 3}, 5},
		{"Exactly chunks size", make([]int, CHUNKS), 0},
		{"Large data", makeLargeTestData(100), 999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func makeLargeTestData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}

	if size > 0 {
		data[size-1] = 999
	}
	return data
}

func TestIntegration(t *testing.T) {
	t.Run("Integration test", func(t *testing.T) {
	
		data := generateRandomElements(1000)
		require.Len(t, data, 1000)
		
		max1 := maximum(data)
		assert.True(t, max1 >= 0)
		
		max2 := maxChunks(data)
		assert.Equal(t, max1, max2)
	})
}
