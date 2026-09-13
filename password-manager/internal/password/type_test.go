package password

import "testing"

func TestNewPasswordEncryptsAndDecryptsPassword(t *testing.T) {
	key := []byte("01234567890123456789012345678901")

	credential, err := NewPassword(
		"GitHub",
		"usuario@example.com",
		"senha-secreta",
		key,
	)
	if err != nil {
		t.Fatalf("não esperava erro ao criar password: %v", err)
	}

	if credential.Name != "GitHub" {
		t.Fatalf("nome esperado %q, recebido %q", "GitHub", credential.Name)
	}

	if credential.Email != "usuario@example.com" {
		t.Fatalf("email esperado %q, recebido %q", "usuario@example.com", credential.Email)
	}

	if len(credential.Ciphertext) == 0 {
		t.Fatal("esperava conteúdo criptografado")
	}

	plaintext, err := credential.Decrypt(key)
	if err != nil {
		t.Fatalf("não esperava erro ao descriptografar password: %v", err)
	}

	if plaintext != "senha-secreta" {
		t.Fatalf("senha esperada %q, recebida %q", "senha-secreta", plaintext)
	}
}

func TestPasswordDecryptWithWrongKey(t *testing.T) {
	key := []byte("01234567890123456789012345678901")
	wrongKey := []byte("11111111111111111111111111111111")

	credential, err := NewPassword(
		"GitHub",
		"usuario@example.com",
		"senha-secreta",
		key,
	)
	if err != nil {
		t.Fatalf("não esperava erro ao criar password: %v", err)
	}

	_, err = credential.Decrypt(wrongKey)
	if err == nil {
		t.Fatal("esperava erro ao descriptografar com chave incorreta")
	}
}
