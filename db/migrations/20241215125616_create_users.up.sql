CREATE TABLE users (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    firebase_uid VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
