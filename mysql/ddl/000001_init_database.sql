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

create table account_refresh_token
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
    `id`         integer unsigned      not null auto_increment,
    `created_at` datetime     not null default current_timestamp,
    `updated_at` datetime     not null default current_timestamp on update current_timestamp,
    `deleted_at` datetime              default null,
    `account_id` integer unsigned      not null ,
    `email`      varchar(512) not null,
    `name`       varchar(512) not null,
    `is_banned`  bool         not null default false,
    foreign key (`account_id`) references account (`id`) on delete cascade,
    primary key (id)

);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

drop table users;
drop table account_refresh_token;
drop table account;
