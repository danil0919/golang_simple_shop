CREATE TABLE IF NOT EXISTS products(
    id bigserial not null primary key,
    created timestamp null,
    name varchar null,
    price decimal null,
    description text null
);