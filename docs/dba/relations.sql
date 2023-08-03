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

 Date: 03/08/2023 19:26:32
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for relations
-- ----------------------------
DROP TABLE IF EXISTS `relations`;
CREATE TABLE `relations`  (
  `relation_id` int(0) UNSIGNED NOT NULL COMMENT 'relation_id ',
  `follow_id` int(0) UNSIGNED NOT NULL COMMENT 'follow_id 被关注者id',
  `follower_id` int(0) UNSIGNED NOT NULL COMMENT 'follower_id 关注者id',
  `is_follow` int(0) NOT NULL COMMENT 'is_follow 关注关系是否还存在',
  `create_time` datetime(0) NOT NULL COMMENT 'create_time 首次创建填充',
  `update_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT 'update_time 修改刷新',
  PRIMARY KEY (`relation_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
