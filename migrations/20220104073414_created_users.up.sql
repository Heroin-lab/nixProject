CREATE TABLE users (
    id int not null primary key,
    email varchar(255) not null unique,
    password varchar(255) not null
);