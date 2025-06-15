set sql_mode=(select replace(@@sql_mode,'NO_ZERO_IN_DATE,NO_ZERO_DATE',''));

-- ----------------------------
-- Table structure for account_user_analytics
-- ----------------------------
DROP TABLE IF EXISTS `account_user_analytics`;
CREATE TABLE `account_user_analytics`  (
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `user_spend` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '消费总额',
  `user_refund` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '退款总额',
  `user_order_buy_num` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效订单数量',
  `user_order_return_num` int(11) UNSIGNED NOT NULL COMMENT '退单数量',
  `user_product_buy_num` int(11) UNSIGNED NOT NULL COMMENT '购买数量',
  `user_product_return_num` int(11) UNSIGNED NOT NULL COMMENT '退货数量',
  `version` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户角色信息表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_base
-- ----------------------------
DROP TABLE IF EXISTS `account_user_base`;
CREATE TABLE `account_user_base`  (
  `user_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `user_account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `user_password` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码',
  `user_salt` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'salt值',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `user_account`(`user_account`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户基本信息表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of account_user_base
-- ----------------------------
INSERT INTO `account_user_base` VALUES (10001, 'admin', 'ff64b03e55429443e29a3fff3abbe00b', '77c96449b196487d92214bd08121ac62');

-- ----------------------------
-- Table structure for account_user_bind_connect
-- ----------------------------
DROP TABLE IF EXISTS `account_user_bind_connect`;
CREATE TABLE `account_user_bind_connect`  (
  `bind_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '绑定标记',
  `bind_type` smallint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '绑定类型(EMUN):1-mobile;  2-email;   13-weixin公众号',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `bind_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '绑定时间',
  `bind_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名称',
  `bind_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户图标',
  `bind_openid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '访问编号',
  `bind_unionid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'unionid',
  `bind_active` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否激活(BOOL):0-未激活;1-激活',
  PRIMARY KEY (`bind_id`) USING BTREE,
  UNIQUE INDEX `bind_type`(`bind_type`, `bind_id`, `user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户绑定表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_block
-- ----------------------------
DROP TABLE IF EXISTS `account_user_block`;
CREATE TABLE `account_user_block`  (
  `user_black_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '黑名单编号',
  `user_id` int(11) UNSIGNED NOT NULL COMMENT '用户编号',
  `black_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '拉黑用户编号',
  `user_black_time` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '拉黑时间',
  PRIMARY KEY (`user_black_id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `black_id`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户黑名单表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_card
-- ----------------------------
DROP TABLE IF EXISTS `account_user_card`;
CREATE TABLE `account_user_card`  (
  `card_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '会员卡编号',
  `card_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '会员卡名称',
  `user_level_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员等级',
  `card_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '状态(BOOL):0-禁用;1-开启',
  `card_pirce` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '会费（元/期）',
  `card_market_price` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '原会费（元/期）',
  `card_save_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '预计年省金额',
  `card_year` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效期年',
  `card_month` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效期月',
  `card_day` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效期日',
  `card_time` bigint(13) NOT NULL DEFAULT 0 COMMENT '上线时间',
  `card_sort` smallint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `user_type_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '会员类型(ENUM):1-普通用户;2-扩展用户',
  `transport_type_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '物流模板编号',
  `card_bind_addr` bit(1) NOT NULL DEFAULT b'0' COMMENT '唯一地址(BOOL):0-否;1-是',
  `card_is_exp` bit(1) NOT NULL DEFAULT b'0' COMMENT '体验卡(BOOL):0-否;1-是',
  `card_exp_month` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '体验月',
  `card_exp_day` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '体验日',
  `card_exp_fee` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '体验费用',
  `card_exp_status` bit(1) NOT NULL DEFAULT b'0' COMMENT '体验状态(BOOL):0-禁用;1-开启',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `activity_id_discount` int(10) NOT NULL DEFAULT 0 COMMENT '活动会员专享折扣编号',
  `activity_id_gitfbag` int(10) NOT NULL DEFAULT 0 COMMENT '活动会员开卡赠品编号',
  `activity_id_member_day` int(10) NOT NULL DEFAULT 0 COMMENT '活动会员日（实物）编号',
  `activity_id_birthday` int(10) NOT NULL DEFAULT 0 COMMENT '活动生日专享编号',
  `user_fx_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '允许分销(BOOL):1-启用分销;0-禁用分销',
  `user_fx_rate` decimal(6, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '分销佣金比例',
  `card_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '分享标题',
  `card_desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '分享描述',
  `card_keywords` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '搜索关键字',
  `card_img` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '分享图标',
  `card_fx_active` bit(1) NOT NULL DEFAULT b'0' COMMENT '开通推广员(BOOL):1-开通;0-不开通',
  `card_img_bg` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '背景图片',
  PRIMARY KEY (`card_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户会员卡' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_channel
-- ----------------------------
DROP TABLE IF EXISTS `account_user_channel`;
CREATE TABLE `account_user_channel`  (
  `user_channel_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '渠道编号',
  `user_channel_name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '渠道名称',
  `user_channel_prefix` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '渠道前缀',
  `user_channel_enable` smallint(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `user_channel_buildin` tinyint(1) UNSIGNED NOT NULL COMMENT '是否内置(BOOL):0-非内置;1-内置',
  PRIMARY KEY (`user_channel_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户渠道表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_channel_code
-- ----------------------------
DROP TABLE IF EXISTS `account_user_channel_code`;
CREATE TABLE `account_user_channel_code`  (
  `ucc_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '渠道编号',
  `ucc_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '渠道名称',
  `ucc_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员编号',
  `user_channel_id` int(11) UNSIGNED NOT NULL COMMENT '渠道编号',
  `ucc_used` smallint(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否使用(BOOL):0-未使用;1-已使用',
  `ucc_use_time` bigint(14) NOT NULL COMMENT '使用时间',
  `ucc_create_time` bigint(14) NOT NULL COMMENT '创建时间',
  `ucc_is_unique` smallint(11) UNSIGNED NOT NULL DEFAULT 1 COMMENT '唯一使用(BOOL):0-不限制;1-可用一次',
  `ucc_use_num` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '使用次数',
  PRIMARY KEY (`ucc_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户渠道邀请码表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_delivery_address
-- ----------------------------
DROP TABLE IF EXISTS `account_user_delivery_address`;
CREATE TABLE `account_user_delivery_address`  (
  `ud_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '地址编号',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `ud_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `ud_intl` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家编码',
  `ud_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `ud_telephone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `ud_province_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '省编号',
  `ud_province` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '省份',
  `ud_city_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '市编号',
  `ud_city` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '市',
  `ud_county_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '县',
  `ud_county` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '县区',
  `ud_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细地址',
  `ud_postalcode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮政编码',
  `ud_tag_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址标签(ENUM):1001-家里;1002-公司',
  `ud_longitude` double NOT NULL DEFAULT 0 COMMENT '经度',
  `ud_latitude` double NOT NULL DEFAULT 0 COMMENT '纬读',
  `ud_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `ud_is_default` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否默认(BOOL):0-非默认;1-默认',
  PRIMARY KEY (`ud_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `ud_is_default`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户地址表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_distribution
-- ----------------------------
DROP TABLE IF EXISTS `account_user_distribution`;
CREATE TABLE `account_user_distribution`  (
  `user_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `user_parent_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级用户编号',
  `user_partner_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属城市合伙人',
  `user_team_count` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '团队人数',
  `user_province_team_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属省公司',
  `user_city_team_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属市公司',
  `user_county_team_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属区公司',
  `role_level_id` mediumint(11) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '角色等级',
  `ucc_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '渠道编号',
  `activity_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动编号',
  `user_time` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `user_fans_num` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '粉丝数量:冗余',
  `user_is_sp` bit(1) NOT NULL DEFAULT b'0' COMMENT '服务商(BOOL):0-否;1-是;',
  `user_is_da` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '区代理(BOOL):0-否;1-是为区Id;',
  `user_is_ca` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '市代理(BOOL):0-否;1-是为市Id;',
  `user_is_pa` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '省代理(BOOL):0-否;1-是为省Id;',
  `user_is_pt` bit(1) NOT NULL DEFAULT b'0' COMMENT '城市合伙人(BOOL):0-否;1-是;',
  `user_active` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否有效(BOOL):0-未生效;1-有效',
  `user_voucher_ids` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分销优惠券',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `user_time`(`user_time`) USING BTREE,
  INDEX `user_parent_id`(`user_parent_id`, `user_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '粉丝来源关系表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_friend
-- ----------------------------
DROP TABLE IF EXISTS `account_user_friend`;
CREATE TABLE `account_user_friend`  (
  `user_friend_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(11) UNSIGNED NOT NULL COMMENT '用户编号',
  `friend_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '好友编号',
  `friend_note` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注名称',
  `user_friend_addtime` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `friend_state` tinyint(2) UNSIGNED NOT NULL DEFAULT 1 COMMENT '关注状态(ENUM):1-单向关注;2-双向关注',
  `friend_invite` tinyint(2) UNSIGNED NOT NULL DEFAULT 0 COMMENT '邀请状态(ENUM):0-新邀请;2-处理完成后邀请',
  PRIMARY KEY (`user_friend_id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `friend_id`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户好友关系表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_group
-- ----------------------------
DROP TABLE IF EXISTS `account_user_group`;
CREATE TABLE `account_user_group`  (
  `group_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '好友组编号',
  `group_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组名称',
  `group_type` tinyint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群组类型(ENUM):0-临时组上限100人;  1-普通组上限300人; 2-VIP组 上限500人',
  `user_id` int(11) UNSIGNED NOT NULL COMMENT '管理员',
  `group_num` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '数量',
  PRIMARY KEY (`group_id`) USING BTREE,
  INDEX `group_name`(`group_name`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '好友管理组' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_group_rel
-- ----------------------------
DROP TABLE IF EXISTS `account_user_group_rel`;
CREATE TABLE `account_user_group_rel`  (
  `group_rel_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '好友组编号',
  `group_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '组名称',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户',
  PRIMARY KEY (`group_rel_id`) USING BTREE,
  UNIQUE INDEX `group_name`(`group_id`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '好友组关系' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_history
-- ----------------------------
DROP TABLE IF EXISTS `account_user_history`;
CREATE TABLE `account_user_history`  (
  `user_id` int(11) UNSIGNED NOT NULL COMMENT '用户编号',
  `user_history_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '历史密码',
  `user_history_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '历史 IP',
  `user_history_time` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户密码历史信息表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_info
-- ----------------------------
DROP TABLE IF EXISTS `account_user_info`;
CREATE TABLE `account_user_info`  (
  `user_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `user_account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `user_state` mediumint(8) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态(ENUM):0-锁定;1-已激活;2-未激活;',
  `user_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号码(mobile)',
  `user_intl` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家编码',
  `user_gender` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '性别(ENUM):0-保密;1-男;  2-女;',
  `user_birthday` date NOT NULL DEFAULT '2000-01-01' COMMENT '生日(DATE)',
  `user_email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱(email)',
  `user_level_id` smallint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '等级编号',
  `user_realname` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `user_idcard` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '身份证',
  `user_idcard_images` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '身份证图片(DTO)',
  `user_is_authentication` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '认证状态(ENUM):0-未认证;1-待审核;2-认证通过;3-认证失败',
  `tag_ids` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户标签(DOT)',
  `user_from` smallint(4) UNSIGNED NOT NULL DEFAULT 2310 COMMENT '用户来源(ENUM):2310-其它;2311-pc;2312-H5;2313-APP;2314-小程序;2315-公众号',
  `user_new` bit(1) NOT NULL DEFAULT b'1' COMMENT '新人标识(BOOL):0-不是;1-是',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户详细信息表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of account_user_info
-- ----------------------------
INSERT INTO `account_user_info` VALUES (10001, 'admin', '系统管理员', 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231031/8aa85cf4119746f2be5854c8cb081ea5.png', 1, '13488888888', '+86', 2, '2023-10-26', 'suishang@2018.com', 1001, '随商', '341222199310239898', 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231016/2cca497ab8344d588e353bdafe531992.png,https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231016/6a3e27ae488b4d408e4d454957755126.png', 2, '', 2310, b'1');

-- ----------------------------
-- Table structure for account_user_invoice
-- ----------------------------
DROP TABLE IF EXISTS `account_user_invoice`;
CREATE TABLE `account_user_invoice`  (
  `user_invoice_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '发票编号',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户',
  `invoice_title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '发票抬头',
  `invoice_company_code` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '纳税人识别号',
  `invoice_content` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '发票内容',
  `invoice_is_company` bit(1) NOT NULL DEFAULT b'1' COMMENT '公司开票(BOOL):0-个人;1-公司',
  `invoice_is_electronic` bit(1) NOT NULL DEFAULT b'1' COMMENT '电子发票(ENUM):0-纸质发票;1-电子发票',
  `invoice_type` smallint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '发票类型(ENUM):1-普通发票;2-增值税专用发票',
  `invoice_datetime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '添加时间',
  `invoice_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '单位地址',
  `invoice_phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '单位电话',
  `invoice_bankname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '开户银行',
  `invoice_bankaccount` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '银行账号',
  `invoice_contact_mobile` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票人手机',
  `invoice_contact_email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票人邮箱',
  `invoice_is_default` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否默认',
  `invoice_contact_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票人',
  `invoice_contact_area` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票人地区',
  `invoice_contact_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票详细地址',
  PRIMARY KEY (`user_invoice_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户发票管理表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_level
-- ----------------------------
DROP TABLE IF EXISTS `account_user_level`;
CREATE TABLE `account_user_level`  (
  `user_level_id` smallint(4) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '等级编号',
  `user_level_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '等级名称',
  `user_level_exp` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '升级经验值',
  `user_level_spend` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '累计消费',
  `user_level_logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '等级图标',
  `user_level_rate` smallint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '折扣率百分比',
  `user_level_time` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  `user_level_is_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):0-否;1-是',
  PRIMARY KEY (`user_level_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1004 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户等级表-平台' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of account_user_level
-- ----------------------------
INSERT INTO `account_user_level` VALUES (1001, 'vip1', 100, 0, 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230914/a20b208b4da04b13af76d8f9703054ca.png', 100, 20230829143055, b'1');
INSERT INTO `account_user_level` VALUES (1002, 'vip2', 0, 2, 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230914/ddd36d38be35474691ef1dedf807dcf0.png', 95, 20230829143058, b'1');
INSERT INTO `account_user_level` VALUES (1003, 'vip3', 0, 3, 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230914/6bb643e537854d1ab9bd53e6f801183a.png', 90, 20230829140104, b'1');

-- ----------------------------
-- Table structure for account_user_level_log
-- ----------------------------
DROP TABLE IF EXISTS `account_user_level_log`;
CREATE TABLE `account_user_level_log`  (
  `user_level_log_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志编号',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `user_level_id` smallint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前等级ID',
  `user_level_pre_id` smallint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '之前等级ID',
  `operate_remark` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作备注',
  `operate_user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作人',
  `create_time` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作时间',
  PRIMARY KEY (`user_level_log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户等级记录表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_login
-- ----------------------------
DROP TABLE IF EXISTS `account_user_login`;
CREATE TABLE `account_user_login`  (
  `user_id` int(11) UNSIGNED NOT NULL COMMENT '用户编号',
  `user_active_time` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '激活时间',
  `user_reg_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '注册IP',
  `user_reg_date` date NOT NULL COMMENT '注册日期',
  `user_reg_time` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `user_count_login` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录次数',
  `user_lastlogin_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '上次登录IP',
  `user_lastlogin_time` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上次登录时间',
  `user_clientid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'APPCID(DOT)',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `user_reg_time`(`user_reg_time`) USING BTREE,
  INDEX `user_reg_date`(`user_reg_date`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户登录信息表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_login_history
-- ----------------------------
DROP TABLE IF EXISTS `account_user_login_history`;
CREATE TABLE `account_user_login_history`  (
  `user_login_history_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志编号',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `user_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `user_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录IP',
  `user_login_time` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录时间',
  `user_login_state` bit(1) NOT NULL DEFAULT b'1' COMMENT '登录状态(BOOL):0-登录失败;1-登录成功',
  PRIMARY KEY (`user_login_history_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `user_login_ip`(`user_login_ip`, `user_login_state`, `user_login_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户登录信息表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_message
-- ----------------------------
DROP TABLE IF EXISTS `account_user_message`;
CREATE TABLE `account_user_message`  (
  `message_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '消息编号',
  `message_parent_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级编号',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 1 COMMENT '所属用户:发送者或者接收者，如果message_kind=1则为当前用户发送的消息。',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `message_kind` smallint(4) UNSIGNED NOT NULL DEFAULT 1 COMMENT '消息种类(ENUM):1-发送消息;2-接收消息',
  `user_other_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '相关用户:发送者或者接收者',
  `user_other_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '相关昵称:发送者或者接收者',
  `message_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息标题',
  `message_content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息内容',
  `message_time` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送时间',
  `message_is_read` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否读取(BOOL):0-未读;1-已读',
  `message_is_delete` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除(BOOL):0-正常状态;1-删除状态',
  `message_type` smallint(4) UNSIGNED NOT NULL DEFAULT 2 COMMENT '消息类型(ENUM):1-系统消息;2-用户消息',
  `message_cat` enum('text','img','video','file','location','redpack') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'text' COMMENT '消息类型(ENUM):text-文本消息;img-图片消息;video-视频消息;file:文件;location:位置;redpack:红包',
  `message_data_type` smallint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息分类(ENUM):0-默认消息;1-公告消息;2-订单消息;3-商品消息;4-余额卡券;5-服务消息',
  `message_data_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息数据:商品编号|订单编号',
  `message_length` mediumint(8) NOT NULL DEFAULT 0 COMMENT '消息长度',
  `message_w` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '图片宽度',
  `message_h` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '图片高度',
  PRIMARY KEY (`message_id`) USING BTREE,
  INDEX `from_member_id`(`user_other_id`) USING BTREE,
  INDEX `to_member_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '短消息-聊天记录' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_sns
-- ----------------------------
DROP TABLE IF EXISTS `account_user_sns`;
CREATE TABLE `account_user_sns`  (
  `user_id` int(11) UNSIGNED NOT NULL COMMENT '用户编号',
  `user_blog` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '微博数量',
  `user_friend` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '好友数量',
  `user_fans` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '粉丝数量',
  `user_growth` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '成长值',
  `user_report` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否可以举报(BOOL):0-不可以;1-可以',
  `user_buy` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否可以购买商品(BOOL):0-不可以;1-可以',
  `user_comment` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否允许发表言论(BOOL):0-不可以;1-可以',
  `user_fans_store` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '推广店铺数量',
  `user_story` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子数量',
  `user_story_comment` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论数量',
  `user_favorites_store` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏店铺',
  `user_favorites_item` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏商品',
  `user_favorites_brand` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏品牌',
  `user_story_collection` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏帖子',
  `user_story_like` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子点赞',
  `user_story_forward` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子转发',
  `version` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户SNS信息表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_tag_base
-- ----------------------------
DROP TABLE IF EXISTS `account_user_tag_base`;
CREATE TABLE `account_user_tag_base`  (
  `tag_id` int(50) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '标签编码',
  `tag_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签标题',
  `tag_group_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属分类',
  `tag_sort` smallint(4) UNSIGNED NOT NULL DEFAULT 255 COMMENT '配置排序:从小到大',
  `tag_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `tag_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  PRIMARY KEY (`tag_id`) USING BTREE,
  INDEX `tag_title`(`tag_title`) USING BTREE,
  INDEX `tag_group_id`(`tag_group_id`, `tag_sort`, `tag_enable`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户标签表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of account_user_tag_base
-- ----------------------------
INSERT INTO `account_user_tag_base` VALUES (31, '新客户', 22, 10, b'0', b'0');
INSERT INTO `account_user_tag_base` VALUES (32, '重点客户', 22, 20, b'1', b'0');
INSERT INTO `account_user_tag_base` VALUES (33, '内部员工', 22, 50, b'1', b'0');
INSERT INTO `account_user_tag_base` VALUES (34, '待召回', 22, 255, b'1', b'0');
INSERT INTO `account_user_tag_base` VALUES (35, '华东', 23, 255, b'1', b'0');
INSERT INTO `account_user_tag_base` VALUES (36, '华中', 23, 255, b'1', b'0');
INSERT INTO `account_user_tag_base` VALUES (37, '华南', 23, 255, b'1', b'0');
INSERT INTO `account_user_tag_base` VALUES (38, '华北', 23, 255, b'1', b'0');
INSERT INTO `account_user_tag_base` VALUES (39, '电子数码', 24, 255, b'1', b'0');
INSERT INTO `account_user_tag_base` VALUES (40, '美妆达人', 24, 255, b'1', b'0');
INSERT INTO `account_user_tag_base` VALUES (42, '大客户', 22, 1, b'1', b'0');
INSERT INTO `account_user_tag_base` VALUES (44, '不活跃用户', 22, 0, b'1', b'0');

-- ----------------------------
-- Table structure for account_user_tag_group
-- ----------------------------
DROP TABLE IF EXISTS `account_user_tag_group`;
CREATE TABLE `account_user_tag_group`  (
  `tag_group_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分组编号',
  `tag_group_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分组名称',
  `tag_group_sort` smallint(4) UNSIGNED NOT NULL DEFAULT 255 COMMENT '分组排序:从小到大',
  `tag_group_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否有效(BOOL):0-禁用;1-启用',
  `tag_group_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_time` timestamp(0) NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  PRIMARY KEY (`tag_group_id`) USING BTREE,
  INDEX `index_name`(`tag_group_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '标签分组表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of account_user_tag_group
-- ----------------------------
INSERT INTO `account_user_tag_group` VALUES (22, '客户级别', 10, b'1', b'0', '2023-09-14 08:22:05', '0000-00-00 00:00:00');
INSERT INTO `account_user_tag_group` VALUES (23, '客户地区', 20, b'1', b'0', '2023-09-14 08:24:12', '0000-00-00 00:00:00');
INSERT INTO `account_user_tag_group` VALUES (24, '购买类型', 30, b'1', b'0', '2023-09-14 08:25:49', '0000-00-00 00:00:00');

-- ----------------------------
-- Table structure for account_user_zone
-- ----------------------------
DROP TABLE IF EXISTS `account_user_zone`;
CREATE TABLE `account_user_zone`  (
  `zone_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '好友组编号',
  `zone_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '群组名称',
  `zone_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群组类型(ENUM):0-临时组上限100人;  1-普通组上限300人; 2-VIP组 上限500人',
  `zone_permission` tinyint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '申请加入模式(ENUM): 0-默认直接加入; 1-需要身份验证; 2-私有群组',
  `zone_declared` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '群组公告',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员',
  `zone_bind_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '第三方群组编号',
  `zone_user_num` int(11) UNSIGNED NOT NULL DEFAULT 2 COMMENT '人数',
  PRIMARY KEY (`zone_id`) USING BTREE,
  UNIQUE INDEX `group_name`(`zone_name`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '群组' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_zone_message
-- ----------------------------
DROP TABLE IF EXISTS `account_user_zone_message`;
CREATE TABLE `account_user_zone_message`  (
  `zone_message_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '短消息索引编号',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 1 COMMENT '发送用户',
  `zone_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群组编号',
  `zone_message_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '短消息内容',
  `zone_message_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '短消息发送时间',
  `zone_message_is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '短消息状态(BOOL):0-正常状态;1-删除状态',
  `zone_message_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 2 COMMENT '消息类型(ENUM):1-系统消息;2-用户消息;3-私信',
  PRIMARY KEY (`zone_message_id`) USING BTREE,
  INDEX `from_member_id`(`zone_id`) USING BTREE,
  INDEX `to_member_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '群组消息-聊天记录' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for account_user_zone_rel
-- ----------------------------
DROP TABLE IF EXISTS `account_user_zone_rel`;
CREATE TABLE `account_user_zone_rel`  (
  `zone_rel_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '关系编号',
  `zone_id` int(11) UNSIGNED NOT NULL COMMENT '群组编号',
  `user_id` int(11) UNSIGNED NOT NULL COMMENT '用户',
  `zone_rel_permission` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '申请加入模式(ENUM): 0-加入; 1-待验证;',
  PRIMARY KEY (`zone_rel_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '群组用户关系表' ROW_FORMAT = Compact;
