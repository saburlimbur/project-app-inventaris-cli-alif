CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    name VARCHAR(150) NOT NULL UNIQUE,
    price NUMERIC(12,2) NOT NULL CHECK (price > 0),
    purchase_date DATE NOT NULL,
    usage_days INT NOT NULL DEFAULT 0 CHECK (usage_days >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ============================================================
--  VIEW: items_need_replace (barang yg udah dipakai lebih dari 100 hari)
-- ============================================================

CREATE OR REPLACE VIEW items_need_replace AS
SELECT *
FROM items
WHERE usage_days > 100;

-- ============================================================
--  VIEW: items_investment (laporan depresiasi & nilai saat ini)
-- ============================================================

CREATE OR REPLACE VIEW items_investment AS
SELECT 
    id,
    name,
    price AS initial_price,
    EXTRACT(YEAR FROM AGE(NOW(), purchase_date)) AS years_used,
    price * POWER(0.8, EXTRACT(YEAR FROM AGE(NOW(), purchase_date))) AS current_value
FROM items;

-- ============================================================
-- INDEX: search nama barang
-- ============================================================

CREATE INDEX IF NOT EXISTS idx_items_name 
ON items USING gin (to_tsvector('simple', name));

-- ============================================================
-- INSERT DEFAULT DUMMY DATA (KATEGORI + BARANG)
-- ============================================================

INSERT INTO categories (name, description)
VALUES
('Elektronik', 'Perangkat elektronik kantor'),
('Furniture', 'Perabot kantor seperti kursi dan meja'),
('ATK', 'Alat tulis kantor'),
('Kebersihan', 'Peralatan kebersihan gedung')
ON CONFLICT (name) DO NOTHING;

INSERT INTO items (category_id, name, price, purchase_date, usage_days)
VALUES
(1, 'Laptop Lenovo ThinkPad', 12000000, '2024-01-10', 120),
(1, 'Monitor Dell 24 inch', 3500000, '2024-03-01', 80),
(2, 'Kursi Ergonomis', 1500000, '2023-12-10', 200),
(3, 'Stapler Besar', 75000, '2024-05-12', 40),
(4, 'Vacuum Cleaner', 950000, '2023-11-15', 180)
ON CONFLICT (name) DO NOTHING;
