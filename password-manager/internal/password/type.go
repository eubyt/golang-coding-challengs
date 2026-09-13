package password

import "github.com/eubyt/codingchallenges/password-manager/internal/util"

type Password struct {
	Name       string
	Email      string
	Ciphertext []byte
}

func NewPassword(name, email, plaintext string, key []byte) (*Password, error) {
	ciphertext, err := util.Encrypt([]byte(plaintext), key)
	if err != nil {
		return nil, err
	}

	return &Password{
		Name:       name,
		Email:      email,
		Ciphertext: ciphertext,
	}, nil
}

func (p *Password) Decrypt(key []byte) (string, error) {
	plaintext, err := util.Decrypt(p.Ciphertext, key)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
