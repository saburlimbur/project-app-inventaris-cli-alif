TABLE categories {
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
}

TABLE items {
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL REFERENCES categories(id),
    name VARCHAR(150) NOT NULL UNIQUE,
    price NUMERIC(12,2) NOT NULL CHECK (price > 0),
    purchase_date DATE NOT NULL,
    usage_days INT NOT NULL DEFAULT 0 CHECK (usage_days >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
}
