create table if not exists replies(
    id serial primary key,
    post_id int references posts(id) on delete cascade not null,
    replied_user_id int references users(id) on delete cascade not null,
    reply varchar not null,
    created_at timestamp default now()
);