package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
)

type Vector struct {
	Data     []any
	dataType reflect.Type
}

type Matrix struct {
	Data      map[string]Vector
	dataTypes map[string]reflect.Type
}

type Dataset struct {
	Target Vector
	Data   Matrix
	Source string
}

var TypeTranslate = map[string]reflect.Type{
	"int":    reflect.TypeOf(math.MaxInt),
	"string": reflect.TypeOf("string"),
	"bool":   reflect.TypeOf(true),
	"float":  reflect.TypeOf(math.MaxFloat64),
}

func printMatrix(matrix [][]string) {
	for _, record := range matrix {
		fmt.Println(record, len(record))
	}
}

func guessDatasetTypes(guessRows [][]string) []string {
	printMatrix(guessRows)
	rows, cols := len(guessRows), len(guessRows[0])

	fmt.Println(rows, cols)

	predicts := [][]string{}
	for _, row := range guessRows {
		rowTypePredict := make([]string, 0)
		for _, val := range row {
			_, errFloat := strconv.ParseFloat(val, 64)
			_, errInt := strconv.Atoi(val)
			if errInt == nil && errFloat == nil { // All floats are ints, but not all ints are floats
				rowTypePredict = append(rowTypePredict, "int")
				continue
			} else if errInt == nil {
				rowTypePredict = append(rowTypePredict, "int")
				continue
			} else if errFloat == nil {
				rowTypePredict = append(rowTypePredict, "float")
				continue
			}

			_, err := strconv.ParseBool(val)
			if err == nil {
				rowTypePredict = append(rowTypePredict, "bool")
				continue
			}

			rowTypePredict = append(rowTypePredict, "string")
		}

		predicts = append(predicts, rowTypePredict)
	}

	printMatrix(predicts)

	// Does a voting
	choice, choiceCount, ballot := "", 0, make(map[string]int)
	finalPredicts := []string{}
	for j := range cols {
		for i := range rows {
			ballot[predicts[i][j]]++
		}

		for typer, votes := range ballot {
			if votes > choiceCount {
				choiceCount = votes
				choice = typer
			}
		}

		finalPredicts = append(finalPredicts, choice)

		// Reset vars
		choiceCount = 0
		ballot = make(map[string]int)
	}

	return finalPredicts
}

func ReadCSV(filename string) Dataset {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("ERROR: Error trying to open the file during CSV Parsing: ", err)
	}
	defer csvFile.Close()

	// Read CSV file
	csvValues, err := csv.NewReader(csvFile).ReadAll()
	header := csvValues[0]
	csvValues = csvValues[1:]


	// Types treatment
	dataTypes := guessDatasetTypes(csvValues[:math.Ceil(len(csvValues) * 0.05)]) // Samples 5% of dataset
	types := map[string]reflect.Type{}
	for i := range len(header) {
		types[header[i]] = TypeTranslate[dataTypes[i]]
	}

	// Put in the correct containers
	var dummy any 
	rows, cols := len(csvValues), len(csvValues[0])
	matrixData := [][]Vector{}
	for j := range cols { // cols
		vecData := make([]any, rows)
		for i := range rows { // rows
			dummy = csvValues[i][j]
		}

		vec := Vector{
			Data: vecData,
			dataType: TypeTranslate[dataTypes[j]],
		}
		matrixData = append(matrixData, vec)
	}



	



	return Dataset{
		dataTypes: 
	}
}

func DatasetFromCSV(filename string, sep rune, hasHeader bool) Dataset {
	return Dataset{}
}
