package schnorrq

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lfourq-qubic
#include "fourq-qubic.h"
*/
import "C"
import (
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	// Get the path of the current file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	// Check if LD_LIBRARY_PATH is already set
	currLDPath := os.Getenv("LD_LIBRARY_PATH")
	if currLDPath == "" {
		// If LD_LIBRARY_PATH is not set, set it to the path of the shared library
		os.Setenv("LD_LIBRARY_PATH", dir)
	} else {
		// If LD_LIBRARY_PATH is already set, append the path of the shared library
		os.Setenv("LD_LIBRARY_PATH", currLDPath+":"+dir)
	}
}

func Sign(subseed [32]byte, pubKey [32]byte, messageDigest [32]byte) [64]byte {
	// Call the C function to sign the transaction
	var signature [64]C.uchar
	C.sign((*C.uchar)(&subseed[0]), (*C.uchar)(&pubKey[0]), (*C.uchar)(&messageDigest[0]), &signature[0])

	// Convert the signature from C type to Go type
	var signatureGo [64]byte
	for i := 0; i < len(signatureGo); i++ {
		signatureGo[i] = byte(signature[i])
	}

	return signatureGo
}

func Verify(pubKey [32]byte, messageDigest [32]byte, signature [64]byte) bool {
	ok := C.verify((*C.uchar)(&pubKey[0]), (*C.uchar)(&messageDigest[0]), (*C.uchar)(&signature[0]))

	return bool(ok)
}
