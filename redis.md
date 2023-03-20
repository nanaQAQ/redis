# Redis

## 通用命令

- KEYS：查看符合模板的所有key，不建议在生产环境中使用；
- DEL：删除一个指定的key，尽可能删除，在删除没有的key时不报错；
- EXISTS：判断一个key是否存在；
- EXPIRE：设置一个key的有效期；EXPIRE key seconds
- TTL：查看一个key的剩余有效期

## 数据结构

## Key

可以层级结构

set xdu:user:1 {json}

### String

字符串类型

底层时字节数组，最大空间不超过512m。

### Hash

无序字典

### List

双向链表

有序，插入和删除快，查询速度一般

保存对顺序由要求的数据

### Set

hashset

无序，不可重复，查找快，支持交集、并集、差集等功能

### SortedSet

跳表加哈希表

可排序，不重复，查询速度快

排行榜