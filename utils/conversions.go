package utils

func ConvertRow[T any](table [][]string, conversionFunc func(string) T, row int) []T {
	rowData := table[row]
	result := make([]T, len(rowData))
	for i, value := range rowData {
		result[i] = conversionFunc(value)
	}
	return result
}

// ConvertColumn processes a specific column from a 2D string slice based on a conversion function and returns the transformed column.
// table represents the input 2D slice of strings where the column is extracted.
// conversionFunc is a function that converts a string value into the generic type T.
// column specifies the index of the column to extract and transform, starting at 0.
func ConvertColumn[T any](table [][]string, conversionFunc func(string) T, column int) []T {
	columnData := make([]string, len(table))
	for i := range table {
		columnData[i] = table[i][column]
	}
	result := make([]T, len(columnData))
	for i, value := range columnData {
		result[i] = conversionFunc(value)
	}
	return result
}

// Wraps stdlib functions to allow generic usage for common conversions that return errors.
func StringConversionWrapper[T any](conversionFN func(x string) (result T, err error)) func(val string) T {
	return func(val string) T {
		if result, err := conversionFN(val); err == nil {
			return result
		} else {
			panic(err)
		}
	}
}
