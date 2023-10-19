CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100),
    role VARCHAR(100) CHECK (role IN ('admin', 'user')),
    is_verified BOOLEAN DEFAULT false
)