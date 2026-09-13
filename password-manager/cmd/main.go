package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eubyt/codingchallenges/password-manager/internal/util"
	"github.com/eubyt/codingchallenges/password-manager/internal/vault"
)

const (
	senhaPessoal  = "senha-pessoal"
	senhaTrabalho = "senha-trabalho"
)

// Iniciar o repositório de vaults e retornar a função para adicionar novos vaults
func initRepository() (*vault.Repository, func(name, passwordHash string) *vault.Vault) {
	repository := vault.NewRepository()

	addRepository := func(name, passwordHash string) *vault.Vault {
		v, err := vault.NewVault(name, passwordHash)
		if err != nil {
			fmt.Println("Erro ao criar o vault:", err)
			return nil
		}
		repository.Add(v)
		return v
	}

	return repository, addRepository
}

func readLine(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(line), nil
}

func addPassword(v *vault.Vault, name, email, plaintext string, key []byte) {
	if v == nil {
		return
	}

	if _, err := v.AddPassword(name, email, plaintext, key); err != nil {
		fmt.Println("Erro ao adicionar a senha:", err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	repository, addRepository := initRepository()

	pessoal := addRepository("pessoal", senhaPessoal)
	trabalho := addRepository("trabalho", senhaTrabalho)

	keyPessoal := util.DeriveKey(senhaPessoal)
	addPassword(pessoal, "github", "eu@exemplo.com", "g1thub-s3nh4", keyPessoal)
	addPassword(pessoal, "netflix", "eu@exemplo.com", "n3tfl1x-s3nh4", keyPessoal)
	addPassword(pessoal, "banco", "eu@exemplo.com", "b4nc0-s3nh4", keyPessoal)

	keyTrabalho := util.DeriveKey(senhaTrabalho)
	addPassword(trabalho, "gitlab", "eu@empresa.com", "g1tl4b-s3nh4", keyTrabalho)
	addPassword(trabalho, "jira", "eu@empresa.com", "j1r4-s3nh4", keyTrabalho)

	fmt.Println("Lista de vaults:")

	for i, item := range repository.List() {
		fmt.Printf("%d - %s (%d senhas)\n", i, item.Name, len(item.ListPasswords()))
	}

	option, err := readLine(reader, "Escolha uma opção: ")
	if err != nil {
		fmt.Println("Erro ao ler a opção:", err)
		return
	}

	optionVault, err := strconv.Atoi(option)
	if err != nil || optionVault < 0 || optionVault >= len(repository.List()) {
		fmt.Println("Opção inválida.")
		return
	}

	selected := repository.List()[optionVault]

	senha, err := readLine(reader, fmt.Sprintf("Senha do vault %s: ", selected.Name))
	if err != nil {
		fmt.Println("Erro ao ler a senha:", err)
		return
	}

	if !selected.CheckPassword(senha) {
		fmt.Println("Senha incorreta.")
		return
	}

	fmt.Printf("Senhas do vault %s:\n", selected.Name)

	for i, item := range selected.ListPasswords() {
		fmt.Printf("%d - %s (%s)\n", i, item.Name, item.Email)
	}
}
