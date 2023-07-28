select *,
       case
           when score >= 80 and course = '数学' then '优秀学生'
           when score >= 90 and course = '语文' then '三好学生'
           else '坏学生'
           end '称号'
from (select c.course, c.score, b.classname, a.stuname
      from A a
               join B b
               join C c
                    on a.classid = b.classid and a.stuid = c.stuid) view;
