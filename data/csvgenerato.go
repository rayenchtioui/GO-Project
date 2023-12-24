package data

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

func WriteCSV[T any](fileName string, data []T) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	t := reflect.TypeOf(data[0])
	var headers []string
	for i := 0; i < t.NumField(); i++ {
		headers = append(headers, t.Field(i).Name)
	}
	if err := writer.Write(headers); err != nil {
		log.Fatalf("Could not write headers to CSV: %v", err)
	}

	// Write data
	for _, record := range data {
		var row []string
		v := reflect.ValueOf(record)
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			var val string
			switch field.Kind() {
			case reflect.String:
				val = field.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				val = strconv.FormatInt(field.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				val = strconv.FormatUint(field.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				val = strconv.FormatFloat(field.Float(), 'f', -1, 64)
			case reflect.Struct:
				if t, ok := field.Interface().(time.Time); ok {
					val = t.Format(time.RFC3339)
				}
			}
			row = append(row, val)
		}
		if err := writer.Write(row); err != nil {
			log.Fatalf("Could not write record to CSV: %v", err)
		}
	}
}
