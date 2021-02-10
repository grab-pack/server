create table if not exists package
(
    name        varchar(128)    not null    unique,
    maintainer  varchar(64)     not null,
    repository  varchar(256)                unique,

    constraint pk_package
        primary key (name),
    constraint fk_maintainer
        foreign key (maintainer)
            references maintainer (handle)
);

-- package.repository is nullable since some packages may not
-- be open-source and will not be able to provide a link to
-- the public repository.
