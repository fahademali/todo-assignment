ALTER TABLE todos
ADD creation_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
-- FK to list relation
ALTER TABLE todos
ADD COLUMN list_id INTEGER REFERENCES list(id);