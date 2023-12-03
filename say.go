package gomodproject

import "fmt"

// SayHi say Hi to someone
func SayHi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
