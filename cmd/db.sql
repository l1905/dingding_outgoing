CREATE TABLE `outgoing` (
 `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
 `user_id` varchar(100) NOT NULL DEFAULT '' COMMENT '新建用户ID',
 `nickname` varchar(200) NOT NULL DEFAULT '' COMMENT '新建用户昵称',
 `chatbot_corp_id` varchar(200) NOT NULL DEFAULT '' COMMENT '机器人公司ID',
 `tag` varchar(200) NOT NULL DEFAULT '' COMMENT '标签',
 `url` varchar(1000) NOT NULL DEFAULT '' COMMENT '文档链接地址',
 `item_desc` varchar(1000) NOT NULL DEFAULT '' COMMENT '详细内容',
 `creation_time` datetime DEFAULT NULL,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COMMENT='钉钉outgoing机器表'