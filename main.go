package main

import "fmt"

func main() {

	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Println(a)
		a, b = b, a+b
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
