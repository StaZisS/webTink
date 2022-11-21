CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "users" (
                         "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
                         "name" VARCHAR NOT NULL,
                         "surname" VARCHAR NOT NULL,
                         "username" VARCHAR NOT NULL,
                         "email" VARCHAR NOT NULL,
                         "grade" INTEGER NOT NULL,
                         "photo" VARCHAR NOT NULL DEFAULT 'default.jpg',
                         "verified" BOOLEAN NOT NULL DEFAULT FALSE,
                         "password" VARCHAR NOT NULL,
                         "role" VARCHAR NOT NULL DEFAULT 'user',
                         "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         "updated_at" TIMESTAMP(3) NOT NULL,

                         CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
CREATE UNIQUE INDEX "users_username_key" ON "users" ("username");
CREATE TABLE "posts"(
                            "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
                            "title" VARCHAR NOT NULL,
                            "content" VARCHAR NOT NULL,
                            "author_id" UUID NOT NULL,
                            "photo" VARCHAR NOT NULL,
                            "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            "updated_at" TIMESTAMP(3) NOT NULL,

                            CONSTRAINT "posts_pkey" PRIMARY KEY ("id"),
                            CONSTRAINT "posts_author_id_fkey" FOREIGN KEY ("author_id") REFERENCES "users" ("id") ON DELETE CASCADE
);