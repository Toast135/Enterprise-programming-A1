package main

import (
	"web/setup"
	"web/templates/notepage"
)

func main() {
	setup.CreateDB()
	setup.CreateTables()
    notepage.SelectNote()
}