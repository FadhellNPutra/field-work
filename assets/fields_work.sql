CREATE DATABASE field_work_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE role_type AS ENUM('Customer', 'Admin');

CREATE TABLE users(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	username VARCHAR(50) NOT NULL UNIQUE,
	PASSWORD VARCHAR(128) NOT NULL,
	address TEXT,
	role role_type DEFAULT 'Customer',
	created_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE size_type AS ENUM('S', 'M', 'L', 'XL');

CREATE TABLE products(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_name VARCHAR(100) NOT NULL,
	quantity BIGINT NOT NULL,
	price BIGINT NOT NULL,
	material VARCHAR(50),
	summary VARCHAR(250),
	description TEXT,
	created_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_images (
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_id uuid NOT NULL,
	file_name VARCHAR(100),
	FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TYPE status_type AS ENUM('Received', 'Process', 'Sending', 'Done');

CREATE TABLE transactions(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	user_id uuid NOT NULL,
	product_id uuid NOT NULL,
	total_price BIGINT NOT NULL,
	total_quantity BIGINT NOT NULL,
	STATUS status_type NOT NULL,
	size size_type NOT NULL,
	color VARCHAR(50) NOT NULL,
	customer_message TEXT,
	created_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE blogs(
	id uuid DEFAULT uuid_generate_v4(),
	title VARCHAR(100) NOT NULL,
	article TEXT NOT NULL,
	created_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP
);