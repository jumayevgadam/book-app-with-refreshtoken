-- authors table
CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    bio TEXT,
    avatar VARCHAR(100),
    password VARCHAR(100)
);

-- books table
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    author_id INT REFERENCES authors (id) ON DELETE CASCADE,
    title VARCHAR(100),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);