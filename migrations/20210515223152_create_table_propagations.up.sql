create table if not exists propagations
(
    id      serial primary key not null,
    node    varchar(256)       not null,
    message varchar(512)       not null,
    unique (node, message)
);
