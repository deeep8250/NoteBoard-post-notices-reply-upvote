create table if not exists users(
    id serial primary key,
    name varchar(200) not null,
    email varchar(200) unique not null,
    hashed_pass varchar not null,
    created_at timestamp default now()
);