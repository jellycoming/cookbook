tornado是一个基于Epoll的单进程单线程的异步网络IO模型，它在底层用一个IOLoop通过调用Epoll来对多个socket链接（文件描述符）进行监测，一旦有socket链接准备好，  
就对其进行相应的读写操作。在操作某一个请求的时候，整个线程是阻塞的，也就是说我们在某个RequestHandler中进行了阻塞的操作，将会导致整个服务的阻塞。  
如下，如果在处理first reqeust的时候执行了某些阻塞的操作，那么其他客户端在进行second request的时候也会被阻塞。

request1`curl localhost:8001/first` 会sleep10秒，  
request2`curl localhost:8001/second` 会被first请求block10秒。

```py
# coding=utf-8
import time
import tornado.web
import tornado.ioloop
import tornado.gen

class FirstHandler(tornado.web.RequestHandler):
    def get(self):
        self.sleep()
        self.write('handle first request\n')

    def sleep(self):
        time.sleep(10)
        return 'blocking several seconds\n'

class SecondHandler(tornado.web.RequestHandler):
    def get(self):
        self.write('handle second request\n')


if __name__ == '__main__':
    application = tornado.web.Application([
        (r'/first', FirstHandler),
        (r'/second', SecondHandler),
    ], **{'debug': True})
application.listen(8001)
tornado.ioloop.IOLoop.instance().start()
```

如何实现异步非阻塞呢？借助线程或协程。

```py
# 多线程方式

# 基本的思想是将阻塞的任务(sleep)放入单独的线程中跑，这样做request1与request2都不会被block
class FirstHandler(tornado.web.RequestHandler):
    def get(self):
        threading.Thread(target=self.sleep).start()
        self.write('handle first request\n')

    def sleep(self):
        time.sleep(10)
        msg = 'blocking several seconds\n'
        print msg
        return msg

# tornado提供的解决方案(python2.7环境future模块需要安装)
from tornado.concurrent import run_on_executor
from concurrent.futures import ThreadPoolExecutor

class FirstHandler(tornado.web.RequestHandler):
    executor = ThreadPoolExecutor(8)

    def get(self):
        self.sleep()
        self.write('handle first request\n')

    @run_on_executor
    def sleep(self):
        time.sleep(10)
        msg = 'blocking several seconds\n'
        print msg
        return msg

# 如果需要将阻塞任务的返回值作为相应返回给客户端，则需借助tornado的两个装饰器：
# @tornado.web.asynchronous将请求变为长连接
# @tornado.gen.coroutine

class FirstHandler(tornado.web.RequestHandler):
    executor = ThreadPoolExecutor(8)

    @tornado.web.asynchronous
    @tornado.gen.coroutine
    def get(self):
        res = yield self.sleep()
        self.write('handle first request\n')
        self.write(res)
        self.finish()

    @run_on_executor
        def sleep(self):
        time.sleep(10)
        msg = 'blocking several seconds\n'
        print msg
        return msg
```

对于最后一个Handler：

`curl localhost:8001/first` 会阻塞10s后输出：

> handle first request
>
> blocking several seconds

`curl localhost:8001/second`会立即输出：

> handle second request



