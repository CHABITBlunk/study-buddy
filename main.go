package main

import (
	"github.com/chabitblunk/study-buddy/cmd"
	"github.com/chabitblunk/study-buddy/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
