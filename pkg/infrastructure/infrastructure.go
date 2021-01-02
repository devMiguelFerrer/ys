package infrastructure

import (
	"fmt"

	"github.com/devMiguelFerrer/ys/pkg/shared"
)

var (
	mongooseTypes = map[string]string{
		"string":  "String",
		"number":  "Number",
		"date":    "Date",
		"boolean": "Boolean",
		"array":   "Array",
	}
)

// GenerateController method create ADD Use Case
func GenerateController(recipe []shared.Recipe) string {
	return shared.ReplaceByRecipe(baseController(), recipe)
}

// GenerateIndex typeORM repository
func GenerateIndex(recipe []shared.Recipe) string {
	return shared.ReplaceByRecipe(baseIndex, recipe)
}

// GenerateTypeORMModel typeORM repository
func GenerateTypeORMModel(fileLoaded map[string]interface{}, recipe []shared.Recipe) string {
	base := handleTypeORMModel()

	base += generateTypeORMColumns(fileLoaded)
	return shared.ReplaceByRecipe(base, recipe)
}

// GenerateTypeORMCriteria typeORM repository
func GenerateTypeORMCriteria(recipe []shared.Recipe) string {
	return shared.ReplaceByRecipe(handleTypeORMCriteria(), recipe)
}

// GenerateTypeORMRepository typeORM repository
func GenerateTypeORMRepository(recipe []shared.Recipe) string {
	return shared.ReplaceByRecipe(handleTypeORMRepository(), recipe)
}

// GenerateMongooseModel Mongoose repository
func GenerateMongooseModel(fileLoaded map[string]interface{}, recipe []shared.Recipe) string {
	base := handleMongooseModel(generateMongooseColumns(fileLoaded))

	// base += generateMongooseColumns(fileLoaded)
	return shared.ReplaceByRecipe(base, recipe)
}

// GenerateMongooseCriteria Mongoose repository
func GenerateMongooseCriteria(recipe []shared.Recipe) string {
	return shared.ReplaceByRecipe(handleMongooseCriteria(), recipe)
}

// GenerateMongooseRepository Mongoose repository
func GenerateMongooseRepository(recipe []shared.Recipe) string {
	return shared.ReplaceByRecipe(handleMongooseRepository(), recipe)
}

func generateTypeORMColumns(rawColumns map[string]interface{}) string {
	formatedColumns := ""
	formatedConstructor := "  constructor(__var__: I__name__) {\n"

	for k, v := range rawColumns {
		formatedColumns += "  @Column()\n"
		formatedColumns += fmt.Sprintln("  "+k+":", v)
		formatedColumns += "\n"

		formatedConstructor += fmt.Sprintln("    this." + k + " = " + "__var__." + k)
	}
	return formatedColumns + formatedConstructor + "  }\n}\n"
}

func generateMongooseColumns(rawColumns map[string]interface{}) string {
	formatedColumns := ""

	for k, v := range rawColumns {
		formatedColumns += "	  " + k + ": {\n"
		formatedColumns += fmt.Sprintln("			type: " + mongooseTypes[fmt.Sprintf("%v", v)] + ",")
		formatedColumns += "			required: true\n"
		formatedColumns += "		},\n"
	}
	return formatedColumns
}
