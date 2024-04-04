#ifndef FOURQSIGN_H
#define FOURQSIGN_H

void signTransaction(const unsigned char* subseed, const unsigned char* publicKey, const unsigned char* messageDigest, unsigned char* signature);

#endif