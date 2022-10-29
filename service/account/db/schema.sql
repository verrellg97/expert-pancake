CREATE SCHEMA IF NOT EXISTS account;

create table account.users
(
    id           text not null
        primary key,
    fullname     text not null,
    nickname     text not null,
    email        text,
    phone_number text not null,
    created_at   timestamp with time zone default now(),
    updated_at   timestamp with time zone default now()
);

create table account.user_passwords
(
    user_id    text not null
        primary key,
    password   text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create table account.user_infos
(
    user_id    text not null,
    key        text not null,
    value      text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    primary key (user_id, key)
);