package files

import (
	"os"
	"samurenkoroma/services/internal/hashpass/outputs"

	"github.com/fatih/color"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		outputs.PrintError(err)
		color.Blue("Будет создана новая база данных")
		return nil, err
	}

	return data, nil

}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		outputs.PrintError(err)
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		outputs.PrintError(err)
		return
	}

	color.Blue("Запись добавлена")
}
