CREATE TABLE IF NOT EXISTS store_product(
upc                  varchar(12) PRIMARY KEY,
selling_price        decimal(13,4) NOT NULL,
promotional_product  boolean NOT NULL,
product_number       int NOT NULL,
upc_prom varchar(12) references store_product(upc) ON UPDATE CASCADE ON DELETE SET NULL NULL,
fk_id_product int references product(id_product) ON UPDATE CASCADE ON DELETE NO ACTION NOT NULL
-- CONSTRAINT upc_prom
   --  FOREIGN KEY(upc)
      --  REFERENCES store_product(upc)
       -- ON UPDATE CASCADE ON DELETE SET NULL,

 -- CONSTRAINT fk_id_product
    -- FOREIGN KEY(id_product)
       -- REFERENCES product(id_product)
       -- ON UPDATE CASCADE ON DELETE NO ACTION
);