package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

// To generate a RSA Key pair (public and private keys)
func main() {
	key := savePrivateKey();
	result := savePublicKey(key);

	if (result) {
		log.Println("Successfully saved the private and public keys in private.pem and public.pem files!")
	} else {
		log.Println("Something went wrong in creating the private and public keys.")
	}
}

func savePrivateKey() *rsa.PrivateKey {
	// Generate a new RSA key pair
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalln("Error in generating RSA key-pair", err)
	}

	// Print the private key in PEM format
	privateKey := key;
	privateKeyPEM := x509.MarshalPKCS1PrivateKey(privateKey)
	// Block of PEM-encoded data.
	privateKeyBlock := pem.Block {
		Type: "RSA PRIVATE KEY",
		Bytes: privateKeyPEM,
	}
	privateKeyPEMEncoded := pem.EncodeToMemory(&privateKeyBlock)

	// Write the private key to a file.
	err = ioutil.WriteFile("private.pem", privateKeyPEMEncoded, 0600)
	if err != nil {
		log.Fatalln("Error in writing private key to a .pem file: ", err)
	}

	return key;
}

func savePublicKey(key *rsa.PrivateKey) bool {
	// Print the public key in PEM format
	publicKey := &key.PublicKey
	publicKeyPEM, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatalln("Error in generating public key in PEM format: ", err)
	}
	// Block of PEM-encoded data.
	publicKeyBlock := pem.Block {
		Type: "RSA PUBLIC KEY",
		Bytes: publicKeyPEM,
	}

	publicKeyPEMEncoded := pem.EncodeToMemory(&publicKeyBlock)
	// Write the private key to a file.
	err = ioutil.WriteFile("public.pem", publicKeyPEMEncoded, 0600)
	if err != nil {
		log.Fatalln("Error in writing private key to a .pem file: ", err)
		return false;
	}
	return true;
}