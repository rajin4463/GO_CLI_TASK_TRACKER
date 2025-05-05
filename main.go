package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to Gopher Track!. A CLI task tracker")
	fmt.Println(csvread("tasks.csv"))
}