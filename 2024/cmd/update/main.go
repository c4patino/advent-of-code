package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	const mainTemplate = `package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"cpatino.com/advent-of-code/2024/cmd/generate"

	%s
)

var days = map[int]func(string) (int, int){
	%s
}

var rootCmd = &cobra.Command{
	Use:   "advent-of-code",
	Short: "Advent of Code 2024",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		day, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		file, _ := cmd.Flags().GetString("file")
		if file == "" {
			file = fmt.Sprintf("./day%%02d/input.txt", day)
		}

		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Fatalf("file does not exist: %%s", file)
		}

		log.Printf("Day: %%d, Input File: %%s\n", day, file)

		part1, part2 := days[day](file)

		log.Printf("Part 1: %%d\n", part1)
		log.Printf("Part 2: %%d\n", part2)
	},
}

func init() {
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.Flags().String("file", "", "Path to a file")
}

func Execute() error {
	if len(os.Args) == 1 {
		return rootCmd.Execute()
	}

	subCmd := os.Args[1]
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == subCmd {
			return cmd.Execute()
		}
	}

	return rootCmd.Execute()
}
`

	var imports, mappings strings.Builder

	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() || !strings.HasPrefix(info.Name(), "day") {
			return nil
		}

		dayStr := strings.TrimPrefix(info.Name(), "day")

		day, err := strconv.Atoi(dayStr)
		if err != nil {
			log.Fatal(err)
		}

		imports.WriteString(fmt.Sprintf("\t\"cpatino.com/advent-of-code/2024/day%02d\"\n", day))
		mappings.WriteString(fmt.Sprintf("\t%d: day%02d.Run,\n", day, day))

		return nil
	})

	importsStr := strings.TrimSpace(imports.String())
	mappingsStr := strings.TrimSpace(mappings.String())

	if err != nil {
		fmt.Println("Error walking the path:", err)
		os.Exit(1)
	}

	mainCode := fmt.Sprintf(mainTemplate, importsStr, mappingsStr)

	os.WriteFile("./cmd/root.go", []byte(mainCode), 0644)

	cmd := exec.Command("go", "fmt", "./cmd/root.go")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running go fmt:", err)
		os.Exit(1)
	}
}
