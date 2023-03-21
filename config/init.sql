drop table if exists `household`;
drop table if exists `record`;
drop table if exists `room`;
drop table if exists `user`;


create table if not exists household
(
    id         bigint auto_increment comment 'auto increment id, primary key, unique id each elder'
    primary key,
    user_id    bigint               not null comment 'owner of the house',
    room_id    bigint               not null comment 'id of the room where the elder lives in',
    age        int                  not null comment 'age of the elder',
    height     int                  not null comment 'height of the elder',
    wheelchair tinyint(1) default 0 null comment 'true if the elder uses a wheelchair, false by default
',
    deleted    tinyint(1) default 0 not null comment 'true if the elder is deleted, false by default'
    )
    comment 'information about the elders';

create table if not exists record
(
    id        bigint auto_increment comment 'auto increment id, primary key, unique id for each record'
    primary key,
    user_id   bigint                         not null comment 'owner of the room which the record is related to',
    room_id   bigint                         not null comment 'which room this record is related to',
    risk_type enum ('low', 'medium', 'high') not null comment 'how important this record is',
    title     varchar(255)                   null comment 'title of the record',
    content   text                           null comment 'content of the record',
    deleted   tinyint(1) default 0           not null comment 'if this record is deleted, false by default'
    );

create table if not exists room
(
    id        bigint auto_increment comment 'auto increment id, primary key, unique id for each room'
    primary key,
    room_name varchar(255)         not null comment 'name of the room',
    user_id   bigint               null comment 'owner of this room',
    city      varchar(255)         null comment 'which city is the room in',
    deleted   tinyint(1) default 0 not null comment 'true if the room is deleted, false by default.'
    );

create table if not exists `user`
(
    id        bigint auto_increment comment 'auto increment id, primary key, unique id for each user
'
    primary key,
    user_name varchar(255) not null comment 'email of the user'
    );

