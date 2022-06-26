## 第二章 搜索class文件

### 类路径

https://blog.csdn.net/qq_37284484/article/details/80776810 本书详细介绍的类路径

#### 类分类

- 启动类路径(boostrap classpath)：启动类路径默认对应jre/lib目录，Java标准库（大部分在rt.jar里）位于该路径。
- 扩展类路径(extension classpath)：扩展类路径默认对应jre/lib/ext目录，使用Java扩展机制的类位于这个路径。
- 用户类路径(user classpath)：我们自己实现的类，以及第三方类库位于用户路径，可以通过-Xbootclasspath选项修改启动类的路径（一般不会用到）。

#### java -cp

https://zhuanlan.zhihu.com/p/214093661 这文章cp命令写反了。。

这个和Java jar还是有一点区别的，jar命令会寻找当前jar的Main-class,如果你的MANIFEST.MF文件中没有Main-Class，就会提示Cant load main-class之类的错误。所以在导出jar包的时候一定要指定main-class。

对于java -cp就不需要指定Main-Class来指定入口。因为第一个参数就指定了你的入口类，第二个参数就是你的jar包。它会根据你的jar包找到第一个参数指定的Test类，来输出HelloWorld。

