package schnorrq

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lfourq-qubic -Wl,-rpath,$ORIGIN
#include "fourq-qubic.h"
*/
import "C"
import (
	"github.com/pkg/errors"
	"github.com/qubic/go-node-connector/types"
)

func Sign(seed string, pubKey [32]byte, messageDigest [32]byte) ([64]byte, error) {
	subseed, err := types.GetSubSeed(seed)
	if err != nil {
		return [64]byte{}, errors.Wrap(err, "getting subseed")
	}
	// Call the C function to sign the transaction
	var signature [64]C.uchar
	C.sign((*C.uchar)(&subseed[0]), (*C.uchar)(&pubKey[0]), (*C.uchar)(&messageDigest[0]), &signature[0])

	// Convert the signature from C type to Go type
	var signatureGo [64]byte
	for i := 0; i < len(signatureGo); i++ {
		signatureGo[i] = byte(signature[i])
	}

	return signatureGo, nil
}

func Verify(pubKey [32]byte, messageDigest [32]byte, signature [64]byte) bool {
	ok := C.verify((*C.uchar)(&pubKey[0]), (*C.uchar)(&messageDigest[0]), (*C.uchar)(&signature[0]))

	return bool(ok)
}
