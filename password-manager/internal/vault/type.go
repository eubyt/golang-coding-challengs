package vault

import (
	"fmt"

	"github.com/eubyt/codingchallenges/password-manager/internal/password"
	"github.com/eubyt/codingchallenges/password-manager/internal/util"
)

type Vault struct {
	Name         string
	PasswordHash string
	Passwords    []*password.Password
	createdAt    int64
	updatedAt    int64
}

func NewVault(name, passwordHash string) (*Vault, error) {
	passwordHash, err := util.HashPassword(passwordHash)

	if err != nil {
		return nil, fmt.Errorf("Erro ao criar o hash da senha: %v", err)
	}

	return &Vault{
		Name:         name,
		PasswordHash: passwordHash,
		Passwords:    []*password.Password{},
		createdAt:    util.DateTimeNow(),
		updatedAt:    0,
	}, nil
}

func (v *Vault) CheckPassword(password string) bool {
	return util.CheckPassword(password, v.PasswordHash)
}

func (v *Vault) AddPassword(name, email, plaintext string, key []byte) (*password.Password, error) {
	newPassword, err := password.NewPassword(name, email, plaintext, key)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar password: %w", err)
	}

	v.Passwords = append(v.Passwords, newPassword)
	return newPassword, nil
}

func (v *Vault) ListPasswords() []*password.Password {
	return v.Passwords
}
