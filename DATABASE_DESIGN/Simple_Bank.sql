CREATE TABLE "accouts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "accout_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_accout_id" bigint NOT NULL,
  "to_accout_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "entries" ADD FOREIGN KEY ("accout_id") REFERENCES "accouts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_accout_id") REFERENCES "accouts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_accout_id") REFERENCES "accouts" ("id");

CREATE INDEX ON "accouts" ("owner");

CREATE INDEX ON "entries" ("accout_id");

CREATE INDEX ON "transfers" ("from_accout_id");

CREATE INDEX ON "transfers" ("to_accout_id");

CREATE INDEX ON "transfers" ("from_accout_id", "to_accout_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';
