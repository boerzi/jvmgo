## 第一章 命令行工具

我直接用的开源的 https://github.com/guonaihong/clop，之前解析curl用的就是这个，比这本书用的系统自带的简单很多。

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



## 第三章 解析class文件

使用010 editor可以很方便查看到class信息，这一节基本就是在搬砖（对照010 editor来搬砖就好了），代码也有很多地方可以优化，自己手打太辛苦了。

https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html 官方说明

这里有个很有意思的地方，class文件中用的不是标准的utf-8，用的是MUTF-8。

![image-20220701222242533](/Users/leiyuchen/OneDrive/笔记/图片/image-20220701222242533.png)

这边发现一个奇怪的问题，long类型为什么会有一个continued

![image-20220710210445676](/Users/leiyuchen/OneDrive/笔记/图片/image-20220710210445676.png)
