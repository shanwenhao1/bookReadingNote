# 锁

## 锁使用原则
- 尽量减少锁的持续时间
    - 细化锁的粒度, 避免在持有锁的时候做各种耗时的操作
    - 不要在持有锁的时候做IO操作, 用锁来保护IO操作需要的资源而不是IO操作本身

## 原子性操作

- [锁使用示例](mutexUse.go)
    - [tryLock](../../../infra/tool/lock/tryLock_test.go)
- 利用锁实现[单例模式](../../../infra/tool/singleton/singleton.go)