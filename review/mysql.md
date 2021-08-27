# Tips

```sql
--插入记录，如果记录已存在，并且插入的col2的值比原值更大，则更新col2的值。
insert into table_name(col1,col2) values(1,2) on duplicate key update col2=if(col2>values(col2),col2,values(col2));
insert into table_name(col1,col2) select 1,2 from dual on duplicate key update col2=if(col2>values(col2),col2,values(col2));
-- 截取字符子串并转为数字类型
select cast(substring_index('p123','p',-1) as UNSIGNED),substring('v123',2)+0 from dual;
-- 时间日期相关,now()='2017-11-21 15:58:17'
select now(),current_date,date_sub(current_date,interval 1 day),date_add(current_date,interval 1 day),date_format(date_sub(now(),interval 1 hour),'%Y-%m-%d %H:%i:%S') from dual;
-- 2017-11-21 15:58:17 | 2017-11-21 | 2017-11-20 | 2017-11-22| 2017-11-21 14:58:17
```

# dump and reload

* dump

`mysqldump -h127.0.0.1 -uuser_name --password database_name [table_name] > /tmp/database_name.sql`

* reload

`mysql -h127.0.0.1 -uuser_name -ppasswd -e "source /tmp/database_name.sql" database_name`

* reload

`mysql -h127.0.0.1 -uuser_name --password database_name < /tmp/database_name.sql`

# load data

```sql
load data local infile '/home/script_runner/idfa_history.csv' ignore|replace
into table canon_ioscn_prod.idfa_history 
fields terminated by '\t' 
lines terminated by '\n' 
(@col1,@col2) set appid=@col1,idfa=@col2;
-- (@col1,@col2)代表csv文件中的两列，后面的appid、idfa是mysql表中的两列，一一对应。
```

# 查看各个表总记录数及占用空间大小,单位K

```sql
SELECT table_name, sum(TABLE_ROWS),(sum(DATA_LENGTH)+sum(INDEX_LENGTH))/1024 as num
FROM information_schema.TABLES where TABLE_SCHEMA='databasename' 
group by table_name
order by num desc;
```

# 慢查询

```sql
-- 查看是否开启慢查询日志
show variables like '%slow_query_log%';
-- 开启慢查询日志(重启后会失效)
set global slow_query_log=1;
-- 或者修改my.cnf，增加配置并重启MySQL，该配置会永久生效
slow_query_log =1
slow_query_log_file=/tmp/mysql_slow.log
-- 查看及设置慢查询时间阈值(需重建会话链接)
show variables like '%long_query_time%';
set global long_query_time=4;
-- 查看及设置日志存储方式
-- log_output='FILE'表示将日志存入文件，默认值是'FILE'。
-- log_output='TABLE'表示将日志存入数据库mysql.slow_log表中
-- select * from mysql.slow_log
-- MySQL数据库支持同时两种日志存储方式，配置的时候以逗号隔开log_output='FILE,TABLE'。
show variables like '%log_output%';
set global log_output='TABLE';
-- 设置未使用索引的查询也记录到慢查询日志中
show variables like '%log_queries_not_using_indexes%';
set global log_queries_not_using_indexes=1;
-- 查看有多少条慢查询记录
show global status like '%Slow_queries%';
```



