-- 20240529133229 - adding-table migration

CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    title TEXT, 
    completed integer DEFAULT 0,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);