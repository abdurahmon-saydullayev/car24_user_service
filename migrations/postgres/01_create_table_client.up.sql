CREATE TABLE IF NOT EXISTS "client"(
    "id" UUID PRIMARY KEY,
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(13) NOT NULL,
    "driving_license_number" VARCHAR(50) NOT NULL,
    "passport_number" VARCHAR(50) NOT NULL,
    "photo" VARCHAR(50),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "driving_number_given_place" VARCHAR(50) NOT NULL,
    "driving_number_given_date" TIMESTAMP NOT NULL,
    "driving_number_expired" TIMESTAMP NOT NULL,
    -- 0 - not blocked, 1 - blocked
    "is_blocked" SMALLINT DEFAULT 0,
    "propiska" VARCHAR(50) NOT NULL,
    "additional_phone_numbers" VARCHAR,
    "passport_pinfl" VARCHAR(14) NOT NULL,,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
 
-- CREATE TABLE IF NOT EXISTS "comments" (
--     "id" UUID PRIMARY KEY,
--     "order_id" UUID NOT NULL,
--     "client_id" UUID NOT NULL,
--     "message" VARCHAR(50) NOT NULL,
--     "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     "updated_at" TIMESTAMP,
--     FOREIGN KEY ("client_id") REFERENCES "client" ("id")
-- );

-- CREATE TABLE IF NOT EXISTS "branch"(
--     "id" UUID PRIMARY KEY,
--     "name" VARCHAR(50) NOT NULL,
--     "phone_number" VARCHAR(50) NOT NULL
-- );

CREATE TABLE IF NOT EXISTS "time_zone"(
    "id" UUID PRIMARY KEY,
    "country_id" UUID,
    "name" jsonb,
    "short_name" VARCHAR(50),
    "gmt_offet" REAL NOT NULL
);