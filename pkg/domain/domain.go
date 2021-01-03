package domain

import (
	"fmt"

	"github.com/devMiguelFerrer/ys/pkg/shared"
	"github.com/spf13/viper"
)

// Ifilter string custom var
const Ifilter = `export interface IFilter {
  readonly filterField: string;
  readonly filterOperator: string;
  readonly filterValue: string;
}

export interface ICriteria {
  filters: IFilter[];
  order: string;
  offset: number;
  limit: number;
}
`

// GenerateClass create the Class to Domain
func GenerateClass(fileLoaded map[string]interface{}, name string) string {
	varName := shared.ConvertToVar(name)
	classFormated := "import { I" + name + " } from '.';\n\n"
	classFormated += "export class " + name + " implements I" + name + " {\n"
	classProperties := ""
	classAssignations := "  constructor(" + varName + ": I" + name + ") {\n"
	for k, v := range fileLoaded {
		classProperties += fmt.Sprintln("  "+k+":", v)
		classAssignations += fmt.Sprintln("    this."+k, "=", varName+"."+k)
	}
	return classFormated + classProperties + classAssignations + "  }\n}\n"
}

// GenerateInterface create the Interface to Domain
func GenerateInterface(fileLoaded map[string]interface{}, name string) string {
	interfaceFormated := "export interface I" + name + " {\n"
	for k, v := range fileLoaded {
		interfaceFormated += fmt.Sprintln("  " + k + ": " + fmt.Sprintf("%v", v) + ";")
	}
	return interfaceFormated + "}\n"
}

// LoadConfigFile load properties from file
func LoadConfigFile(entityPath string) map[string]interface{} {
	viper.SetConfigFile(entityPath)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return viper.AllSettings()
}

// GenerateDomainRepository method
func GenerateDomainRepository(recipe []shared.Recipe) string {
	const base = `import { ICriteria, I__name__, __name__ } from ".";

  export interface I__name__Repository {
    add(__var__: __name__): Promise<void>;
    update(__var__: __name__): Promise<void>;
    remove(id: string): Promise<void>;
    query(criteria: ICriteria): Promise<{ data: I__name__[]; count: number }>;
  }
`
	return shared.ReplaceByRecipe(base, recipe)
}

// GenerateDomainIndex method create INDEX to export Use Cases
func GenerateDomainIndex(recipe []shared.Recipe) string {
	const base = `export * from './ICriteria';
export * from './I__name__Repository';
export * from './I__name__';
export * from './__name__';
`
	return shared.ReplaceByRecipe(base, recipe)
}
