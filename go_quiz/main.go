package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func readCSV(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

func parseCSV(data []byte) (*csv.Reader, error) {
	r := csv.NewReader(bytes.NewReader(data))
	return r, nil
}

func processCSV(reader *csv.Reader) []string {
	var vals []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading CSV data:", err)
			break
		}
		// fmt.Println(record)
		vals = append(vals, record...)
	}
	return vals
}

func main() {
	var filename string
	fmt.Println("Enter the name of file: ")
	fmt.Scanln(&filename)

	data, err := readCSV(filename)
	if err != nil {
		log.Fatal(err)
	}

	reader, err := parseCSV(data)
	if err != nil {
		log.Fatal(err)
	}

	record := processCSV(reader)

	var count int = 0

	for i := 0; i < len(record); i += 2 {
		var userans string
		fmt.Printf("%s?: ", record[i])
		fmt.Scanln(&userans)
		if userans == record[i+1] {
			count++
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Total questions asked: %d", len(record)/2)
	fmt.Printf("\nMarks Scored: %d", count)

}
