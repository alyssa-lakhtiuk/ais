CREATE TABLE IF NOT EXISTS sale
(
    product_number    int NOT NULL,
    selling_price     decimal(13,4) NOT NULL,

    CONSTRAINT fk_upc
        FOREIGN KEY(upc)
            REFERENCES store_product(upc)
            ON UPDATE CASCADE ON DELETE NO ACTION,

    CONSTRAINT fk_check_number
        FOREIGN KEY(check_number)
            REFERENCES check(check_number)
            ON UPDATE CASCADE ON DELETE CASCADE,
    PRIMARY KEY (fk_upc, fk_check_number)
);