-- Remove the unique constraint from the "username" column
ALTER TABLE users DROP CONSTRAINT IF EXISTS unique_username;
-- Remove the unique constraint from the "email" column
ALTER TABLE users DROP CONSTRAINT IF EXISTS unique_email;
-- Set the "password" column to allow null values again
ALTER TABLE users
ALTER COLUMN password DROP NOT NULL;