package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "io/ioutil"
)

func main() {
    privateKeyPEM, err := ioutil.ReadFile("private.pem")
    if err != nil {
        panic(err)
    }
    privateKeyBlock, _ := pem.Decode(privateKeyPEM)
    privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
    if err != nil {
        panic(err)
    }

    ciphertext := []byte{0x88, 0xaa, 0x63, 0x24, 0x2d, 0x48, 0xfd, 0xb1, 0x63, 0x71, 0x33, 0x17, 0x2a, 0x01, 0xce, 0x15, 0x1b, 0x25, 0xac, 0xcd, 0x35, 0xc1, 0x7c, 0x2a, 0x48, 0x58, 0x79, 0xae, 0x73, 0xf3, 0x5e, 0xc9, 0x89, 0xa7, 0x8a, 0x92, 0xa4, 0x3f, 0x3d, 0xb3, 0x43, 0x1d, 0x01, 0x74, 0xee, 0xd1, 0x1e, 0x95, 0x2b, 0x4f, 0x42, 0x46, 0x0b}
    plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Decrypted: %s\n", plaintext)
}
