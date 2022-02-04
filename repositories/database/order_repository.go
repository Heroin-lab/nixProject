package database

type OrderRepose struct {
	storage *Storage
}

//SELECT o.id AS 'id_for_prod', product_name, category_name, title
//FROM orders_to_products
//INNER JOIN products p on orders_to_products.products_id = p.id
//INNER JOIN orders o on orders_to_products.order_id = o.id
//INNER JOIN categories c on p.category_id = c.id
//INNER JOIN suppliers s on p.supplier_id = s.id
//WHERE user_id=3
//ORDER BY o.id
//DESC;
//
//SELECT id AS 'order_id', paid_status, adress, price FROM orders WHERE user_id=3
