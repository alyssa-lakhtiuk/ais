CREATE TABLE IF NOT EXISTS product
(
    id_product      int PRIMARY KEY,
    product_name    varchar(50) NOT NULL,
    description     varchar(100) NOT NULL,

    CONSTRAINT fk_category_number
        FOREIGN KEY(category_number)
            REFERENCES category(category_number)
            ON UPDATE CASCADE ON DELETE NO ACTION
);
-- "characteristics" already booked word by SQL, so we use "description"