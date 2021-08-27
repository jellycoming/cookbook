# basic

### set
* `set -e`/`set -o errexit` "Exit immediately if a simple command exits with a non-zero status.",也就是说，在"set -e"之后出现的代码，一旦出现了返回值非零，整个脚本就会立即退出。
* `set -u`/`set -o nounset` 当执行时使用到未定义过的变量，则报错并停止执行。
* `set -x`/`set -o xtrace` 执行指令后，会先显示该指令及参数，常用于复杂脚本的调试
* `set -o pipefail` 管道中只要一个子命令失败，整个管道命令就失败，脚本就会终止执行。
```bash
current_dir=$(cd $(dirname $0);pwd)  # 获取当前目录并赋值给变量current_dir
cd $current_dir  # 进入当前目录(使用变量)
```
```bash
# ssh后以$username用户执行svn命令
ssh $host "runuser -c 'svn up $path' $username"
```

# if语句常用条件

* `[ -a FILE ]`  如果 FILE 存在则为真。  
* `[ -b FILE ]`  如果 FILE 存在且是一个块特殊文件则为真。  
* `[ -c FILE ]`  如果 FILE 存在且是一个字特殊文件则为真。  
* `[ -d FILE ]`  如果 FILE 存在且是一个目录则为真。  
* `[ -e FILE ]`  如果 FILE 存在则为真。  
* `[ -f FILE ]`  如果 FILE 存在且是一个普通文件则为真。  
* `[ -g FILE ]`  如果 FILE 存在且已经设置了SGID则为真。  
* `[ -h FILE ]`  如果 FILE 存在且是一个符号连接则为真。  
* `[ -k FILE ]`  如果 FILE 存在且已经设置了粘制位则为真。  
* `[ -p FILE ]`  如果 FILE 存在且是一个名字管道\(F如果O\)则为真。  
* `[ -r FILE ]`  如果 FILE 存在且是可读的则为真。  
* `[ -s FILE ]`  如果 FILE 存在且大小不为0则为真。  
* `[ -t FD   ]`  如果文件描述符 FD 打开且指向一个终端则为真。  
* `[ -u FILE ]`  如果 FILE 存在且设置了SUID \(set user ID\)则为真。  
* `[ -w FILE ]`  如果 FILE 如果 FILE 存在且是可写的则为真。  
* `[ -x FILE ]`  如果 FILE 存在且是可执行的则为真。
* `[ -o OPTIONNAME ]`  如果 shell选项 “OPTIONNAME” 开启则为真。  
* `[ -z STRING ]`  “STRING” 的长度为零则为真。  
* `[ -n STRING ]` or `[ STRING ]`  “STRING” 的长度为非零则为真。  
* `[ STRING1 ==|!=|<|> STRING2 ]`  两个字符串比较。

# Linux控制字符

* `;` 用这个符号连接的命令将被按顺序依次执行。
* `&&` 用这个符号连接的命令,表示当前一个命令执行成功后\(exit code=0\),才继续执行下一个命令。 
* `|` 管道,前一个命令的输出作为后一个命令的输入。
* `&` 以这个字符结尾的命令将在后台执行。
* `>` redirects output to a file, overwriting the file.
* `>>` redirects output to a file appending the redirected output at the end.
* `$#` 参数个数，不包含命令或函数本身。
* `$$` 当前bash shell进程ID。
* `$?` 上一个命令的退出状态。
* `$@` 命令行或函数参数列表，不包括命令本身。
* `$*` 同`$@`，但是是以一个单字符串显示所有向脚本传递的参数。

# Get a list of open ports in Linux

`netstat -lntup`

* `-l` = only services which are listening on some port
* `-n` = show port number, don't try to resolve the service name
* `-t` = tcp ports
* `-u` = udp ports
* `-p` = name of the program

# axel: 贼JB快多线程下载工具

```
# 10个线程下载
axel -n 10 [URL]
```

# cut: 文件及字符串纵向切割

* `-d` 分隔符。
* `-f` 要选择的fields。

`cut -d'' -f1,2 file`  
`echo a.b.c|cut -d'.' -f1,2`

> output:  a.b

# split: 文件拆分

* `-l n` 指定每个输出文件有n行。
* `-b n[KMG]` 指定每个输出文件的字节大小。
* `-d` 使用数字后缀代替字母后缀。

`split -l 1000 -d file out`

> output files: out00,out01,out02...

# date

```
# 昨天,YYYY-MM-DD格式  
date --date="-1 day" +"%Y-%m-%d"
```

```
# 在2017-01-01与2016-01-01之间按月迭代,-I == +"%Y-%m-%d"
day='2017-01-01'
while [ "$day" != '2016-01-01' ]
do
        echo $day
        day=$(date -I -d "$day - 1 month")
done
```

# curl

```
# POST FILE
curl -X POST -F "file=@/tmp/upload.sql" http://ooxx.com/upload
```

# Array loop

```
a=(1 2 3 4)
echo "elements: ${a[@]}"
echo "elements length: ${#a[@]}"
echo "first elements: ${a[0]}"
echo "second elements: ${a[1]}"
for i in ${a[@]}
    do
        echo $i
    done
```

# while loop

```
declare -i i=1
while ((i<=3));do
    echo "do b.sh..."
    bash b.sh
    if [ $? -gt 0 ];then
        let i+=1
        sleep 5
    else
        let i+=5
    fi
done
if [ $i -le 5 ];then
    echo "do b.sh error.."
    exit 1
fi
```

# rsync

```
# rsyncd.conf
#######################
use chroot = no
lock file = /var/run/rsync.lock
log file = /var/log/rsyncd.log
pid file = /var/run/rsyncd.pid
[data_sync]
uid = www 
gid = www 
path = /path/to # 这个目录权限属于上面配置的用户
comment = The folder for data sync
read only = no
auth users = datasync # 客户端使用，服务端不必存在该用户
secrets file = /etc/rsyncd.secrets # datasync:12345
#hosts allow = 10.* 
#hosts deny=*
#######################
# 服务端
$ rsync --daemon
# 客户端 passwd文件内容:12345, -a:归档模式 -z:传输过程中压缩 -v:详细输出
# push
$ rsync -avz --password-file=passwd /data/src/ datasync@server-ip::data_sync
# pull
$ rsync -avzP --password-file=passwd --port=9527 datasync@server-ip::data_sync/file.txt /tmp/
```

# du && sort

```
# 查看各个目录所占磁盘容量，深度为1，单位为M,以排序后的第一列（转为数字）做比较因子
du -m --max-depth=1|sort -nk1
```

# 字符及文件处理

```bash
#!/usr/bin/env bash

# 格式化mysqldump数据
current=$(cd $(dirname $0);pwd)
for i in $(ls ${current})
do
    # 匹配所有以uid_mapping开头的文件
    if [[ ${i} == "uid_mapping"* ]];then
        echo 'format '${i}
        # 匹配包含INSERT INTO的那一行，并依次替换指定符号
        awk '$0 ~ /INSERT INTO/{print $0}' ${i}|sed -e "s/),(/\n/g"|sed -e "s/,/\t/g"|sed -e "s/[';]//g"|sed -e "s/[()]/\n/g"|awk 'NR!=1{print $0}' > ${i}.txt
    fi
done


# 格式化hive load语句
current=$(cd $(dirname $0);pwd)
for i in $(ls ${current})
do
    # 匹配以uid_mapping开头,以.txt结尾的文件
    if [[ ${i} == "uid_mapping"* ]] && [ ${i##*.} == txt ];then
        table=${i%.txt} # 获取文件名.txt之前的部分
        t=${table:19} # 获取table字符串index为19以后的子串
        echo "LOAD DATA LOCAL INPATH '$i' INTO TABLE table_name partition (type='$t');"
    fi
done
```

# find

```
$ find [path...] [expression]
$ find . -name "*.sh"  # 在当前目录及其子目录下搜索所有以.sh结尾的文件
$ find . -size 0 -type f -mtime +5 -exec rm {} \;  # 在当前目录及其子目录下搜索大小为0、更改时间在5日以上d额文件并删除之
$ find . ! -type d -print  # 在当前目录下查找除目录以外的所有类型的文件 
$ find . -size +1000000c –print 在当前目录下查找文件长度大于1M字节的文件
# -mtime -n +n 按照文件的更改时间来查找文件, -n表示文件更改时间距现在n天以内, +n表示文件更改时间距现在n天以前
# -mmin -n +n 单位为分钟
# -type f[普通文件] d[目录] l[符号链接文件] p[管道文件] c[字符设备文件] b[块设备文件]
$ find . -perm 755 -print  # 在当前目录及其子目录下搜索文件权限为755的文件
$ find . -user jellycoming -print  # 在当前目录及其子目录下搜索属主为jellycoming用户的文件
```

# log

```
# shell定义log函数及调用,$$指shell本身PID,$@指所有参数列表
function log(){
    echo "log $$ [$(date +'%Y-%m-%d %H:%M:%S')] $@" >>log.${user}.$(date +%Y-%m-%d).log 2>&1
}

log "hello world" "exec log"
```

# crontab

```
#####################################################################################
SHELL=/bin/bash
PATH=/sbin:/bin:/usr/sbin:/usr/bin
MAILTO=root
HOME=/

# For details see man 4 crontabs

# Example of job definition:
# .---------------- minute (0 - 59)
# |  .------------- hour (0 - 23)
# |  |  .---------- day of month (1 - 31)
# |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
# |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
# |  |  |  |  |
# *  *  *  *  * user-name command to be executed
######################################################################################
*/2   *   *   *   *              command      #每两分钟执行
0   */2   *   *   *              command      #每两小时执行
0 6,12,18   *   *   *            command      #每天6点、12点、18点执行
*/2   14-15   *   *   1,2,3,4,5  command      #周一到周五每天14-15点每隔2分钟执行一次
```

# 检测任务是否通过crontab调用

```
# ps -C [cmd]列出指定命令的状况
function check_crontab(){
    local ppid=$$
    local not_crontab=1
    while  true
    do
        ppid=$(grep PPid /proc/$ppid/status |awk '{print $2}') # 找出父进程ID
        local filter_list=$(ps -o pid,command -C crond|grep "$ppid " )
        if [ -n "$filter_list" ]
        then
           not_crontab=0
        fi 
        if [ "$ppid" -eq 1 ] # 查找到跟进程时退出
        then
            break
        fi
    done
    if [ "$not_crontab" -eq 1 ]
    then
       log "not invoked from crontab, exit, $msg" 
       echo "not invoked from crontab, current_pid: $$, exit, $msg"
       echo "pstree: "
       pstree -p |grep "$$" -A 15 -B 15
       exit 3
   fi
}
```
# 查看显卡占用情况
```
# 每0.2秒刷新
$ watch -n 0.2 "nvidia-smi"
```



