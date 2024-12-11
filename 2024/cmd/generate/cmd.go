package generate

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func handleErr(err error, folder string) {
	os.RemoveAll(folder)
	log.Fatal(err)
}

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new day directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		day, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		folder := fmt.Sprintf("./day%02d", day)
		if err := os.Mkdir(folder, 0755); err != nil {
			handleErr(err, folder)
		}

		if err := generateMainCode(day); err != nil {
			handleErr(err, folder)
		}

		if err := generateTestCode(day); err != nil {
			handleErr(err, folder)
		}

		if err := retrieveInput(day); err != nil {
			handleErr(err, folder)
		}

		filename := fmt.Sprintf("./day%02d/test.txt", day)
		if err := os.WriteFile(filename, []byte{}, 0755); err != nil {
			handleErr(err, folder)
		}
	},
}
