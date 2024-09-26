-- Select Users
SELECT * FROM Users;

-- Select Products
SELECT * FROM Products;

-- Select orders by user
SELECT * FROM Orders WHERE user_id=5;

-- Select user statistics
SELECT SUM(Orders.total_amount) AS total_order_amount,
    AVG(Products.price) AS avg_product_price
    FROM Users
    JOIN Orders ON 5 = Orders.user_id
    JOIN OrderProducts ON Orders.id = OrderProducts.order_id
    JOIN Products ON OrderProducts.product_id = Products.id;