-- name: create-locations-table
CREATE TABLE IF NOT EXISTS locations (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    parking_lots TEXT [] NOT NULL,
    address TEXT
);


-- name: get-locations
SELECT id, name, parking_lots, address FROM locations;

--name: get-location-by-id
SELECT * FROM LOCATIONS WHERE ID = ? ;