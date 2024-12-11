package submit

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func handleErr(err error, folder string) {
	os.RemoveAll(folder)
	log.Fatal(err)
}

var SubmitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit solution to Advent of Code 2024",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		day, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		part, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}

		answer, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}

		if err := submitInput(day, part, answer); err != nil {
			log.Fatal(err)
		}
	},
}
