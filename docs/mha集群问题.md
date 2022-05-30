```text
# 超管用户
root  uos12345

# 业务用户
uos Uos12345

# 主从数据复制用户
replication replication

# mha 管理账户
mha uos12345

```

### 数据库配置

主节点配置:

```text
vim /etc/my.cnf.d/mysql-server.conf

[mysqld]
datadir=/var/lib/mysql
socket=/var/lib/mysql/mysql.sock
log-error=/var/log/mysql/mysqld.log
pid-file=/run/mysqld/mysqld.pid
# ---------------------------新增配置
server-id=118
# 开启binlog，并指定binlog文件的名称
log_bin=mysql_bin
# 开启relay_log，并指定relay_log文件的名称
relay_log=relay_bin
# 将relaylog的同步内容记录到binlog中
log_slave_updates=on
# 开启GTID复制模式
gtid_mode=ON
enforce_gtid_consistency=1


```

slave-1 节点:

```text
[mysqld]
datadir=/var/lib/mysql
socket=/var/lib/mysql/mysql.sock
log-error=/var/log/mysql/mysqld.log
pid-file=/run/mysqld/mysqld.pid
server-id=117
log_bin=mysql_bin
relay_log=relay_bin
log_slave_updates=on
gtid_mode=ON
enforce_gtid_consistency=1


```

slave-2 节点配置:

```text
[mysqld]
datadir=/var/lib/mysql
socket=/var/lib/mysql/mysql.sock
log-error=/var/log/mysql/mysqld.log
pid-file=/run/mysqld/mysqld.pid
server-id=189
log_bin=mysql_bin
relay_log=relay_bin
log_slave_updates=on
gtid_mode=ON
enforce_gtid_consistency=1


```

三个节点都需要执行:  `systemctl restart mysqld` 重启生效。

### 启动slave出错:

```text
start slave 出现错误:
ERROR 1872 (HY000): Slave failed to initialize relay log info structure from the repository

执行:

reset slave;
start slave;

 成功解决。


```

### Slave_IO_Running: Connecting 问题

```text
# 189 节点执行 show slave status\G; 发现有一个状态异常

Slave_IO_Running: Connecting
Slave_SQL_Running: Yes

# 排查思路:
- 网络连通: ping master 没问题，无丢包，排除
- 复制用户远程连接: mysql -h <master> -u replication -p 连接成功，排除
- iptables -nL 查看防火墙，没问题排除

此时再次仔细查看有问题的slave状态发现一个提示:
Last_IO_Errno: 2061
Last_IO_Error: error connecting to master 'reolication@10.20.15.118:3306' - retry-time: 60 retries: 18 message: Authentication plugin 'caching_sha2_password' reported error: Authentication requires secure connection.


```

重新创建slave的复制用户，然后连接:

```text
set global validate_password.policy=0;

create user 'replication'@'%' identified with mysql_native_password by 'replication';

grant replicaton slave on *.* to 'replication'@'%';

flush privileges;

stop slave;

reset slave;

change master to master_host='10.20.15.118', master_port=3306, master_user='replication', master_password='replication', master_auto_position=1;

start  slave;

show slave status\G;

正常了.

```

