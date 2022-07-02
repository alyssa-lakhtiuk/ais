CREATE TABLE IF NOT EXISTS category (
    category_number  int PRIMARY KEY,
    category_name    varchar(50) unique NOT NULL
    );

CREATE TABLE IF NOT EXISTS employee (
    id_employee      varchar(10) PRIMARY KEY,
    empl_surname     varchar(50) NOT NULL,
    empl_name        varchar(50) NOT NULL,
    empl_patronymic  varchar(50),
    empl_role        varchar(10) NOT NULL,
    salary           decimal(13,4) NOT NULL,
    date_of_birth    date NOT NULL,
    date_of_start    date NOT NULL,
    phone_number     varchar(13) NOT NULL,
    city             varchar(50) NOT NULL,
    street           varchar(50) NOT NULL,
    zip_code         varchar(9) NOT NULL
    );

CREATE TABLE IF NOT EXISTS customer_card (
    card_number      varchar(13) PRIMARY KEY,
    cust_surname     varchar(50) NOT NULL,
    cust_name        varchar(50) NOT NULL,
    cust_patronymic  varchar(50),
    phone_number     varchar(13) NOT NULL,
    city             varchar(50) NOT NULL,
    street           varchar(50) NOT NULL,
    zip_code         varchar(9) NOT NULL,
    discount         int NOT NULL
    );

-- "percent" already booked word by SQL, so we use "discount"