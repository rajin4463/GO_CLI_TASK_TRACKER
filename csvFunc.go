package main

import (
	"encoding/csv"
	"os"
)

func csvread(fileName string) [][]string {
	tasks, _ := os.Open(fileName)
	task, err := csv.NewReader(tasks).ReadAll()
	if err != nil {}
	return task
}