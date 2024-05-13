package main

import (
	"fmt"
    "strings"
)

// TODO: CLI colors
// TODO: Settings file

func main() {
	JsonFilename := "data.json" // FIXME: hardcoded name
	CsvFilename := "__printers.csv" // FIXME: hardcoded name

	// App title
	fmt.Println("== Zebra serial number extractor\n")

	// Read data from json
	printers, err := readJSONFile(JsonFilename)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)

		return // Terminate if we can't read the json file
	}

	// Create CSV
	writer, err := csvCreateFileWithHeaders(CsvFilename)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)

		return // Terminate if we can't create a csv file
	}

	// The loop
	for _, printer := range printers {
		// Get the serial number
        commandSerial := "! U1 getvar \"device.unique_id\"\r\n"
		serialNumber, err := getSocketAnswer(printer.IP, commandSerial)

        // Get the product name
        commandProduct := "~HI"
		productString, err := getSocketAnswer(printer.IP, commandProduct)
		productName := splitProductName(productString)

		// Print a stats text
		if err == nil {
			fmt.Printf("[OK] Location: %s, Serial number: %s, Product: %s, IP: %s\n", printer.Location, serialNumber, productName, printer.IP)
			// fmt.Printf(serialNumber) // Serial number
		} else {
			// fmt.Printf("[ERROR]: getSocketAnswer, %s\n", err) // [ERROR]: getSocketAnswer, dial tcp4 172.18.191.20:9100: i/o timeout
			fmt.Printf("[ERROR] Location: %s, IP: %s\n", printer.Location, printer.IP)
			serialNumber = "Nil" // This holding serial number if we fail to get it
		}

		// Use the writer to add data to the CSV file
		err = csvInsertRow(writer, printer.Location, serialNumber, productName, printer.IP)
		if err != nil {
			fmt.Println("Error adding to CSV file:", err)

			return // Terminate if we can't add data to the csv file
		}
	}

	// Close csv the file
	writer.Flush() // Write everything to the csv file, until this point csv file is empty
	err = writer.Error()
	if err != nil {
		fmt.Println("Error flushing CSV writer:", err)

		return // Terminate if we can't close the csv file
	}
	fmt.Println("\nCSV file created successfully:", CsvFilename)
}

//
// Helpers
//

func splitProductName(productString string) string {
	productNameParts := strings.Split(productString, "-")
	//fmt.Println(productNameParts[0])
	var productName string

	if len(productNameParts) > 0 {
		productName = productNameParts[0]
		//fmt.Println(productName) // ZD620
	} else {
		productName = productString
	}

	return productName
}
