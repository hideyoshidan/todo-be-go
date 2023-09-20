package main

import (
	"flag"
	"fmt"

	"golang.org/x/exp/slices"
	"todo.com/cmd/database"
)

// main accept command for any this service operations
func main() {
	var (
		opes = map[string]bool{
			"MIGDRO": false,
			"SEED":   false,
		}
		dbOperate = flag.String("d", "", "migrate flag")
		dbSeed    = flag.String("s", "", "seeder flag")
	)
	flag.Parse()

	var err error
	// DB operation
	tableOperateFlags := []string{database.DB_TYPE_MIGRATE, database.DB_TYPE_DROP}
	if slices.Contains(tableOperateFlags, *dbOperate) {
		opes["MIGDROP"] = true
		err = dbOperation(*dbOperate)
	}

	seederOperateFlags := []string{database.DB_SEED}
	if slices.Contains(seederOperateFlags, *dbSeed) {
		opes["SEED"] = true
		err = dbOperation(*dbSeed)
	}

	if err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m", "Failed.")
		return
	}

	if !opes["MIGDROP"] && !opes["SEED"] {
		fmt.Printf("Nothing to run. Please set some option. Allowed options are '%s'\n", "-d")
	}

	fmt.Printf("\x1b[32m%s\x1b[0m finished to run your operation.\n", "Successfully")
}

func dbOperation(flag string) error {
	if err := database.Run(flag); err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m While DB operation %s : %v", "Error occured.", flag, fmt.Errorf("Error has occurred. %v", err))
		return err
	}

	return nil
}
