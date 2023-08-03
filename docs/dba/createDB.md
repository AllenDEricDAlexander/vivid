# 建表说明

字符集 utf8mb4

全部非空，int类型就设置 默认值0，string类型 看情况 设置默认值，时间类型 设置为系统时间

id非自增，程序介入或手动传，demo可以做成uuid 或者 随机数 先占位

# User表

```sql
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `user_id` int(0) UNSIGNED NOT NULL COMMENT 'user_id，非自增，后端处理传进来',
  `user_name` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'user_name，not null 默认生成UUID填充',
  `follow_count` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'follow_count，not null ，默认填充0',
  `follower_count` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'follower_count，not null，填充0',
  `avatarURL` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'avatarURL，not null 默认填充url',
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'background_image，notnull 默认填充url',
  `signature` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT 'signature，not null，默认填充“暂未输入”',
  `total_favorite` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'total_favorite,not null ，默认填充0',
  `work_count` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'work_count，not null，默认填充0',
  `favorite_count` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'favorite_count,not null 默认填充0',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'password，not null ',
  `is_delete` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'is_delete，not null 默认填充0',
  `create_time` datetime(0) NOT NULL COMMENT 'create_time not null 第一次填充系统时间',
  `update_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT 'update_time not null 更新刷新',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;
SET FOREIGN_KEY_CHECKS = 1;
```

