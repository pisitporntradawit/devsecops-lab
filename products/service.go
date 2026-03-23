package products

import(
	"context"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service{
	return &Service{
		Repo : repo,
	}
}


func (s *Service) GetProducts(ctx context.Context) ([]ProductsModel, error){
	return s.Repo.GetProducts(ctx)
}

func(s *Service) CreateProduct(ctx context.Context, InsertProduct *ProductsModel) error {
	return s.Repo.CreatProduct(ctx, InsertProduct)
}