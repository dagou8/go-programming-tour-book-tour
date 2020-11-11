package cmd

import (
	"log"

	"github.com/dagou8/go-programming-tour-book-tour/internal/sql2struct"
	"github.com/spf13/cobra"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql conversion and processing",
	Long:  "sql conversion and processing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql convert to struct",
	Long:  "sql convert to struct",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}

		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "Please enter the database account name")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "Please enter the database account password")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1", "Please enter the database host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "Please enter the character set of the database")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "Please enter the type of database")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "Please enter the database name")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "Please enter the table name")
}
