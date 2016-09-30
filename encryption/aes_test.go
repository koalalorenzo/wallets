package encryption

import "testing"

func TestAES(t *testing.T) {
	text := "Oh my Glob!"
	key := []byte("WoW! It is Mathematical!")

	encrypted, err := AESEncrypt(key, text)
	if err != nil {
		t.Error(err)
	}

	decrypted, err := AESDecrypt(key, encrypted)
	if err != nil {
		t.Error(err)
	}

	if text != decrypted {
		t.Error("AES Encryption failed")
	}
}
