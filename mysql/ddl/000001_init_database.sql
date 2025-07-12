-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table account
(
    id           integer unsigned      not null auto_increment,
    created_at   datetime     not null default current_timestamp,
    updated_at   datetime     not null default current_timestamp on update current_timestamp,
    deleted_at   datetime              default null,
    password     varchar(256) not null,
    signed_up_at datetime     not null default current_timestamp,
    primary key (id)
) engine=innodb default charset=utf8mb4;

create table account_refresh_tokens
(
    id              integer unsigned      not null auto_increment,
    created_at      datetime     not null default current_timestamp,
    updated_at      datetime     not null default current_timestamp on update current_timestamp,
    deleted_at      datetime              default null,
    account_id      integer unsigned      not null,
    token           varchar(512) not null,
    access_token_id varchar(256) not null,
    foreign key (account_id) references account (id) on delete cascade,
    primary key (id)
) engine=innodb default charset=utf8mb4;

create table users
(
    `id`             integer unsigned      not null auto_increment,
    `created_at`     datetime     not null default current_timestamp,
    `updated_at`     datetime     not null default current_timestamp on update current_timestamp,
    `deleted_at`     datetime              default null,
    `account_id`     integer unsigned      not null,
    `email`          varchar(512) not null,
    `name`           varchar(512) not null,
    `is_banned`      bool         not null default false,
    `is_first_login` bool         not null,
    foreign key (`account_id`) references account (`id`) on delete cascade,
    primary key (id)

);

create table user_facebook_logins
(
    `id`         integer unsigned      not null auto_increment,
    `created_at` datetime     not null default current_timestamp,
    `updated_at` datetime     not null default current_timestamp on update current_timestamp,
    `deleted_at` datetime              default null,
    `user_id`    integer unsigned      not null,
    `fb_id`      varchar(512) not null,
    `email`      varchar(512) not null,
    `name`       varchar(512) not null,
    foreign key (`user_id`) references users (`id`) on delete cascade,
    primary key (id)

);

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


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table user_follows;
drop table comment_likes;
drop table post_likes;
drop table comments;
drop table posts;
drop table user_facebook_logins;
drop table users;
drop table account_refresh_tokens;
drop table accounts;
