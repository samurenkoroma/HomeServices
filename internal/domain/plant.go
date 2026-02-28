package domain

type Family struct {
	ID   int64
	Name string
}

type Species struct {
	ID     int64
	Name   string
	Family Family
}

type Plant struct {
	ID      int64
	Name    string
	Species Species
}
