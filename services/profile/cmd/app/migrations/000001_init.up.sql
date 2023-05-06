CREATE TABLE "user" (
    "id" uuid primary key,
    "login" varchar unique not null,
    "password" varchar not null,
    "name" varchar not null,
    "middle_name" varchar not null,
    "surname" varchar not null,
    "phone" varchar not null,
    "city" varchar not null,
    "role" varchar not null,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
)