create index idx_posts_user_id on posts(user_id);
create index idx_replies_post_id on replies(post_id);
create index idx_replies_user_id on replies(replied_user_id);