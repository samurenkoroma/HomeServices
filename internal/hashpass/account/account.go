package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"pass"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc Account) Display() {
	color.Green("(%s) - [ %s : %s ] \n", acc.Url, acc.Login, acc.Password)
}

func (acc *Account) generatePass(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] += letters[rand.IntN(len(letters))]
	}
	acc.Password = string(res)
}

func New(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("WRONG_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("WRONG_URL")
	}

	acc := Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
	}

	if acc.Password == "" {
		acc.generatePass(12)
	}

	return &acc, nil
}
