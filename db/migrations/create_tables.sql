-- Create Markets Table
CREATE TABLE IF NOT EXISTS markets (
    market_id SERIAL PRIMARY KEY,
    market_name VARCHAR(100) NOT NULL
);

-- Create Submarkets Table
CREATE TABLE IF NOT EXISTS submarkets (
    submarket_id SERIAL PRIMARY KEY,
    submarket_name VARCHAR(100) NOT NULL,
    market_id INT REFERENCES markets(market_id)
);

-- Create Orders Table
CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    market_id INT REFERENCES markets(market_id),
    submarket_id INT REFERENCES submarkets(submarket_id)
);
