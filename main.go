package main

import "fmt"

func findMax(s []interface{}) interface{} {
	max := s[0]
	switch s[0].(type) {
	case int:
		for _, v := range s {
			if max.(int) < v.(int) {
				max = v
			}
		}
	case string:
		for _, v := range s {
			if max.(string) < v.(string) {
				max = v
			}
		}
	}
	return max
}

func main() {
	s := []interface{}{1, 2, 3}
	fmt.Println(findMax(s))
}
