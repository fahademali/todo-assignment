-- Add unique constraints to the "username" and "email" columns
ALTER TABLE users
ADD CONSTRAINT unique_username UNIQUE (username),
    ADD CONSTRAINT unique_email UNIQUE (email);
-- Set the "password" column to not allow null values
ALTER TABLE users
ALTER COLUMN password
SET NOT NULL;