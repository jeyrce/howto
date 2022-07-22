CREATE TABLE test
(
    id   int auto_increment primary key,
    name varchar(255) null default 'x' comment 'name',
    x    varchar(255) null default '',
    y    integer      NULL default 0,
    z    timestamp    NULL default now()
);


delimiter $$$
create procedure zqtest()
begin
    declare i int default 0;
    set i = 0;
    start transaction;
    while i < 80000
        do
            INSERT INTO test VALUES (0, name = 'x', x = 'ccc', y = 5, z = now());
            set i = i + 1;
        end while;
    commit;
end
$$$
delimiter;
call zqtest();
