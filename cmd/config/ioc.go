package config

import (
	"service-routes/internal/domain/repository"
	"service-routes/internal/domain/usecase"
	"service-routes/internal/infra/storage"
)

func genIoc(config *Configuration) map[string]interface{} {
	ioc := make(map[string]interface{})
	//database
	db := GetDB()
	repoComment := repository.NewRepositoryComments(db)
	repoResource := repository.NewRepositoryResource(db)
	repoRoutes := repository.NewRepositoryRoutes(db)
	repoSteps := repository.NewRepositorySteps(db)

	client := storage.InitStorage(GetStorageClient(), config.Credential.Gcbucket)
	//ioc
	ioc["comment"] = usecase.NewCommentsUseCase(repoComment)
	ioc["resource"] = usecase.NewResourcesUseCase(repoResource, client)
	ioc["routes"] = usecase.NewRoutesUseCase(repoRoutes, client)
	ioc["steps"] = usecase.NewStepsUseCase(repoSteps, client)
	return ioc
}
