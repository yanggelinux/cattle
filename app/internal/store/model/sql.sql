create
database if not exists cattle DEFAULT CHARSET utf8 COLLATE utf8_general_ci;


CREATE TABLE IF NOT EXISTS `test`
(
    `id`           bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`         varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
    `data`         json COMMENT '数据',
    `is_deleted`   tinyint(4)  NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time` datetime    NOT NULL COMMENT '删除时间',
    `updated_time` datetime    NOT NULL COMMENT '更新时间',
    `created_time` datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_name_isdel_dtime` (`name`, `is_deleted`, `deleted_time`),
    KEY            `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '测试表';


-- 用户权限相关
CREATE TABLE IF NOT EXISTS `user`
(
    `id`              bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_name`       varchar(50)  NOT NULL DEFAULT '' COMMENT '用户名',
    `password`        varchar(50)  NOT NULL DEFAULT '' COMMENT '密码',
    `email`           varchar(150) NOT NULL DEFAULT '' COMMENT '邮箱',
    `display_name`    varchar(50)  NOT NULL DEFAULT '' COMMENT '展示名',
    `dept_name`       varchar(50)  NOT NULL DEFAULT '' COMMENT '部门',
    `origin`          tinyint(4)   NOT NULL DEFAULT 1 COMMENT '1 系统用户 2 ldap用户',
    `is_deleted`      tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `last_login_time` datetime COMMENT '最后登录时间',
    `deleted_time`    datetime     NOT NULL COMMENT '删除时间',
    `updated_time`    datetime     NOT NULL COMMENT '更新时间',
    `created_time`    datetime     NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY               `idx_created_time` (`created_time`),
    UNIQUE KEY `uniq_uname_isdel_dtime` (`user_name`, `is_deleted`, `deleted_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';


CREATE TABLE IF NOT EXISTS `role`
(
    `id`           bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `role_name`    varchar(50) NOT NULL DEFAULT '' COMMENT '角色名',
    `display_name` varchar(50) NOT NULL DEFAULT '' COMMENT '展示名',
    `is_super`     tinyint(4)  NOT NULL DEFAULT '0' COMMENT '是否超级管理员 0 否 1 是',
    `is_deleted`   tinyint(4)  NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time` datetime    NOT NULL COMMENT '删除时间',
    `updated_time` datetime    NOT NULL COMMENT '更新时间',
    `created_time` datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_rname_isdel_dtime` (`role_name`, `is_deleted`, `deleted_time`),
    KEY            `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '角色表';

CREATE TABLE IF NOT EXISTS `team`
(
    `id`           bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`         varchar(50) NOT NULL DEFAULT '' COMMENT '团队名称',
    `leader`       varchar(50) NOT NULL DEFAULT '' COMMENT '组长',
    `director`     varchar(50) NOT NULL DEFAULT '' COMMENT '总监',
    `is_deleted`   tinyint(4)  NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time` datetime    NOT NULL COMMENT '删除时间',
    `updated_time` datetime    NOT NULL COMMENT '更新时间',
    `created_time` datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_name_isdel_dtime` (`name`, `is_deleted`, `deleted_time`),
    KEY            `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '团队表';



CREATE TABLE IF NOT EXISTS `user_role_rel`
(
    `id`           bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`      bigint(20) NOT NULL COMMENT '角色',
    `role_id`      bigint(20) NOT NULL COMMENT '角色ID',
    `created_time` datetime NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_role_id` (`role_id`),
    KEY            `idx_user_id` (`user_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户角色关联表';

-- 权限
CREATE TABLE IF NOT EXISTS `permission`
(
    `id`           bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `parent_id`    bigint(20)  NOT NULL DEFAULT '0' COMMENT '父id',
    `name`         varchar(50) NOT NULL DEFAULT '' COMMENT '权限名称',
    `code`         varchar(50) NOT NULL DEFAULT '' COMMENT '权限编码',
    `uri`          varchar(50) NOT NULL DEFAULT '' COMMENT '权限uri',
    `method`       varchar(50) NOT NULL DEFAULT '' COMMENT '方法',
    `project`      varchar(50) NOT NULL DEFAULT '' COMMENT '权限所属项目',
    `perm_type`    tinyint(4)  NOT NULL DEFAULT '1' COMMENT '权限类型 1 菜单 2 api',
    `is_enabled`   tinyint(4)  NOT NULL DEFAULT '1' COMMENT '是否启用 0 否 1 是',
    `sort`         bigint(20)  NOT NULL DEFAULT '0' COMMENT '排序',
    `is_deleted`   tinyint(4)  NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time` datetime    NOT NULL COMMENT '删除时间',
    `updated_time` datetime    NOT NULL COMMENT '更新时间',
    `created_time` datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_name` (`name`),
    KEY            `idx_created_time` (`created_time`),
    UNIQUE KEY `uniq_code_isdel_dtime` (`code`, `is_deleted`, `deleted_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '权限表';


CREATE TABLE IF NOT EXISTS `role_perm_rel`
(
    `id`           bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `role_id`      bigint(20) NOT NULL COMMENT '角色ID',
    `perm_id`      bigint(20) NOT NULL COMMENT '权限ID',
    `created_time` datetime NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_role_id` (`role_id`),
    KEY            `idx_perm_id` (`perm_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '角色权限关联表';

-- 架构图相关
-- 架构组
CREATE TABLE IF NOT EXISTS `arch_group`
(
    `id`           bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `parent_id`    bigint(20)  NOT NULL DEFAULT '0' COMMENT '父id',
    `group_name`   varchar(50) NOT NULL DEFAULT '' COMMENT '图组名称',
    `is_deleted`   tinyint(4)  NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time` datetime    NOT NULL COMMENT '删除时间',
    `updated_time` datetime    NOT NULL COMMENT '更新时间',
    `created_time` datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_created_time` (`created_time`),
    KEY            `idx_parentid` (`parent_id`),
    KEY            `idx_groupname` (`group_name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '流程组表';

-- 架构图
CREATE TABLE IF NOT EXISTS `arch_graph`
(
    `id`           bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `group_id`     bigint(20)  NOT NULL DEFAULT 0 COMMENT '图组id',
    `graph_key`    varchar(50) NOT NULL DEFAULT '' COMMENT '图key',
    `graph_name`   varchar(50) NOT NULL DEFAULT '' COMMENT '图名称',
    `graph_label`  varchar(50) NOT NULL DEFAULT '' COMMENT '图标签',
    `node_data`    json COMMENT '点json',
    `edge_data`    json COMMENT '边json',
    `image_data`   mediumtext COMMENT '存放图片的base64',
    `owner`        varchar(50) NOT NULL DEFAULT '' COMMENT '归属人',
    `status`       tinyint(4)  NOT NULL DEFAULT '1' COMMENT '图状态 0 未审批 1 审批中 2 审批成功 3审批失败 4 审批成功失效',
    `is_shared`    tinyint(4)  NOT NULL DEFAULT '0' COMMENT '是否分享 0 否 1 是',
    `is_deleted`   tinyint(4)  NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time` datetime    NOT NULL COMMENT '删除时间',
    `updated_time` datetime    NOT NULL COMMENT '更新时间',
    `created_time` datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_created_time` (`created_time`),
    KEY            `idx_groupid_name` (`group_id`, `graph_name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '架构图表';

-- 架构图记录表 实现快照功能
CREATE TABLE IF NOT EXISTS `arch_graph_record`
(
    `id`           bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `graph_id`     bigint(20) NOT NULL DEFAULT 0 COMMENT '图id',
    `node_data`    json COMMENT '点json',
    `edge_data`    json COMMENT '边json',
    `image_data`   mediumtext COMMENT '存放图片的base64',
    `record_type`  tinyint(4) NOT NULL DEFAULT '1' COMMENT '记录类型 1 save保存数据 2 sync同步的数据 3 审批通过数据',
    `created_time` datetime NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_graph_id` (`graph_id`),
    KEY            `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '架构图记录表';

CREATE TABLE IF NOT EXISTS `arch_graph_review`
(
    `id`           bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `graph_id`     bigint(20) NOT NULL DEFAULT 0 COMMENT '图id',
    `graph_key`    varchar(50)   NOT NULL DEFAULT '' COMMENT '图key',
    `reviewer`     varchar(50)   NOT NULL DEFAULT '' COMMENT '评审人',
    `notify_party` varchar(50)   NOT NULL DEFAULT '' COMMENT '被通知人',
    `content`      varchar(1000) NOT NULL DEFAULT '' COMMENT '评审内容',
    `created_time` datetime      NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_graph_id` (`graph_id`),
    KEY            `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '架构图评审记录表';

CREATE TABLE IF NOT EXISTS `process_order`
(
    `id`                 bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `graph_id`           bigint(20)   NOT NULL DEFAULT 0 COMMENT '图id',
    `order_id`           bigint(20)   NOT NULL DEFAULT 0 COMMENT '工单id',
    `title`              varchar(50)   NOT NULL DEFAULT '' COMMENT '工单标题',
    `env`                varchar(50)   NOT NULL DEFAULT '' COMMENT '环境',
    `order_name`         varchar(50)   NOT NULL DEFAULT '' COMMENT '工单',
    `graph_name`         varchar(50)   NOT NULL DEFAULT '' COMMENT '图名称',
    `demand_name`        varchar(50)   NOT NULL DEFAULT '' COMMENT '请求名称',
    `owner`              varchar(50)   NOT NULL DEFAULT '' COMMENT '工单所属人',
    `order_field`        json COMMENT '工单需要填的字段',
    `order_info`         json COMMENT '工单需要填的信息',
    `order_process`      json COMMENT '工单流程，显示相关流程',
    `image_hash`         varchar(100)  NOT NULL DEFAULT '' COMMENT '存放图片的base64 hash值',
    `enabled_image_hash` varchar(100)  NOT NULL DEFAULT '' COMMENT '存放已经生效图片的base64 hash值',
    `order_type`         tinyint(4)   NOT NULL DEFAULT '1' COMMENT '工单类型 1 架构图申请工单 2 架构图变更工单 3 请求类资源工单 4 请求类非资源工单 5 非请求工单',
    `description`        varchar(1000) NOT NULL DEFAULT '' COMMENT '描述',
    `task_status`        tinyint(4)   NOT NULL DEFAULT '0' COMMENT '任务执行状态 0 未执行 1 执行中 2 执行成功 3 执行失败',
    `task_result`        json COMMENT '工单任务执行结果',
    `status`             tinyint(4)   NOT NULL DEFAULT '1' COMMENT '审批状态 0 未审批 1 审批中 2 审批成功 3 审批失败',
    `is_deleted`         tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time`       datetime      NOT NULL COMMENT '删除时间',
    `updated_time`       datetime      NOT NULL COMMENT '更新时间',
    `created_time`       datetime      NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY                  `idx_graph_id` (`graph_id`),
    KEY                  `idx_name` (`graph_name`),
    KEY                  `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '流程工单表';


CREATE TABLE IF NOT EXISTS `process_arch`
(
    `id`                 bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `image_hash`         varchar(100) NOT NULL DEFAULT '' COMMENT '存放图片的base64 hash值',
    `enabled_image_hash` varchar(100) NOT NULL DEFAULT '' COMMENT '存放已经生效图片的base64 hash值',
    `image_data`         mediumtext COMMENT '存放图片的base64',
    `enabled_image_data` mediumtext COMMENT '存放已经生效图片的base64',
    `created_time`       datetime     NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY                  `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '流程工单架构信息表';

-- 审批表，记录了审批记录
CREATE TABLE IF NOT EXISTS `process_approval`
(
    `id`           bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `order_id`     bigint(20)  NOT NULL DEFAULT '0' COMMENT '工单id',
    `approver`     varchar(50)  NOT NULL DEFAULT '' COMMENT '审批人',
    `action`       varchar(50)  NOT NULL DEFAULT '' COMMENT '审批动作',
    `opinion`      varchar(500) NOT NULL DEFAULT '' COMMENT '审批意见',
    `status`       tinyint(4)   NOT NULL DEFAULT '1' COMMENT '审批状态 0 未审批 1 审批中 2 审批成功 3 审批失败',
    `created_time` datetime     NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_order_id` (`order_id`),
    KEY            `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '流程审批表';


CREATE TABLE IF NOT EXISTS `process`
(
    `id`           bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`         varchar(50) NOT NULL DEFAULT '' COMMENT '流程名',
    `proc_info`    json COMMENT '流程信息',
    `node_data`    json COMMENT '点json',
    `edge_data`    json COMMENT '边json',
    `status`       tinyint(4)   NOT NULL DEFAULT '1' COMMENT '流程状态 0 失效 1 生效',
    `is_deleted`   tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time` datetime    NOT NULL COMMENT '删除时间',
    `updated_time` datetime    NOT NULL COMMENT '更新时间',
    `created_time` datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_name` (`name`),
    KEY            `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '流程表';


CREATE TABLE IF NOT EXISTS `order`
(
    `id`           bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`         varchar(50) NOT NULL DEFAULT '' COMMENT '工单名',
    `process_id`   bigint(20) NOT NULL DEFAULT '0' COMMENT '流程id',
    `group_id`     bigint(20) NOT NULL DEFAULT '0' COMMENT '工单组id',
    `order_type`   tinyint(4)   NOT NULL DEFAULT '1' COMMENT '工单类型 1 请求类资源工单 2 请求类非资源工单 3 非请求工单',
    `node_type`    varchar(50) NOT NULL DEFAULT '' COMMENT '请求类资源工单关联架构图节点类型',
    `label`        varchar(50) NOT NULL DEFAULT '' COMMENT '标签',
    `layout`       tinyint(4)   NOT NULL DEFAULT '1' COMMENT '字段布局类型 1 单列、2双列、3列、4列',
    `task_url`     varchar(50) NOT NULL DEFAULT '' COMMENT '执行任务的url 把表单发送给对面',
    `task_method`  varchar(50) NOT NULL DEFAULT 'post' COMMENT '执行任务的请求方法 默认post',
    `sort`         bigint(20)  NOT NULL DEFAULT '0' COMMENT '排序',
    `status`       tinyint(4)   NOT NULL DEFAULT '1' COMMENT '工单状态 0 失效 1 生效 ',
    `is_deleted`   tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time` datetime    NOT NULL COMMENT '删除时间',
    `updated_time` datetime    NOT NULL COMMENT '更新时间',
    `created_time` datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_name` (`name`),
    KEY            `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '工单表';

CREATE TABLE IF NOT EXISTS `order_group`
(
    `id`           bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`         varchar(50) NOT NULL DEFAULT '' COMMENT '工单组名',
    `sort`         bigint(20)  NOT NULL DEFAULT '0' COMMENT '排序',
    `status`       tinyint(4)   NOT NULL DEFAULT '1' COMMENT '工单状态 0 失效 1 生效',
    `is_deleted`   tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time` datetime    NOT NULL COMMENT '删除时间',
    `updated_time` datetime    NOT NULL COMMENT '更新时间',
    `created_time` datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_name` (`name`),
    KEY            `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '工单分组表';

CREATE TABLE IF NOT EXISTS `order_field`
(
    `id`            bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `order_id`      bigint(20) NOT NULL DEFAULT '0' COMMENT '工单id',
    `name`          varchar(50)  NOT NULL DEFAULT '' COMMENT '字段名',
    `key`           varchar(50)  NOT NULL DEFAULT '' COMMENT '字段Key',
    `component`     varchar(50)  NOT NULL DEFAULT '' COMMENT '组件类型 输入框 数字输入框 文本域 选择器 选择器多选等',
    `placeholder`   varchar(50)  NOT NULL DEFAULT '' COMMENT '占位',
    `ver_rule`      tinyint(4)   NOT NULL DEFAULT '1' COMMENT '验证规则，1 无规则 2 仅小写字母 3 仅大写字母 4仅包含中文 等',
    `default_val`   varchar(50)  NOT NULL DEFAULT '' COMMENT '默认值',
    `is_required`   tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否必填 0 否 1 是',
    `is_title`      tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否是生成title字段 0 否 1 是',
    `is_edit`       tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否审批可编辑 0 否 1 是',
    `is_clear`      tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否审批不通过清空 0 否 1 是',
    `display_field` varchar(50)  NOT NULL DEFAULT '' COMMENT '根据哪个字段展示',
    `display_val`   varchar(50)  NOT NULL DEFAULT '' COMMENT '根据哪个字段值展示，和display_field一块出现',
    `description`   varchar(500) NOT NULL DEFAULT '' COMMENT '描述',
    `enum`          text COMMENT '枚举值，分为列表 list 形式,字符串分隔',
    `group_name`    varchar(50)  NOT NULL DEFAULT '' COMMENT '分组名称',
    `sort`          bigint(20)  NOT NULL DEFAULT '0' COMMENT '排序',
    `status`        tinyint(4)   NOT NULL DEFAULT '1' COMMENT '字段状态 0 失效 1 生效',
    `is_deleted`    tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time`  datetime     NOT NULL COMMENT '删除时间',
    `updated_time`  datetime     NOT NULL COMMENT '更新时间',
    `created_time`  datetime     NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY             `idx_order_id` (`order_id`),
    KEY             `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '工单字段表';

CREATE TABLE IF NOT EXISTS `demand`
(
    `id`             bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`           varchar(50)   NOT NULL DEFAULT '' COMMENT '需求名称',
    `demand_type`    tinyint(4)   NOT NULL DEFAULT '1' COMMENT '需求类型  1 常规 2 紧急',
    `order_no`       varchar(50)   NOT NULL DEFAULT '' COMMENT '对应的oa单号',
    `biz`            varchar(50)   NOT NULL DEFAULT '' COMMENT '业务组',
    `owner`          varchar(50)   NOT NULL DEFAULT '' COMMENT '所属人',
    `description`    varchar(1000) NOT NULL DEFAULT '' COMMENT '需求描述',
    `opinion`        varchar(1000) NOT NULL DEFAULT '' COMMENT '评审意见',
    `review_process` json COMMENT '评审流程，显示相关流程',
    `evaluation`     json COMMENT '评价',
    `is_evaluate`    tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否评价 0 否 1 是',
    `status`         tinyint(4)   NOT NULL DEFAULT '1' COMMENT '需求状态 0 未评审 1 评审中 2 评审成功 3 评审失败',
    `is_deleted`     tinyint(4)   NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
    `deleted_time`   datetime      NOT NULL COMMENT '删除时间',
    `updated_time`   datetime      NOT NULL COMMENT '更新时间',
    `created_time`   datetime      NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY              `idx_name` (`name`),
    KEY              `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '需求表';

CREATE TABLE IF NOT EXISTS `demand_approval`
(
    `id`           bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `demand_id`    varchar(50)  NOT NULL DEFAULT '' COMMENT '需求id',
    `approver`     varchar(50)  NOT NULL DEFAULT '' COMMENT '评审人',
    `action`       varchar(50)  NOT NULL DEFAULT '' COMMENT '评审动作',
    `opinion`      varchar(500) NOT NULL DEFAULT '' COMMENT '评审意见',
    `status`       tinyint(4)   NOT NULL DEFAULT '1' COMMENT '评审批状态 0 未评审 1 评审中 2 评审成功 3 评审失败',
    `created_time` datetime     NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_demand_id` (`demand_id`),
    KEY            `idx_created_time` (`created_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT = '需求评审记录表';

--创建管理员角色
insert into role values (1,'admin','管理员',1,0,'2025-01-01 00:00:00','2025-05-01 12:00:00','2025-05-01 12:00:00');
--创建管理员用户
insert into user values (1,'admin','12345678','admin@e-cattle.com','管理员','运维团队',0,'2025-01-01 00:00:00','2025-05-01 12:00:00','2025-05-01 12:00:00');
--将管理员角色关联到管理员用户
insert into user_role_rel values (1,1,1,'2025-05-01 12:00:00');