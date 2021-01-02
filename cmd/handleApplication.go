package cmd

import (
	"strings"

	"github.com/devMiguelFerrer/ys/pkg/application"
	"github.com/devMiguelFerrer/ys/pkg/shared"
)

func generateApplication(recipes []shared.Recipe) {

	applicationPath := fullPath + "/Application"

	createDirectory(applicationPath)

	str := application.GenerateGet(recipes)
	saveFile(applicationPath+"/get"+className+".ts", strings.NewReader(str))

	str = application.GenerateAdd(recipes)
	saveFile(applicationPath+"/add"+className+".ts", strings.NewReader(str))

	str = application.GenerateUpdate(recipes)
	saveFile(applicationPath+"/update"+className+".ts", strings.NewReader(str))

	str = application.GenerateDelete(recipes)
	saveFile(applicationPath+"/remove"+className+".ts", strings.NewReader(str))

	str = application.GenerateApplicationIndex(recipes)
	saveFile(applicationPath+"/index.ts", strings.NewReader(str))
}
