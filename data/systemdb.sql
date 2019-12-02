-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 2019-12-02 01:29:52
-- 服务器版本： 5.5.53
-- PHP Version: 7.2.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `systemdb`
--

-- --------------------------------------------------------

--
-- 表的结构 `system_log`
--

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
(1, '', '/redirect', 'layout/Layout', '', '', '', '', 0, 0, 0, 1, 1, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(2, '', '/redirect/:path*', 'views/redirect/index', '', '', '', '', 0, 0, 0, 1, 0, 1, 0, 1, 0, '2019-11-07 16:22:29'),
(3, '', '/login', 'views/login/index', '', '', '', '', 0, 0, 0, 1, 1, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(4, '', '/auth-redirect', 'views/login/auth-redirect', '', '', '', '', 0, 0, 0, 1, 1, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(5, '', '/404', 'views/error-page/404', '', '', '', '', 0, 0, 0, 1, 1, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(6, '', '/401', 'views/error-page/401', '', '', '', '', 0, 0, 0, 1, 1, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(7, '', '', 'layout/Layout', 'dashboard', '', '', '', 0, 0, 0, 1, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(8, 'Dashboard', 'dashboard', 'views/dashboard/index', '', '', 'dashboard', '', 0, 0, 0, 1, 0, 7, 0, 1, 0, '2019-11-07 16:22:29'),
(9, '', '/documentation', 'layout/Layout', '', '', '', '', 0, 0, 0, 1, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(10, 'Documentation', 'index', 'views/documentation/index', '', '', 'documentation', '', 0, 0, 0, 1, 0, 9, 0, 1, 0, '2019-11-07 16:22:29'),
(11, '', '/guide', 'layout/Layout', '/guide/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(12, 'Guide', 'index', 'views/guide/index', '', '', 'guide', '', 0, 0, 0, 1, 0, 11, 0, 1, 0, '2019-11-07 16:22:29'),
(13, '', '/permission', 'layout/Layout', '/permission/index', '', 'permission', 'lock', 0, 1, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(14, 'PagePermission', 'page', 'views/permission/page', '', '', 'pagePermission', '', 0, 0, 0, 2, 0, 13, 0, 1, 0, '2019-11-07 16:22:29'),
(15, 'DirectivePermission', 'directive', 'views/permission/directive', '', '', 'directivePermission', '', 0, 0, 0, 2, 0, 13, 0, 1, 0, '2019-11-07 16:22:29'),
(16, 'RolePermission', 'role', 'views/permission/role', '', '', 'rolePermission', '', 0, 0, 0, 2, 0, 13, 0, 1, 0, '2019-11-07 16:22:29'),
(17, '', '/icon', 'layout/Layout', '', '', '', '', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(18, 'Icons', 'index', 'views/icons/index', '', '', 'icons', '', 0, 0, 0, 2, 0, 17, 0, 1, 0, '2019-11-07 16:22:29'),
(19, '', '/components', 'layout/Layout', 'noRedirect', '', 'components', 'component', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(20, 'TinymceDemo', 'tinymce', 'views/components-demo/tinymce', '', '', 'tinymce', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(21, 'MarkdownDemo', 'markdown', 'views/components-demo/markdown', '', '', 'markdown', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(22, 'JsonEditorDemo', 'json-editor', 'views/components-demo/json-editor', '', '', 'jsonEditor', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(23, 'SplitpaneDemo', 'split-pane', 'views/components-demo/split-pane', '', '', 'splitPane', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(24, 'AvatarUploadDemo', 'avatar-upload', 'views/components-demo/avatar-upload', '', '', 'avatarUpload', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(25, 'DropzoneDemo', 'dropzone', 'views/components-demo/dropzone', '', '', 'dropzone', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(26, 'StickyDemo', 'sticky', 'views/components-demo/sticky', '', '', 'sticky', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(27, 'CountToDemo', 'count-to', 'views/components-demo/count-to', '', '', 'countTo', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(28, 'ComponentMixinDemo', 'mixin', 'views/components-demo/mixin', '', '', 'componentMixin', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(29, 'BackToTopDemo', 'back-to-top', 'views/components-demo/back-to-top', '', '', 'backToTop', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(30, 'DragDialogDemo', 'drag-dialog', 'views/components-demo/drag-dialog', '', '', 'dragDialog', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(31, 'DragSelectDemo', 'drag-select', 'views/components-demo/drag-select', '', '', 'dragSelect', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(32, 'DndListDemo', 'dnd-list', 'views/components-demo/dnd-list', '', '', 'dndList', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(33, 'DragKanbanDemo', 'drag-kanban', 'views/components-demo/drag-kanban', '', '', 'dragKanban', '', 0, 0, 0, 2, 0, 19, 0, 1, 0, '2019-11-07 16:22:29'),
(34, '', '/charts', 'layout/Layout', 'noRedirect', '', 'charts', 'chart', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(35, 'KeyboardChart', 'keyboard', 'views/charts/keyboard', '', '', 'keyboardChart', '', 0, 0, 0, 2, 0, 34, 0, 1, 0, '2019-11-07 16:22:29'),
(36, 'LineChart', 'line', 'views/charts/line', '', '', 'lineChart', '', 0, 0, 0, 2, 0, 34, 0, 1, 0, '2019-11-07 16:22:29'),
(37, 'MixChart', 'mixchart', 'views/charts/mixChart', '', '', 'mixChart', '', 0, 0, 0, 2, 0, 34, 0, 1, 0, '2019-11-07 16:22:29'),
(38, '', '/nested', 'layout/Layout', '/nested/menu1/menu1-1', '', 'nested', 'nested', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(39, 'Menu1', 'menu1', 'views/nested/menu1/index', '', '', 'menu1', '', 0, 0, 0, 2, 0, 38, 0, 1, 0, '2019-11-07 16:22:29'),
(40, 'Menu2', 'menu2', 'views/nested/menu2/index', '', '', 'menu2', '', 0, 0, 0, 2, 0, 38, 0, 1, 0, '2019-11-07 16:22:29'),
(41, '', '/example', 'layout/Layout', '/example/list', '', 'example', 'example', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(42, 'CreateArticle', 'create', 'views/example/create', '', '', 'createArticle', '', 0, 0, 0, 2, 0, 41, 0, 1, 0, '2019-11-07 16:22:29'),
(43, 'EditArticle', 'edit/:id(\\d+)', 'views/example/edit', '', '', 'editArticle', '', 0, 0, 0, 2, 0, 41, 0, 1, 0, '2019-11-07 16:22:29'),
(44, 'ArticleList', 'list', 'views/example/list', '', '', 'articleList', '', 0, 0, 0, 2, 0, 41, 0, 1, 0, '2019-11-07 16:22:29'),
(45, '', '/tab', 'layout/Layout', '', '', '', '', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(46, 'Tab', 'index', 'views/tab/index', '', '', 'tab', '', 0, 0, 0, 2, 0, 45, 0, 1, 0, '2019-11-07 16:22:29'),
(47, '', '/error', 'layout/Layout', 'noRedirect', '', 'errorPages', '404', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(48, 'Page401', '401', 'views/error-page/401', '', '', 'page401', '', 0, 0, 0, 2, 0, 47, 0, 1, 0, '2019-11-07 16:22:29'),
(49, 'Page404', '404', 'views/error-page/404', '', '', 'page404', '', 0, 0, 0, 2, 0, 47, 0, 1, 0, '2019-11-07 16:22:29'),
(50, '', '/error-log', 'layout/Layout', 'noRedirect', '', '', '', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(51, 'ErrorLog', 'log', 'views/error-log/index', '', '', 'errorLog', '', 0, 0, 0, 2, 0, 50, 0, 1, 0, '2019-11-07 16:22:29'),
(52, '', '/excel', 'layout/Layout', '/excel/export-excel', '', 'excel', 'excel', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(53, 'ExportExcel', 'export-excel', 'views/excel/export-excel', '', '', 'exportExcel', '', 0, 0, 0, 2, 0, 52, 0, 1, 0, '2019-11-07 16:22:29'),
(54, 'SelectExcel', 'export-selected-excel', 'views/excel/select-excel', '', '', 'selectExcel', '', 0, 0, 0, 2, 0, 52, 0, 1, 0, '2019-11-07 16:22:29'),
(55, 'MergeHeader', 'export-merge-header', 'views/excel/merge-header', '', '', 'mergeHeader', '', 0, 0, 0, 2, 0, 52, 0, 1, 0, '2019-11-07 16:22:29'),
(56, 'UploadExcel', 'upload-excel', 'views/excel/upload-excel', '', '', 'uploadExcel', '', 0, 0, 0, 2, 0, 52, 0, 1, 0, '2019-11-07 16:22:29'),
(57, '', '/zip', 'layout/Layout', '/zip/download', '', 'zip', 'zip', 0, 1, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(58, 'ExportZip', 'download', 'views/zip/index', '', '', 'exportZip', '', 0, 0, 0, 2, 0, 57, 0, 1, 0, '2019-11-07 16:22:29'),
(59, '', '/pdf', 'layout/Layout', '/pdf/index', '', '', '', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(60, 'PDF', 'index', 'views/pdf/index', '', '', 'pdf', '', 0, 0, 0, 2, 0, 59, 0, 1, 0, '2019-11-07 16:22:29'),
(61, '', '/pdf/download', 'views/pdf/download', '', '', '', '', 0, 0, 0, 2, 1, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(62, '', '/theme', 'layout/Layout', 'noRedirect', '', '', '', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(63, 'Theme', 'index', 'views/theme/index', '', '', 'theme', '', 0, 0, 0, 2, 0, 62, 0, 1, 0, '2019-11-07 16:22:29'),
(64, '', '/clipboard', 'layout/Layout', 'noRedirect', '', '', '', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(65, 'ClipboardDemo', 'index', 'views/clipboard/index', '', '', 'clipboardDemo', '', 0, 0, 0, 2, 0, 64, 0, 1, 0, '2019-11-07 16:22:29'),
(66, '', '/i18n', 'layout/Layout', '', '', '', '', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(67, 'I18n', 'index', 'views/i18n-demo/index', '', '', 'i18n', '', 0, 0, 0, 2, 0, 66, 0, 1, 0, '2019-11-07 16:22:29'),
(68, '', 'external-link', 'layout/Layout', '', '', '', '', 0, 0, 0, 2, 0, 0, 0, 1, 0, '2019-11-07 16:22:29'),
(69, '', 'https://github.com/PanJiaChen/vue-element-admin', '', '', '', 'externalLink', '', 0, 0, 0, 2, 0, 68, 0, 1, 0, '2019-11-07 16:22:29'),
(70, '', '*', '', '/404', '', '', '', 0, 0, 0, 3, 1, 0, 0, 1, 0, '2019-11-07 16:22:29');

-- --------------------------------------------------------

--
-- 表的结构 `system_role`
--

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

CREATE TABLE `system_role_menu` (
  `id` int(11) NOT NULL COMMENT '主键',
  `system_role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色主键',
  `system_menu_id` int(11) NOT NULL DEFAULT '0' COMMENT '菜单主键'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色与菜单关联表';

--
-- 转存表中的数据 `system_role_menu`
--

INSERT INTO `system_role_menu` (`id`, `system_role_id`, `system_menu_id`) VALUES
(88, 1, 13),
(89, 1, 14),
(90, 1, 16),
(91, 1, 17),
(92, 1, 18),
(93, 1, 19),
(94, 1, 20),
(100, 2, 13),
(101, 2, 17),
(102, 2, 18),
(95, 3, 17),
(96, 3, 18);

-- --------------------------------------------------------

--
-- 表的结构 `system_user`
--

CREATE TABLE `system_user` (
  `id` int(11) NOT NULL COMMENT '主键',
  `name` varchar(50) NOT NULL COMMENT '姓名',
  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '用户登录名',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `salt` varchar(4) NOT NULL COMMENT '盐',
  `phone` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(300) NOT NULL DEFAULT '' COMMENT '头像',
  `introduction` varchar(300) NOT NULL DEFAULT '' COMMENT '简介',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态（0 停止1启动）',
  `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `last_login_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '上次登录时间',
  `last_login_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '最近登录IP',
  `ctime` datetime NOT NULL COMMENT '注册时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理账户表';

--
-- 转存表中的数据 `system_user`
--

INSERT INTO `system_user` (`id`, `name`, `nickname`, `password`, `salt`, `phone`, `avatar`, `introduction`, `status`, `utime`, `last_login_time`, `last_login_ip`, `ctime`) VALUES
(1, 'admin', 'admin', '297f8efd64f95e37a7d792d926a7b5db47c58403', 'MbBQ', '11111111111', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '', 1, '2019-11-01 07:02:33', '0000-00-00 00:00:00', '', '2019-10-24 20:20:34');

-- --------------------------------------------------------

--
-- 表的结构 `system_user_role`
--

CREATE TABLE `system_user_role` (
  `id` int(11) NOT NULL COMMENT '主键',
  `system_user_id` int(11) NOT NULL COMMENT '用户主键',
  `system_role_id` int(11) NOT NULL COMMENT '角色主键'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账户和角色关联表';

--
-- 转存表中的数据 `system_user_role`
--

INSERT INTO `system_user_role` (`id`, `system_user_id`, `system_role_id`) VALUES
(1, 1, 1),
(2, 1, 2),
(3, 1, 3);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `system_log`
--
ALTER TABLE `system_log`
  ADD PRIMARY KEY (`id`),
  ADD KEY `RELATION_ID` (`relation_id`),
  ADD KEY `SYSTEM_USER_ID` (`system_user_id`),
  ADD KEY `CTIME` (`ctime`),
  ADD KEY `RELATION_TABLE` (`relation_table`);

--
-- Indexes for table `system_menu`
--
ALTER TABLE `system_menu`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_list` (`pid`,`sort`,`status`) USING BTREE,
  ADD KEY `path` (`path`);

--
-- Indexes for table `system_role`
--
ALTER TABLE `system_role`
  ADD PRIMARY KEY (`id`),
  ADD KEY `TYPE` (`type`),
  ADD KEY `STATUS` (`status`);

--
-- Indexes for table `system_role_menu`
--
ALTER TABLE `system_role_menu`
  ADD PRIMARY KEY (`id`),
  ADD KEY `system_role_id` (`system_role_id`,`system_menu_id`);

--
-- Indexes for table `system_user`
--
ALTER TABLE `system_user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `NICKNAME` (`nickname`),
  ADD KEY `PASSWORD` (`password`);

--
-- Indexes for table `system_user_role`
--
ALTER TABLE `system_user_role`
  ADD PRIMARY KEY (`id`),
  ADD KEY `system_user_id` (`system_user_id`,`system_role_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `system_log`
--
ALTER TABLE `system_log`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `system_menu`
--
ALTER TABLE `system_menu`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=71;

--
-- 使用表AUTO_INCREMENT `system_role`
--
ALTER TABLE `system_role`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `system_role_menu`
--
ALTER TABLE `system_role_menu`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=103;

--
-- 使用表AUTO_INCREMENT `system_user`
--
ALTER TABLE `system_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `system_user_role`
--
ALTER TABLE `system_user_role`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
