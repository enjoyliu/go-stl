# OrderedMap

参考C++的STL库实现的有序Map，支持插入、删除、查找、遍历等操作。

[cpp_map](https://en.cppreference.com/w/cpp/container/map)


## 1. 介绍

map的底层数据结构使用红黑树,可以参考
[rbTree](https://github.com/sakeven/RbTree)

## 2. 待议
如何确保通用性的同时，保证输入的key是有序的？