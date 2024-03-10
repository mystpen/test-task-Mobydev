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

INSERT INTO videos (title, type_project, category, created_year, description)
VALUES ('How to code', 'educational', 'programming', 2020, 'The film shows how to work as a programmer'),
    ('Popular memes', 'entertaining', 'comedy', 2023, 'The most popular memes of 2023'),
    ('Working with SQL', 'educational', 'programming', 2019, 'Basics of working with SQL queries');