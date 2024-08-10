package controller

import "fmt"

type DBError struct {
	Action string
}
type ScanError struct {
	Model string
}
type TmplError struct {
	Path string
}

func (db DBError) Error(err error) {
	fmt.Printf("error with '%s' database action\n", db.Action)
	panic(err)
}

func (tmpl TmplError) Error(err error) {
	fmt.Printf("error parsing file '%s'\n", tmpl.Path)
	panic(err)
}

func (scan ScanError) Error(err error) {
	fmt.Println("error scanning data from result to", scan.Model)
	panic(err)
}
