* ##### 新建git项目

```
Create a new repository
$ git clone https://github.com/jellycoming/project.git
$ cd project
$ touch README.md
$ git add README.md
$ git commit -m "add README"
$ git push -u origin master

Existing folder or Git repository
$ cd project
$ git init
$ git remote add origin https://github.com/jellycoming/project.git
$ git add .
$ git commit
$ git push -u origin master
```

* ##### 分支

```
# 基于当前分支新建分支，同时切换到新分支
$ git checkout -b master-bugfix
或者：
$ git branch master-bugfix
$ git checkout master-bugfix

# 基于tag新建分支
$ git checkout -b newbranch v1.0

# 将当前分支回退到指定tag版本
$ git reset --hard v1.0

# 比较两个分支差异
$ git diff master..master-bugfix

# 将本地分支推送到远程服务器
$ git push -u origin master-bugfix

# 合并本地分支到master
$ git checkout master
$ git merge master-bugfix

# 删除本地分支
$ git branch -d master-bugfix

# 删除远程分支
$ git push origin --delete master-bugfix

# 从远程服务器检出指定分支到本地
$ git checkout -b master-bugfix origin/master-bugfix
$ git checkout --track origin/master-bugfix

# 基于commit hash检出分支
$ git checkout -b branch_name <commit-hash>
```

* ##### 标签

```
# 添加本地标签
$ git tag -a v1.0 -m 'release tag'
# 将指定本地标签推送到远程服务器
$ git push origin v1.0
# 将所有本地标签推送到远程服务器
$ git push origin –-tags
# 删除远程服务器标签
$ git push --delete origin tagname
# 删除本地标签
$ git tag --delete tagname
# 查看标签v1.1.0的前10行注释内容
$ git tag -l -n10 v1.1.0
```

* ##### 更换本地Git用户

```
# 查看并更换remote url
$ git config --get remote.origin.url
$ git remote set-url origin https://username:'passwd'@gitserver.com/prod/project.git
```

* ##### 打包

```
git archive --format zip --output /tmp/target.zip [branch|tag]
```

* ##### reset origin/master to specific commit

```
git checkout master
git reset --hard e3f1e37
git push --force origin master
# Then to prove it (it won't print any diff)
git diff master..origin/master
```

* #### 用dev分支覆盖master分支

```
git checkout master
git pull
git checkout dev
git merge -s ours master
git checkout master
git merge dev
```

* ##### 查看tag何时被创建
```shell script
git log --tags --simplify-by-decoration --pretty="format:%ai %d"
```

* #### 新版本授权
```shell script
git remote set-url origin https://{token}@github.com/{username}/cookbook.git
```
