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
	dataToSign := "This is Binh"
	mySig := Sign(privKey, dataToSign)

	assert.NotNil(mySig.SignHash)
	assert.NotNil(mySig.R)
	assert.NotNil(mySig.S)

	// verify
	ok := Verify(&pubKey, mySig)
	assert.True(ok)
}

func TestVerify2(t *testing.T) {
	assert := assert.New(t)

	privKey := NewKeyPair()

	_ = privKey.PublicKey

	// sign data
	dataToSign := "This is Binh"
	mySig := Sign(privKey, dataToSign)

	assert.NotNil(mySig.SignHash)
	assert.NotNil(mySig.R)
	assert.NotNil(mySig.S)

	// another pubkey
	privKey2 := NewKeyPair()
	pubKey2 := privKey2.PublicKey

	// verify
	ok := Verify(&pubKey2, mySig)
	assert.False(ok)

}
