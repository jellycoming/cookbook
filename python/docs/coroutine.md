```python
# coding=utf-8
def hello():
    while True:
        # r 的值是通过该生成器的send方法发送而来,如果没有coroutine.send(msg)调用, r 会始终为None
        # yield 后表达式的值将会在coroutine.next()方法被调用时返回
        r = yield simulate_http_request()
        print 'r=' + str(r)


def simulate_http_request():
    import random
    print 'simulate http request...'
    return 'OK' if random.randint(1, 100) > 50 else 'ERROR'


if __name__ == '__main__':
    coroutine = hello()
    for _ in range(3):
        yielded = coroutine.next()  # coroutine.next() 获取yield关键字后表达式的值(simulate_http_request函数的结果)
        print 'yielded=' + yielded
    coroutine.send('message...')  # coroutine.send(msg) 为yield关键字前面的变量赋值(hello中的r)
```



