package account

import (
	"encoding/json"
	"samurenkoroma/services/internal/hashpass/encrypter"
	"samurenkoroma/services/internal/hashpass/outputs"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write(content []byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{
			db:  db,
			enc: enc,
			Vault: Vault{
				Accounts:  []Account{},
				CreatedAt: time.Now(),
			},
		}
	}

	var vault Vault
	data := enc.Decrypt(file)
	err = json.Unmarshal(data, &vault)
	if err != nil {
		outputs.PrintError(err)
	}
	color.Cyan("Найдено записей: %d", len(vault.Accounts))
	return &VaultWithDb{
		enc:   enc,
		db:    db,
		Vault: vault,
	}

}

func (v *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(v)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (v *VaultWithDb) Add(acc Account) {
	v.Accounts = append(v.Accounts, acc)
	v.update()
}

func (v *VaultWithDb) FilterAccounts(filter string, checker func(Account, string) bool) []Account {
	res := []Account{}

	for _, acc := range v.Accounts {
		if checker(acc, filter) {
			res = append(res, acc)
		}
	}
	return res
}

func (v *VaultWithDb) DeleteAccount(url string) bool {
	var accounts []Account
	isDeleted := false

	for _, acc := range v.Accounts {
		if acc.Url != url {
			accounts = append(accounts, acc)
			continue
		}
		isDeleted = true
	}

	v.Accounts = accounts

	v.update()

	return isDeleted
}

func (v *VaultWithDb) update() {
	v.UpdatedAt = time.Now()
	data, err := v.Vault.ToBytes()
	encryptedData := v.enc.Encrypt(data)
	if err != nil {
		outputs.PrintError(err)
	}
	v.db.Write(encryptedData)
}
