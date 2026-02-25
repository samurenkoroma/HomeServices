package main

type Node struct {
	Rank  string `json:"rank"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Items []Node `json:"items"`
}
type Supplier struct {
	Name string `json:"name"`
	Site string `json:"website"`
}

type Data struct {
	Nodes     []Node     `json:"nodes"`
	Suppliers []Supplier `json:"suppliers"`
}
