create table ozline.`user`
(
    `id`               bigint auto_increment not null,
    `username`         varchar(255)                                                                 not null,
    `created_at`       timestamp    default current_timestamp                                       not null,
    `updated_at`       timestamp    default current_timestamp                                       not null on update current_timestamp,
    `deleted_at`       timestamp    default null null,
    constraint `id`
        primary key (`id`)
) engine=InnoDB auto_increment=10000 default charset=utf8mb4;

INSERT INTO west2online.user (username, created_at, updated_at, deleted_at)
VALUES ('ozline', DEFAULT, DEFAULT, null);

INSERT INTO west2online.user (username, created_at, updated_at, deleted_at)
VALUES ('ozline1', DEFAULT, DEFAULT, null);