package cmd

import (
	"strings"

	"github.com/devMiguelFerrer/ys/pkg/domain"
	"github.com/devMiguelFerrer/ys/pkg/shared"
)

func generateDomain(recipes []shared.Recipe) {

	createDirectory(fullPath + "/" + className + "/Domain")

	saveFile(fullPath+"/"+className+"/Domain/ICriteria.ts", strings.NewReader(domain.Ifilter))

	str := domain.GenerateDomainRepository(recipes)
	saveFile(fullPath+"/"+className+"/Domain/I"+className+"Repository.ts", strings.NewReader(str))

	strGroup := domain.LoadConfigFile(selectedInterface)

	str = domain.GenerateClass(strGroup, className)
	saveFile(fullPath+"/"+className+"/Domain/"+className+".ts", strings.NewReader(str))

	str = domain.GenerateInterface(strGroup, className)
	saveFile(fullPath+"/"+className+"/Domain/I"+className+".ts", strings.NewReader(str))

	str = domain.GenerateDomainIndex(recipes)
	saveFile(fullPath+"/"+className+"/Domain/index.ts", strings.NewReader(str))
}
