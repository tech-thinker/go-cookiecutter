package repository

import "github.com/astaxie/beego/orm"

// ArticleRepo is interface for article repository
type ArticleRepo interface {
}

type articleRepo struct {
	db orm.Ormer
}

// NewArticleRepo initializes articleRepo
func NewArticleRepo(db orm.Ormer) ArticleRepo {
	return &articleRepo{
		db: db,
	}
}
