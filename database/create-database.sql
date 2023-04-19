-- User & Database Deletion
drop database if exists fds;
drop schema if exists public cascade;
drop user if exists fds_user;

-- User & Database Creation
create user fds_user with password '1234';
create database fds;
grant all privileges on database fds to fds_user;
create schema public;

-- Create Tables
create table minecraft_skin
(
    id          serial primary key,
    skin_base64 text not null
);

create table mojang_user
(
    id                serial primary key,
    uuid              varchar(32) not null unique,
    name              varchar(16) not null,
    minecraft_skin_id int,
    foreign key (minecraft_skin_id) references minecraft_skin (id)
);

create table hypixel_player
(
    id                  serial primary key,
    mojang_user_id      int       not null,
    latest_player_stats json      not null,
    latest_lookup       timestamp not null,
    first_lookup        timestamp not null,
    tracking            boolean   not null,
    foreign key (mojang_user_id) references mojang_user (id)
);

create table "user"
(
    hypixel_player_id    int         not null unique,
    password             text        not null,
    discord              varchar(37) not null,
    registered_at        timestamp   not null,
    is_confirmed         boolean     not null,
    last_password_change timestamp   not null,
    api_key              varchar(36),
    foreign key (hypixel_player_id) references hypixel_player (id)
);

create table daily_hypixel_stats
(
    hypixel_player_id int not null unique,
    player_data       json,
    foreign key (hypixel_player_id) references hypixel_player (id)
);

create table discord_user
(
    hypixel_player_id   int       not null,
    level               int       not null,
    overflow_xp         int       not null,
    dailies_streak      int       not null,
    xp_from_dailies     int       not null,
    last_daily_claimed  timestamp not null,
    minutes_spent_in_vc int       not null,
    messages_sent       int       not null,
    foreign key (hypixel_player_id) references hypixel_player (id)
);

create table hypixel_game
(
    id       serial primary key,
    type     text    not null,
    verified boolean not null,
    special  boolean not null,
    data     json    not null
);

create table hypixel_game_player
(
    hypixel_player_id int not null,
    hypixel_game_id   int not null,
    primary key (hypixel_player_id, hypixel_game_id),
    foreign key (hypixel_player_id) references hypixel_player (id),
    foreign key (hypixel_game_id) references hypixel_game (id)
);