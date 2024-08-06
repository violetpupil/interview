# transaction

## ACID特性

- 原子性 Atomicity 事务中的操作，要么都执行，要么都不执行
- 一致性 Consistency 不同数据之间的关系保持一致
- 隔离性 Isolation 事务中对数据的读写，可以做到相互隔离
- 持久性 Durability 事务结束后，数据持久保存

## 隔离级别

- 读未提交 read uncommitted 指一个事务还没提交时，它做的变更就能被其他事务看到
- 读提交 read committed 指一个事务提交之后，它做的变更才能被其他事务看到
- 可重复读 repeatable read 指一个事务执行过程中看到的数据，一直跟这个事务启动时看到的数据是一致的，MySQL InnoDB 引擎的默认隔离级别
- 串行化 serializable 会对记录加上读写锁，在多个事务对这条记录进行读写操作时，如果发生了读写冲突的时候，后访问的事务必须等前一个事务执行完成，才能继续执行
