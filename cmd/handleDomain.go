package cmd

import (
	"strings"

	"github.com/devMiguelFerrer/ys/pkg/domain"
	"github.com/devMiguelFerrer/ys/pkg/shared"
)

func generateDomain(recipes []shared.Recipe) {

	domainPath := fullPath + "/Domain"

	createDirectory(domainPath)

	saveFile(domainPath+"/ICriteria.ts", strings.NewReader(domain.Ifilter))

	str := domain.GenerateDomainRepository(recipes)
	saveFile(domainPath+"/I"+className+"Repository.ts", strings.NewReader(str))

	strGroup := domain.LoadConfigFile(selectedInterface)

	str = domain.GenerateClass(strGroup, className)
	saveFile(domainPath+"/"+className+".ts", strings.NewReader(str))

	str = domain.GenerateInterface(strGroup, className)
	saveFile(domainPath+"/I"+className+".ts", strings.NewReader(str))

	str = domain.GenerateDomainIndex(recipes)
	saveFile(domainPath+"/index.ts", strings.NewReader(str))
}
