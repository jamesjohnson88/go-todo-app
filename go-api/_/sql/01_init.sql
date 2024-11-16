-- Create the database (run this when connected to postgres)
CREATE DATABASE go_todo_app;

-- Then connect to the todo_app database before running:
\c go_todo_app

-- Create UUID extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Now create the table
CREATE TABLE todo_items (
                            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                            title VARCHAR NOT NULL,
                            description TEXT,
                            completed BOOLEAN DEFAULT FALSE NOT NULL,
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
                            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
                            due_date TIMESTAMP WITH TIME ZONE,
                            priority INTEGER DEFAULT 0 NOT NULL,
                            user_id UUID
);

-- Create the update trigger
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_todo_items_updated_at
    BEFORE UPDATE ON todo_items
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();