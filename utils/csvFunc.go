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

func CsvWrite(fileName string, data [][]string) {
	// Read the existing CSV file to get the header
	existingData := CsvRead(fileName)

	// Check if the file exists and has a header
	if len(existingData) > 0 {
		header := existingData[0]
		newData := [][]string{header}
		newData = append(newData, data...)
		csvFile, err := os.Create(fileName) // Create/overwrite the file
		if err != nil {
			log.Fatalf("Failed creating file: %s", err)
		}
		defer csvFile.Close()

		csvwriter := csv.NewWriter(csvFile)
		err = csvwriter.WriteAll(newData) // Write all rows at once
		if err != nil {
			log.Fatalf("Error writing CSV data: %s", err)
		}
		csvwriter.Flush()
	} else {
		defaultHeader := []string{"ID", "Title", "Priority", "Status", "Due Date", "Assigned To"}
		newData := [][]string{defaultHeader}
		newData = append(newData, data...)
		csvFile, err := os.Create(fileName) // Create/overwrite the file
		if err != nil {
			log.Fatalf("Failed creating file: %s", err)
		}
		defer csvFile.Close()

		csvwriter := csv.NewWriter(csvFile)
		err = csvwriter.WriteAll(newData) // Write all rows at once
		if err != nil {
			log.Fatalf("Error writing CSV data: %s", err)
		}
		csvwriter.Flush()
	}
}

func CsvDel(fileName string) {
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

	CsvWrite("tasks.csv",headerOnly)
}
