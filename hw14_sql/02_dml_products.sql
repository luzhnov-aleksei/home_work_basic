-- Insert products
INSERT INTO Products(id, name, price) VALUES
    (1, 'milk', 5),
    (2, 'bread', 2.13),
    (3, 'tomato', 0.5),
    (4, 'apple', 0.7),
    (5, 'cake', 4.1),
    (6, 'banana', 0.3),
    (7, 'donut', 2);

-- Update products
UPDATE Products SET name='milk', price=6.5 WHERE id=1;

-- Delete product
DELETE FROM OrderProducts WHERE product_id=3;