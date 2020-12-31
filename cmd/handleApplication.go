package cmd

import (
	"strings"

	"github.com/devMiguelFerrer/ys/pkg/application"
	"github.com/devMiguelFerrer/ys/pkg/shared"
)

func generateApplication(recipes []shared.Recipe) {

	createDirectory(fullPath + "/" + className + "/Application")

	str := application.GenerateGet(recipes)
	saveFile(fullPath+"/"+className+"/Application/get"+className+".ts", strings.NewReader(str))

	str = application.GenerateAdd(recipes)
	saveFile(fullPath+"/"+className+"/Application/add"+className+".ts", strings.NewReader(str))

	str = application.GenerateUpdate(recipes)
	saveFile(fullPath+"/"+className+"/Application/update"+className+".ts", strings.NewReader(str))

	str = application.GenerateDelete(recipes)
	saveFile(fullPath+"/"+className+"/Application/remove"+className+".ts", strings.NewReader(str))

	str = application.GenerateApplicationIndex(recipes)
	saveFile(fullPath+"/"+className+"/Application/index.ts", strings.NewReader(str))
}
