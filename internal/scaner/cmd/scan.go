package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"samurenkoroma/services/pkg/db"
	"samurenkoroma/services/pkg/entities"
	"samurenkoroma/services/pkg/repositories"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Сканирует указанную папку",
	Long:  `Сканирует указанную папку на наличие книг`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		database := db.NewDb(conf)
		bookRepo := repositories.NewBookRepo(database)
		fmt.Printf("scan %s\n", args[0])
		dirname := args[0]
		var book_files []string

		extSlicses := []string{".fb2", ".epub", ".djvu", ".pdf"}
		err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				fmt.Println(err)
				return nil
			}

			if !info.IsDir() {

				if slices.Contains(extSlicses, filepath.Ext(path)) {
					filename := info.Name()
					p := strings.TrimPrefix(path, dirname)
					bookRepo.Create(
						&entities.Book{
							Title: filename[:len(filename)-len(filepath.Ext(filename))],
							Resources: []entities.Resource{
								{
									Type: entities.DocumentType,
									Meta: "",
									File: p,
								},
							},
						})
					book_files = append(book_files, path)
				}
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}

		// for _, file := range epub_f {
		// fmt.Println(file)
		// }
		fmt.Printf("books: %d\n", len(book_files))

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
