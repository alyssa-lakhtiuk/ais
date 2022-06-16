CREATE TABLE IF NOT EXISTS bill
(
    bill_number    varchar(10) PRIMARY KEY,
    print_date     timestamp NOT NULL,
    sum_total      decimal(13,4) NOT NULL,
    vat            decimal(13,4) NOT NULL,

    CONSTRAINT fk_id_employee
        FOREIGN KEY(id_employee)
            REFERENCES employee(id_employee)
            ON UPDATE CASCADE ON DELETE NO ACTION,

    CONSTRAINT fk_card_number
        FOREIGN KEY(card_number)
            REFERENCES customer_card(card_number)
            ON UPDATE CASCADE ON DELETE CASCADE
);

-- "check" already booked word, so we use "bill"