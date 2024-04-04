package schnorrq

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L./lib -lfourqsign -Wl,-rpath=./lib
#include "fourqsign.h"

void signTransaction(const unsigned char* subseed, const unsigned char* publicKey, const unsigned char* messageDigest, unsigned char* signature);
*/
import "C"

func Sign(subseed [32]byte, pubKey [32]byte, messageDigest [32]byte) [64]byte {
	// Call the C function to sign the transaction
	var signature [64]C.uchar
	C.signTransaction((*C.uchar)(&subseed[0]), (*C.uchar)(&pubKey[0]), (*C.uchar)(&messageDigest[0]), &signature[0])

	// Convert the signature from C type to Go type
	var signatureGo [64]byte
	for i := 0; i < len(signatureGo); i++ {
		signatureGo[i] = byte(signature[i])
	}

	return signatureGo
}

//func Verify(pubKey [32]byte, messageDigest [32]byte, signature [64]byte) bool {
//	ok := C.verify((*C.uchar)(&pubKey[0]), (*C.uchar)(&messageDigest[0]), (*C.uchar)(&signature[0]))
//
//	return bool(ok)
//}
