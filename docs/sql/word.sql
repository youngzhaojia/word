-- ----------------------------
-- Table structure for t_group
-- ----------------------------
DROP TABLE IF EXISTS `t_group`;
CREATE TABLE `t_group` (
  `FuiGroupId` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `FstrName` varchar(64) NOT NULL DEFAULT '' COMMENT '分组名称',
  `FuiUserId` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `FuiCreateTime` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `FuiUpdateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`FuiGroupId`),
  KEY `idx_FuiUserId` (`FuiUserId`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='单词分组表';

-- ----------------------------
-- Table structure for t_word
-- ----------------------------
DROP TABLE IF EXISTS `t_word`;
CREATE TABLE `t_word` (
  `FuiWordId` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `FstrContent` varchar(64) NOT NULL DEFAULT '' COMMENT '单词',
  `FuiUserId` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `FuiGroupId` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '单词分组ID',
  `FuiCreateTime` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`FuiWordId`),
  KEY `idx_FuiGroupId` (`FuiGroupId`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='单词表';
