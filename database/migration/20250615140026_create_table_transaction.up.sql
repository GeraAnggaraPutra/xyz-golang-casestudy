CREATE TABLE IF NOT EXISTS "transactions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "customer_guid" varchar NOT NULL,
  "contract_no" varchar(50) UNIQUE NOT NULL,
  "otr" decimal(18,2) NOT NULL,
  "admin_fee" decimal(18,2) NOT NULL,
  "installment_amount" decimal(18,2) NOT NULL,
  "interest_amount" decimal(18,2) NOT NULL,
  "asset_name" varchar NOT NULL,
  "asset_type" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar
);

INSERT INTO "transactions" (
    "customer_guid", "contract_no", "otr", "admin_fee", "installment_amount", "interest_amount", "asset_name", "asset_type"
) VALUES (
    (SELECT guid FROM "customers" WHERE nik = '3273010101900001'),
    'KP-WG-20250615-0001',   
    90000.00,              
    5000.00,              
    100000.00,            
    5000.00,            
    'Smart Refrigerator XYZ', 
    'White Goods'
),(
    (SELECT guid FROM "customers" WHERE nik = '3273020202920002'),
    'KP-MOT-20250620-0002',
    1000000.00,
    10000.00,
    1200000.00,
    50000.00,
    'Motorcycle Honda Beat',
    'Motor'
)
ON CONFLICT ("guid") DO NOTHING;