package database

import "github.com/Heroin-lab/nixProject/repositories/models"

type ProductRepose struct {
	storage *Storage
}

func (r *ProductRepose) GetByCategory(cat_name string) ([]*models.ForSelectProducts, error) {
	u := &models.ForSelectProducts{}

	allProdSql, err := r.storage.db.Query("SELECT product_name, category_name, price, prod_desc, amount_left, title\n"+
		"FROM products\n"+
		"INNER JOIN categories c on products.category_id = c.id\n"+
		"INNER JOIN suppliers s on products.supplier_id = s.id\n"+
		"WHERE category_name = ?", cat_name)
	if err != nil {
		return nil, err
	}
	defer allProdSql.Close()

	allProdArr := make([]*models.ForSelectProducts, 0)

	for allProdSql.Next() {

		err = allProdSql.Scan(
			&u.Product_name,
			&u.Category_name,
			&u.Price,
			&u.Prod_desc,
			&u.Amount_left,
			&u.Title,
		)
		if err != nil {
			return nil, err
		}

		allProdArr = append(allProdArr, &models.ForSelectProducts{
			Product_name:  u.Product_name,
			Category_name: u.Category_name,
			Price:         u.Price,
			Prod_desc:     u.Prod_desc,
			Amount_left:   u.Amount_left,
			Title:         u.Title,
		})
	}
	return allProdArr, nil
}

func (r *ProductRepose) InsertItem(p *models.Products) (*models.Products, error) {
	_, err := r.storage.db.Exec("INSERT INTO products (product_name, category_id, price, prod_desc, amount_left, supplier_id)\n "+
		"VALUES (?, ?, ?, ?, ?, ?)",
		p.Product_name,
		p.Category_id,
		p.Price,
		p.Prod_desc,
		p.Amount_left,
		p.Supplier_id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
