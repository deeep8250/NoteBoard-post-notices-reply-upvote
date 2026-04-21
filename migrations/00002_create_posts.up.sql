create table if not exists posts(
    id serial primary key,
    user_id int references users(id) on delete cascade not null,
    title varchar(250) not null,
    content text not null,
    created_at timestamp default now()
);





