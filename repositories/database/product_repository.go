package database

import (
	models "github.com/Heroin-lab/nixProject/repositories/models"
)

type ProductRepose struct {
	storage *Storage
}

func (r *ProductRepose) GetByCategory(category string) ([]*models.SelectProducts, error) {
	u := &models.SelectProducts{}

	allProdSql, err := r.storage.db.Query(
		"SELECT products.id, product_name, category_name, price, description, amount_left, title\n"+
			"FROM products\n"+
			"INNER JOIN categories c on products.category_id = c.id\n"+
			"INNER JOIN suppliers s on products.supplier_id = s.id\n"+
			"WHERE category_name = ?", category)
	if err != nil {
		return nil, err
	}

	allProdArr := make([]*models.SelectProducts, 0)

	for allProdSql.Next() {

		err = allProdSql.Scan(
			&u.Id,
			&u.Product_name,
			&u.Category_name,
			&u.Price,
			&u.Description,
			&u.Amount_left,
			&u.Title,
		)
		if err != nil {
			return nil, err
		}

		allProdArr = append(allProdArr, &models.SelectProducts{
			Id:            u.Id,
			Product_name:  u.Product_name,
			Category_name: u.Category_name,
			Price:         u.Price,
			Description:   u.Description,
			Amount_left:   u.Amount_left,
			Title:         u.Title,
		})
	}
	return allProdArr, nil
}
