package schnorrq

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
*/
import "C"
import (
	"path/filepath"
	"runtime"
	"unsafe"
)

// Declare function prototypes for dynamically loaded functions
type signFunc func(subseed, publicKey, messageDigest, signature unsafe.Pointer)
type verifyFunc func(publicKey, messageDigest, signature unsafe.Pointer) C.int

var (
	sign   signFunc
	verify verifyFunc
)

func init() {
	// Get the path of the current file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	// Load the shared library dynamically
	lib := C.dlopen(C.CString(filepath.Join(dir, "libfourq-qubic.so")), C.RTLD_LAZY)
	if lib == nil {
		panic("Failed to load shared library")
	}

	// Load the necessary functions from the shared library
	signTransactionSymbol := C.CString("sign")
	verifySymbol := C.CString("verify")
	sign = (signFunc)(C.dlsym(lib, signTransactionSymbol))
	verify = (verifyFunc)(C.dlsym(lib, verifySymbol))

	// Free the C strings used for symbol names
	C.free(unsafe.Pointer(signTransactionSymbol))
	C.free(unsafe.Pointer(verifySymbol))
}

// Sign function calls the dynamically loaded sign function
func Sign(subseed [32]byte, pubKey [32]byte, messageDigest [32]byte) [64]byte {
	var signature [64]byte
	sign(unsafe.Pointer(&subseed[0]), unsafe.Pointer(&pubKey[0]), unsafe.Pointer(&messageDigest[0]), unsafe.Pointer(&signature[0]))
	return signature
}

// Verify function calls the dynamically loaded verify function
func Verify(pubKey [32]byte, messageDigest [32]byte, signature [64]byte) bool {
	ok := verify(unsafe.Pointer(&pubKey[0]), unsafe.Pointer(&messageDigest[0]), unsafe.Pointer(&signature[0]))
	return ok != 0
}
