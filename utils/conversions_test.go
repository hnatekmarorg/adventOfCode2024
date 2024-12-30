package utils

import (
	"reflect"
	"strconv"
	"testing"
)

func TestConvertRow(t *testing.T) {
	matrix := [][]string{{"1", "2"}, {"3", "4"}}
	expected := []int{1, 3}
	result := ConvertColumn(matrix, func(cell string) int {
		val, err := strconv.Atoi(cell)
		if err != nil {
			panic(err)
		}
		return val
	}, 0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
