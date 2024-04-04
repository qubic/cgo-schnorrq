#include <stdbool.h>
#include <stdlib.h>

#ifndef FOURQ_QUBIC_H
#define FOURQ_QUBIC_H

void sign(const unsigned char* subSeed, const unsigned char* publicKey, const unsigned char* messageDigest, unsigned char* signature);
bool verify(const unsigned char* publicKey, const unsigned char* messageDigest, const unsigned char* signature);

#endif // End of include guard