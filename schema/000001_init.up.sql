CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS cars (
    id SERIAL PRIMARY KEY,
    regNum VARCHAR(20) NOT NULL,
    mark VARCHAR(100),
    model VARCHAR(100),
    year INTEGER,
    owner_id INTEGER REFERENCES people(id)
);

