package utils

import (
	"encoding/csv"
	"os"
	"log"
)

func CsvRead(fileName string) [][]string {
	tasks, _ := os.Open(fileName)
	task, err := csv.NewReader(tasks).ReadAll()
	if err != nil {}
	return task
}

func CsvWrite(fileName string, data []string){
	csvFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	for _, task := range data{
		_ = csvwriter.Write([]string{task})
	}
}