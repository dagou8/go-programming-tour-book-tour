package cmd

import (
	"log"
	"strings"

	"github.com/dagou8/go-programming-tour-book-tour/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var str string
var mode int8
var desc = strings.Join([]string{
	"This subcommand supports various word format conversions, and the mode as follows: ",
	"1: Convert all words to uppercase",
	"2: Convert all words to lowercase",
	"3: Convert underscore to uppercase camel case",
	"4: Convert underscore to lowercase camel case",
	"4: Convert camel case to underscore",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "Word format conversion",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("This type conversion is not currently supported, please execute `help word` to view the help document.")
		}

		log.Printf("The output result is: %v", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "Please enter a word.")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "Please enter ther conversion mode.")
}
