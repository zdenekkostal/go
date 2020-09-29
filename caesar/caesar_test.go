package caesar_test

import (
	"github.com/zdenekkostal/go/caesar"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestItEncryptsAndDecryptsCorrectly(t *testing.T) {
	input := "THIS IS A TEST MESSAGE"
	cipher := "CIPHER"

	encrypted := caesar.Encode(input, caesar.GetCipherGetter(cipher))
	decrypted := caesar.Decode(encrypted, caesar.GetCipherGetter(cipher))

	// Using testify
	assert.Equal(t, input, decrypted)
}
