package database

import (
	"github.com/Heroin-lab/nixProject/models"
)

type OrderRepose struct {
	storage *Storage
}

func (r *OrderRepose) GetAllUserOrders(uid string) ([]*models.Order, error) {
	orderModel := &models.Order{}
	productModel := &models.ForSelectProducts{}

	ordersSet, err := r.storage.DB.Query("SELECT id AS 'order_id', paid_status, address, price \n"+
		"FROM orders \n"+
		"WHERE user_id=?", uid)
	if err != nil {
		return nil, err
	}

	productsSet, err := r.storage.DB.Query("SELECT o.id AS 'id_for_prod', product_name, category_name, title, quantity\n"+
		"FROM orders_to_products\n"+
		"INNER JOIN products p on orders_to_products.products_id = p.id\n"+
		"INNER JOIN orders o on orders_to_products.order_id = o.id\n"+
		"INNER JOIN categories c on p.category_id = c.id\n"+
		"INNER JOIN suppliers s on p.supplier_id = s.id\n"+
		"WHERE user_id=?\n"+
		"ORDER BY o.id\n"+
		"DESC;", uid)
	if err != nil {
		return nil, err
	}

	ordersArray := make([]*models.Order, 0)
	productsArray := make([]models.ForSelectProducts, 0)

	for productsSet.Next() {
		err = productsSet.Scan(
			&productModel.Id,
			&productModel.Product_name,
			&productModel.Category_name,
			&productModel.Title,
			&productModel.Quantity,
		)
		if err != nil {
			return nil, err
		}

		productsArray = append(productsArray, models.ForSelectProducts{
			Id:            productModel.Id,
			Product_name:  productModel.Product_name,
			Category_name: productModel.Category_name,
			Title:         productModel.Title,
			Quantity:      productModel.Quantity,
		})
	}

	for ordersSet.Next() {
		err = ordersSet.Scan(
			&orderModel.Id,
			&orderModel.Paid_status,
			&orderModel.Address,
			&orderModel.Price,
		)
		if err != nil {
			return nil, err
		}

		ordersArray = append(ordersArray, &models.Order{
			Id:          orderModel.Id,
			Paid_status: orderModel.Paid_status,
			Address:     orderModel.Address,
			Price:       orderModel.Price,
		})
	}

	for i := 0; i < len(ordersArray); i++ {
		for j := 0; j < len(productsArray); j++ {
			if ordersArray[i].Id == productsArray[j].Id {
				ordersArray[i].ProductArr = append(ordersArray[i].ProductArr, productsArray[j])
			}
		}
	}

	return ordersArray, nil
}
