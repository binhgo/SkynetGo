package crypto

import (
	"crypto/dsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"math/big"
	"os"
)

type MySignature struct {
	SignHash []byte
	R        *big.Int
	S        *big.Int
}

func NewKeyPair() *dsa.PrivateKey {

	params := new(dsa.Parameters)

	// see http://golang.org/pkg/crypto/dsa/#ParameterSizes
	if err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	privateKey := new(dsa.PrivateKey)
	privateKey.PublicKey.Parameters = *params
	dsa.GenerateKey(privateKey, rand.Reader) // this generates a public & private key pair

	return privateKey
}

func Sign(privateKey *dsa.PrivateKey, dataToSign string) *MySignature {

	// Sign
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, "This is the message to be signed and verified!")
	signHash := h.Sum(nil)

	r, s, err := dsa.Sign(rand.Reader, privateKey, signHash)
	if err != nil {
		fmt.Println(err)
	}

	// sig := r.Bytes()
	// sig = append(sig, s.Bytes()...)
	// fmt.Printf("Signature : %x\n", sig)

	return &MySignature{SignHash: signHash, R: r, S: s}
}

func Verify(pubKey *dsa.PublicKey, mySig *MySignature) bool {

	verifyStatus := dsa.Verify(pubKey, mySig.SignHash, mySig.R, mySig.S)
	fmt.Println(verifyStatus) // should be true

	return verifyStatus
}
