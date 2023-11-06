-- Drop the 'list_id' foreign key constraint named 'todos_list_id_fkey'
ALTER TABLE todos
DROP CONSTRAINT IF EXISTS todos_list_id_fkey;

-- Remove the 'list_id' column
ALTER TABLE todos
DROP COLUMN IF EXISTS list_id;

-- Remove the 'creation_time' column
ALTER TABLE todos
DROP COLUMN IF EXISTS creation_time;
