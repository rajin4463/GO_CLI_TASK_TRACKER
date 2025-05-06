package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func CsvRead(fileName string) [][]string {
	tasks, _ := os.Open(fileName)
	task, err := csv.NewReader(tasks).ReadAll()
	if err != nil {
		fmt.Printf("Error while reading CSV: %v\n", err)
	}
	return task
}

func CsvWrite(fileName string, data []string) {
	csvFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	for _, task := range data {
		_ = csvwriter.Write([]string{task})
	}
	csvwriter.Flush()
}

func CsvDel(fileName string, id int) {
	// data := CsvRead(fileName)
	// var NewData [][]string
	// for i := 1; i < len(data); i++ { // Start from 1 to skip the header
	// 	Listid, _ := strconv.Atoi(data[i][0])
	// 	if Listid != id {
	// 		NewData = append(NewData, data[i])
	// 	}
	// }
	// fmt.Println(NewData)
	data := CsvRead(fileName)
	if len(data) == 0 {
		fmt.Errorf("empty or invalid CSV file")
		return 
	}

	// Keep only the header row
	headerOnly := [][]string{data[0]}

	// Create the file anew
	csvFile, err := os.Create(fileName)
	if err != nil {
		fmt.Errorf("failed creating file: %s", err)
		return 
	}
	defer csvFile.Close()

	// Write only the header back to the file
	csvwriter := csv.NewWriter(csvFile)
	err = csvwriter.WriteAll(headerOnly)
	if err != nil {
		fmt.Errorf("error writing CSV data: %s", err)
		return 
	}

	csvwriter.Flush()
	return 
}
