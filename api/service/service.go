package service

import (
	"github.com/mrasif/gomvc/api"
	"github.com/mrasif/gomvc/instance"
	"github.com/mrasif/gomvc/repository"
)

// Services is interface for all service entrypoint
type Services interface {
	ArticleService() api.ArticleService
}

type services struct {
	articleService api.ArticleService
}

func (svc *services) ArticleService() api.ArticleService {
	return svc.articleService
}

// Init initializes services repo
func Init() Services {
	db := instance.DB()
	return &services{
		articleService: api.NewArticleService(
			repository.NewArticleRepo(db),
		),
	}
}
