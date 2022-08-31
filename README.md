# users micro service by grpc

#  deploy

## get the code

```
git clone https://github.com/kalaGN/usermservice
```

## edit the config.ini

```
cd usermservice/ && cp config.example config.ini
```
config the port you want to start it,default value is 8322
```
server.port=8322
```
config the database link 
```
database.host=127.0.0.1
database.port=3306
database.user=root
database.password=pwd123
database.dbname=test
```


## start  service by dokcer

```
docker build -t usermservice:v1 .
```
```
docker run  -p 8322:8322 --name ums  usermservice:v1
```

## connect test by client
```
cd usermservice/client/v1/ && go run client.go
```

# struct of table
```
CREATE TABLE `accounts` (
  `accountid` int(10) NOT NULL AUTO_INCREMENT COMMENT '自增账号编号',
  `username` varchar(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '密码',
  `corpid` int(10) NOT NULL COMMENT '租户编号',
  `createTime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `lastLogin` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `isdelete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`accountid`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='账号表';
```