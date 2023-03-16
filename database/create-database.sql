-- Setup & Creation
set @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0;
drop schema if exists fds;
create schema if not exists fds default character set utf8;
use fds;

-- Create Tables
create table fds.mojangUser
(
    uuid varchar(32) not null unique,
    name varchar(16) not null,
    skin text        not null,
    primary key (uuid)
);

create table fds.user
(
    mojangUser_uuid varchar(32) not null unique,
    password        text        not null,
    discord         varchar(37) not null,
    registeredAt    datetime    not null,
    isConfirmed     boolean     not null,
    foreign key (mojangUser_uuid) references fds.mojangUser (uuid)
);

create table fds.dailyHypixelStats
(
    mojangUser_uuid varchar(32) not null unique,
    playerData      json,
    foreign key (mojangUser_uuid) references fds.mojangUser (uuid)
);

create table fds.hypixelPlayer
(
    mojangUser_uuid   varchar(32) not null unique,
    latestPlayerStats json        not null,
    latestLookup      datetime    not null,
    firstLookup       datetime    not null,
    tracking          boolean     not null,
    foreign key (mojangUser_uuid) references fds.mojangUser (uuid)
);

create table fds.discordUser
(
    mojangUser_uuid  varchar(32) not null,
    level            int         not null default 0,
    overflowXp       int         not null default 0,
    dailiesStreak    int         not null default 0,
    lastDailyClaimed datetime    not null,
    minutesSpentInVc int         not null default 0,
    messagesSent     int         not null default 0,
    foreign key (mojangUser_uuid) references fds.mojangUser (uuid)
);

create table fds.hypixelGamePlayer
(
    hypixelPlayer_uuid varchar(32) not null,
    hypixelGame_id     int         not null,
    primary key (hypixelPlayer_uuid, hypixelGame_id),
    foreign key (hypixelPlayer_uuid) references fds.hypixelPlayer (mojangUser_uuid),
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