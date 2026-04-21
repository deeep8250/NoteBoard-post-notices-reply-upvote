create table if not exists upvotes(
    id serial primary key,
    post_id int references posts(id) on delete cascade not null,
    user_id int references users(id) on delete cascade not null,
    created_at timestamp default now(),
    unique(post_id,user_id)
);