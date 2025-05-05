package main

import (
	// "fmt"
	// "io"
	// "strings"
	"log"
	"os"
	// "encoding/csv"
)

func csvread(fileName string) string{
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}