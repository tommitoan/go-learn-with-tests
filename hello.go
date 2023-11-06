//package main
//
//import "fmt"
//
//func Hello(name string) string {
//	return "Hello, " + name
//}
//
//func main() {
//	fmt.Println(Hello("world"))
//}

package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
