package main

import (
	"fmt"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{
			name:     "Положительный размер",
			size:     10,
			expected: 10,
		},
		{
			name:     "Большой размер",
			size:     1000,
			expected: 1000,
		},
		{
			name:     "Размер 1",
			size:     1,
			expected: 1,
		},
		{
			name:     "Нулевой размер",
			size:     0,
			expected: 0,
		},
		{
			name:     "Отрицательный размер",
			size:     -5,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			if len(result) != tt.expected {
				t.Errorf("generateRandomElements(%d) вернул слайс длины %d, ожидалось %d",
					tt.size, len(result), tt.expected)
			}

			if tt.size > 0 {
				for i, val := range result {
					if val < 0 || val >= 1000 {
						t.Errorf("Элемент по индексу %d равен %d, ожидался диапазон [0, 999]", i, val)
					}
				}

				if result == nil {
					t.Error("generateRandomElements() вернул nil слайс для положительного размера")
				}
			} else {
				if result != nil && len(result) != 0 {
					t.Error("generateRandomElements() должен возвращать пустой слайс для неположительного размера")
				}
			}
		})
	}
}

func TestGenerateRandomElements_DiapazonZnacheniy(t *testing.T) {
	size := 10000
	result := generateRandomElements(size)

	uvidennyeZnacheniya := make(map[int]bool)
	vseVRamkah := true

	for _, val := range result {
		if val < 0 || val >= 1000 {
			vseVRamkah = false
			break
		}
		uvidennyeZnacheniya[val] = true
	}

	if !vseVRamkah {
		t.Errorf("Не все значения в диапазоне [0, 999]")
	}

	if len(uvidennyeZnacheniya) <= 1 && size > 10 {
		t.Errorf("Ожидалось разнообразие случайных значений, получено только %d уникальных значений", len(uvidennyeZnacheniya))
	}
}

func TestGenerateRandomElements_Soglasovannost(t *testing.T) {
	for i := 0; i < 5; i++ {
		size := 50 + i*10
		result := generateRandomElements(size)
		if len(result) != size {
			t.Errorf("Итерация %d: ожидалась длина %d, получено %d", i, size, len(result))
		}
	}
}

func TestGenerateRandomElements_KraynieSluchai(t *testing.T) {
	kraynieSluchai := []struct {
		size     int
		expected int
		desc     string
	}{
		{size: -1, expected: 0, desc: "Минус один"},
		{size: -100, expected: 0, desc: "Большое отрицательное"},
		{size: 0, expected: 0, desc: "Ноль"},
		{size: 1, expected: 1, desc: "Один"},
		{size: 2, expected: 2, desc: "Два"},
	}

	for _, tc := range kraynieSluchai {
		t.Run(tc.desc, func(t *testing.T) {
			result := generateRandomElements(tc.size)
			if len(result) != tc.expected {
				t.Errorf("Для размера %d: ожидалась длина %d, получено %d", tc.size, tc.expected, len(result))
			}
		})
	}
}

func BenchmarkGenerateRandomElements(b *testing.B) {
	razmery := []int{10, 100, 1000, 10000, 100000}

	for _, size := range razmery {
		b.Run(fmt.Sprintf("Размер_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				generateRandomElements(size)
			}
		})
	}
}
