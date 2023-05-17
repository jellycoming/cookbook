死锁的本质原因用一个字概括，就是“等”

* 循环调用 - 等自己，在获取锁后还未释放的前提下再次尝试获取同一把锁

```py
class MyThread(threading.Thread):
    def __init__(self, lock):
        super(MyThread, self).__init__()
        self.lock = lock
        self.name = self.getName()

    def run(self):
        print self.name + ' run'
        if self.lock.acquire():
            print self.name + ' got lock first time'
            # 在未释放的前提下再次尝试获取相同的锁,程序将在这里进入无限的等待中
            if self.lock.acquire():
                print self.name + ' got lock second time'
                self.lock.release()
            self.lock.release()

if __name__ == '__main__':
    lock = threading.Lock()
    my_thread = MyThread(lock)
    my_thread.start()
    print '{} done.'.format(threading.currentThread().getName())

# output
# Thread-1 run
# MainThread done.
# Thread-1 got lock first time
# ...
```

* 交叉调用 - 等别人，两个线程相互等待对方释放资源

```py
class ThreadA(threading.Thread):
    def __init__(self, lock_a, lock_b):
        super(ThreadA, self).__init__()
        self.lock_a = lock_a
        self.lock_b = lock_b
        self.name = self.getName()

    def run(self):
        print self.name + ' run'
        if self.lock_a.acquire():
            print self.name + ' got lock a'
            time.sleep(2) # 模拟阻塞,给线程B足够的时间获取锁b
            # 尝试获取锁b的时候,锁b已经被线程B获取并未被释放
            print self.name + ' waiting for lock b'
            if self.lock_b.acquire():
                print self.name + ' got lock b'
                self.lock_b.release()
            self.lock_a.release()


class ThreadB(threading.Thread):
    def __init__(self, lock_a, lock_b):
        super(ThreadB, self).__init__()
        self.lock_a = lock_a
        self.lock_b = lock_b
        self.name = self.getName()

    def run(self):
        print self.name + ' run'
        if self.lock_b.acquire():
            print self.name + ' got lock b'
            time.sleep(2) # 模拟阻塞,给线程A足够的时间获取锁a
            # 尝试获取锁a的时候,锁a已经被线程A获取并未被释放,此时A线程也在尝试获取锁b,两个线程同时等待对方释放锁,造成死锁
            print self.name + ' waiting for lock a'
            if self.lock_a.acquire():
                print self.name + ' got lock a'
                self.lock_a.release()
            self.lock_b.release()

if __name__ == '__main__':
    lock_a = threading.Lock()
    lock_b = threading.Lock()
    thread_a = ThreadA(lock_a, lock_b)
    thread_b = ThreadB(lock_a, lock_b)
    thread_a.start()
    thread_b.start()
    print '{} done.'.format(threading.currentThread().getName())

# output
# Thread-1 run
# Thread-1 got lock a
# Thread-2 run
# Thread-2 got lock b
# MainThread done.
# Thread-1 waiting for lock b
# Thread-2 waiting for lock a
# ...
```