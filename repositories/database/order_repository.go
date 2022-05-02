package database

import (
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/models"
	"strconv"
	"strings"
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

	productsSet, err := r.storage.DB.Query("SELECT o.id AS 'id_for_prod', product_name, product_type, supp_name, quantity, p.price, p.img\n"+
		"FROM orders_to_products\n"+
		"INNER JOIN products p on orders_to_products.products_id = p.id\n"+
		"INNER JOIN orders o on orders_to_products.order_id = o.id\n"+
		"INNER JOIN products_types c on p.type_id = c.id\n"+
		"INNER JOIN suppliers s on p.supplier_id = s.id\n"+
		"WHERE user_id=?\n"+
		"ORDER BY o.id\n"+
		"DESC;", uid)
	if err != nil {
		return nil, err
	}

	ordersArray := make([]*models.Order, 0)
	productsArray := make([]*models.ForSelectProducts, 0)

	for productsSet.Next() {
		err = productsSet.Scan(
			&productModel.Id,
			&productModel.Product_name,
			&productModel.Prod_type_name,
			&productModel.Supplier,
			&productModel.Quantity,
			&productModel.Price,
			&productModel.Img,
		)
		if err != nil {
			return nil, err
		}

		productsArray = append(productsArray, &models.ForSelectProducts{
			Id:             productModel.Id,
			Product_name:   productModel.Product_name,
			Prod_type_name: productModel.Prod_type_name,
			Supplier:       productModel.Supplier,
			Quantity:       productModel.Quantity,
			Price:          productModel.Price,
			Img:            productModel.Img,
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
				ordersArray[i].ProductArr = append(ordersArray[i].ProductArr, *productsArray[j])
			}
		}
	}

	logger.Info("All orders list was sent to user with id=" + uid)
	return ordersArray, nil
}

func (r *OrderRepose) AddOrder(insert *models.OrderForInsert) error {
	orderResult, err := r.storage.DB.Exec("INSERT INTO orders (paid_status, address, price, user_id, phone, first_name, second_name)"+
		" VALUES (?, ?, ?, ?, ?, ?, ?);",
		insert.Paid_status,
		insert.Address,
		insert.Price,
		insert.User_id,
		insert.Phone,
		insert.First_name,
		insert.Second_name)
	if err != nil {
		return err
	}

	lastOne, _ := orderResult.LastInsertId()
	newValues := strings.ReplaceAll(insert.ProductArr, "?", strconv.FormatInt(lastOne, 10))

	_, err = r.storage.DB.Exec("INSERT INTO orders_to_products (order_id, products_id, quantity)\n" +
		"VALUES" + newValues)
	if err != nil {
		return err
	}

	logger.Info("User with id: " + insert.User_id + "was successfully make an order")
	return nil
}

func (r *OrderRepose) DeleteOrder(deleteId string) error {
	_, err := r.storage.DB.Exec("DELETE FROM orders WHERE id=?", deleteId)
	if err != nil {
		return err
	}

	logger.Info("Order with 'order_id='" + deleteId + "was successfully deleted!")
	return nil
}
