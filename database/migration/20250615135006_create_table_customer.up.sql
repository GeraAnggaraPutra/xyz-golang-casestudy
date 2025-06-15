CREATE TABLE IF NOT EXISTS "customers" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "nik" varchar(16) NOT NULL,
  "full_name" varchar(50) NOT NULL,
  "legal_name" varchar(50) NOT NULL,
  "birth_place" varchar NOT NULL,
  "birth_date" date NOT NULL,
  "salary" float NOT NULL,
  "photo_ktp" varchar NOT NULL,
  "photo_selfie" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar
);

INSERT INTO "customers" (
    "nik", "full_name", "legal_name", "birth_place", "birth_date", "salary", "photo_ktp", "photo_selfie"
) VALUES (
    '3273010101900001',
    'Budi Santoso',
    'Budi Santoso',
    'Jakarta',
    '1990-01-15',
    7500000.00,
    'image/2025-03-27/1743055334_ktp_budi.jpg',
    'image/2025-03-27/1743055334_selfie_budi.jpg'
),
(
    '3273020202920002',  
    'Annisa Putri',
    'Annisa Putri',
    'Bandung',
    '1992-05-20',
    12000000.00,      
    'image/2025-03-27/1743055334_ktp_annisa.jpg',
    'image/2025-03-27/1743055334_selfie_annisa.jpg'
) ON CONFLICT DO NOTHING;
