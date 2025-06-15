CREATE TABLE "users" (
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

CREATE TABLE "sessions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "user_guid" varchar NOT NULL,
  "access_token" text NOT NULL,
  "access_token_expired_at" timestamp NOT NULL,
  "refresh_token" text NOT NULL,
  "refresh_token_expired_at" timestamp NOT NULL,
  "ip_address" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "customers" (
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

CREATE TABLE "customer_tenor_limits" (
  "customer_guid" varchar NOT NULL,
  "tenor_months" int NOT NULL,
  "limit_amount" decimal(18,2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar,
  PRIMARY KEY ("customer_guid", "tenor_months")
);

CREATE TABLE "transactions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "customer_guid" varchar NOT NULL,
  "contract_no" varchar(50) UNIQUE NOT NULL,
  "otr" decimal(18,2) NOT NULL,
  "admin_fee" decimal(18,2) NOT NULL,
  "installment_amount" decimal(18,2) NOT NULL,
  "interest_amount" decimal(18,2) NOT NULL,
  "asset_name" varchar NOT NULL,
  "asset_type" varchar NOT NULL,
  "tenor_months" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar
);

ALTER TABLE "customer_tenor_limits" ADD FOREIGN KEY ("customer_guid") REFERENCES "customers" ("guid");

ALTER TABLE "transactions" ADD FOREIGN KEY ("customer_guid") REFERENCES "customers" ("guid");
