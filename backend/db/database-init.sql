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
CREATE TABLE IF NOT EXISTS photos (
    id           SERIAL PRIMARY KEY,
    url          VARCHAR(255) NOT NULL,
    category     VARCHAR(50)  NOT NULL,
    title        VARCHAR(100),
    description  TEXT,
    uploaded_at  TIMESTAMP    NOT NULL DEFAULT NOW()
);

-- 4. products table
CREATE TABLE IF NOT EXISTS products (
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(150)  NOT NULL,
    description  TEXT,
    price        DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    stock        INT           NOT NULL DEFAULT 0,
    product_url  VARCHAR(255)  NOT NULL,
    created_at   TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP     NOT NULL DEFAULT NOW()
);
