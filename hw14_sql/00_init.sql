CREATE TABLE Users (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    password VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX idx_email ON Users (email);

CREATE TABLE Orders (
    id serial PRIMARY KEY,
    user_id INT,
    order_date DATE,
    total_amount FLOAT NOT NULL,
    FOREIGN KEY (user_id)  REFERENCES Users (id)
);

CREATE INDEX idx_orders ON Orders (user_id);

CREATE TABLE Products (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL
);

CREATE INDEX idx_products ON Products (name);

CREATE TABLE OrderProducts (
    order_id INT PRIMARY KEY,
    product_id INT NOT NULL,
    amount INT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES Orders(id),
    FOREIGN KEY (product_id) REFERENCES Products(id)
);