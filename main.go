package main

import (
	"fmt"
)

func main() {
	var x [5]int
	for i := 0; i <= 2; i++ {
		fmt.Println("Enter the value of x[", i, "]")
		if x[0] == x[4] && x[2] == x[3] {
			fmt.Println(fmt.Sprintf("if %d == %d && %d == %d", x[0], x[4], x[2], x[3]))
		}

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
