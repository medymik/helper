package helper_test

import (
	"log"
	"testing"

	"github.com/medymik/helper"
)

func TestCanIHashPassword(t *testing.T) {
	plainPassword := "s3cr3t"
	hash, err := helper.HashPassword(plainPassword)

	if err != nil {
		log.Fatalln(err)
	}

	ok := helper.CheckPassword(plainPassword, hash)
	if !ok {
		t.Errorf("expect password check to be %v but go %v", true, ok)
	}
}
