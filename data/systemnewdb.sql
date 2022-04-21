-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2022-04-21 15:29:23
-- 服务器版本： 8.0.26
-- PHP 版本： 7.3.24-(to be removed in future macOS)

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `systemnewdb`
--

-- --------------------------------------------------------

--
-- 表的结构 `system_article`
--

CREATE TABLE `system_article` (
  `id` bigint UNSIGNED NOT NULL COMMENT '主键',
  `author` int NOT NULL DEFAULT '0' COMMENT '作者',
  `importance` tinyint UNSIGNED NOT NULL DEFAULT '1' COMMENT '重要级别',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态(0:draft，1:published,10:deleted)',
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `content_short` varchar(500) NOT NULL DEFAULT '' COMMENT '摘要',
  `source_uri` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '来源',
  `ctime` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `image_uri` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `comment_disabled` tinyint NOT NULL DEFAULT '0' COMMENT '是否展示评论',
  `display_time` datetime NOT NULL COMMENT '发布时间',
  `mtime` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;

-- --------------------------------------------------------

--
-- 表的结构 `system_log`
--

CREATE TABLE `system_log` (
  `id` bigint NOT NULL COMMENT '主键',
  `system_user_id` int DEFAULT '0' COMMENT '主键',
  `title` varchar(300) NOT NULL DEFAULT '' COMMENT '日志标题',
  `content` text COMMENT '日志内容记录SQL',
  `relation_id` bigint NOT NULL DEFAULT '0' COMMENT '相关对应表主键',
  `relation_table` int NOT NULL DEFAULT '1' COMMENT '对应表(1 system_user,2 system_menu,3 system_role)',
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'ip',
  `url` varchar(500) NOT NULL DEFAULT '',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='操作日志';

-- --------------------------------------------------------

--
-- 表的结构 `system_menu`
--

CREATE TABLE `system_menu` (
  `id` int NOT NULL COMMENT '主键',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
  `alias` varchar(100) NOT NULL DEFAULT '' COMMENT '别名',
  `path` varchar(50) NOT NULL DEFAULT '' COMMENT '路径',
  `component` varchar(100) NOT NULL DEFAULT '' COMMENT '组件',
  `redirect` varchar(200) NOT NULL DEFAULT '' COMMENT '重定向',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT '访问url',
  `meta_title` varchar(50) NOT NULL DEFAULT '' COMMENT 'meta标题',
  `meta_icon` varchar(50) NOT NULL DEFAULT '' COMMENT 'meta icon',
  `meta_i18n` tinyint NOT NULL DEFAULT '0' COMMENT '是否国家化（1:是 0:否）',
  `meta_showlink` tinyint NOT NULL DEFAULT '0' COMMENT '是否总是显示（1:是0：否）',
  `meta_showparent` tinyint NOT NULL COMMENT '是否显示父级菜单1是0否',
  `meta_rank` tinyint NOT NULL DEFAULT '0' COMMENT '排序',
  `meta_keepalive` tinyint NOT NULL DEFAULT '0' COMMENT '是否开启缓存（1开启0关闭)',
  `type` tinyint NOT NULL DEFAULT '2' COMMENT '类型(1:固定,2:权限配置3特殊)',
  `meta_framesrc` varchar(300) NOT NULL DEFAULT '0' COMMENT '内嵌的iframe链接',
  `transition_name` tinyint NOT NULL DEFAULT '0' COMMENT '是否显示动画（1是0否）',
  `transition_enter` varchar(50) NOT NULL DEFAULT '' COMMENT '进入动画',
  `transition_leave` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '离开动画',
  `dynamiclevel` tinyint NOT NULL DEFAULT '3' COMMENT '动态路由可打开的最大数量',
  `refreshredirect` varchar(50) NOT NULL DEFAULT '' COMMENT '刷新重定向',
  `extraicon_svg` tinyint NOT NULL DEFAULT '0' COMMENT '额外图标svg(1是0否)',
  `extraicon_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图片名称',
  `pid` int NOT NULL DEFAULT '0' COMMENT '父ID',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '状态（0禁止1启动）',
  `level` tinyint NOT NULL DEFAULT '0' COMMENT '层级',
  `ctime` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限';

-- --------------------------------------------------------

--
-- 表的结构 `system_role`
--

CREATE TABLE `system_role` (
  `id` int NOT NULL COMMENT '主键',
  `name` varchar(100) NOT NULL COMMENT '角色名称',
  `alias_name` varchar(50) NOT NULL DEFAULT '' COMMENT '别名',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '角色状态（0无效1有效）',
  `type` int NOT NULL DEFAULT '1' COMMENT '属于哪个应用',
  `ctime` datetime NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色';

-- --------------------------------------------------------

--
-- 表的结构 `system_role_menu`
--

CREATE TABLE `system_role_menu` (
  `id` int NOT NULL COMMENT '主键',
  `system_role_id` int NOT NULL DEFAULT '0' COMMENT '角色主键',
  `system_menu_id` int NOT NULL DEFAULT '0' COMMENT '菜单主键'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色与菜单关联表';

-- --------------------------------------------------------

--
-- 表的结构 `system_user`
--

CREATE TABLE `system_user` (
  `id` int NOT NULL COMMENT '主键',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '登录名',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `salt` varchar(4) NOT NULL COMMENT '盐',
  `phone` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(300) NOT NULL DEFAULT '' COMMENT '头像',
  `introduction` varchar(300) NOT NULL DEFAULT '' COMMENT '简介',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '状态（0 停止1启动）',
  `utime` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `last_login_time` datetime NOT NULL COMMENT '上次登录时间',
  `last_login_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '最近登录IP',
  `ctime` datetime NOT NULL COMMENT '注册时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理账户表';

-- --------------------------------------------------------

--
-- 表的结构 `system_user_role`
--

CREATE TABLE `system_user_role` (
  `id` int NOT NULL COMMENT '主键',
  `system_user_id` int NOT NULL COMMENT '用户主键',
  `system_role_id` int NOT NULL COMMENT '角色主键'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='账户和角色关联表';

--
-- 转储表的索引
--

--
-- 表的索引 `system_article`
--
ALTER TABLE `system_article`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `system_log`
--
ALTER TABLE `system_log`
  ADD PRIMARY KEY (`id`),
  ADD KEY `RELATION_ID` (`relation_id`),
  ADD KEY `SYSTEM_USER_ID` (`system_user_id`),
  ADD KEY `CTIME` (`ctime`),
  ADD KEY `RELATION_TABLE` (`relation_table`);

--
-- 表的索引 `system_menu`
--
ALTER TABLE `system_menu`
  ADD PRIMARY KEY (`id`),
  ADD KEY `path` (`path`);

--
-- 表的索引 `system_role`
--
ALTER TABLE `system_role`
  ADD PRIMARY KEY (`id`),
  ADD KEY `TYPE` (`type`),
  ADD KEY `STATUS` (`state`);

--
-- 表的索引 `system_role_menu`
--
ALTER TABLE `system_role_menu`
  ADD PRIMARY KEY (`id`),
  ADD KEY `system_role_id` (`system_role_id`,`system_menu_id`);

--
-- 表的索引 `system_user`
--
ALTER TABLE `system_user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `NICKNAME` (`nickname`),
  ADD KEY `PASSWORD` (`password`);

--
-- 表的索引 `system_user_role`
--
ALTER TABLE `system_user_role`
  ADD PRIMARY KEY (`id`),
  ADD KEY `system_user_id` (`system_user_id`,`system_role_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `system_article`
--
ALTER TABLE `system_article`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `system_log`
--
ALTER TABLE `system_log`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `system_menu`
--
ALTER TABLE `system_menu`
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `system_role`
--
ALTER TABLE `system_role`
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `system_role_menu`
--
ALTER TABLE `system_role_menu`
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `system_user`
--
ALTER TABLE `system_user`
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `system_user_role`
--
ALTER TABLE `system_user_role`
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键';
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
