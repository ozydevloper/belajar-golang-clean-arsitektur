CREATE SEQUENCE IF NOT EXISTS users_id_seq
AS INT
INCREMENT BY 1
START 100
MINVALUE 100;

CREATE TABLE IF NOT EXISTS users (
    id int not null primary key default nextval('users_id_seq'),
    email varchar(100) not null unique,
    password varchar(100) not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

ALTER SEQUENCE users_id_seq OWNED BY users.id;
