package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/devMiguelFerrer/ys/pkg/shared"
	"github.com/spf13/viper"
)

func createDirectory(path string) {
	fmt.Println(path)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		panic("Error: createDirs -> " + path)
	}
}

func saveFile(fullPath string, r io.Reader) {
	newFile, err := os.Create(fullPath)
	if err != nil {
		panic("Error: os.Create -> " + err.Error())
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, r)
	if err != nil {
		panic("Error: io.Copy -> " + err.Error())
	}
	fmt.Printf("%s %3v %d bytes.\n", fullPath, string('📦'), bytesCopied)
	totalBytes += bytesCopied
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func handleGlobalVars(args []string) []shared.Recipe {
	if selectedDB != typeORM && selectedDB != mongoose {
		panic("Error selectedDB")
	}
	splitted := strings.Split(args[0], "/")
	fullPath = args[0]
	className = splitted[len(splitted)-1]
	varName = shared.ConvertToVar(className)
	routeName = convertToRoute(className)
	splitted = removeIndex(splitted, len(splitted)-1)
	partialPath = strings.Join(splitted, "/")

	return []shared.Recipe{
		{Key: name, Value: className},
		{Key: variable, Value: varName},
		{Key: route, Value: routeName},
		{Key: repo, Value: selectedDB},
	}
}

func loadConfigFile(entityPath string) map[string]interface{} {
	viper.SetConfigFile(entityPath)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return viper.AllSettings()
}

func convertToRoute(s string) string {
	splitted := strings.Split(s, "")

	for i, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) && i > 0 {
			splitted[i] = "-" + splitted[i]
		}
	}
	joinned := strings.Join(splitted, "")

	return strings.ToLower(joinned)
}
