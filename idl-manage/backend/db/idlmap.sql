CREATE TABLE "idlmap" (
  "id" bigserial PRIMARY KEY,
  "svcname" varchar NOT NULL,
  "idl" varchar NOT NULL,
  "type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "idlmap" ("svcname");
