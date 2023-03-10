create table
    `user`(
        id varchar(32) not null PRIMARY KEY,
        username VARCHAR(40) not null
    );
CREATE table 
    `account`(
        id varchar(32) not null PRIMARY KEY,
        user_id varchar(32) not null,
        `name` VARCHAR(40) not null,
        `password` VARCHAR(32) not null
    );
CREATE table 
    `role`(
        id varchar(32) not null PRIMARY KEY,
        `code` VARCHAR(40) not null, 
        `name` VARCHAR(40) not null
    );
CREATE table
    `user_role`(
        user_id varchar(32) not null,
        role_id varchar(32) not null
    );

-- create default user and admin role
INSERT INTO `user` (id, username) VALUES ('1', 'admin');
INSERT INTO `role` (id, `code`, `name`) VALUES ('1', 'admin', 'admin');
INSERT INTO `user_role` (user_id, role_id) VALUES ('1', '1');
-- password save as md5({{.name}}{{.password}})
INSERT INTO `account` (id, user_id, `name`, `password`) VALUES ('1', '1', 'admin', '21232f297a57a5a743894a0e4a801fc3');