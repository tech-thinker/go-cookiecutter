package api

import "github.com/mrasif/gomvc/repository"

// ArticleService is interface for articleService
type ArticleService interface {
}

type articleService struct {
	articleRepo repository.ArticleRepo
}

// NewArticleService initializes articleService
func NewArticleService(
	articleRepo repository.ArticleRepo,
) ArticleService {
	return &articleService{
		articleRepo: articleRepo,
	}
}
