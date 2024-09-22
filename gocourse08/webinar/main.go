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

func PtrOf[T any](v T) *T {
	return &v
}

func FindIn[T comparable](s []T, target T) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}

func Unique[T comparable](s []T) []T {
	set := make(map[T]struct{})
	for _, v := range s {
		set[v] = struct{}{}
	}
	unique := make([]T, 0, len(set))

	for key := range set {
		unique = append(unique, key)
	}

	return unique
}

func main() {
	fmt.Println(PtrOf(10.01))
	fmt.Println(PtrOf(5))
	fmt.Println(FindIn([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(FindIn([]string{"s", "v", "a", "k", "o"}, "o"))
	fmt.Println(Unique([]any{1, 2, 3, 4, 5, 3, 3, 4, 6, 1, "st", "st", "1", "2"}))
}
