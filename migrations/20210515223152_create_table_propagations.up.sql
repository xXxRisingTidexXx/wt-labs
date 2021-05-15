create table if not exists propagations
(
    id   serial primary key not null,
    node varchar(256)       not null,
    ip   inet               not null,
    unique (node, ip)
);
