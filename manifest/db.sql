USE `gf`;

create Table menu (
    mid int AUTO_INCREMENT COMMENT "唯一标识",
    parent_mid int COMMENT "父节点id",
    id VARCHAR(20) NOT NULL COMMENT "id",
    title VARCHAR(30) COMMENT "标题",
    icon VARCHAR(255) COMMENT "图标",
    type int COMMENT "类型",
    open_type varchar(255) COMMENT "展开类型",
    href VARCHAR(255) COMMENT "链接",
    PRIMARY KEY (`mid`)
) engine innodb DEFAULT charset=utf8mb4;


/*Table structure for table `book` */
drop table `user`;
CREATE TABLE `user` (
    `id` INT(10) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name` VARCHAR(50) NOT NULL COMMENT '用户名称',
    `age` int(10) NOT NULL COMMENT '用户年龄',
    `gender` BOOLEAN default 1 COMMENT '性别',
    PRIMARY KEY (`id`)
) engine=innodb DEFAULT charset=utf8;


CREATE TABLE `role` (
    `id` INT(10) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name` VARCHAR(50) NOT NULL COMMENT '角色名称',
    `information` VARCHAR(50) COMMENT '角色信息',
    PRIMARY KEY (`id`)
) engine=innodb DEFAULT charset=utf8;

CREATE TABLE `permission` (
    `id` INT(10) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name` VARCHAR(50) NOT NULL COMMENT '权限名称',
    `information` VARCHAR(50) COMMENT '权限信息',
    PRIMARY KEY (`id`)
) engine=innodb DEFAULT charset=utf8;

CREATE Table `many_role_permission` (
  `id` INT(10) NOT NULL AUTO_INCREMENT COMMENT "id",
  `rid` INT(10) COMMENT "角色 id",
  `pid` INT(10) COMMENT "权限 id",
  PRIMARY KEY (`id`)
) engine = innodb DEFAULT charset = utf8;

DROP TABLE IF EXISTS `book`;

CREATE TABLE `book` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` VARCHAR(50) NOT NULL COMMENT '书名',
  `author` VARCHAR(30) NOT NULL COMMENT '作者',
  `price` DOUBLE NOT NULL COMMENT '价格',
  `publish_time` DATE COMMENT '出版时间',
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

/*Data for the table `book` */

INSERT  INTO `book`(`id`,`name`,`author`,`price`) VALUES 
(1,'MySQL数据库从入门到精通','王飞飞',59.8),
(2,'设计模式','刘伟',45),
(3,'数据库原理及应用','刘亮',33),
(4,'Linux驱动开发入门与实践','郑强',69),
(5,'Linux驱动开发入门与实践','郑强',69),
(6,'Linux驱动开发入门与实践','郑强',69);

/*Table structure for table `dept` */

DROP TABLE IF EXISTS `dept`;

CREATE TABLE `dept` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` INT(10) UNSIGNED DEFAULT NULL COMMENT '上级部门ID',
  `name` VARCHAR(30) DEFAULT NULL COMMENT '部门名称',
  `leader` VARCHAR(20) DEFAULT NULL COMMENT '部门领导',
  `phone` VARCHAR(11) DEFAULT NULL COMMENT '联系电话',
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8;

/*Data for the table `dept` */

INSERT  INTO `dept`(`id`,`pid`,`name`,`leader`,`phone`) VALUES 
(100,0,'哪都通','赵方旭','10000000000'),
(101,100,'华北大区','徐四','10000000001'),
(102,100,'东北大区','高廉','10000000002'),
(103,100,'华东大区','窦乐','10000000003'),
(104,100,'华中大区','任菲','10000000004'),
(105,100,'华南大区',NULL,NULL),
(106,100,'西北大区','华风','10000000005'),
(107,100,'西南大区','郝意','10000000006');

/*Table structure for table `emp` */

DROP TABLE IF EXISTS `emp`;

CREATE TABLE `emp` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `dept_id` INT(10) UNSIGNED NOT NULL COMMENT '所属部门',
  `name` VARCHAR(30) NOT NULL COMMENT '姓名',
  `gender` TINYINT(1) DEFAULT NULL COMMENT '性别: 0=男 1=女',
  `phone` VARCHAR(11) DEFAULT NULL COMMENT '联系电话',
  `email` VARCHAR(50) DEFAULT NULL COMMENT '邮箱',
  `avatar` VARCHAR(100) DEFAULT NULL COMMENT '照片',
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

/*Data for the table `emp` */

INSERT  INTO `emp`(`id`,`dept_id`,`name`,`gender`,`phone`,`email`) VALUES 
(1,100,'赵方旭',0,'10000000000','zhaofx@nadoutong.com'),
(2,100,'毕游龙',0,'10000000007','biyoulong@nadoutong.com'),
(3,100,'黄伯仁',0,'10000000008','huangboren@nadoutong.com'),
(4,101,'徐四',0,'10000000001','xusi@nadoutong.com'),
(5,101,'徐三',0,'10000000009','xusan@nadoutong.com'),
(6,101,'冯宝宝',1,'10000000010','fengbaobao@nadoutong.com'),
(7,101,'张楚岚',0,'10000000011','zhangchulan@nadoutong.com'),
(8,102,'高廉',0,'10000000002','gaolian@nadoutong.com'),
(9,102,'高二壮',1,'10000000012','gaoerzhuang@nadoutong.com'),
(10,103,'窦乐',0,'10000000003','doule@nadoutong.com'),
(11,103,'肖自在',0,'10000000013','xiaozizai@nadoutong.com'),
(12,104,'任菲',0,'10000000004','renfei@nadoutong.com'),
(13,106,'华风',0,'10000000005','huafeng@nadoutong.com'),
(14,107,'郝意',0,'10000000006','huafeng@nadoutong.com');

DROP TABLE IF EXISTS `hobby`;

CREATE TABLE `hobby` (  
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `emp_id` INT UNSIGNED NOT NULL COMMENT 'EmpID',
  `hobby` VARCHAR(50) COMMENT '爱好',
  PRIMARY KEY (`id`) 
) ENGINE=INNODB CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO `hobby` (`id`, `emp_id`, `hobby`) VALUES
(1, 6, '埋人'),
(2, 4, '看美女'),
(3, 7, '月下遛鸟');

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (  
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` VARCHAR(20) NOT NULL COMMENT '用户名',
  `nickname` VARCHAR(30) COMMENT '昵称',
  `password` VARCHAR(32) COMMENT '密码',
  `avatar` VARCHAR(100) COMMENT '头像',
  `created_at` DATETIME COMMENT '创建时间',
  PRIMARY KEY (`id`) 
) ENGINE=INNODB CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO 
`user` (`id`, `username`, `nickname`, `password`, `avatar`, `created_at`) VALUES
(1, 'libai', '李白', '123456', '', '2023-10-08 16:57:24'),
(2, 'dufu', '杜甫', '123456', '', '2023-10-08 16:57:24'),
(3, 'baijuyi', '白居易', '123456', '', '2023-10-08 16:57:24');