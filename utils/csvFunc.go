package utils

import (
	"encoding/csv"
	"strconv"
	"fmt"
	"log"
	"os"
)

func CsvRead(fileName string) [][]string {
	tasks, _ := os.Open(fileName)
	task, err := csv.NewReader(tasks).ReadAll()
	if err != nil {
	}
	return task
}

func CsvWrite(fileName string, data []string) {
	csvFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	for _, task := range data {
		_ = csvwriter.Write([]string{task})
	}
}

func CsvDel(fileName string, id int) {
	data := CsvRead(fileName)
	var NewData [][]string
	for i := 1; i < len(data); i++ { // Start from 1 to skip the header
		fmt.Println(data[i])
		Listid, _ := strconv.Atoi(data[i][0])
		if Listid == id{
			
		}
	}
}
