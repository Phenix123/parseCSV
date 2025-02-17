package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// https://github.com/dr5hn/countries-states-cities-database/blob/master/csv/countries.csv

func main() {
	// Open the CSV file
	file, err := os.Open("countries.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	nameIndex := 1

	for _, record := range records[1:] {
		fmt.Printf("\"%s\",\n", record[nameIndex])
	}
}
