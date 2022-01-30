CREATE TABLE `user_info` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `phone_num` varchar(36) NOT NULL DEFAULT '' COMMENT '手机号',
  `password` varchar(36) NOT NULL DEFAULT '' COMMENT '密码',
  `create_time` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modify_time` int NOT NULL DEFAULT '0' COMMENT '修改时间',
  `delete_time` int NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除@：0否@：1是',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

CREATE TABLE `user_login` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `device_id` varchar(36) NOT NULL DEFAULT '' COMMENT '设备唯一码',
  `access_token` varchar(36) NOT NULL DEFAULT '' COMMENT 'access_token',
  `refresh_token` varchar(36) NOT NULL DEFAULT '' COMMENT 'refresh_token',
  `access_time` int NOT NULL DEFAULT '0' COMMENT '最后刷新access_token时间',
  `create_time` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modify_time` int NOT NULL DEFAULT '0' COMMENT '修改时间',
  `delete_time` int NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除@：0否@：1是',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户登录表';