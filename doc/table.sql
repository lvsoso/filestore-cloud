create database fileserver DEFAULT character set utf8;

CREATE TABLE `tbl_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `file_sha1` char(40) NOT NULL DEFAULT '',
  `file_name` varchar(256) NOT NULL DEFAULT '',
  `file_size` bigint(20) DEFAULT '0',
  `file_addr` varchar(1024) NOT NULL DEFAULT '',
  `create_at` datetime default NOW(),
  `update_at` datetime default NOW() on update current_timestamp(),
  `status` int(11) NOT NULL DEFAULT '0',
  `ext1` int(11) DEFAULT '0',
  `ext2` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_file_hash` (`file_sha1`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT 0,
  `user_name` varchar(64) NOT NULL DEFAULT '',
  `user_pwd` varchar(256) NOT NULL DEFAULT '',
  `email` varchar(64) DEFAULT '',
  `phone` varchar(128) DEFAULT '',
  `signup_at` datetime default NOW(),
  `last_active` datetime default NOW() on update current_timestamp(),
  `profile` TEXT,
  `status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `tbl_user_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT 0,
  `file_sha1` varchar(64) NOT NULL DEFAULT '',
  `file_name` varchar(256) NOT NULL DEFAULT '',
  `channel_id` int(11) NOT NULL DEFAULT 0,
  `upload_at` datetime default NOW(),
  `last_update` datetime default NOW() on update current_timestamp(),
  `status` int(11) NOT NULL DEFAULT '0',
  `ext1` int(11) DEFAULT '0',
  `ext2` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_file` (`user_id`,`file_name`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
