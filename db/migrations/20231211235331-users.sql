
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
create table users(
    id bigint auto_increment,
    name varchar(50) not null,
    email varchar(100) not null,
    no_hp varchar(50) not null,
    password varchar(100) not null,
    primary key (id)
)engine = InnoDB;


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table if exists users;