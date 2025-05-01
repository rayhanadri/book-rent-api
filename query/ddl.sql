
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255),
    status VARCHAR(255) NOT NULL DEFAULT 'ACTIVE',
    balance INT DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publisher VARCHAR(255) NOT NULL,
    published_at DATE NOT NULL,
    isbn VARCHAR(255) NOT NULL UNIQUE,
    category VARCHAR(255) NOT NULL,
    stock INT NOT NULL,
    available BOOLEAN NOT NULL,
    price INT NOT NULL,
    description TEXT
);

CREATE TABLE rents (
    id SERIAL PRIMARY KEY,
    book_id INT NOT NULL REFERENCES books(id),
    user_id INT NOT NULL REFERENCES users(id),
    quantity INT NOT NULL,
    total_price INT NOT NULL,
    rent_start_date DATE,
    rent_end_date DATE,
    rent_status VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    transaction_type VARCHAR(255) NOT NULL,
    payment_method VARCHAR(255) NOT NULL,
    amount INT NOT NULL,
    status VARCHAR(255) NOT NULL,
    description TEXT,
    user_id INT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    invoice_id VARCHAR(255),
    invoice_url VARCHAR(255),
    rent_id INT NULL REFERENCES rents(id)
);