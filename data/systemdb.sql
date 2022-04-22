-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2020-06-20 23:02:14
-- 服务器版本： 8.0.12
-- PHP 版本： 7.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `systemdb`
--

-- --------------------------------------------------------

--
-- 表的结构 `system_article`
--

DROP TABLE IF EXISTS `system_article`;
CREATE TABLE `system_article` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `author` int(10) NOT NULL DEFAULT '0' COMMENT '作者',
  `importance` tinyint(4) UNSIGNED NOT NULL DEFAULT '1' COMMENT '重要级别',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态(0:draft，1:published,10:deleted)',
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `content_short` varchar(500) NOT NULL DEFAULT '' COMMENT '摘要',
  `source_uri` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '来源',
  `ctime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `image_uri` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `comment_disabled` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否展示评论',
  `display_time` datetime NOT NULL COMMENT '发布时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `system_article`
--

INSERT INTO `system_article` (`id`, `author`, `importance`, `status`, `title`, `content`, `content_short`, `source_uri`, `ctime`, `image_uri`, `comment_disabled`, `display_time`, `mtime`) VALUES
(1, 1, 2, 1, 'fasd', '<p>fasdffasdf</p>', 'a发生的发生', '', 1585670400, '20200407/ftxqfEY9EjWGyHS7CFaDWkYYh2SLYUIY.jpeg', 0, '2020-04-10 00:00:00', '2020-04-09 15:14:31'),
(2, 1, 2, 1, 'fasd', '<p>fasdffasdf</p>', 'a', 'https://www.baidu.com/', 1585670400, '20200407/jbsW7XXW5nFcj6rNEydZhs9ma4BerLpB.jpeg', 0, '2020-04-01 00:00:00', '2020-03-31 16:00:00'),
(3, 1, 2, 1, '测试', '<p><img class=\"wscnph\" src=\"http://localhost:8090/showimage?imgname=upload/20200413/8GMq0TEmC7rZcXLrDoSM0gNEYsUcdc6z.png\" />asdfasdf<img class=\"wscnph\" src=\"http://localhost:8090/showimage?imgname=upload/20200413/CBgqld3ZFuR4uBs3w9XJM9xPtQlhjflT.png\" /></p>', 'fasd', '', 1586707200, '20200413/yjJiBlZMs42Ag1YOWxW6PuuBU3Ut9XSe.png', 0, '2020-04-13 00:00:00', '2020-04-13 15:06:59'),
(4, 1, 2, 1, '测试', '<p><img class=\"wscnph\" src=\"http://localhost:8090/showimage?imgname=upload/20200413/y0hBc2CFylSBcxIFOFhFkloXewmrViVg.png\" />asdfasdf<img class=\"wscnph\" src=\"http://localhost:8090/showimage?imgname=upload/20200413/wXVgmVtr5WLdV6BMR6frCOkqAU5XElW9.png\" /></p>', 'fasd', '', 1586707200, '20200413/SRiwahWF5RqPZWMrRUeT3CmX5Bxveyko.png', 0, '2020-04-13 00:00:00', '2020-04-12 16:00:00'),
(5, 1, 3, 1, 'fasd', '<p>fasdfasdf</p>', 'fasd', '', 1588780800, '20200515/KRUo3oEEZt51nh3ygS02Fc9JWvMb8ZuM.jpeg', 0, '2020-05-07 00:00:00', '2020-05-06 16:00:00');

-- --------------------------------------------------------

--
-- 表的结构 `system_log`
--

DROP TABLE IF EXISTS `system_log`;
CREATE TABLE `system_log` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `system_user_id` int(11) DEFAULT '0' COMMENT '主键',
  `title` varchar(300) NOT NULL DEFAULT '' COMMENT '日志标题',
  `content` text COMMENT '日志内容记录SQL',
  `relation_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '相关对应表主键',
  `relation_table` int(4) NOT NULL DEFAULT '1' COMMENT '对应表(1 system_user,2 system_menu,3 system_role)',
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'ip',
  `url` varchar(500) NOT NULL DEFAULT '',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作日志';

-- --------------------------------------------------------

--
-- 表的结构 `system_menu`
--

DROP TABLE IF EXISTS `system_menu`;
CREATE TABLE `system_menu` (
  `id` int(11) NOT NULL COMMENT '主键',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
  `path` varchar(50) NOT NULL DEFAULT '' COMMENT '路径',
  `component` varchar(100) NOT NULL DEFAULT '' COMMENT '组件',
  `redirect` varchar(200) NOT NULL DEFAULT '' COMMENT '重定向',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT '访问url',
  `meta_title` varchar(50) NOT NULL DEFAULT '' COMMENT 'meta标题',
  `meta_icon` varchar(50) NOT NULL DEFAULT '' COMMENT 'meta icon',
  `meta_nocache` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否缓存（1:是 0:否）',
  `alwaysshow` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否总是显示（1:是0：否）',
  `meta_affix` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否加固（1:是0：否）',
  `type` tinyint(4) NOT NULL DEFAULT '2' COMMENT '类型(1:固定,2:权限配置3特殊)',
  `hidden` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否隐藏（0否1是）',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父ID',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态（0禁止1启动）',
  `level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '层级',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限';

--
-- 转存表中的数据 `system_menu`
--

INSERT INTO `system_menu` (`id`, `name`, `path`, `component`, `redirect`, `url`, `meta_title`, `meta_icon`, `meta_nocache`, `alwaysshow`, `meta_affix`, `type`, `hidden`, `pid`, `sort`, `status`, `level`, `ctime`) VALUES
(1, '系统管理', '#', '#', '', '#', '系统管理', 'fafa-adjust', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-12-02 06:14:15'),
(2, '用户管理', '/system/user', '/system/user/index', '', '/user', '用户管理', '#', 0, 0, 0, 2, 0, 1, 0, 1, 1, '2019-12-02 00:00:00'),
(3, '菜单管理', '/system/menu', '/system/menu/index', '', '/menu', '菜单管理', '#', 0, 0, 0, 2, 0, 1, 0, 1, 1, '2019-12-02 00:00:00'),
(26, '角色管理', '/system/role', '/system/role/index', '', '/roles', '角色管理', '#', 0, 1, 0, 0, 0, 1, 0, 1, 1, '2019-12-25 19:44:16'),
(27, '添加用户', '/system/user/create', '/system/user/create/index', '', '/user/create', '添加用户', '#', 0, 1, 0, 0, 0, 2, 0, 1, 2, '2019-12-25 20:43:21'),
(28, '用户列表', '/system/user/list', '/system/user/list/index', '', '/user/index', '用户列表', '#', 0, 0, 0, 0, 0, 2, 0, 1, 2, '2019-12-31 09:16:43'),
(29, '用户编辑', '/system/user/edit/:id(\\d+)', '/system/user/edit/index', '', '/user/edit', '用户编辑', '#', 0, 1, 0, 0, 1, 2, 0, 1, 2, '2019-12-31 09:17:41'),
(30, '内容管理', '#', '#', '', '/article', '内容管理', '#', 0, 1, 0, 0, 0, 0, 0, 1, 0, '2019-12-31 09:49:54'),
(31, '创建文章', '/articles/create', '/articles/create/index', '', '/articles/create', '创建文章', '#', 0, 1, 0, 0, 0, 30, 0, 1, 1, '2019-12-31 09:51:12'),
(32, '文章编辑', '/articles/edit/:id(\\d+)', '/articles/edit/index', '', '/articles/edit', '文章编辑', '#', 0, 1, 0, 0, 1, 30, 0, 1, 1, '2019-12-31 09:51:56'),
(33, '文章列表', '/articles/list', '/articles/list/index', '', '/articles/list', '文章列表', '#', 0, 1, 0, 0, 0, 30, 0, 1, 1, '2019-12-31 09:52:36'),
(34, '上传图片', '/upload/image', '/upload/image', '', '/upload/image', '上传图片', '#', 0, 0, 0, 0, 1, 30, 0, 1, 1, '2020-06-19 18:24:32'),
(35, '删除上传图片', '/del/image', '/del/image', '', '/del/image', '删除上传图片', '#', 0, 0, 0, 0, 1, 30, 0, 1, 1, '2020-06-19 18:26:14');

-- --------------------------------------------------------

--
-- 表的结构 `system_role`
--

DROP TABLE IF EXISTS `system_role`;
CREATE TABLE `system_role` (
  `id` int(11) NOT NULL COMMENT '主键',
  `name` varchar(100) NOT NULL COMMENT '角色名称',
  `alias_name` varchar(50) NOT NULL DEFAULT '' COMMENT '别名',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '角色状态（0无效1有效）',
  `type` int(4) NOT NULL DEFAULT '1' COMMENT '属于哪个应用',
  `ctime` datetime NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色';

--
-- 转存表中的数据 `system_role`
--

INSERT INTO `system_role` (`id`, `name`, `alias_name`, `description`, `status`, `type`, `ctime`) VALUES
(1, 'admin', 'admin', '超级管理员具有所有权限', 1, 1, '2019-11-07 16:22:29'),
(2, 'editor', 'editor', '运营者', 1, 1, '2019-11-07 16:22:29'),
(3, 'normal', 'normal', '普通管理员', 0, 0, '0001-01-01 00:00:00');

-- --------------------------------------------------------

--
-- 表的结构 `system_role_menu`
--

DROP TABLE IF EXISTS `system_role_menu`;
CREATE TABLE `system_role_menu` (
  `id` int(11) NOT NULL COMMENT '主键',
  `system_role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色主键',
  `system_menu_id` int(11) NOT NULL DEFAULT '0' COMMENT '菜单主键'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色与菜单关联表';

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

DROP TABLE IF EXISTS `system_user`;
CREATE TABLE `system_user` (
  `id` int(11) NOT NULL COMMENT '主键',
  `name` varchar(50) CHARACTER SET utf8mb4  NOT NULL COMMENT '登录名',
  `nickname` varchar(50) CHARACTER SET utf8mb4  NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `salt` varchar(4) NOT NULL COMMENT '盐',
  `phone` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(300) NOT NULL DEFAULT '' COMMENT '头像',
  `introduction` varchar(300) NOT NULL DEFAULT '' COMMENT '简介',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态（0 停止1启动）',
  `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `last_login_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上次登录时间',
  `last_login_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '最近登录IP',
  `ctime` datetime NOT NULL COMMENT '注册时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理账户表';

--
-- 转存表中的数据 `system_user`
--

INSERT INTO `system_user` (`id`, `name`, `nickname`, `password`, `salt`, `phone`, `avatar`, `introduction`, `status`, `utime`, `last_login_time`, `last_login_ip`, `ctime`) VALUES
(1, 'admin', 'admin', '297f8efd64f95e37a7d792d926a7b5db47c58403', 'MbBQ', '11111111111', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '', 1, '2019-11-01 07:02:33', '0001-01-01 00:00:00', '', '2019-10-24 20:20:34'),
(3, 'admin1', 'admin1', '297f8efd64f95e37a7d792d926a7b5db47c58403', 'MbBQ', '11111111111', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '', 1, '2020-03-18 15:15:55', '0000-00-00 00:00:00', '', '2019-10-24 20:20:34'),
(4, 'admin2', 'admin12', '297f8efd64f95e37a7d792d926a7b5db47c58403', 'MbBQ', '11111111111', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '', 1, '2020-03-18 15:16:01', '0000-00-00 00:00:00', '', '2019-10-24 20:20:34'),
(5, 'admin3', 'admin123', '297f8efd64f95e37a7d792d926a7b5db47c58403', 'MbBQ', '11111111111', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '', 1, '2020-03-18 15:16:06', '0000-00-00 00:00:00', '', '2019-10-24 20:20:34');

-- --------------------------------------------------------

--
-- 表的结构 `system_user_role`
--

DROP TABLE IF EXISTS `system_user_role`;
CREATE TABLE `system_user_role` (
  `id` int(11) NOT NULL COMMENT '主键',
  `system_user_id` int(11) NOT NULL COMMENT '用户主键',
  `system_role_id` int(11) NOT NULL COMMENT '角色主键'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账户和角色关联表';

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
  ADD KEY `idx_list` (`pid`,`sort`,`status`) USING BTREE,
  ADD KEY `path` (`path`);

--
-- 表的索引 `system_role`
--
ALTER TABLE `system_role`
  ADD PRIMARY KEY (`id`),
  ADD KEY `TYPE` (`type`),
  ADD KEY `STATUS` (`status`);

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
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `system_log`
--
ALTER TABLE `system_log`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `system_menu`
--
ALTER TABLE `system_menu`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=37;

--
-- 使用表AUTO_INCREMENT `system_role`
--
ALTER TABLE `system_role`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `system_role_menu`
--
ALTER TABLE `system_role_menu`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=27;

--
-- 使用表AUTO_INCREMENT `system_user`
--
ALTER TABLE `system_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `system_user_role`
--
ALTER TABLE `system_user_role`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
