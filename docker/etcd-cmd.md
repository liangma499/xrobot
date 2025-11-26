etcdctl get --prefix "" --列出etcd所有key
etcdctl endpoint status --write-out=table  --列出所有集群状态
可以使用 etcdctl endpoint status 命令查看该节点的存储空间使用情况,包括数据库大小等指标。