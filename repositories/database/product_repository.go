package database

import "github.com/Heroin-lab/nixProject/repositories/models"

type ProductRepose struct {
	storage *Storage
}

func (r *ProductRepose) GetByCategory() ([]*models.Products, error) {
	u := &models.Products{}

	allProdSql, err := r.storage.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	allProdArr := make([]*models.Products, 0)

	for allProdSql.Next() {

		err = allProdSql.Scan(
			&u.Id,
			&u.Product_name,
			&u.Category_id,
			&u.Price,
			&u.Description,
			&u.Amount_left,
			&u.Supplier_id,
		)
		if err != nil {
			return nil, err
		}

		allProdArr = append(allProdArr, &models.Products{
			Id:           u.Id,
			Product_name: u.Product_name,
			Category_id:  u.Category_id,
			Price:        u.Price,
			Description:  u.Description,
			Amount_left:  u.Amount_left,
			Supplier_id:  u.Supplier_id,
		})
	}
	return allProdArr, nil
}
