CREATE TABLE IF NOT EXISTS `admin` (
      `uid` varchar(36) NOT NULL,
      `phone` varchar(11) not null default '' comment '用户手机号',
      `password` varchar(36) not null default '' comment '用户密码',
      `account` varchar(36) not null default '' comment '账户',
      `avatar` varchar(255) not null default '' comment '头像路径',
      `is_disable` int(5)  not null default 1 comment '是否禁用 1正常 2 禁用',
      `emil` varchar(36) not null default '' comment '邮箱',
      `we_chat_account` varchar(36) not null default '' comment '微信号',
      `sex` int(5)  not null default 0 comment '0全部 1.男。2女 3保密',
      `post` varchar(36) not null default '' comment '职位',
      `name` varchar(36) not null default '' comment '姓名',
      `role_tag` char(20) not null default '' comment '角色标签',
      `created_time` int(11)  not null default 0,
      `updated_time` int(11)  not null default 0,
      `deleted_time` int(11)  not null default 0,
      PRIMARY KEY (`uid`) USING BTREE ,
      unique key `number_unique` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `application` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` char(50) not null  comment '应用名称',
    `sso_protocol` tinyint(1) not null comment '1 SAML 2 OAUTH2 3 表单代填',
    `desc` char(50) not null default '' comment '应用描述',
    `state` tinyint(1) not null default 0 comment '禁用状态 2 禁用 1 启用',
    `level_id` int(10) not null comment '信用等级',
    `source_ip` varchar(255) not null default '' comment '源站IP地址',
    `icon` varchar(255) not null default '' comment 'icon地址',
	`address` varchar(2048) not null default '' comment '应用地址',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `name_unique` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `application_account` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_id` int(10) not null comment '应用ID',
    `app_account` char(50) not null comment '应用账户',
    `app_account_pass` varchar(255) not null comment '应用账户密码',
    `zt_accounts` json not null comment '零信任平台账户 json列表',
    `desc` char(50) not null default '' comment '应用描述',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `appid_account_unique` (`app_id`,`app_account`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `application_grant_relation` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_id` int(10) not null comment '应用ID',
    `org` json not null comment '组织架构ID json',
    `user` json not null comment '用户ID json',
    `group` json not null comment '用户组ID json',
    `desc` char(50) not null comment '描述',
    `type` tinyint(1) not null comment '应用授权类型 1组织架构 2用户组 3用户',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `credit_level` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` char(10)  NOT NULL comment '信用等级名称',
    `score` tinyint(3) NOT NULL comment '最低信用分数',
    `desc` char(50) not null default '' comment '描述',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `name_unique` (`name`) USING BTREE,
    UNIQUE KEY `score_unique` (`score`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `device` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` char(50) not null default '' comment '设备名称',
    `os` char(64) not null default '' comment '设备操作系统',
    `version` char(15)  not null default '0.0.0.0' comment '版本信息',
    `ip` char(32) not null default '' comment 'ip地址',
    `sn` char(36) not null comment '设备SN 特征码',
    `device_id` char(36) not null comment '设备ID',
    `state` tinyint(1) not null default 0 comment '设备禁用状态 1 启用 2 禁用 ',
    `type` tinyint(1) not null default 1 comment '客户端种类 1 client 2 ...', 
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    `deleted_time` int(11)  not null default 0 ,
    `mac` char(128) NOT NULL DEFAULT '' comment '设备MAC地址',
    PRIMARY KEY (`id`) USING BTREE,
    unique key `sn_unique` (`sn`),
    unique key `device_id_unique` (`device_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `device_his` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `device_id` int(10) unsigned  NOT NULL comment '设备主键ID',
    `uid` char(36) NOT NULL comment '用户ID',
    `user_name` char(36) not null default '' comment '用户姓名',
    `user_account` char(20) not null default '' comment '用户账号',
    `phone` char(20) not null default '' comment '用户手机号',
    `emil` char(30) not null default '' comment '用户邮箱',
    `last_login_time` int(11) not null default 0 comment '最后一次登陆时间',
	`last_ip` char(50) not null default '' comment '上次登录ip',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `gateway` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` char(50) not null default '' comment '网关名称',
    `version` char(15)  not null default '0.0.0.0' comment '版本号',
    `desc` char(50) not null default '' comment '网关描述信息',
    `sn` char(50) not null comment '网关特征码 如 099bbc254624d67b0cbd5c88e0842ae5',
    `state` tinyint(1) not null default 2 comment '网关状态 1 上线 2 下线',
    `is_online` tinyint(1) not null default 2 comment '是否在线 1在线 2离线',
    `addr` char(50) not null comment '网关地址 如 https://154.210.13.174:443',
    `ip` char(15) not null default '' comment '网关ip地址',
    `port` char(5) not null comment '端口',
    `info` char(50) not null comment '网关信息 如192.168.1.163/24/fa:16:3e:5f:8e:b2',
    `hw_info` char(50) not null comment '网关硬件信息 如 处理器16核, 内存31G, 磁盘2088G',
    `last_online_time` int(11) not null default 0 comment '最近上线时间',
    -- `sdp_ip` char(20) not null default '' comment '网关所属控制器IP',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `sn_unique` (`sn`),
    UNIQUE KEY `name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `gateway_grant_relation` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `gateway_set_id` int(10) not null comment '网关组ID',
    `org` json not null comment '组织架构ID json',
    `user` json not null comment '用户ID json',
    `group` json not null comment '用户组ID json',
    `desc` char(50) not null comment '描述',
    `type` tinyint(1) not null comment '网关组授权类型 1组织架构 2用户组 3用户',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `gateway_set` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` char(20) not null default '' comment '网关组名称',
    `count` int(10) not null default 0 comment '网关数量',
    `desc` char(50) not null default '' comment '网关组描述信息',
    `gateway` json not null comment '网关ID json列表',
    -- `sdp_ip` char(15) not null default '' comment '所属控制器IP',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `log` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS `organization` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` char(20) NOT NULL DEFAULT '' COMMENT '组织名称：必填项，文本输入框，字符限制20；名称不能重复，需要判断。',
    `desc` char(30) NOT NULL DEFAULT '' COMMENT '描述：非必填，文本域，限制字符30',
    `pid` int(11) NOT NULL DEFAULT 0 COMMENT '父ID 默认0 为根结点的子节点',
    `source` tinyint(1) not null default 0 comment '账户来源 1自建 2LDAP/AD',
	`source_id` int(11) not null default 0 comment '账号源 0 自建 >0 具体的source',
	`entry_uuid` varchar(100) not null default '' comment '源同步来的uuid',
    `created_time` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间戳',
    `updated_time` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间戳',
	UNIQUE KEY `pid_name_unique` (`name`,`pid`) USING BTREE,
	-- unique key `source_id_entry_uuid_unique` (`source_id`, `entry_uuid`) USING BTREE,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS `permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '父级ID',
  `url` json not null COMMENT 'api路径,用以服务器校验权限',
  `menu_name` varchar(50) NOT NULL DEFAULT '' COMMENT '菜单名',
  `menu_type` char(1) NOT NULL DEFAULT '' COMMENT '菜单类型(M:目录，C:菜单，F:按钮)',
  `route_name` char(32) NOT NULL DEFAULT '' COMMENT '路由别名 唯一，不能重复',
  `route_path` varchar(120) NOT NULL DEFAULT '' COMMENT '路由地址',
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '图标链接',
  `component` char(120) NOT NULL DEFAULT '' COMMENT '组件',
  `perms` char(120) NOT NULL DEFAULT '' COMMENT '权限标识',
  `no_cache` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否不缓存当前页面，0：否  1：是',
  `hidden` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否隐藏当前路由，0：否  1：是',
  `children_name` varchar(255) NOT NULL DEFAULT '' COMMENT '子集组件名称，用于页面嵌套页面时左侧导航栏显示(以‘;’分割)\n',
  `parent_name` varchar(255) NOT NULL DEFAULT '' COMMENT '归属的父页面名称',
  `sort` int(10) NOT NULL DEFAULT 0 COMMENT '排序值 ',
  `is_allot` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否可以分配，0：否    1：是',
  `no_column` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否隐藏侧边栏, 0否 1是',
  `level_hidden` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否显示在菜单中显示隐藏路由, 0否1是, 默认1',
  `is_except` tinyint(1) NOT NULL DEFAULT 0 COMMENT '业务超管 是否排除该权限 0否 1是',
  PRIMARY KEY (`id`) USING BTREE
--   UNIQUE KEY `idx_url` (`url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



CREATE TABLE IF NOT EXISTS `policy` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `type` tinyint unsigned NOT NULL COMMENT '策略类型：0/1 认证/终端策略',
  `subtype` tinyint NOT NULL COMMENT '策略子类型：0/1/2/3/ 认证策略下的账号/密码/会话等策略',
  `nodetype` tinyint NOT NULL COMMENT '子策略下的成员策略：1/2/3认证策略下的账号策略下的长时间未登录/强二次验证等。',
  `config` json NOT NULL COMMENT '成员策略具体配置',
  `action` tinyint NOT NULL COMMENT '策略触发动作：0/1/2 二次认证/账号锁定等',
  `enable` tinyint(1) NOT NULL COMMENT '策略状态：是否启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `utype` (`type`,`subtype`,`nodetype`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `resource` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` char(20) not null comment '资源名称',
    `pattern` char(50) not null comment '资源标识',
    `desc` char(50) not null default '' comment '描述',
    `app_id` int(10) not null comment '应用ID',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `appid_name_unique` (`app_id`,`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `resource_grant_relation` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `res_id` int(10) not null comment 'resource_id',
    `org` json not null comment '组织架构ID json',
    `user` json not null comment '用户ID json',
    `group` json not null comment '用户组ID json',
    `desc` char(50) not null comment '描述',
    `type` tinyint(1) not null comment '资源组授权类型 1组织架构 2用户组 3用户',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS `role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL DEFAULT '' COMMENT '名称',
  `tag` char(20) NOT NULL DEFAULT '' COMMENT '标签，用于权限匹配，该值不可重复',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `perms` text COMMENT '权限列表',
  `is_super` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是超级管理员角色, 0:普通用户 1:系统超管 2:业务超管',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE,
  UNIQUE KEY `idx_tag` (`tag`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `rule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) NOT NULL COMMENT '规则名称',
  `type` tinyint NOT NULL COMMENT '对应策略类型。0/1 认证/终端策略',
  `policy` tinyint NOT NULL COMMENT '实际触发策略。0/1/2/3 账号/密码/会话等策略',
  `score` tinyint unsigned NOT NULL COMMENT '触发规则扣除分数',
  `enable` tinyint(1) NOT NULL COMMENT '启用状态',
  `created_time` int(11) NOT NULL DEFAULT 0 COMMENT '触发创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`name`),
  UNIQUE KEY `upolicy` (`policy`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `rule_history` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uid` int unsigned NOT NULL COMMENT '用户ID',
  `now_score` tinyint unsigned NOT NULL COMMENT '当前用户信誉分数\n',
  `sub_score` tinyint NOT NULL COMMENT '当前用户总扣除分数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `rule_history_desc` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `hid` int unsigned NOT NULL COMMENT '扣分历史记录ID',
  `type` tinyint(4) NOT NULL COMMENT '策略类型',
  `policy` tinyint unsigned NOT NULL COMMENT '触发策略',
  `sub_score` tinyint unsigned DEFAULT NULL COMMENT '此次所扣分数',
  `ip` varchar(32) NOT NULL COMMENT '客户端IP',
  `created_time` datetime NOT NULL COMMENT '触发时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `source` (
    `id` int(10) unsigned not null AUTO_INCREMENT,
    `name`  char(20) not null comment '账号源名称',
    `desc`  char(50) not null comment '账号源描述',
    `type`  tinyint(1) not null comment '账号源类型 1AD 2LDAP 3企业微信 4钉钉',
    `state` tinyint(1) not null default 0 comment '账号源状态  0禁用 1启用',
    `sync_pass` varchar(255) not null comment '同步密码,base64',
    `account_field` char(64) not null comment '账号字段',
    `username_field` char (64) not null comment '用户姓名字段',
    `phone_field` char (64) not null comment '手机号码字段',
    `email_field`char (64) not null comment '电子邮箱字段',
    `org_id` int(10) not null comment '组织ID', 
    `config` json not null comment '源配置',
    `created_time` int(11) not null default 0 comment '创建时间',
    `updated_time` int(11) not null default 0 comment '更新时间',
    `deleted_time` int(11) not null default 0 comment '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `sso_form` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_id` int(10) not null comment '应用 ID',
    `login_url` varchar(500) not null comment '登录地址',
    `username_field` char(20) not null comment '登录名属性字段名称',
    `pass_field` char(20) not null comment '登录密码字段名称',
    `method` tinyint(1) not null comment '1POST 2GET',
    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `app_id_unique` (`app_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `sso_oauth` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_id` int(10) not null comment '应用 ID',
    
    `client_id` varchar(100) not null comment '客户端身份标识',
    `client_secret` varchar(100) not null comment '客户端密钥',
    `redirect_url` varchar(500) not null comment 'OAuth2 Redirect URI,请以http或https开头',
    `sp_home_url` varchar(100) not null comment '应用首页地址,支持手动发起SSO。',
    `grant_type` tinyint(1) not null comment '授权类型 1:implict 简化模式(在redirect_uri的hash传递token) 2:authorization_code 授权码模式(即先登录获取code,再获取token)',
    `token_expire_time` int(10) not null default 7200 comment 'Access_token的有效时长(单位:秒),默认为7200(2小时)',
    `token_refresh_time` int(10) not null default 604800 comment 'Refresh_token的有效时长(单位:秒),默认为604800(7天)',

    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `app_id_unique` (`app_id`) USING BTREE,
	UNIQUE KEY `client_id_unique` (`client_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `sso_saml` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_id` int(10) not null comment '应用 ID',
    
    `ldp_entity_id` varchar(100) not null comment '用于标识SAML IdP的信息',
    `sp_entity_id` varchar(100) not null comment '用于标识SAML SP的信息',
    `sp_acs_url` varchar(500) not null comment 'SAML SSO Location',
    `sp_logout_url` varchar(500) not null default '' comment 'SP登出地址',
    `assertion_attr` json comment '自定义的断言属性',
    `custom_attr` json comment '自定义属性',

    `created_time` int(11)  not null default 0,
    `updated_time` int(11)  not null default 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `app_id_unique` (`app_id`) USING BTREE,
	UNIQUE KEY `sp_entity_id_unique` (`sp_entity_id`) USING BTREE,
	UNIQUE KEY `sp_acs_url_unique` (`sp_acs_url`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `sync` (
    `id` int(10) unsigned not null AUTO_INCREMENT,
    `source_id` int(10) unsigned not null comment '账号源ID',
    `source_name` char(20) not null comment '账号源名称 冗余',
    `state` tinyint(1) not null default 0 comment '同步状态 1正在同步 2同步完成',
	`exception_log` text not null comment '同步异常日志',
    `created_time` int(11) not null default 0 comment '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `sync_log` (
    `id` int(10) unsigned not null AUTO_INCREMENT,
    `sync_id` int(11) unsigned not null  comment '同步ID',
    `source_id` int(10) unsigned not null comment '账号源ID',
    `desc` varchar(255) not null comment '同步描述(数据类型)',
    `result` tinyint(1) not null comment '同步结果 0失败 1成功',
    `exception` varchar(255) not null default '' comment '同步异常原因', 
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `system_setting` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '唯一系统配置代号',
  `value` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '对应代号配置的参数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ucode` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user` (
      `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
      `uid` char(36) NOT NULL,  
      `account` char(20) not null default '' comment '账户',
      `phone` varchar(20) not null default '' comment '用户手机号',
      `name` char(36) not null default '' comment '姓名',
      `password` char(32) not null default '' comment '用户密码',
      `emil` char(30) not null default '' comment '邮箱',
    --   `avatar` varchar(255) not null default '' comment '头像路径',
      `disabled` tinyint(1)  not null default 2 comment '是否禁用 1禁用 2 正常',
      `remark` char(50) not null default '' comment '备注',
      `source` tinyint(1) not null default 0 comment '账户来源 1自建 2LDAP/AD',
      `org_id`  int(11) not null default 1 comment '组织ID',
      `group` json not null comment '用户组',
    --   `we_chat_account` varchar(36) not null default '' comment '微信号',
    --   `sex` int(5)  not null default 0 comment '0全部 1.男。2女 3保密',
    --   `post` varchar(36) not null default '' comment '职位',
      `last_ip` char(50) not null default '' comment '上次登录ip',
      `last_device` varchar(255) not null default '' comment '上次登录设备',
      `last_login_time` int(11) not null default 0 comment '上次登录时间',
      `last_operate_time` int(11) not null default 0 comment '上次操作时间 用来判断是否在线',
      `online` tinyint(1) not null default 2 comment '是否在线 1在线 2离线',
      `created_time` int(11)  not null default 0,
      `updated_time` int(11)  not null default 0,
      `deleted_time` int(11)  not null default 0,
      `otp` tinyint(1) NOT NULL DEFAULT 2 COMMENT '是否开启otp： 1 启用 2 禁用',
      `last_edit_password_time` int(11)  not null default 0 comment '上次修改密码的时间',
      `password_level` char(5) not null default '' COMMENT '密码强度等级+长度\nLevel-length',
	  `source_id` int(11) not null default 0 comment '账号所属源ID',
	  `code` char(50) not null default '' comment '用户编号 可以为空 且唯一',
	  `entry_uuid` varchar(100) not null default '' comment '源同步来的uuid',
      `secret` varchar(64) NOT NULL DEFAULT '' COMMENT 'otp 私钥',
      `sdp_ip` varchar(50) NOT NULL DEFAULT '' comment '当前已连接的控制器IP',
      `gateway_ip` varchar(50) NOT NULL default '' comment '当前已连接的网关IP',
      PRIMARY KEY (`id`) USING BTREE,
      UNIQUE KEY `uid_unique` (`uid`) USING BTREE,
      unique key `number_unique` (`account`) USING BTREE
	--   unique key `code_unique` (`code`) USING BTREE,
	--   unique key `source_id_entry_uuid_unique` (`source_id`, `entry_uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_group` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` char(20) NOT NULL DEFAULT '' COMMENT '用户组名称 20个字符',
    `user_count` int(11) NOT NULL DEFAULT 0 COMMENT '用户数量',
    `desc` char(30) NOT NULL DEFAULT '' COMMENT '用户组描述 30个字符',
    `created_time` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间戳',
    `updated_time` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间戳',
    `deleted_time` int(11) NOT NULL DEFAULT 0 COMMENT '删除时间戳',
    PRIMARY KEY (`id`),
    unique key `name_unique` (`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;