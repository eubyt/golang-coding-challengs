package vault

import "testing"

func TestVaultAddsAndListsPasswords(t *testing.T) {
	key := []byte("01234567890123456789012345678901")

	v, err := NewVault("pessoal", "senha-do-vault")
	if err != nil {
		t.Fatalf("não esperava erro ao criar vault: %v", err)
	}

	credential, err := v.AddPassword(
		"GitHub",
		"usuario@example.com",
		"senha-secreta",
		key,
	)
	if err != nil {
		t.Fatalf("não esperava erro ao adicionar password: %v", err)
	}

	passwords := v.ListPasswords()
	if len(passwords) != 1 {
		t.Fatalf("esperava 1 password, recebido %d", len(passwords))
	}

	if passwords[0] != credential {
		t.Fatal("a password adicionada não foi encontrada no vault")
	}

	plaintext, err := passwords[0].Decrypt(key)
	if err != nil {
		t.Fatalf("não esperava erro ao descriptografar password: %v", err)
	}

	if plaintext != "senha-secreta" {
		t.Fatalf("senha esperada %q, recebida %q", "senha-secreta", plaintext)
	}
}
