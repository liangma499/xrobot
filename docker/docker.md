docker-compose -f docker-compose-card-prd.yaml up -d
docker exec --it 180dd1f3b197 /bin/bash



删除未使用的容器 删除所有停止的容器：

docker container prune
删除未使用的镜像 删除未使用的镜像：

docker image prune
删除所有未使用的镜像（包括未使用的中间层）：

docker image prune -a
删除未使用的网络 删除未使用的网络：

docker network prune
删除未使用的卷 删除未使用的卷：

docker volume prune
系统清理 执行全面清理：

docker system prune
选择删除所有未使用的资源，包括停止的容器、未使用的网络、未使用的镜像（无关的）、未使用的卷：

docker system prune -a --volumes

日志文件：
/var/lib/docker/containers/<container_id>/<container_id>-json.log。
