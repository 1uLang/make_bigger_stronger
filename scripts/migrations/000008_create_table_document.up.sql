CREATE TABLE IF NOT EXISTS `document` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` char(100) NOT NULL default '' comment '文件名称',
    `desc` char(255) NOT NULL default '' comment '文件描述',
    `file_size` char(20) not null comment '文件大小',
    `url` char(255) NOT NULL default '' comment '访问链接',
    `created_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;