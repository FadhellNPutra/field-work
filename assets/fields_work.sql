CREATE DATABASE IF NOT EXISTS field_work_db;
CREATE EXTENSION IF NOT EXISTS 'uuid-ossp';
CREATE EXTENSION IF NOT EXISTS 'pgcrypto';

CREATE TYPE  role_type AS ENUM('customer', 'admin');

CREATE TABLE users(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    address TEXT NOT NULL,
    role role_type DEFAULT 'customer';
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP 
);

CREATE TYPE  size_type AS ENUM('S', 'M', 'L', 'XL');

CREATE TABLE products(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    quantity BIGINT NOT NULL,
    price BIGINT NOT NULL,
    material VARCHAR(50),
    description TEXT,
    photo VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP 
);

CREATE TYPE status_type AS ENUM('received','process', 'sending', 'done')

CREATE TABLE transaction(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id uuid NOT NULL,
    product_id uuid NOT NULL,
    total_price BIGINT NOT NULL,
    total_quantity BIGINT NOT NULL,
    status status_type NOT NULL,
    size size_type NOT NULL,
    color VARCHAR(50) NOT NULL,
    customer_message TEXT,
    FOREIGN KEY user_id REFERENCES users('id')
    FOREIGN KEY products_id REFERENCES products('id')
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP 
);

CREATE TABLE blog(
    id uuid DEFAULT uuid_generate_v4(),
    title VARCHAR(100) NOT NULL,
    article TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);