package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string
	Authors   []Author `gorm:"many2many:authors_books;"`
	Resources []Resource
}

type Author struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"`
	Books []Book `gorm:"many2many:authors_books;"`
}

type Resource struct {
	gorm.Model
	Type   ResourceType
	File   string
	Meta   string
	BookID uint
}

type ResourceType int8

const (
	ArchiveType ResourceType = iota
	DocumentType
	AudioType
)
