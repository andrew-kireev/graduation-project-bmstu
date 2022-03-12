CREATE TABLE IF NOT EXISTS Profiles
(
    profiles_id        bigserial not null primary key,
    email              varchar   not null unique,
    nickname           varchar   not null unique,
    first_name         varchar   not null,
    second_name        varchar   not null,
    encrypted_password varchar   not null,
    avatar             varchar   not null,
    favorite_genre     text[]    NOT NULL DEFAULT '{}'::text[]
);

CREATE TABLE IF NOT EXISTS admin_profiles
(
    profile_id         bigserial not null primary key,
    login              varchar   not null unique,
    encrypted_password varchar   not null
);


create index profiles_nickname on profiles using btree (nickname);
create index admin_profiles_login on admin_profiles using btree (login);