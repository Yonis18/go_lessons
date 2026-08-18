package main

import "fmt"

func main() {
	agree("abdi")
	sum(265, 1)

}

func agree(name string) {
	fmt.Println("hi", name)

}

func sum(a int, b int) {
	result := a - b

	fmt.Println("the result", result)
}
