-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- Posts table
create table posts
(
    id          integer unsigned      not null auto_increment,
    created_at  datetime     not null default current_timestamp,
    updated_at  datetime     not null default current_timestamp on update current_timestamp,
    deleted_at  datetime              default null,
    user_id     integer unsigned      not null,
    title       varchar(255) not null,
    content     text         not null,
    image_url   varchar(512)          default null,
    is_public   bool         not null default true,
    view_count  integer      not null default 0,
    like_count  integer      not null default 0,
    comment_count integer    not null default 0,
    foreign key (user_id) references users (id) on delete cascade,
    primary key (id)
) engine=innodb default charset=utf8mb4;

-- Comments table
create table comments
(
    id         integer unsigned      not null auto_increment,
    created_at datetime     not null default current_timestamp,
    updated_at datetime     not null default current_timestamp on update current_timestamp,
    deleted_at datetime              default null,
    post_id    integer unsigned      not null,
    user_id    integer unsigned      not null,
    parent_id  integer unsigned               default null,
    content    text         not null,
    like_count integer      not null default 0,
    foreign key (post_id) references posts (id) on delete cascade,
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (parent_id) references comments (id) on delete cascade,
    primary key (id)
) engine=innodb default charset=utf8mb4;

-- Post likes table
create table post_likes
(
    id        integer unsigned      not null auto_increment,
    created_at datetime     not null default current_timestamp,
    updated_at datetime     not null default current_timestamp on update current_timestamp,
    deleted_at datetime              default null,
    post_id   integer unsigned      not null,
    user_id   integer unsigned      not null,
    foreign key (post_id) references posts (id) on delete cascade,
    foreign key (user_id) references users (id) on delete cascade,
    primary key (id),
    unique key unique_post_user_like (post_id, user_id)
) engine=innodb default charset=utf8mb4;

-- Comment likes table
create table comment_likes
(
    id         integer unsigned      not null auto_increment,
    created_at datetime     not null default current_timestamp,
    updated_at datetime     not null default current_timestamp on update current_timestamp,
    deleted_at datetime              default null,
    comment_id integer unsigned      not null,
    user_id    integer unsigned      not null,
    foreign key (comment_id) references comments (id) on delete cascade,
    foreign key (user_id) references users (id) on delete cascade,
    primary key (id),
    unique key unique_comment_user_like (comment_id, user_id)
) engine=innodb default charset=utf8mb4;

-- User follows table (for following/followers feature)
create table user_follows
(
    id           integer unsigned      not null auto_increment,
    created_at   datetime     not null default current_timestamp,
    updated_at   datetime     not null default current_timestamp on update current_timestamp,
    deleted_at   datetime              default null,
    follower_id  integer unsigned      not null,
    following_id integer unsigned      not null,
    foreign key (follower_id) references users (id) on delete cascade,
    foreign key (following_id) references users (id) on delete cascade,
    primary key (id),
    unique key unique_follow_relationship (follower_id, following_id)
) engine=innodb default charset=utf8mb4;

-- Create indexes for better performance
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

drop table user_follows;
drop table comment_likes;
drop table post_likes;
drop table comments;
drop table posts;