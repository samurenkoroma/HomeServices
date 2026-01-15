package main

import "fmt"

type Bookmarks map[string]string

const (
	SHOW = iota + 1
	ADD
	DELETE
	QUIT
)

func main() {
	fmt.Println("Закладки")
	var bookmarks = Bookmarks{}
App:
	for {
		switch getMenu() {
		case SHOW:
			listBookmarks(bookmarks)
		case ADD:
			addTo(bookmarks)
		case DELETE:
			deleteFrom(bookmarks)
		case QUIT:
			{
				fmt.Println("Прощенья просим")
				break App
			}
		}

	}

}

func listBookmarks(m Bookmarks) {
	fmt.Println(m)
}

func addTo(m Bookmarks) {
	fmt.Println("Добавление закладки")
	var key, value string
	
	fmt.Println("Введите ключ")
	fmt.Scan(&key)

	fmt.Println("Введите значение")
	fmt.Scan(&value)
	
	m[key] = value
	fmt.Println("Закладка успешно добавлена")
}

func deleteFrom(m Bookmarks) {
	fmt.Println("Удаление закладки")
	var key string
	fmt.Println("Введите ключ")
	fmt.Scan(&key)
	delete(m, key)
}

func getMenu() int {
	var variant int
	fmt.Println("1. Лицезреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. УдОлить закладку")
	fmt.Println("4. Исход")
	fmt.Scan(&variant)
	return variant
}
