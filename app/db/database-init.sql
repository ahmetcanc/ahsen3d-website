-- 1. contacts table
CREATE TABLE IF NOT EXISTS contacts (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100)  NOT NULL,
    email       VARCHAR(100)  NOT NULL,
    subject     VARCHAR(255),
    message     TEXT          NOT NULL,
    created_at  TIMESTAMP     NOT NULL DEFAULT NOW()
);

-- 2. home_contents table
CREATE TABLE IF NOT EXISTS home_contents (
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(150) NOT NULL,
    description  TEXT         NOT NULL,
    created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP    NOT NULL DEFAULT NOW()
);

-- 3. photos table
CREATE TABLE IF NOT EXISTS products (
    id           SERIAL PRIMARY KEY,
    url          VARCHAR(255) NOT NULL,
    category     VARCHAR(50)  NOT NULL,
    title        VARCHAR(100),
    description  TEXT,
    product_url  VARCHAR(255),
    uploaded_at  TIMESTAMP    NOT NULL DEFAULT NOW()
);

-- 5. users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(150) UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

