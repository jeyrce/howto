create table A
(
    stuid   int auto_increment
        primary key,
    classid varchar(32) null,
    stuname varchar(32) null
);

create table B
(
    classid   varchar(32) null,
    classname varchar(32) null,
    id        int auto_increment
        primary key
);

create table C
(
    stuid  int         not null
        primary key,
    course varchar(32) null,
    score  int         null
);

INSERT INTO dev.A (stuid, classid, stuname) VALUES (1, 'A', 'chen');
INSERT INTO dev.A (stuid, classid, stuname) VALUES (2, 'B', 'liu');

INSERT INTO dev.B (classid, classname, id) VALUES ('A', '一班', 1);
INSERT INTO dev.B (classid, classname, id) VALUES ('B', '二班', 2);

INSERT INTO dev.C (stuid, course, score) VALUES (1, '数学', 80);
INSERT INTO dev.C (stuid, course, score) VALUES (2, '语文', 95);
