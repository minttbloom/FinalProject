#--------------
# Setup
# -------------

SET @StockCountPiperony = 100,
    @StockCountChees = 50,
    @StockCountKetchup = 200,
    @StockCountTesto = 70,
    @StockCountOvoshi = 120,
    @StockCountMeet = 100;

#-------------
# Insert Data
#--------------

INSERT INTO Pizza (name, cook_time_ms, difficult)
VALUES ('Piperony', 600000, 5),
       ('Foxy', 300000, 4),
       ('Margarita', 100000, 2),
       ('Derevenskaya', 1200000, 9);

INSERT INTO Stock (name, quantity)
VALUES ('Piperony', @StockCountPiperony),
       ('Chees', @StockCountChees),
       ('Ketchup', @StockCountKetchup),
       ('Testo', @StockCountTesto),
       ('Ovoshi', @StockCountOvoshi),
       ('Meet', @StockCountMeet);

# Piperony Ingridients
INSERT INTO PizzaIngredients (pizza_id, stock_id, quantity)
VALUES (0, 0, 2),
       (0, 1, 1),
       (0, 3, 1),
       (0, 2, 2);

# Foxy Ingridients
INSERT INTO PizzaIngredients (pizza_id, stock_id, quantity)
VALUES (1, 1, 1),
       (1, 3, 1),
       (1, 5, 1),
       (1, 2, 2);

# Margarita Ingridients
INSERT INTO PizzaIngredients (pizza_id, stock_id, quantity)
VALUES (2, 1, 2),
       (2, 3, 1),
       (2, 2, 2);

# Derevenskaya Ingridients
INSERT INTO PizzaIngredients (pizza_id, stock_id, quantity)
VALUES (3, 1, 2),
       (3, 3, 1),
       (3, 5, 2),
       (3, 4, 2),
       (3, 2, 2);

