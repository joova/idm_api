package crypto

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	has, err := HashPassword("password")
	if err != nil && has != "$2a$14$M1jWfTHh3C0cAJBoUP7O1OHbuG4dY48UI7A2CQNiGygRPG6KTPFMq" {
		t.Error("Hash result not valid")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	if !CheckPasswordHash("password", "$2a$14$M1jWfTHh3C0cAJBoUP7O1OHbuG4dY48UI7A2CQNiGygRPG6KTPFMq") {
		t.Error("Password & hash not match.")
	}
}
