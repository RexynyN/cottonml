package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Vector []any

type Matrix map[string]Vector

type Dataset struct {
	Target Vector
	Data   Matrix
	Source string
}

func ReadCSV(filename string) Dataset {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("ERROR: Error trying to open the file during CSV Parsing: ", err)
	}
	defer csvFile.Close()

	csvFile.Seek(0, 0) // Reset the file pointer
	csvValues, err := csv.NewReader(csvFile).ReadAll()

	header := csvValues[0]
	fmt.Println("Header: ", header)
	csvValues = csvValues[1:]

	for _, record := range csvValues {
		fmt.Println(record)
	}

	return Dataset{}
}

func DatasetFromCSV(filename string, sep rune, hasHeader bool) Dataset {
	return Dataset{}
}
