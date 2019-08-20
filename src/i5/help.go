package i5

import (
	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/io/console"
)

func PrintHelp() {
	console.Println(`
Usage:

	i5 [options] [file]

options:

	--help      print help
	`)
	errors.Exit(0)
}