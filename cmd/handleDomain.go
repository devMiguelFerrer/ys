package cmd

import (
	"strings"

	"github.com/devMiguelFerrer/ys/pkg/domain"
	"github.com/devMiguelFerrer/ys/pkg/shared"
)

func generateDomain(recipes []shared.Recipe) {
	if selectedValidations {
		generateDomainWithVO(recipes)
	} else {
		generateDomainWithoutVO(recipes)
	}
}

func generateDomainWithoutVO(recipes []shared.Recipe) {

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

func generateDomainWithVO(recipes []shared.Recipe) {

	domainPath := fullPath + "/Domain"
	domainVOPath := domainPath + "/Validations"

	createDirectory(domainPath)
	createDirectory(domainVOPath)

	saveFile(domainPath+"/ICriteria.ts", strings.NewReader(domain.Ifilter))

	str := domain.GenerateDomainRepository(recipes)
	saveFile(domainPath+"/I"+className+"Repository.ts", strings.NewReader(str))

	strGroup := domain.LoadConfigFile(selectedInterface)

	str = domain.GenerateClassVO(strGroup, className)
	saveFile(domainPath+"/"+className+".ts", strings.NewReader(str))

	str = domain.GenerateInterface(strGroup, className)
	saveFile(domainPath+"/I"+className+".ts", strings.NewReader(str))

	str = domain.GenerateDomainIndex(recipes)
	saveFile(domainPath+"/index.ts", strings.NewReader(str))

	str = domain.GenerateDomainBaseValidations()
	saveFile(domainVOPath+"/BaseTypes.ts", strings.NewReader(str))
}
