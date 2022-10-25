CREATE TABLE accounts (
  id VARCHAR(36) PRIMARY KEY NOT NULL,
  name VARCHAR NOT NULL,
  document VARCHAR NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE transactions (
  id VARCHAR(36) PRIMARY KEY NOT NULL,
  account_id VARCHAR NOT NULL,
  amount float8 NOT NULL,
  description VARCHAR NOT NULL,
  payment_method VARCHAR NOT NULL,
  card_number VARCHAR(4) NOT NULL,
  card_owner VARCHAR NOT NULL,
  card_expiration_date VARCHAR NOT NULL,
  card_cvv VARCHAR NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");


CREATE TABLE payables (
  id VARCHAR(36) PRIMARY KEY NOT NULL,
  transaction_id VARCHAR NOT NULL,
  account_id VARCHAR NOT NULL,
  status VARCHAR NOT NULL,
  fee VARCHAR NOT NULL,
  payment_date DATE NOT NULL,
  amount_paid float8 NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "payables" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");
ALTER TABLE "payables" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
