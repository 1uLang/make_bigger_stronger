CREATE TABLE IF NOT EXISTS `authorization` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `key` varchar(2048) not null comment '续订密钥',
    `created_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;