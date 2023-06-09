INSERT IGNORE INTO `role` (`id`, `name`, `tag`, `remark`, `perms`, `is_super`) VALUES (17, '超级管理员', 'superadmin', '', '[]', 1);

INSERT IGNORE INTO `admin` (`uid`, `phone`, `password`, `account`, `avatar`, `is_disable`, `emil`, `we_chat_account`, `sex`, `post`, `name`, `role_tag`, `created_time`, `updated_time`, `deleted_time`) VALUES ('a37d9c17-bd20-4d40-97d6-6e88895385ac', '', 'a924c2f18a5d42429bbc1f14358506cf', 'admin', '', 1, '', '', 1, '0', '超级管理员', 'superadmin', 1660291670, 1660291670, 0);
