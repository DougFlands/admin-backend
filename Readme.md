ADMIN 后端服务
# 简介
一个 golang 基于微服务开发的后台管理系统，并有相应的 k8s 配置文件，对应有 react 开发的前端项目  

包含功能:  
* 用户注册与登录
* 文件上传下载，重命名
* 使用腾讯云OSS存储文件
* 使用kafka作为消息队列
* 使用consul做服务注册中心

# 开发
可能用到以下命令，对proto进行解析  
https://www.jianshu.com/p/1a3f1c3031b5  
```
protoc --micro_out=. *.proto 
protoc --go_out=. *.proto
```

# 表设计
fileserver
``` sql
-- 创建文件表
CREATE TABLE `tbl_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `file_sha1` char(40) NOT NULL DEFAULT '' COMMENT '文件hash',
  `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',
  `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小',
  `file_addr` varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
  `create_at` datetime default NOW() COMMENT '创建日期',
  `update_at` datetime default NOW() on update current_timestamp() COMMENT '更新日期',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态(可用/禁用/已删除等状态)',
  `ext1` int(11) DEFAULT '0' COMMENT '备用字段1',
  `ext2` text COMMENT '备用字段2',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_file_hash` (`file_sha1`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 创建用户表
CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT '用户encoded密码',
  `email` varchar(64) DEFAULT '' COMMENT '邮箱',
  `phone` varchar(128) DEFAULT '' COMMENT '手机号',
  `email_validated` tinyint(1) DEFAULT 0 COMMENT '邮箱是否已验证',
  `phone_validated` tinyint(1) DEFAULT 0 COMMENT '手机号是否已验证',
  `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
  `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间戳',
  `profile` text COMMENT '用户属性',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '账户状态(启用/禁用/锁定/标记删除等)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`user_name`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建用户token表
CREATE TABLE `tbl_user_token` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `user_token` char(40) NOT NULL DEFAULT '' COMMENT '用户登录token',
    PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建用户文件表
CREATE TABLE `tbl_user_file` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL,
  `file_sha1` varchar(64) NOT NULL DEFAULT '' COMMENT '文件hash',
  `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小',
  `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',
  `file_rename` varchar(256) NOT NULL DEFAULT '' COMMENT '重命名文件名',
  `upload_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `last_update` datetime DEFAULT CURRENT_TIMESTAMP 
          ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '文件状态(0正常1已删除2禁用)',
  UNIQUE KEY `idx_user_file` (`user_name`, `file_sha1`),
  KEY `idx_status` (`status`),
  KEY `idx_user_id` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

```

# docker
打包镜像
```
docker build --network host -t account ./ -f ./deploy/docker/account.dockerfile
docker build --network host -t apigw ./ -f ./deploy/docker/apigw.dockerfile
docker build --network host -t dbproxy ./ -f ./deploy/docker/dbproxy.dockerfile
docker build --network host -t transfer ./ -f ./deploy/docker/transfer.dockerfile
docker build --network host -t upload ./ -f ./deploy/docker/upload.dockerfile
```

# K8S
在 `/deploy/k8s` 目录下执行 `kubectl apply -f xx.yalm`  
其中 `config.yaml` 需要自己配置，项目所需 config 存放于 k8s 的字典里
配置示例
```yaml
Kafka:
  Topic: "oss"
  Host: "127.0.0.1:9002"
Consul:
  Addr: "127.0.0.1:8500"
Oss:
  Path: "https://xx.cos.ap-guangzhou.myqcloud.com"
  AccessKeyId: ""
  AccessKeySecret: ""
Mysql:
  Host: "user:pwd@tcp(host)/db?charset=utf8"
Redis:
  Host: "127.0.0.1:30379"
  Passwd: ""
Service:
# 本地用8080，云上用80
  ApiGatewayHost: "0.0.0.0:8080"
  UploadHost: "0.0.0.0:8081"
# 密码混淆字串
PwdSalt: ""
```