CREATE TABLE IF NOT EXISTS roles (
    empl_id  varchar(10) PRIMARY KEY,
    role    varchar(50) NOT NULL,
    password    varchar(100) NOT NULL,
    phone    varchar(13) NOT NULL
);
