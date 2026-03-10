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
	int | int8 | int16 | int32 | int64 | float32 | float64 | string | bool
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Label interface {
	int64 | float64 | string
}

type Row struct {
	Data     []any
	dataType reflect.Type
}

type Vector[T Data] []T

type TypedDataFrame[T Data] map[string][]T

type Matrix[T Data] [][]T

type DataFrame struct {
	Data      map[string][]any
	dataTypes map[string]reflect.Type
}

func ColumnAsType[T Data](d *DataFrame, col string) Vector[T] {
	data, ok := d.Data[col]
	if !ok {
		panic("ERROR: " + col + " is not a column in the DataFrame")
	}

	vector := make(Vector[T], 0)
	for idx, value := range data {
		aux, ok := value.(T)
		if !ok {
			panic(fmt.Sprintf("ERROR: %a could not be converted to given type at index %d", value, idx))
		}

		vector = append(vector, aux)
	}

	return vector
}

func ToTypedDataFrame[T Data](d *DataFrame) TypedDataFrame[T] {
	tData := make(TypedDataFrame[T])
	for key, data := range d.Data {
		newData := make([]T, 0)
		for idx, value := range data {
			newValue, ok := value.(T)
			if !ok {
				panic(fmt.Sprintf("ERROR: %a could not be converted to given type at column %s, index %d", value, key, idx))
			}
			newData = append(newData, newValue)
		}
		tData[key] = newData
	}

	return tData
}

func (d *DataFrame) AsFloat(col string) Vector[float64] {
	return d.AsFloat64(col)
}

func (d *DataFrame) AsFloat32(col string) Vector[float32] {
	return ColumnAsType[float32](d, col)
}

func (d *DataFrame) AsFloat64(col string) Vector[float64] {
	return ColumnAsType[float64](d, col)
}

func (d *DataFrame) AsString(col string) Vector[string] {
	return ColumnAsType[string](d, col)
}

func (d *DataFrame) AsInt(col string) Vector[int] {
	return ColumnAsType[int](d, col)
}

func (d *DataFrame) AsInt8(col string) Vector[int8] {
	return ColumnAsType[int8](d, col)
}

func (d *DataFrame) AsInt16(col string) Vector[int16] {
	return ColumnAsType[int16](d, col)
}

func (d *DataFrame) AsInt32(col string) Vector[int32] {
	return ColumnAsType[int32](d, col)
}

func (d *DataFrame) AsInt64(col string) Vector[int64] {
	return ColumnAsType[int64](d, col)
}

func (d *DataFrame) AsBool(col string) Vector[bool] {
	return ColumnAsType[bool](d, col)
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

// func main() {
// 	// Create a sample Matrix with int64 and float64 data
// 	matrix := &Matrix{
// 		Data: [][]any{
// 			{1, 2.5, "a"},
// 			{2, 3.5, "b"},
// 			{3, 4.5, "c"},
// 		},
// 	}

// 	// Get the first column and convert to int64
// 	column, err := matrix.GetColumnAsType(0, "float")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Column 0:", column)

// 	// Get the second column and convert to float64
// 	column, err = matrix.GetColumnAsType(1, "int")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Column 1:", column)

// 	// Print matrix
// 	matrix.ToString()
// }
