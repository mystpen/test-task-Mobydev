CREATE TABLE IF NOT EXISTS videos (
    id bigserial PRIMARY KEY,  
    title TEXT NOT NULL,
    type_project TEXT,
    category TEXT,
    age_category_from INTEGER,
    age_category_to INTEGER,
    created_year INTEGER,
    description TEXT
);