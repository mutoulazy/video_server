# video_server
go video server

## database

``` mysql
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `video_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `author_id` int(10) NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL,
  `time` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions`  (
  `session_id` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `TTL` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL,
  `login_name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `login_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `pwd` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UK_login_name`(`login_name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for video_info
-- ----------------------------
DROP TABLE IF EXISTS `video_info`;
CREATE TABLE `video_info`  (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `author_id` int(10) NULL DEFAULT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL,
  `display_ctime` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL,
  `create_time` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
```