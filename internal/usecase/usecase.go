package usecase

import (
	"context"

	"github.com/santekno/learn-golang-restful/internal/models"
)

//go:generate mockery --name=ArticleUsecase --filename=usecase_mock.go --inpackage

type ArticleUsecase interface {
	GetAll(ctx context.Context) ([]models.ArticleResponse, error)
	GetByID(ctx context.Context, id int64) (models.ArticleResponse, error)
	Update(ctx context.Context, article models.ArticleUpdateRequest) (models.ArticleResponse, error)
	Store(ctx context.Context, article models.ArticleCreateRequest) (models.ArticleResponse, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
