package backend

import (
	"fmt"
	"strings"
)

type Seller struct {
	Id int
	Name string
	Price float32
}

func GetSellersByQuery(repo Repository, q string) ([]Seller, error) {
	query := "SELECT s.id, s.name FROM seller s WHERE LOWER(s.name) like '%' || $1 || '%'"
	rows, err := repo.DB.Query(query, strings.ToLower(q))
	if err != nil {
		return nil, err
	}
	sellers := make([]Seller, 0)
	for rows.Next() {
		var s Seller
		if rows.Scan(&s.Id, &s.Name) != nil {
			continue
		}
		sellers = append(sellers, s)
	}
	return sellers, nil
}

func GetSellersForProduct(repo Repository, productId int) ([]Seller, error) {
	query := "SELECT s.id, s.name, sp.price::money::numeric::float32 FROM seller s, seller_product sp WHERE s.id = sp.seller_id AND sp.product_id = $1"
	rows, err := repo.DB.Query(query, productId)
	if err != nil {
		return nil, err
	}
	sellers := make([]Seller, 0)
	for rows.Next() {
		var s Seller
		err := rows.Scan(&s.Id, &s.Name, &s.Price)
		if err != nil {
			fmt.Println("err: ", err.Error())
			continue
		}
		sellers = append(sellers, s)
	}
	return sellers, nil
}
