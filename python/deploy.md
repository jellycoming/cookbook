# supervisor + nginx + gunicorn

```bash
# supervisor
$ pip install supervisor
$ echo_supervisord_conf > supervisor.conf   # 生成 supervisor 默认配置文件
$ vim supervisor.conf                       # 修改 supervisor 配置文件，添加 gunicorn 进程管理
$ supervisord -c supervisor.conf                             通过配置文件启动supervisor
$ supervisorctl -c supervisor.conf status                    察看supervisor的状态
$ supervisorctl -c supervisor.conf reload                    重新载入 配置文件
$ supervisorctl -c supervisor.conf start [all]|[appname]     启动指定/所有 supervisor管理的程序进程
$ supervisorctl -c supervisor.conf stop [all]|[appname]      关闭指定/所有 supervisor管理的程序进程

# gunicorn
# gunicorn,-w 表示开启多少个 worker，-b 表示 gunicorn 开发的访问地址
$ gunicron -w4 -b0.0.0.0:8000 app:wsgi_app
```

```
# nginx配置
upstream oranges { 
    server 127.0.0.1:8080;
}
server {
        listen 80; ## listen for ipv4; this line is default and implied
        #root /usr/share/nginx/www;
        #index index.html index.htm;
        # Make site accessible from http://localhost/
        server_name orange.com;

        # 精确匹配 /, 主机名后面不能带任何字符串
        # 直接匹配网站根，通过域名访问网站首页比较频繁，使用这个会加速处理
        location = / {
            root /data/web;
            index index.html;
        }

        # 处理静态文件请求
        # 匹配任何以 /static/ 开头的地址，匹配符合以后，停止往下搜索正则，采用这一条。
        location ^~ /static/{
            root /data/web/static;
            access_log off;
            expires max;
        }

        # 匹配以指定字符结尾的请求，不区分大小写
        # 但是所有以 /static/ 开头的静态文件都会被上一条规则处理，达不到这条规则
        location ~* \.(gif|jpg|jpeg|png|css|js|ico)$ {
            root /data/web/static;
        }

        # 因为所有的地址都以 / 开头，所以这条规则将匹配到所有请求
        # 但是正则和最长字符串会优先匹配
        location / { 
            proxy_pass_header Server;    
            proxy_set_header Host $http_host;    
            proxy_set_header X-Real-IP $remote_addr;    
            proxy_set_header X-Scheme $scheme;    
            proxy_pass http://oranges;    
            proxy_next_upstream error;
        }   

}
```



