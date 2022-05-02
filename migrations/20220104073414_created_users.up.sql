CREATE TABLE `basket` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `paid_status` tinyint(1) NOT NULL DEFAULT '0',
                          `adress` varchar(255) NOT NULL,
                          `price` varchar(128) NOT NULL,
                          PRIMARY KEY (`id`)
);

CREATE TABLE `categories` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `category_name` varchar(255) NOT NULL,
                              PRIMARY KEY (`id`)
);

CREATE TABLE `suppliers_type` (
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `type_name` varchar(255) NOT NULL,
                                  PRIMARY KEY (`id`)
);

CREATE TABLE `suppliers` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `title` int DEFAULT NULL,
                             `type_id` int NOT NULL,
                             `working_time` varchar(255) NOT NULL,
                             PRIMARY KEY (`id`),
                             KEY `suppliers_to_type` (`type_id`),
                             CONSTRAINT `suppliers_to_type` FOREIGN KEY (`type_id`) REFERENCES `suppliers_type` (`id`)
);

CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `email` varchar(128) NOT NULL,
                         `password` int NOT NULL,
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
);

CREATE TABLE `products` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `product_name` varchar(255) NOT NULL,
                            `category_id` int NOT NULL,
                            `price` varchar(128) NOT NULL,
                            `description` text NOT NULL,
                            `amount_left` int DEFAULT NULL,
                            `supplier_id` int NOT NULL,
                            PRIMARY KEY (`id`),
                            KEY `products_to_categories` (`category_id`),
                            KEY `products_to_supplier` (`supplier_id`),
                            CONSTRAINT `products_to_categories` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
                            CONSTRAINT `products_to_supplier` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`)
);

CREATE TABLE `orders` (
                          `user_id` int NOT NULL,
                          `basket_id` int NOT NULL,
                          KEY `order_to_user` (`user_id`),
                          KEY `order_to_basket` (`basket_id`),
                          CONSTRAINT `order_to_basket` FOREIGN KEY (`basket_id`) REFERENCES `basket` (`id`),
                          CONSTRAINT `order_to_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `basket_to_products` (
                                      `basket_id` int NOT NULL,
                                      `products_id` int NOT NULL,
                                      KEY `basket_to_products` (`basket_id`),
                                      KEY `basket_to_product_id` (`products_id`),
                                      CONSTRAINT `basket_to_product_id` FOREIGN KEY (`products_id`) REFERENCES `products` (`id`),
                                      CONSTRAINT `basket_to_products` FOREIGN KEY (`basket_id`) REFERENCES `basket` (`id`)
);