# index

## 按存储的数据分类

聚簇索引：叶子节点存放实际数据，索引键的选择顺序：

1) 优先使用主键
2) 没有主键，使用第一个不可为NULL的唯一列
3) 都没有的话，生成一个隐式自增id列

二级索引：叶子节点存放主键值

### 应用

如果查询的数据能从二级索引中获取，称为覆盖索引，效率比较高

否则，需要先从二级索引获取主键值，然后从聚簇索引获取数据，称为回表

## 按字段特性分类

普通索引

前缀索引：对字符类型字段前几个字符建立索引

唯一索引：索引列的值唯一，允许有NULL值

主键索引：索引列的值唯一，不允许有NULL值，只能有一个主键索引
