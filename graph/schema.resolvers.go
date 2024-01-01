package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	// "fmt"
	"strconv"

	"deepaks2112.github.io/marketplace/backend"
	"deepaks2112.github.io/marketplace/graph/model"
)

// GetProducts is the resolver for the getProducts field.
func (r *queryResolver) GetProducts(ctx context.Context, q *string) ([]*model.Product, error) {
	// panic(fmt.Errorf("not implemented: GetProducts - getProducts"))
	products, err := backend.GetProducts(r.repo, *q)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Product, len(products))
	for i, p := range products {
		sellers, err := backend.GetSellersForProduct(r.repo, p.Id)
		if err != nil {
			continue
		}
		sellerForProduct := make([]*model.Seller, len(sellers))
		for j, s := range sellers {
			sellerForProduct[j] = &model.Seller{
				ID: strconv.FormatInt(int64(s.Id), 10),
				Name: s.Name,
				Price: strconv.FormatFloat(float64(s.Price), 'e', -1, 32),
			}
		}
		result[i] = &model.Product{
			ID: strconv.FormatInt(int64(p.Id), 10),
			Name: p.Name,
			Price: strconv.FormatFloat(float64(sellers[0].Price), 'e', -1, 32),
			Sellers: sellerForProduct,
		}
	}
	return result, nil
}

// GetSellers is the resolver for the getSellers field.
func (r *queryResolver) GetSellers(ctx context.Context, q *string) ([]*model.Seller, error) {
	// panic(fmt.Errorf("not implemented: GetSellers - getSellers"))
	sellers, err := backend.GetSellersByQuery(r.repo, *q)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Seller, len(sellers))
	for i, s := range sellers {
		result[i] = &model.Seller{
			ID: strconv.FormatInt(int64(s.Id), 10),
			Name: s.Name,
			Price: strconv.FormatFloat(float64(s.Price), 'e', -1, 32),
		}
	}
	return result, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { 
	return &queryResolver{r} 
}

type queryResolver struct{ *Resolver }