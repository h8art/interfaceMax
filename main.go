package main

import "fmt"

func findMax(s []interface{}, f func(interface{}, interface{}) interface{}) interface{} {
	max := s[0]
	for _, v := range s {
		max = f(max, v)
	}
	return max
}

func main() {
	s := []interface{}{1, 2, 3}

	fmt.Println(findMax(s, func(max interface{}, v interface{}) interface{} {
		if max.(int) < v.(int) {
			max = v
		}
		return max
	}))
}
