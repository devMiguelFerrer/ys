package cmd

import (
	"strings"

	"github.com/devMiguelFerrer/ys/pkg/infrastructure"
	"github.com/devMiguelFerrer/ys/pkg/shared"
)

func generateInfrasctructure(recipes []shared.Recipe, selectedRepository string) {

	infrastructurePath := fullPath + "/Infrastructure"

	createDirectory(infrastructurePath + "/persistence/" + selectedDB)
	createDirectory(infrastructurePath + "/controllers")

	str := infrastructure.GenerateIndex(recipes)
	saveFile(infrastructurePath+"/persistence/"+selectedDB+"/index.ts", strings.NewReader(str))

	str = infrastructure.GenerateController(recipes)
	saveFile(infrastructurePath+"/controllers/"+varName+"Controller.ts", strings.NewReader(str))

	switch selectedRepository {
	case typeORM:
		generateInfrasctructureTypeORM(infrastructurePath, recipes)
	case mongoose:
		// TODO:
		generateInfrasctructureMongoose(infrastructurePath, recipes)
	default:
		// TODO:
		generateInfrasctructureTypeORM(infrastructurePath, recipes)
	}
}

func generateInfrasctructureTypeORM(infrastructurePath string, recipes []shared.Recipe) {
	str := infrastructure.GenerateTypeORMModel(globalEntity, recipes)
	saveFile(infrastructurePath+"/persistence/"+selectedDB+"/"+className+"Model.ts", strings.NewReader(str))

	str = infrastructure.GenerateTypeORMCriteria(recipes)
	saveFile(infrastructurePath+"/persistence/"+selectedDB+"/"+className+"Criteria.ts", strings.NewReader(str))

	str = infrastructure.GenerateTypeORMRepository(recipes)
	saveFile(infrastructurePath+"/persistence/"+selectedDB+"/"+className+"Repository.ts", strings.NewReader(str))
}

func generateInfrasctructureMongoose(infrastructurePath string, recipes []shared.Recipe) {
	str := infrastructure.GenerateMongooseModel(globalEntity, recipes)
	saveFile(infrastructurePath+"/persistence/"+selectedDB+"/"+className+"Model.ts", strings.NewReader(str))

	str = infrastructure.GenerateMongooseCriteria(recipes)
	saveFile(infrastructurePath+"/persistence/"+selectedDB+"/"+className+"Criteria.ts", strings.NewReader(str))

	str = infrastructure.GenerateMongooseRepository(recipes)
	saveFile(infrastructurePath+"/persistence/"+selectedDB+"/"+className+"Repository.ts", strings.NewReader(str))
}
