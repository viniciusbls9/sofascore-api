CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    fav_position VARCHAR(100),
    biography TEXT,
    image_url VARCHAR(255),
    age VARCHAR(3),
    height VARCHAR(5),
    preferred_foot VARCHAR(10),
    shirt_number VARCHAR(3),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
