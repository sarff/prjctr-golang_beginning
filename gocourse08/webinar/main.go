/*
написати генерік функції:
створення посилання на простий тип PtrOf() (int → *int)
пошук у слайсі елементу FindIn()
видалення зайвих елементів у слайсі з метою залишити лише унікальні Unique()

	func PtrOf(v int) *int {
	    return &v
	}
*/
package main

import "fmt"

func PtrOf[T int | float64](v T) *T {
	return &v
}

func FindIn[T comparable](sl []T, find T) int {
	for i, v := range sl {
		if v == find {
			return i
		}
	}
	return -1
}

func Unique[T comparable](sl []T) []T {
	array := make(map[T]struct{})
	for _, v := range sl {
		array[v] = struct{}{}
	}
	uniqueSlice := make([]T, 0, len(array))

	for key := range array {
		uniqueSlice = append(uniqueSlice, key)
	}

	return uniqueSlice
}

func main() {
	fmt.Println(PtrOf(10.01))
	fmt.Println(PtrOf(5))
	fmt.Println(FindIn([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(FindIn([]string{"s", "v", "a", "k", "o"}, "o"))
	fmt.Println(Unique([]any{1, 2, 3, 4, 5, 3, 3, 4, 6, 1, "st", "st", "1", "2"}))
}
