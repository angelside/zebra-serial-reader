package main

import (
	"encoding/csv"
	"os"
)

// csvCreateFileWithHeaders is responsible for creating a new CSV file with the specified filename.
// And it creates the header row in the CSV file and returns the writer for further use.
// If the file already exists, it will be overwritten with an empty file. If the file doesn't exist, a new file will be created.
func csvCreateFileWithHeaders(filename string) (*csv.Writer, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(file)

	// Write header row
	err = writer.Write([]string{"Location", "Serial number", "Product", "IP"})
	if err != nil {
		writer.Flush()
		file.Close()

		return nil, err
	}

	return writer, nil
}

// csvInsertRow adds a new data row to the CSV file using the provided writer.
func csvInsertRow(writer *csv.Writer, location, serialNumber, product, ip string) error {
	err := writer.Write([]string{location, serialNumber, product, ip})
	if err != nil {
		return err
	}

	return nil
}
