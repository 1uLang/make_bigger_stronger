CREATE TABLE IF NOT EXISTS `application_his` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `uid` varchar(36) NOT NULL comment '用户uuid',
    `user_name` varchar(36) NOT NULL comment '用户名',
    `app_id` int(10) NOT NULL comment '访问应用ID',
    `app_name` varchar(50) NOT NULL default '' COMMENT '应用名称',
    `access_url` varchar(2000) NOT NULL comment '访问地址',
    `created_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `uid` (`uid`) USING BTREE,
    KEY `app_id` (`app_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;