package test

import (
	"github.com/agilistikmal/wallet-go/helper"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	match := helper.CheckPassword("Michaleisadi", "$2a$10$rH/2PIWYazSbaAf6DuJQlerGPH5c7Njf7ZWMBh3nZy6/rzTI89oe6")
	if match {
		t.Log("Password confirmed!")
	} else {
		t.Error("Wrong password")
	}
}
