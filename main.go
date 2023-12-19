package main

/*
	#cgo CFLAGS: -I./src
	#cgo LDFLAGS: -L./lib -ldecode -Wl,-rpath=./lib -lsox
	#include <stdlib.h>
	#include <sox.h>
   	#include "libdecode.h"
*/
import "C"
import (
	"fmt"
	"io/ioutil"
	"unsafe"
)

func ulaw_to_wav(fileBuffer []byte, outputFileName string) int {
	cOutputFileName := C.CString(outputFileName)
	defer C.free(unsafe.Pointer(cOutputFileName))

	// Convert Go byte slice to C char*
	cBuffer := C.CBytes(fileBuffer)
	defer C.free(unsafe.Pointer(cBuffer))

	// Call the C function
	result := int(C.decode((*C.char)(cBuffer), C.size_t(len(fileBuffer)), cOutputFileName))

	return result
}

func main() {
	// Read the contents of the .ul file
	filePath := "./examples/input.ul"
	outputFileName := "./examples/output.wav"
	fileBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	result := ulaw_to_wav(fileBuffer, outputFileName)

	// Print the result
	fmt.Printf("Result from C library: %d\n", result)
}
