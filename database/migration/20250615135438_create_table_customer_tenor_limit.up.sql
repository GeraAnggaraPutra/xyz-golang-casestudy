CREATE TABLE IF NOT EXISTS "customer_tenor_limits" (
    "customer_guid" varchar NOT NULL,
    "tenor_months" int NOT NULL,
    "limit_amount" decimal(18,2) NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "created_by" varchar NOT NULL DEFAULT 'system',
    "updated_at" timestamp,
    "updated_by" varchar,
    PRIMARY KEY ("customer_guid", "tenor_months"),
    FOREIGN KEY ("customer_guid") REFERENCES customers("guid") ON DELETE CASCADE
);

WITH CustomerGUIDs AS (
    SELECT guid, nik FROM customers WHERE nik IN ('3273010101900001', '3273020202920002')
)
INSERT INTO "customer_tenor_limits" (
    "customer_guid", "tenor_months", "limit_amount"
) VALUES
(
    (SELECT guid FROM CustomerGUIDs WHERE nik = '3273010101900001'),
    1,
    100000.00
),
(
    (SELECT guid FROM CustomerGUIDs WHERE nik = '3273010101900001'),
    2,
    200000.00
),
(
    (SELECT guid FROM CustomerGUIDs WHERE nik = '3273010101900001'),
    3,
    500000.00
),
(
    (SELECT guid FROM CustomerGUIDs WHERE nik = '3273010101900001'),
    4,
    700000.00
),
(
    (SELECT guid FROM CustomerGUIDs WHERE nik = '3273020202920002'),
    1,
    1000000.00
),
(
    (SELECT guid FROM CustomerGUIDs WHERE nik = '3273020202920002'),
    2,
    1200000.00
),
(
    (SELECT guid FROM CustomerGUIDs WHERE nik = '3273020202920002'),
    3,
    1500000.00
),
(
    (SELECT guid FROM CustomerGUIDs WHERE nik = '3273020202920002'),
    4,
    2000000.00
)
ON CONFLICT ("customer_guid", "tenor_months") DO NOTHING;
