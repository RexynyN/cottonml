package main

import (
	"fmt"
	"reflect"
)

var TypeTranslate = map[string]reflect.Type{
	"int":    reflect.TypeOf(int64(0)),
	"string": reflect.TypeOf("string"),
	"bool":   reflect.TypeOf(true),
	"float":  reflect.TypeOf(float64(0.0)),
}

type Data interface {
	int64 | float64 | string | bool
}

type Number interface {
	int64 | float64
}

type Label interface {
	int64 | float64 | string
}

type Vector struct {
	Data     []any
	dataType reflect.Type
}

type DataFrame struct {
	Data      map[string]any
	dataTypes map[string]reflect.Type
}

type Matrix struct {
	Data      [][]any
	dataTypes []reflect.Type
}

type Dataset struct {
	Target Vector
	Data   Matrix
	Source string
}

func (m *Matrix) ToString() {
	fmt.Println(len(m.Data), "x", len(m.Data[0]))
	for _, record := range m.Data {
		fmt.Println(record)
	}
}

func (m *Matrix) GetColumn(colIndex int) ([]any, error) {
	if colIndex < 0 || colIndex >= len(m.Data[0]) {
		return nil, fmt.Errorf("column index out of range")
	}

	targetType := m.dataTypes[colIndex]
	return m.getColumnTransform(colIndex, targetType)
}

func (m *Matrix) GetColumnAsType(colIndex int, toType string) ([]any, error) {
	if targetType, ok := TypeTranslate[toType]; ok {
		return m.getColumnTransform(colIndex, targetType)
	}

	return nil, fmt.Errorf("type not mapped, provide a type given in 'TypeTranslate'")
}

func (m *Matrix) getColumnTransform(colIndex int, targetType reflect.Type) ([]any, error) {
	if colIndex < 0 || colIndex >= len(m.Data[0]) {
		return nil, fmt.Errorf("column index out of range")
	}

	column := make([]any, len(m.Data))
	for i, record := range m.Data {
		if colIndex >= len(record) {
			return nil, fmt.Errorf("inconsistent column length at row %d", i)
		}

		val := record[colIndex]
		converted := reflect.ValueOf(val).Convert(targetType).Interface()
		column[i] = converted
	}

	return column, nil
}

func main() {
	// Create a sample Matrix with int64 and float64 data
	matrix := &Matrix{
		Data: [][]any{
			{1, 2.5, "a"},
			{2, 3.5, "b"},
			{3, 4.5, "c"},
		},
	}

	// Get the first column and convert to int64
	column, err := matrix.GetColumnAsType(0, "float")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Column 0:", column)

	// Get the second column and convert to float64
	column, err = matrix.GetColumnAsType(1, "int")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Column 1:", column)

	// Print matrix
	matrix.ToString()
}
