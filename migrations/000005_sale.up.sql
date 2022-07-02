CREATE TABLE IF NOT EXISTS sale
(
    product_number    int NOT NULL,
    selling_price     decimal(13,4) NOT NULL,
    fk_upc varchar(12) references store_product(upc) ON UPDATE CASCADE ON DELETE NO ACTION NOT NULL,
    fk_check_number varchar(10) references bill(bill_number) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL
   -- CONSTRAINT fk_upc
     --   FOREIGN KEY(upc)
       --     REFERENCES store_product(upc)
         --   ON UPDATE CASCADE ON DELETE NO ACTION,

  --  CONSTRAINT fk_check_number
    --    FOREIGN KEY(check_number)
      --      REFERENCES check(check_number)
        --    ON UPDATE CASCADE ON DELETE CASCADE,
  --  PRIMARY KEY (fk_upc, fk_check_number)
);
-- ALTER TABLE subs ADD PRIMARY KEY (id);