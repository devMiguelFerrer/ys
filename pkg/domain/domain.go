package domain

import (
	"fmt"
	"strings"

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
const (
	valString  = "string"
	valNumber  = "number"
	valBoolean = "boolean"
)

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

// GenerateClassVO create the Class to Domain
func GenerateClassVO(fileLoaded map[string]interface{}, name string) string {
	varName := shared.ConvertToVar(name)
	classFormated := "import { I" + name + " } from '.';\n\n"
	classFormated += "export class " + name + " implements I" + name + " {\n"
	classProperties := ""
	classAssignations := "  constructor(" + varName + ": I" + name + ") {\n"
	for k, v := range fileLoaded {
		classProperties += fmt.Sprintln("  "+k+":", v)
		classAssignations += fmt.Sprintln("    this."+k, "= new", strings.Title(k)+"("+varName+"."+k+").value")
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
	const base = `import { ICriteria, I__name__, __name__ } from '.';

export interface I__name__Repository {
	add(__var__: __name__): Promise<void>;
	update(__var__: __name__): Promise<void>;
	remove(id: string): Promise<void>;
	query(criteria: ICriteria): Promise<{ data: I__name__[]; count: number }>;
}
`
	return shared.ReplaceByRecipe(base, recipe)
}

// GenerateDomainIndex method create INDEX to export Domain files
func GenerateDomainIndex(recipe []shared.Recipe) string {
	const base = `export * from './ICriteria';
export * from './I__name__Repository';
export * from './I__name__';
export * from './__name__';
`
	return shared.ReplaceByRecipe(base, recipe)
}

// GenerateDomainBaseValidations method create Bade Validations
func GenerateDomainBaseValidations() string {
	return handleValidationBoolean() + "\n" + handleValidationNumber() + "\n" + handleValidationString()
}

// GenerateDomainExtendedValidations method create Extended Validations by properties
func GenerateDomainExtendedValidations(fileLoaded map[string]interface{}, recipe []shared.Recipe) string {
	base := "import { ValidationBoolean, ValidationNumber, ValidationString } from './BaseTypes';\n"

	for k, v := range fileLoaded {
		switch v {
		case valString:
			base += generateValidationString(k, v)
		case valNumber:
			base += generateValidationNumber(k, v)
		case valBoolean:
			base += generateValidationBoolean(k, v)
		default:
			base += generateValidationString(k, v)
		}
	}

	return shared.ReplaceByRecipe(base, recipe)
}

func generateValidationString(k string, v interface{}) string {
	base := "\nexport class " + strings.Title(k) + ` extends ValidationString {
	constructor(value: ` + fmt.Sprintf("%v", v) + `) {
		super(value);
		if (this.min(1) || this.max(50)) {
			throw new Error('The value should be a number');
		}`
	base += "\n  }\n}\n"
	return base
}

func generateValidationNumber(k string, v interface{}) string {
	base := "\nexport class " + strings.Title(k) + ` extends ValidationNumber {
	constructor(value: ` + fmt.Sprintf("%v", v) + `) {
		super(value);
		if (this.min(0) || this.max(999999)) {
			throw new Error('The value should be a number');
		}`
	base += "\n  }\n}\n"
	return base
}

func generateValidationBoolean(k string, v interface{}) string {
	base := "\nexport class " + strings.Title(k) + ` extends ValidationBoolean {
	constructor(value: ` + fmt.Sprintf("%v", v) + `) {
		super(value);`
	base += "\n  }\n}\n"
	return base
}
