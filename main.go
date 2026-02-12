package main

import (
	"github.com/Get-Blu/blu-code/cmd"
	"github.com/Get-Blu/blu-code/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
