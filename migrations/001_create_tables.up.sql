CREATE TYPE "tariff_type" AS ENUM (
  'fixed',
  'percent'
);

CREATE TYPE "staff_type" AS ENUM (
  'cashier',
  'shop_assistant'
);

CREATE TYPE "payment_type" AS ENUM (
  'card',
  'cash'
);

CREATE TYPE "status_type" AS ENUM (
  'success',
  'cancel'
);

CREATE TYPE "transaction_type" AS ENUM (
  'withdraw',
  'topup'
);

CREATE TYPE "source_type" AS ENUM (
  'sales',
  'bonus'
);

CREATE TABLE "branches" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "adress" varchar NOT NULL,
  "year" int NOT NULL,
  "founded_at" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tariffs" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "type" tariff_type NOT NULL,
  "amount_for_cash" NUMERIC(12, 2),
  "amount_for_card" NUMERIC(12, 2),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "staffs" (
  "id" uuid PRIMARY KEY,
  "name" varchar not null,
  "branch_id" uuid NOT NULL REFERENCES "branches"("id"),
  "tariff_id" uuid NOT NULL REFERENCES "tariffs"("id"),
  "staff_type" staff_type NOT NULL,
  "balance" NUMERIC(12, 2),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sales" (
  "id" uuid PRIMARY KEY,
  "client_name" varchar NOT NULL,
  "branch_id" uuid not null REFERENCES "branches"("id"),
  "shop_assistant_id" uuid REFERENCES "staffs"("id"),
  "cashier_id" uuid not null REFERENCES "staffs"("id"),
  "price" NUMERIC(12, 2),
  "payment_type" payment_type NOT NULL,
  "status" status_type NOT NULL DEFAULT 'success',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" uuid PRIMARY KEY,
  "type" transaction_type NOT NULL,
  "amount" NUMERIC(12, 2),
  "source_type" source_type NOT NULL,
  "text" varchar,
  "sale_id" uuid not null REFERENCES "sales"("id"),
  "staff_id" uuid not null REFERENCES "staffs"("id"),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);
