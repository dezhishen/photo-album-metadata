create table
    `user`(
        id varchar(32) not null PRIMARY KEY,
        username VARCHAR(40) not null,
        `password` VARCHAR(128) not null
    );