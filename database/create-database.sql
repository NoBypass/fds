-- Setup & Creation
set @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0;
drop schema if exists fds;
create schema if not exists fds default character set utf8;
use fds;

-- Create Tables
create table fds.mojangUser
(
    id               int         not null auto_increment primary key,
    uuid             varchar(32) not null unique,
    name             varchar(16) not null,
    minecraftSkin_id int,
    foreign key (minecraftSkin_id) references fds.minecraftSkin (id)
);

create table fds.user
(
    hypixelPlayer_id   int         not null unique,
    password           text        not null,
    discord            varchar(37) not null,
    registeredAt       datetime    not null,
    isConfirmed        boolean     not null,
    lastPasswordChange datetime    not null,
    apiKey             varchar(36),
    foreign key (hypixelPlayer_id) references fds.hypixelPlayer (id)
);

create table fds.dailyHypixelStats
(
    hypixelPlayer_id int not null unique,
    playerData       json,
    foreign key (hypixelPlayer_id) references fds.hypixelPlayer (id)
);

create table fds.hypixelPlayer
(
    id                int      not null auto_increment primary key,
    mojangUser_id     int      not null,
    latestPlayerStats json     not null,
    latestLookup      datetime not null,
    firstLookup       datetime not null,
    tracking          boolean  not null,
    foreign key (mojangUser_id) references fds.mojangUser (id)
);

create table fds.discordUser
(
    hypixelPlayer_id int      not null,
    level            int      not null default 0,
    overflowXp       int      not null default 0,
    dailiesStreak    int      not null default 0,
    xpFromDailies    int      not null default 0,
    lastDailyClaimed datetime not null,
    minutesSpentInVc int      not null default 0,
    messagesSent     int      not null default 0,
    foreign key (hypixelPlayer_id) references fds.hypixelPlayer (id)
);

create table fds.hypixelGamePlayer
(
    hypixelPlayer_id int not null,
    hypixelGame_id   int not null,
    primary key (hypixelPlayer_id, hypixelGame_id),
    foreign key (hypixelPlayer_id) references fds.hypixelPlayer (mojangUser_id),
    foreign key (hypixelGame_id) references fds.hypixelGame (id)
);

create table fds.hypixelGame
(
    id       int     not null auto_increment primary key,
    type     text    not null,
    verified boolean not null,
    special  boolean not null,
    data     json    not null
);

create table fds.minecraftSkin
(
    id int not null auto_increment primary key,
    skinBase64 text not null
);