create table if not exists maintainer
(
    handle  varchar(64)     not null    unique,
    name    varchar(128)    not null,
    email   varchar(320)    not null    unique,

    constraint pk_maintainer
        primary key (handle)
);

-- maintainer.email max length is defined by some standard I
-- found on the internet; correct me if I am wrong, please.
