/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : demo

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 03/08/2023 19:26:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for friends
-- ----------------------------
DROP TABLE IF EXISTS `friends`;
CREATE TABLE `friends`  (
  `friend_id` int(0) UNSIGNED NOT NULL COMMENT 'friend_id 主键',
  `user1_id` int(0) UNSIGNED NOT NULL COMMENT 'user1_id',
  `user2_id` int(0) UNSIGNED NOT NULL COMMENT 'user2_id',
  `is_delete` int(0) NOT NULL COMMENT 'is_delete 关系是否还存在',
  `create_time` datetime(0) NOT NULL COMMENT 'create_time',
  `update_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT 'update_time',
  PRIMARY KEY (`friend_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
