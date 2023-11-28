package main

import "fmt"

func RemoveIndex[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		//索引越界返回原值
		return slice
	}
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func main() {
	s1 := []int{1, 2, 3}
	s1_removed := RemoveIndex(s1, 0)
	fmt.Printf("removed slice is: %v", s1_removed)

}
