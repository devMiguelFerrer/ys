package cmd

import "github.com/spf13/cobra"

const (
	name     = "__name__"
	variable = "__var__"
	route    = "__route__"
	repo     = "__repo__"
	typeORM  = "typeorm"
	mongoose = "mongoose"
)

var (
	selectedDB        string
	selectedInterface string
	className         string
	varName           string
	routeName         string
	partialPath       string
	fullPath          string
	totalBytes        int64
	globalEntity      map[string]interface{}
)

// Execute main function
func Execute() {

	cmdCreate.Flags().StringVarP(&selectedDB, "repository", "r", "typeorm", "mongoose OR typeorm")

	cmdCreate.Flags().StringVarP(&selectedInterface, "entity", "e", "", "entity path like ./src/entity.json")
	cmdCreate.MarkFlagRequired("entity")

	var RootCmd = &cobra.Command{
		Use:   "ys",
		Short: "<ys> is a CLI to generate templates of code ",
		Long:  `<Yes, Sir!> ` + string('🎩') + ` (ys) will help you get more done in less time. ` + string('🤖'),
		Example: `  ys create [ResourceName] -e [EntityFile] -r [ModelType]
  ys create CustomTask -e config.json -r typeorm`,
	}
	RootCmd.AddCommand(cmdCreate, cmdGenerate)
	// cmdCreate.AddCommand(cmdBack)
	RootCmd.Execute()
}
