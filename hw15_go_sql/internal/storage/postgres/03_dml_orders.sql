-- Insert order
INSERT INTO Orders (user_id, order_date, total_amount) VALUES
	(5, now( ), 4);
-- Insert a link between the order and the product
INSERT INTO OrderProducts (order_id, product_id, amount) VALUES
	(1, 7, 2);

-- Insert order
INSERT INTO Orders (user_id, order_date, total_amount) VALUES
	(4, now( ), 10);
-- Insert a link between the order and the product
INSERT INTO OrderProducts (order_id, product_id, amount) VALUES
	(2, 7, 5);

-- Insert order
INSERT INTO Orders (user_id, order_date, total_amount) VALUES
	(5, now( ), 12.78);
-- Insert a link between the order and the product
INSERT INTO OrderProducts (order_id, product_id, amount) VALUES
	(3, 4, 1);

-- Insert order
INSERT INTO Orders (user_id, order_date, total_amount) VALUES
	(1, now( ), 2.1);
-- Insert a link between the order and the product
INSERT INTO OrderProducts (order_id, product_id, amount) VALUES
	(4, 2, 6);

-- Insert order
INSERT INTO Orders (user_id, order_date, total_amount) VALUES
	(3, now( ), 0.3);
-- Insert a link between the order and the product
INSERT INTO OrderProducts (order_id, product_id, amount) VALUES
	(5, 4, 3);

-- Insert order
INSERT INTO Orders (user_id, order_date, total_amount) VALUES
	(5, now( ), 39);
-- Insert a link between the order and the product
INSERT INTO OrderProducts (order_id, product_id, amount) VALUES
	(6, 6, 1);

-- Delete order
DELETE FROM OrderProducts WHERE order_id=3;
DELETE FROM Orders WHERE id=3;
