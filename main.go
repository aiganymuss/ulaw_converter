package main

import (
	"fmt"
	"github.com/aiganymuss/ulaw_converter/converter"
	"io/ioutil"
)

func main() {
	// Read the contents of the .ul file
	filePath := "./examples/input.ul"
	outputFileName := "./examples/output.wav"
	fileBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	result := converter.ConvertToWav(fileBuffer, outputFileName)

	// Print the result
	fmt.Printf("Result from C library: %d\n", result)
}
