package model

type Website struct {
	Address string
	Status  bool
}

type StatusResult struct {
	Status bool
}

var Websites []Website

func init() {
	Websites = make([]Website, 0, 10)
}
