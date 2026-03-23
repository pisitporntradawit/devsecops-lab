package products

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) * Repository{
	return &Repository{
		DB : db,
	}
}
func (r *Repository) GetProducts(ctx context.Context) ([]ProductsModel, error) {
	rows, err := r.DB.Query(ctx, `SELECT id, name From products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []ProductsModel

	for rows.Next() {
		var p ProductsModel

		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func(r *Repository) CreatProduct(ctx context.Context, InsertProduct *ProductsModel) error {
	query := `iNSERT INTO products(name) VALUES($1)`
	_, err := r.DB.Exec(ctx, query, InsertProduct.Name)
	if err != nil {
		return fmt.Errorf("Error %w", err)
	}
	return nil
}