-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2022-03-23 16:30:33
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
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;

--
-- 转存表中的数据 `system_article`
--

INSERT INTO `system_article` (`id`, `author`, `importance`, `state`, `title`, `content`, `content_short`, `source_uri`, `ctime`, `image_uri`, `comment_disabled`, `display_time`, `mtime`) VALUES
(1, 1, 2, 1, 'fasd', '<p>fasdffasdf</p>', 'a发生的发生', '', 1585670400, '20200407/ftxqfEY9EjWGyHS7CFaDWkYYh2SLYUIY.jpeg', 0, '2020-04-10 00:00:00', '2020-04-09 15:14:31'),
(2, 1, 2, 1, 'fasd', '<p>fasdffasdf</p>', 'a', 'https://www.baidu.com/', 1585670400, '20200407/jbsW7XXW5nFcj6rNEydZhs9ma4BerLpB.jpeg', 0, '2020-04-01 00:00:00', '2020-03-31 16:00:00'),
(3, 1, 2, 1, '测试', '<p><img class=\"wscnph\" src=\"http://localhost:8090/showimage?imgname=upload/20200413/8GMq0TEmC7rZcXLrDoSM0gNEYsUcdc6z.png\" />asdfasdf<img class=\"wscnph\" src=\"http://localhost:8090/showimage?imgname=upload/20200413/CBgqld3ZFuR4uBs3w9XJM9xPtQlhjflT.png\" /></p>', 'fasd', '', 1586707200, '20200413/yjJiBlZMs42Ag1YOWxW6PuuBU3Ut9XSe.png', 0, '2020-04-13 00:00:00', '2020-04-13 15:06:59'),
(4, 1, 2, 1, '测试', '<p><img class=\"wscnph\" src=\"http://localhost:8090/showimage?imgname=upload/20200413/y0hBc2CFylSBcxIFOFhFkloXewmrViVg.png\" />asdfasdf<img class=\"wscnph\" src=\"http://localhost:8090/showimage?imgname=upload/20200413/wXVgmVtr5WLdV6BMR6frCOkqAU5XElW9.png\" /></p>', 'fasd', '', 1586707200, '20200413/SRiwahWF5RqPZWMrRUeT3CmX5Bxveyko.png', 0, '2020-04-13 00:00:00', '2020-04-12 16:00:00'),
(5, 1, 3, 1, 'fasd', '<p>fasdfasdf</p>', 'fasd', '', 1588780800, '20200515/KRUo3oEEZt51nh3ygS02Fc9JWvMb8ZuM.jpeg', 0, '2020-05-07 00:00:00', '2020-05-06 16:00:00');

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
  `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '时间'
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

--
-- 转存表中的数据 `system_role`
--

INSERT INTO `system_role` (`id`, `name`, `alias_name`, `description`, `state`, `type`, `ctime`) VALUES
(1, 'admin', 'admin', '超级管理员具有所有权限', 1, 1, '2019-11-07 16:22:29'),
(2, 'editor', 'editor', '运营者', 1, 1, '2019-11-07 16:22:29'),
(3, 'normal', 'normal', '普通管理员', 0, 0, '0001-01-01 00:00:00');

-- --------------------------------------------------------

--
-- 表的结构 `system_role_menu`
--

CREATE TABLE `system_role_menu` (
  `id` int NOT NULL COMMENT '主键',
  `system_role_id` int NOT NULL DEFAULT '0' COMMENT '角色主键',
  `system_menu_id` int NOT NULL DEFAULT '0' COMMENT '菜单主键'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色与菜单关联表';

--
-- 转存表中的数据 `system_role_menu`
--

INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`) VALUES
(14, 1, 1),
(15, 1, 2),
(19, 1, 3),
(20, 1, 26),
(16, 1, 27),
(17, 1, 28),
(18, 1, 29),
(21, 1, 30),
(22, 1, 31),
(23, 1, 32),
(24, 1, 33),
(25, 1, 34),
(26, 1, 35);

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
  `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `last_login_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上次登录时间',
  `last_login_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '最近登录IP',
  `ctime` datetime NOT NULL COMMENT '注册时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理账户表';

--
-- 转存表中的数据 `system_user`
--

INSERT INTO `system_user` (`id`, `name`, `nickname`, `password`, `salt`, `phone`, `avatar`, `introduction`, `state`, `utime`, `last_login_time`, `last_login_ip`, `ctime`) VALUES
(1, 'admin', 'admin', '297f8efd64f95e37a7d792d926a7b5db47c58403', 'MbBQ', '11111111111', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '', 1, '2019-11-01 07:02:33', '0001-01-01 00:00:00', '', '2019-10-24 20:20:34'),
(3, 'admin1', 'admin1', '297f8efd64f95e37a7d792d926a7b5db47c58403', 'MbBQ', '11111111111', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '', 1, '2020-03-18 15:15:55', '0000-00-00 00:00:00', '', '2019-10-24 20:20:34'),
(4, 'admin2', 'admin12', '297f8efd64f95e37a7d792d926a7b5db47c58403', 'MbBQ', '11111111111', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '', 1, '2020-03-18 15:16:01', '0000-00-00 00:00:00', '', '2019-10-24 20:20:34'),
(5, 'admin3', 'admin123', '297f8efd64f95e37a7d792d926a7b5db47c58403', 'MbBQ', '11111111111', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '', 1, '2020-03-18 15:16:06', '0000-00-00 00:00:00', '', '2019-10-24 20:20:34');

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
-- 转存表中的数据 `system_user_role`
--

INSERT INTO `system_user_role` (`id`, `system_user_id`, `system_role_id`) VALUES
(4, 1, 1),
(5, 1, 2);

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
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=6;

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
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `system_role_menu`
--
ALTER TABLE `system_role_menu`
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=27;

--
-- 使用表AUTO_INCREMENT `system_user`
--
ALTER TABLE `system_user`
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `system_user_role`
--
ALTER TABLE `system_user_role`
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
