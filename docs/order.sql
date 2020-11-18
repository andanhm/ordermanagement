DROP DATABASE anglebrokerdb;
CREATE DATABASE anglebrokerdb;

CREATE TABLE tblCustomerOrder (
    id INTEGER AUTO_INCREMENT,
    customer_id INTEGER,
    amount INTEGER,
    status VARCHAR(12),
    constraint pk_order primary key(id)
) ENGINE = InnoDB;

CREATE TABLE tblInvoice (
    id INTEGER AUTO_INCREMENT,
    order_id INTEGER,
    customer_id INTEGER,
    creation_date DATETIME,
    constraint pk_order primary key(id)
) ENGINE = InnoDB;

CREATE TABLE tblProduct (
    id INTEGER AUTO_INCREMENT,
    name VARCHAR(40) not null,
    price DECIMAL(9, 2),
    quantity INTEGER,
    constraint pk_product primary key(id)
) ENGINE = InnoDB;