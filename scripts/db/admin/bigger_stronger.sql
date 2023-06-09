/*
 Navicat Premium Data Transfer

 Source Server         : localhost3306
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:3306
 Source Schema         : bigger_stronger

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 09/06/2023 18:30:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for analyzing
-- ----------------------------
DROP TABLE IF EXISTS `analyzing`;
CREATE TABLE `analyzing` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint unsigned NOT NULL COMMENT '用户ID',
  `qid` bigint unsigned NOT NULL COMMENT '所属问题',
  `text` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `liked` int unsigned NOT NULL COMMENT '点赞数',
  `reply` int unsigned NOT NULL COMMENT '回复数',
  `create_time` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `qid` (`qid`),
  KEY `uid` (`uid`),
  KEY `sort` (`liked`,`reply`,`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for answer_logs
-- ----------------------------
DROP TABLE IF EXISTS `answer_logs`;
CREATE TABLE `answer_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `qid` bigint NOT NULL COMMENT '题目ID',
  `uid` bigint NOT NULL COMMENT '用户ID\n',
  `answer` json NOT NULL COMMENT '回答的答案',
  `right` tinyint NOT NULL COMMENT '是否正确',
  `time` int NOT NULL COMMENT '时间',
  `sid` bigint unsigned NOT NULL COMMENT '所属用户答题统计记录',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `ut` (`uid`,`right`),
  KEY `qid` (`qid`),
  KEY `qr` (`qid`,`right`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for answer_stat_logs
-- ----------------------------
DROP TABLE IF EXISTS `answer_stat_logs`;
CREATE TABLE `answer_stat_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint NOT NULL,
  `stat` varchar(255) NOT NULL,
  `create_time` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `ut` (`uid`,`create_time`),
  KEY `t` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for questions
-- ----------------------------
DROP TABLE IF EXISTS `questions`;
CREATE TABLE `questions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `type` int NOT NULL COMMENT '所属题目类型ID',
  `desc` text NOT NULL COMMENT '题目描述',
  `option` json NOT NULL COMMENT '题目选项',
  `answer` json NOT NULL COMMENT '题目答案',
  `summary` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '题目考点',
  `create_time` bigint NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for reply
-- ----------------------------
DROP TABLE IF EXISTS `reply`;
CREATE TABLE `reply` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `aid` bigint unsigned NOT NULL COMMENT '所属解析回复',
  `rid` bigint NOT NULL COMMENT '所属回复',
  `text` text NOT NULL COMMENT '内容',
  `create_time` int unsigned NOT NULL COMMENT '创建时间',
  `uid` bigint NOT NULL COMMENT '所属用户',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `at` (`aid`,`create_time`),
  KEY `rt` (`rid`,`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for settings
-- ----------------------------
DROP TABLE IF EXISTS `settings`;
CREATE TABLE `settings` (
  `key` char(64) NOT NULL,
  `value` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `id` bigint NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `k` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for types
-- ----------------------------
DROP TABLE IF EXISTS `types`;
CREATE TABLE `types` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `parent` int unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `openId` char(64) NOT NULL COMMENT '用户微信open ID',
  `name` varchar(64) NOT NULL COMMENT '用户名称 默认微信名称',
  `count` int(10) unsigned zerofill NOT NULL COMMENT '答题次数',
  `stat` varchar(255) NOT NULL COMMENT '答题统计：答对数/总数',
  `create_time` int NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `openID` (`openId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

SET FOREIGN_KEY_CHECKS = 1;
