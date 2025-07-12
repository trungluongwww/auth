-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create index account_refresh_token_idx_token on account_refresh_tokens (token);
create index user_idx_email on users (email);
create index idx_posts_user_id on posts (user_id);
create index idx_posts_created_at on posts (created_at);
create index idx_posts_is_public on posts (is_public);
create index idx_comments_post_id on comments (post_id);
create index idx_comments_user_id on comments (user_id);
create index idx_comments_parent_id on comments (parent_id);
create index idx_post_likes_post_id on post_likes (post_id);
create index idx_post_likes_user_id on post_likes (user_id);
create index idx_comment_likes_comment_id on comment_likes (comment_id);
create index idx_comment_likes_user_id on comment_likes (user_id);
create index idx_user_follows_follower_id on user_follows (follower_id);
create index idx_user_follows_following_id on user_follows (following_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop index idx_user_follows_following_id on user_follows;
drop index idx_user_follows_follower_id on user_follows;
drop index idx_comment_likes_user_id on comment_likes;
drop index idx_comment_likes_comment_id on comment_likes;
drop index idx_post_likes_user_id on post_likes;
drop index idx_post_likes_post_id on post_likes;
drop index idx_comments_parent_id on comments;
drop index idx_comments_user_id on comments;
drop index idx_comments_post_id on comments;
drop index idx_posts_is_public on posts;
drop index idx_posts_created_at on posts;
drop index idx_posts_user_id on posts;
drop index user_idx_email on users;
drop index account_refresh_token_idx_token on account_refresh_tokens;