package main

import (
	"encoding/json"
	"os"
)

type Printer struct {
	Location string `json:"location"`
	IP       string `json:"ip"`
}

// readJSONFile function reads all data from json file and assigns data to the []Printer struct
func readJSONFile(JsonFilename string) ([]Printer, error) {
	// Read the JSON file
	data, err := os.ReadFile(JsonFilename)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into a map
	var printerMap map[string]string
	err = json.Unmarshal(data, &printerMap)
	if err != nil {
		return nil, err
	}

	// Convert the map into a slice of Printer structs
	var printers []Printer
	for key, value := range printerMap {
		printer := Printer{
			Location: key,
			IP:       value,
		}
		printers = append(printers, printer)
	}

	return printers, nil
}
