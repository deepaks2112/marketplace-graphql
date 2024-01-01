package backend

import (
	"strings"
)

type Product struct {
	Id int
	Name string
}

func GetProducts(repo Repository, q string) ([]Product, error) {
	query := "SELECT p1.id, p1.name FROM product p1 WHERE LOWER(p1.name) like '%' || $1 || '%'"
	rows, err := repo.DB.Query(query, strings.ToLower(q))
	if err != nil {
		return nil, err
	}
	products := make([]Product, 0)
	for rows.Next() {
		var p Product
		if rows.Scan(&p.Id, &p.Name) != nil {
			continue
		}
		products = append(products, p)
	}
	return products, nil
}