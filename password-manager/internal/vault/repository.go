package vault

type Repository struct {
	vaults []*Vault
}

func NewRepository() *Repository {
	return &Repository{
		vaults: [](*Vault){},
	}
}

func (r *Repository) Add(newVault *Vault) {
	r.vaults = append(r.vaults, newVault)
}

func (r *Repository) List() []*Vault {
	return r.vaults
}
