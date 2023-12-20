package converter

/*
	#cgo CFLAGS: -I../src
	#cgo LDFLAGS: -L../lib -ldecode -Wl,-rpath=./lib -lsox
	#include <stdlib.h>
	#include <sox.h>
   	#include "libdecode.h"
*/
import "C"
import (
	"unsafe"
)

func ConvertToWav(fileBuffer []byte, outputFileName string) int {
	cOutputFileName := C.CString(outputFileName)
	defer C.free(unsafe.Pointer(cOutputFileName))

	// Convert Go byte slice to C char*
	cBuffer := C.CBytes(fileBuffer)
	defer C.free(unsafe.Pointer(cBuffer))

	// Call the C function
	result := int(C.decode((*C.char)(cBuffer), C.size_t(len(fileBuffer)), cOutputFileName))

	return result
}
