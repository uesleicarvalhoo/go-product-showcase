create table if not exists clients (
    "id" uuid not null,
    "name" varchar not null,
    "email" varchar not null,
    "phone" varchar not null,
    "address_id" uuid null,
    "zip_code" varchar null,
    "street" varchar null,
    "city" varchar null,
    constraint clients_pk primary key (id),
    constraint clients_uniq_email unique (email)
);

CREATE UNIQUE INDEX clients_id_idx ON "clients" (id);
CREATE UNIQUE INDEX clients_email_idx ON "clients" (email);

create table if not exists products (
    "id" uuid not null,
    "name" varchar not null,
    "description" varchar not null,
    "code" varchar not null,
    "price" float NOT null,
    "category" varchar,
    "image_url" varchar,
    constraint product_pk primary key (id)
);
