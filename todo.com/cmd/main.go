package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
	"todo.com/cmd/database"
)

const (
	DB_TYPE_MIGRATE = "migrate"
	DB_TYPE_DROP    = "drop"
)

func main() {
	var (
		dbOperate = flag.String("d", "", "migrate flag")
	)
	flag.Parse()

	// DB operation
	tableOperateFlags := []string{DB_TYPE_MIGRATE, DB_TYPE_DROP}
	if !slices.Contains(tableOperateFlags, *dbOperate) {
		fmt.Printf("Nothing to run. Please set some option. Allowed options are '%s'\n", "-d")
	}

	if err := database.Run(*dbOperate); err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m While DB operation: %v", "Error occured.", fmt.Errorf("Error has occurred. %v", err))
		os.Exit(1)
	}

	fmt.Printf("\x1b[32m%s\x1b[0m finished to run your operation.\n", "Successfully")
}
