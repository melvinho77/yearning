package lib

import (
	"Yearning-go/src/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	model.DbInit("../../conf.toml")
}

func TestEncryptToken_Encrypt(t *testing.T) {
	password := "123abc"
	enc := Encrypt(password)
	assert.Equal(t, password, Decrypt(enc))
	fmt.Println(enc)
}

func TestDjangoEncrypt(t *testing.T) {
	enc := DjangoEncrypt("Yearning_admin", "321312312321")
	assert.Equal(t, true, DjangoCheckPassword(&model.CoreAccount{Password: enc},"Yearning_admin"))
}

