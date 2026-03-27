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

func guessDatasetTypes(guessRows [][]string) []string {
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
	percentile5 := int(math.Ceil(float64(len(csvValues)) * 0.05))
	dataTypes := guessDatasetTypes(csvValues[:percentile5]) // Samples 5% of dataset
	types := map[string]reflect.Type{}
	for i := range len(header) {
		types[header[i]] = TypeTranslate[dataTypes[i]]
	}

	// Put in the correct containers
	var dummy any
	rows, cols := len(csvValues), len(csvValues[0])
	matrixData := []Vector{}
	for j := range cols { // cols
		vecData := make([]any, rows)
		for i := range rows { // rows
			dummy = csvValues[i][j]
			vecData = append(vecData, dummy)
		}

		vec := Vector{
			Data:     vecData,
			dataType: TypeTranslate[dataTypes[j]],
		}
		matrixData = append(matrixData, vec)
	}

	return Dataset{}
}

func DatasetFromCSV(filename string, sep rune, hasHeader bool) Dataset {
	return Dataset{}
}
