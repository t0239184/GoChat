CREATE TABLE user (
    id bigint(20) not null auto_increment primary key,
    account varchar(50) not null,
    password varchar(100) not null,
    salt_id bigint(20) not null,
    status varchar(1) not null default '0'
);

CREATE UNIQUE INDEX idx_user_account ON user(account);


CREATE TABLE salt (
    id bigint(20) not null auto_increment primary key,
    salt varchar(100) not null,
    iteration smallint(4) not null
)