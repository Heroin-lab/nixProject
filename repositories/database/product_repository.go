package database

import (
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/models"
)

type ProductRepose struct {
	storage *Storage
}

func (r *ProductRepose) GetByCategory(category string) ([]*models.ForSelectProducts, error) {
	u := &models.ForSelectProducts{}

	allProdSql, err := r.storage.DB.Query("SELECT products.id, product_name, product_type, price, img, supp_name\n"+
		"FROM products\n"+
		"INNER JOIN products_types c on products.type_id = c.id\n"+
		"INNER JOIN suppliers s on products.supplier_id = s.id\n"+
		"WHERE product_type = ?", category)
	if err != nil {
		return nil, err
	}
	defer allProdSql.Close()

	allProdArr := make([]*models.ForSelectProducts, 0)

	for allProdSql.Next() {

		err = allProdSql.Scan(
			&u.Id,
			&u.Product_name,
			&u.Prod_type_name,
			&u.Price,
			&u.Img,
			&u.Supplier,
		)
		if err != nil {
			return nil, err
		}

		allProdArr = append(allProdArr, &models.ForSelectProducts{
			Id:             u.Id,
			Product_name:   u.Product_name,
			Prod_type_name: u.Prod_type_name,
			Price:          u.Price,
			Img:            u.Img,
			Supplier:       u.Supplier,
		})
	}
	return allProdArr, nil
}

func (r *ProductRepose) InsertItem(p *models.Products) (*models.Products, error) {
	_, err := r.storage.DB.Exec("INSERT INTO products (product_name, type_id, price, img, supplier_id)\n "+
		"VALUES (?, ?, ?, ?, ?)",
		p.Product_name,
		p.Type_id,
		p.Price,
		p.Img,
		p.Supplier_id)
	if err != nil {
		return nil, err
	}

	logger.Info("Row with name '" + p.Product_name + "' was successfully added to PRODUCT table!")
	return p, nil
}

func (r *ProductRepose) DeleteItem(stringToDelete string) error {
	_, err := r.storage.DB.Exec("DELETE FROM products WHERE id=?",
		stringToDelete,
	)
	if err != nil {
		return err
	}

	logger.Info("Row with name '" + stringToDelete + "' was successfully deleted from PRODUCTS table!")
	return nil
}

func (r *ProductRepose) UpdateItem(p *models.Products) error {
	rows, err := r.storage.DB.Query("UPDATE products\n"+
		"SET product_name=?, type_id=?,\n"+
		"price=?, img=?,\n"+
		"supplier_id=?\n"+
		"WHERE id=?",
		p.Product_name,
		p.Type_id,
		p.Price,
		p.Img,
		p.Supplier_id,
		p.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	logger.Info("Row in PRODUCTS table was updated! RowID=", p.Id)
	return nil
}
