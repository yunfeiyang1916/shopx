CREATE TABLE IF NOT EXISTS `account_user_base` (
                                     `user_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户编号',
                                     `user_account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
                                     `user_password` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码',
                                     `user_salt` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'salt值',
                                     PRIMARY KEY (`user_id`) USING BTREE,
                                     UNIQUE KEY `user_account` (`user_account`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='用户基本信息表'