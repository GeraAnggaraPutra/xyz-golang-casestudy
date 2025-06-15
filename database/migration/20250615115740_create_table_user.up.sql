CREATE TABLE IF NOT EXISTS "users" (
    "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "created_by" varchar NOT NULL DEFAULT 'system',
    "updated_at" timestamp,
    "updated_by" varchar,
    "deleted_at" timestamp,
    "deleted_by" varchar
);

INSERT INTO users (email, password)
VALUES (
        'admin@gmail.com',
        '$2a$12$LQi1CpKB/dUNMKko2sHd/.umM9hdOYSoMRF7b8JbgiV3ZvSWIEqQC'
        ) ON CONFLICT DO NOTHING;