# 问题收集地

## gateway 压力测试问题

命令：

    ab -n1000 -c100 http://127.0.0.1:5550/ExpenseList

问题：

    CPU 使用率飙升到 100%，并且无法返回数据

解决：

    去掉保存数据包的代码

引入另外一个问题：

    read tcp 127.0.0.1:5550->127.0.0.1:64308: use of closed network connection
