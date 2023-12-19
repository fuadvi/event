
-- +migrate Up
    CREATE TABLE IF NOT EXISTS events(
        id bigint auto_increment,
        title varchar(100) not null,
        thumbnail varchar(255) not null,
        start_date datetime not null,
        end_date datetime not null,
        description text not null,
        max_join smallint not null,
        price double precision not null,
        is_publish boolean not null,
        location boolean not null, # true => online, false => offline
        primary key (id)
    )engine = InnoDB;

-- +migrate Down
DROP TABLE IF EXISTS events;
