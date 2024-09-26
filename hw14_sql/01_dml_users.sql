-- Insert users
INSERT INTO USERS(name, email, password) VALUES
    ('John Doe', 'john@example.com', 'qwerty1'),
    ('Bob', 'bob@example.com', '12345'),
    ('Alice', 'alice@example.com', 'qwerty'),
    ('Tom', 'tom@example.com', '123456'),
    ('Jerry', 'jerry@example.com', '654321');

-- Update user
UPDATE Users
    SET name='Alice Smith', email='alice@example.com', password='gnd3rtw35'
    WHERE id=3;

-- DELETE user
DELETE FROM Users WHERE name='Bob';