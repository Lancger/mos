insert into User(`id`, `username`, `nick_name`, `email`, `phone`, `password`, `created_at`) values (1, "admin", "系统管理员", "admin@onething.net", "1234567890", "U0yIiOh+tpIhrquJaTC0dQ==", 0);

INSERT INTO `Permission` VALUES 
    ('delete_ticket','删除工单','工单权限',1,'2019-05-15 03:03:05','2019-05-15 03:03:05',NULL),
    ('send_ticket','派发工单','工单权限',2,'2019-05-16 08:43:08','2019-05-16 08:43:08',NULL),
    ('change_ticket','转发工单','工单权限',3,'2019-05-16 08:56:04','2019-05-16 08:56:04',NULL),
    ('deal_ticket','处理工单','工单权限',4,'2019-05-16 09:17:20','2019-05-16 09:17:20',NULL),
    ('close_ticket','结单','工单权限',5,'2019-05-16 09:26:22','2019-05-16 09:26:22',NULL);
