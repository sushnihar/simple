package main

import "fmt"

func main() {

	fibonacciIterative(10)

}
func fibonacciIterative(n int) {
	n1, n2 := 0, 1
	if n >= 1 {
		fmt.Println(n1)
	}
	if n >= 2 {
		fmt.Println(n2)
	}
	for i := 3; i <= n; i++ {
		n3 := n1 + n2
		fmt.Println(n3)
		n1 = n2
		n2 = n3
	}
}

// #TODO: datatypes: int,float,string,bool
// #TODO: datastructures: array,slice,map,struct
// #TODO: programflow: if,else,for,switch,case
// #TODO: fibonaccii,palindrome

// ##### INTERMEDIATE #####
// #TODO: structs,pointers, methods, goroutines, gochannel
// #TODO: error handling, panic, defer
// #TODO: file handling, json, xml, csv

// ##### ADVANCED #####
// #TODO: testing, benchmarking
// #TODO: rest api,grpc
