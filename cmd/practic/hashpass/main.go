package main

import (
	"fmt"
	"strings"

	"samurenkoroma/services/internal/hashpass/account"
	"samurenkoroma/services/internal/hashpass/encrypter"
	"samurenkoroma/services/internal/hashpass/files"
	"samurenkoroma/services/internal/hashpass/outputs"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	if err := godotenv.Load(); err != nil {
		outputs.PrintError(err)
	}
	enc := *encrypter.NewEncrypter()
	vault := account.NewVault(files.NewJsonDb("data.vault"), enc)
Menu:
	for {
		variant := promptData(
			"-----------------------------------",
			"1. Создание",
			"2. Поиск по url",
			"3. Поиск по login",
			"4. Удаление",
			"5. Выход",
			"-----------------------------------",
			"Сделайте выбор",
		)

		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	url := promptData("Введите url")
	pass := promptData("Введите пароль")

	acc, err := account.New(login, pass, url)
	if err != nil {
		outputs.PrintError(err)
		return
	}

	vault.Add(*acc)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	query := promptData("Введите url")

	result := vault.FilterAccounts(query, func(acc account.Account, q string) bool {
		return strings.Contains(acc.Url, q)
	})

	outputsResult(&result)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	query := promptData("Введите login")

	result := vault.FilterAccounts(query, func(acc account.Account, q string) bool {
		return strings.Contains(acc.Login, q)
	})

	outputsResult(&result)
}

func deleteAccount(vault *account.VaultWithDb) {
	query := promptData("Введите url")
	isDeleted := vault.DeleteAccount(query)
	if isDeleted {
		color.Green("Запись удалена")
	} else {
		color.Red("Запись %s не найдена", query)
	}
}

func outputsResult(result *[]account.Account) {
	if len(*result) == 0 {
		color.Red("Записей не найдено")
	} else {
		for _, acc := range *result {
			acc.Display()
		}
	}
}

func promptData(prompt ...any) string {

	lastElem := len(prompt) - 1
	for _, p := range prompt[:lastElem] {
		fmt.Println(p)
	}
	color.Yellow("%v: ", prompt[lastElem])
	var res string
	fmt.Scanln(&res)
	return res
}
