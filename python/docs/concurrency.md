# Linux同步原语

`自旋锁 (spinlock)`的设计理念是它仅会被持有非常短的时间;

`信号量(semaphore)`当需要长时间持有一个锁的时候,信号量就是一个很好的解决方案;

# Python threading.Condition\(\)

lock = threading.Condition\(\)

* 适用于两个线程通过状态变换交替执行的业务场景

* lock.wait\(\) 和 lock.notify\(\) 都需要事先调用 lock.acquire\(\) 获得锁对象, lock.notify\(\) 后不会自动释放锁对象，要调用 lock.release\(\),  
  lock.wait\(\)会释放当前锁并进入wait状态，被唤醒后会重新获得锁并继续执行

* 多个持有相同锁对象的线程都进入wait状态时，notify会唤醒最先进入wait状态的线程，事实上Condition对象会维护一个self.\_\_waiters列表，  
  notify方法会对self.\_\_waiters\[:n\]的线程执行唤醒操作

**关于lock.wait\(timeout=1\), 实质上是要在接收到其他线程notify或者timeout后重新获得锁，所以timeout只有在锁闲置时才会起作用。**

如下代码A线程并不能在timeout 1秒后继续执行下面的print语句，因为A线程wait后锁被B线程获得，即使A线程达到了timeout的条件，但B线程并没有release锁，  
所以A线程只能继续挂起，直到B线程执行lock.release\(\)。

```py
class ThreadA(threading.Thread):
    def __init__(self, lock):
        super(ThreadA, self).__init__()
        self.lock = lock

    def run(self):
        self.lock.acquire()
        print '{}:{}'.format(self.getName(), '线程进入就绪状态,等待通知')
        self.lock.wait(timeout=1) # 1秒后并不能继续执行下面的print语句
        print 'accept notify or time out'
        self.lock.release()


class ThreadB(threading.Thread):
    def __init__(self, lock):
        super(ThreadB, self).__init__()
        self.lock = lock

    def run(self):
        self.lock.acquire()
        time.sleep(5)
        print '{}:{}'.format(self.getName(), '12345')
        self.lock.release() # B线程释放锁之后，A线程的wait语句才能获得锁继续执行

if __name__ == '__main__':
    lock = threading.Condition()
    thread_a = ThreadA(lock)
    thread_b = ThreadB(lock)
    thread_a.start()
    thread_b.start()
```

# threading.Semaphore\(\)

信号量同步基于内部计数器，每调用一次acquire\(\)，计数器减1；每调用一次release\(\)，计数器加1.当计数器为0时，acquire\(\)调用被阻塞。  
这是迪科斯彻（Dijkstra）信号量概念P\(\)和V\(\)的Python实现。信号量同步机制适用于访问像服务器这样的有限资源。

# threading.Thread\(\)实例的 join\(\) 与 setDaemon\(\)

`t1.join()`用t1线程阻塞主线程，也就是主线程会等待t1线程执行完之后才会继续执行,输出结果：

> \#sleep 5s
>
> run Thread-1
>
> MainThread done.

`t1.setDaemon(True)`将t1线程设置为守护线程，不会阻塞主线程，当主线程执行完之后会直接kill未完成任务的t1线程，输出结果：

> \#just now
>
> MainThread done.

```py
class ThreadA(threading.Thread):
    def __init__(self):
        super(ThreadA, self).__init__()

    def run(self):
        time.sleep(5)
        print 'run {}'.format(self.getName())


if __name__ == '__main__':
    t1 = ThreadA()
    # t1.setDaemon(True) # 必须在start之前
    t1.start()
    # t1.join() # 必须在start之后
    print '{} done.'.format(threading.currentThread().getName())
```
