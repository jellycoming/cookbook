# Install Python2.7 via source code

> 安装`pip`依赖`openssl`与`zlib`这两个包,如果不预先安装,
> 安装`pip`时需在安装这两个包后重新编译安装`python`,所以最好提前安装好

```bash
yum -y install zlib zlib-devel openssl openssl-devel
```

```bash
wget https://www.python.org/ftp/python/2.7.13/Python-2.7.13.tgz
tar xf Python-2.7.13.tgz
cd Python-2.7.13
./configure --prefix=/usr/local
make && make install
```

安装完成后的执行路径为: `/usr/local/bin/python`
可选: `ln -s /usr/local/bin/python /usr/bin/python`

# Install pip

```bash
wget https://bootstrap.pypa.io/get-pip.py
/usr/local/bin/python get-pip.py
```

安装完成后的pip执行路径为: `/usr/local/bin/pip`

# macOS Mojave\(10.14\)下pip install MySQL-python包时报`ld: library not found for -lssl` 时的解决方式

```bash
$ sudo pip install MySQL-python --global-option=build_ext --global-option="-I/usr/local/opt/openssl/include" --global-option="-L/usr/local/opt/openssl/lib"
```