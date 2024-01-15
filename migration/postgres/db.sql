CREATE TYPE user_role_enum as enum('admin','customer');
CREATE TABLE users(
    id uuid primary key not null,
    fullname varchar(30),
    phone varchar(30) unique not null,
    password_ varchar(30) not null,
    cash int,
    user_role user_role_enum not null

);
CREATE TABLE category(
    id uuid primary key not null,
    name_ varchar(19)
);

CREATE TABLE products(
    id uuid primary key not null,
    name_ varchar(16),
    price int,
    originalprice int,
    quantity int,
    categoryid uuid references category(id)

);
CREATE TABLE basket(
    id uuid primary key not null,
    customerid uuid references users(id),
    totalsum int

);
CREATE TABLE basketproducts(
    id uuid primary key not null,
    basketid uuid references basket(id),
    productid uuid references products(id),
    quantity int
);