package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// https://github.com/dr5hn/countries-states-cities-database/blob/master/csv/states.csv
func main() {
	// Open the CSV file
	file, err := os.Open("states.csv")
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
	countryIndex := 4
	countryName := "United States"

	for _, record := range records[1:] {
		if record[countryIndex] == countryName {
			fmt.Printf("\"%s\",\n", record[nameIndex])
		}
	}
}
