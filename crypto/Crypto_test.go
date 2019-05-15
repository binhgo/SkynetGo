package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKeyPair(t *testing.T) {
	assert := assert.New(t)

	privKey := NewKeyPair()

	assert.NotNil(privKey)
	assert.NotNil(privKey.PublicKey)
}

func TestVerify(t *testing.T) {
	assert := assert.New(t)

	privKey := NewKeyPair()
	pubKey := privKey.PublicKey

	// sign data
	data := "This is Binh"
	mySig := Sign(privKey, data)

	assert.NotNil(mySig.R)
	assert.NotNil(mySig.S)

	// verify
	ok := Verify(&pubKey, mySig, data)
	assert.True(ok)
}

func TestVerify2(t *testing.T) {
	assert := assert.New(t)

	privKey := NewKeyPair()

	_ = privKey.PublicKey

	// sign data
	dataToSign := "This is Binh"
	mySig := Sign(privKey, dataToSign)

	assert.NotNil(mySig.R)
	assert.NotNil(mySig.S)

	// another pubkey
	privKey2 := NewKeyPair()
	pubKey2 := privKey2.PublicKey

	// verify
	ok := Verify(&pubKey2, mySig, dataToSign)
	assert.False(ok)

}

func TestSignFile(t *testing.T) {
	assert := assert.New(t)

	filePath := "./test.txt"

	privKey := NewKeyPair()

	_ = privKey.PublicKey

	// sign data
	mySig, err := SignFile(privKey, filePath)
	assert.Nil(err)

	assert.NotNil(mySig.R)
	assert.NotNil(mySig.S)
}

func TestVerify3(t *testing.T) {
	// verify File
	assert := assert.New(t)

	filePath := "./test.txt"

	privKey := NewKeyPair()

	pubKey := privKey.PublicKey

	// sign data
	mySig, err := SignFile(privKey, filePath)
	assert.Nil(err)

	assert.NotNil(mySig.R)
	assert.NotNil(mySig.S)

	ok, err := VerifyFile(&pubKey, mySig, filePath)
	assert.Nil(err)
	assert.True(*ok)

	// create another key pair
	privKey2 := NewKeyPair()
	pubKey2 := privKey2.PublicKey

	ok, err = VerifyFile(&pubKey2, mySig, filePath)
	assert.Nil(err)
	assert.False(*ok)

}
