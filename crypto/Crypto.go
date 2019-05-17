package crypto

import (
	"crypto/dsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"log"
	"math/big"
	"os"

	"SkynetGo/util"
)

type MySignature struct {
	R *big.Int
	S *big.Int
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

func Sign1() *MySignature {

	var mySig MySignature

	return &mySig
}

func Sign(privateKey *dsa.PrivateKey, dataToSign string) *MySignature {

	// Sign
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, dataToSign)
	signHash := h.Sum(nil)

	log.Printf("SIGN: %x\n", signHash)

	r, s, err := dsa.Sign(rand.Reader, privateKey, signHash)
	if err != nil {
		fmt.Println(err)
	}

	// sig := r.Bytes()
	// sig = append(sig, s.Bytes()...)
	// fmt.Printf("Signature : %x\n", sig)

	return &MySignature{R: r, S: s}
}

func SignFile(privateKey *dsa.PrivateKey, filePath string) (*MySignature, error) {

	// read file content
	fileContent, err := util.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Sign
	mySig := Sign(privateKey, fileContent)

	return mySig, nil
}

func Verify(pubKey *dsa.PublicKey, mySig *MySignature, dataToVerify string) bool {

	// re-compute the signHash from dataToVerify
	var h hash.Hash
	h = md5.New()
	io.WriteString(h, dataToVerify)
	signHash := h.Sum(nil)

	log.Printf("VERIFY: %x\n", signHash)

	// verify data with signature R, S and Pubkey
	ok := dsa.Verify(pubKey, signHash, mySig.R, mySig.S)

	return ok
}

func VerifyFile(pubKey *dsa.PublicKey, mySig *MySignature, filePath string) (*bool, error) {

	// read file content
	fileContent, err := util.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// verify
	ok := Verify(pubKey, mySig, fileContent)
	return &ok, nil
}
