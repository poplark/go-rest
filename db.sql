create table `user` (
  `id` int(10) unsigned not null auto_increment,
  `username` varchar(32) not null default '' comment '用户名',
  `email` varchar(32) not null default '' comment '邮箱',
  `password` varchar(32) not null default '' comment '密码',
  `createdAt` datetime not null default current_timestamp,
  `disableFlag` tinyint not null default 0 comment '是否可用',
  primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='用户表'
