-- name: create-locations-table
CREATE TABLE IF NOT EXISTS locations (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    parking_lots TEXT [] NOT NULL,
    address TEXT
);


-- name: get-locations
SELECT * FROM locations;