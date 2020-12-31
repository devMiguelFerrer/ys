package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdCreate = &cobra.Command{
	Use:     "create",
	Short:   "Create templates of code " + string('🚀'),
	Long:    `Create templates od code like EndPoint with Porst&Adapters Architecture.`,
	Example: "mr create CustomTask -r mongoose",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		recipes := handleGlobalVars(args)
		globalEntity = loadConfigFile(selectedInterface)

		createDirectory(fullPath + "/" + className)
		generateDomain(recipes)
		generateApplication(recipes)
		generateInfrasctructure(recipes, selectedDB)
		fmt.Printf("Total used space: %3v %d bytes.\n", string('📦'), totalBytes)
		fmt.Println(string('🎉')+string('🎉')+string('🎉'), "The Generation Succeeded", string('✅'))
	},
}

var cmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "Is in development " + string('🚷'),
	Long:  `Is in development ` + string('🚷'),
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Is in development ` + string('🚷'))
	},
}
