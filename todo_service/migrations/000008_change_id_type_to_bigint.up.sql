-- Modify columns in the "list" table
ALTER TABLE list
ALTER COLUMN id SET DATA TYPE bigint,
ALTER COLUMN user_id SET DATA TYPE bigint;

-- Modify the "list_id" column in the "todos" table
ALTER TABLE todos
ALTER COLUMN list_id SET DATA TYPE bigint;

-- Modify columns in the "file" table
ALTER TABLE file
ALTER COLUMN id SET DATA TYPE bigint,
ALTER COLUMN todo_id SET DATA TYPE bigint;
