CREATE TABLE "users" (
  "id" varchar(64) PRIMARY KEY,
  "login_id" varchar(128) UNIQUE NOT NULL,
  "password" varchar(128) NOT NULL,
  "status" varchar(32) NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now())
);

CREATE TABLE "profiles" (
  "user_id" varchar NOT NULL,
  "first_name" varchar(64) NOT NULL,
  "last_name" varchar(64) NOT NULL,
  "dob" TIMESTAMP NOT NULL,
  "address_line1" varchar(64) NOT NULL,
  "address_line2" varchar(64) NOT NULL,
  "city" varchar(64) NOT NULL,
  "state" varchar(64) NOT NULL,
  "country" varchar(64) NOT NULL,
  "postal_code" varchar(64) NOT NULL,
  "primary_phone" VARCHAR(20) NOT NULL,
  "secondary_phone" VARCHAR(20) NOT NULL,
  "primary_email" VARCHAR(128) NOT NULL,
  "secondary_email" VARCHAR(128) NOT NULL,
  "created_at" TIMESTAMPZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPZ NOT NULL DEFAULT (now())
);

CREATE TABLE "permissions" (
  "id" VARCHAR(64) PRIMARY KEY,
  "name" VARCHAR(128) UNIQUE NOT NULL,
  "description" VARCHAR(128) NOT NULL,
  "resource_name" VARCHAR(128) NOT NULL,
  "allowed_read" BOOLEAN DEFAULT FALSE,
  "allowed_create" BOOLEAN DEFAULT FALSE,
  "allowed_update" BOOLEAN DEFAULT FALSE,
  "allowed_delete" BOOLEAN DEFAULT FALSE,
  "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE "roles" (
  "id" VARCHAR(64) PRIMARY KEY,
  "name" VARCHAR(128) UNIQUE NOT NULL,
  "description" VARCHAR(128) NOT NULL,
  "is_active" BOOLEAN DEFAULT FALSE,
  "permission_id" VARCHAR(64) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE "user_roles" (
  "user_id" VARCHAR NOT NULL,
  "role_id" VARCHAR NOT NULL,
  "created_at" TIMESTAMPZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPZ NOT NULL DEFAULT (now())
);

ALTER TABLE "profiles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "roles" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");

ALTER TABLE "user_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
