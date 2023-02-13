-- 建表
create database if not exists `test`;
use test;
drop table if exists tb;
create table tb
(
    `id`    int auto_increment primary key,
    `ip`    varchar(20) null,
    `key`   varchar(20) null,
    `value` int         null
) charset 'utf8';

-- 插入测试数据
INSERT INTO test.tb (id, ip, `key`, value) VALUES (1, '10.0.0.1', 'cpu', 4);
INSERT INTO test.tb (id, ip, `key`, value) VALUES (2, '10.0.0.2', 'cpu', 8);
INSERT INTO test.tb (id, ip, `key`, value) VALUES (3, '10.0.0.1', 'mem', 16);
INSERT INTO test.tb (id, ip, `key`, value) VALUES (4, '10.0.0.2', 'mem', 32);
INSERT INTO test.tb (id, ip, `key`, value) VALUES (5, '10.0.0.1', 'disk', 64);
INSERT INTO test.tb (id, ip, `key`, value) VALUES (6, '10.0.0.2', 'disk', 128);


-- 按格式查询输出
select tb.ip, c.cpu, m.mem, d.disk
from tb
    join (select `ip`, `value` as cpu from tb where `key` = 'cpu') as c
    join (select `ip`, `value` as mem from tb where `key` = 'mem') as m
    join (select `ip`, `value` as disk from tb where `key` = 'disk') as d
where c.ip = tb.ip
  and m.ip = tb.ip
  and d.ip = tb.ip
group by tb.ip
order by ip;
