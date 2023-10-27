-- Add the default value of NULL for the 'description' column
ALTER TABLE todos
ALTER COLUMN description
SET DEFAULT NULL;
-- Add the default value of NULL for the 'completion_date' column
ALTER TABLE todos
ALTER COLUMN completion_date
SET DEFAULT '0001-01-01 00:00:00';