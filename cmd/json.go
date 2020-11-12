package cmd

import (
	"log"

	"github.com/dagou8/go-programming-tour-book-tour/internal/json2struct"
	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Conversion and processing json string",
	Long:  "Conversion and processing json string",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var json2StructCmd = &cobra.Command{
	Use:   "struct",
	Short: "Convert json to struct",
	Long:  "Convert json string to struct",
	Run: func(cmd *cobra.Command, args []string) {
		parser, err := json2struct.NewParser(str)
		if err != nil {
			log.Fatalf("Json2Struct.NewParser err: %v", err)
		}
		content := parser.Json2Struct()
		log.Printf("The ouput result is %s", content)
	},
}

func init() {
	jsonCmd.AddCommand(json2StructCmd)
	json2StructCmd.Flags().StringVarP(&str, "str", "s", "", "Please enter the json string")
}
