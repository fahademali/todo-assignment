CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(150),
    due_date TIMESTAMP NOT NULL,
    is_complete BOOLEAN DEFAULT FALSE,
    completion_date TIMESTAMP
)