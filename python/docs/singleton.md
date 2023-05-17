```py
#使用装饰器(decorator),
#这是一种更pythonic,更elegant的方法,
#单例类本身根本不知道自己是单例的,因为他本身(自己的代码)并不是单例的
def singleton(cls, *args, **kw):
     instances = {}
     def _singleton(*args, **kw):
         if cls not in instances:
             instances[cls] = cls(*args, **kw)
         return instances[cls]
     return _singleton
@singleton
class MyClass4(object):
     a = 1
     def __init__(self, x=0):
         self.x = x
```

```py
# 重写__new__方法
class Singleton(object):
    _instance = None
    def __new__(cls, *args, **kwargs):
        if not cls._instance:
            cls._instance = super(Singleton, cls).__new__(cls, *args, **kwargs)
        return cls._instance
```

### 需要单例的对象放在模块中，被其他模块导入,在所有模块中将只有一个对象存在

