CREATE TABLE Orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    client_name TEXT,
    create_time_ms INT,
    complete_time_ms INT,
    delivery_address TEXT,
    phone_number TEXT
);

CREATE TABLE Stock (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name TEXT,
    quantity INT
);

CREATE TABLE PizzaIngredients (
    id INT PRIMARY KEY AUTO_INCREMENT,
    pizza_id INT,
    stock_id INT,
    quantity INT
);

CREATE TABLE OrderedPizza (
    id INT PRIMARY KEY AUTO_INCREMENT,
    pizza_id INT,
    order_id INT,
    status INT DEFAULT 0
);

CREATE TABLE Pizza (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name TEXT NOT NULL,
    cook_time_ms INT,
    difficult INT
);
