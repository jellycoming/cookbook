```bash
# 获取节点上的所有Pod
kubectl get pods -o wide --field-selector spec.nodeName={nodeName}

# 查看Pod内Container的资源占用情况
kubectl top pod {podName} --containers

# 进入到Pod Container内部
kubectl exec -it {podName} -c {containerName} -- /bin/bash

# 删除Pod
kubectl delete pod {podName}

# 滚动重启deployment下的所有pod
kubectl rollout restart deployment {deploymentName}
# 查看deployment状态
kubectl rollout status deployment {deploymentName}
# 更新deployment镜像
kubectl set image deployment/{deploymentName} {containerName}={image}:{tag}


#### 节点维护
# 设置节点不可调度
kubectl cordon {nodeName}
# 驱逐节点上的Pod
kubectl drain {nodeName} --delete-local-data --ignore-daemonsets --force
# 设置节点可调度
kubectl uncordon {nodeName}

# 删除节点
kubectl delete node {nodeName}

# 加入节点
# kubeadm token create --print-join-command
# ?

# 按副本数降序查看Hpa
kubectl get hpa|awk 'NR>1{print $0}'|sort -nr -k 6

# PersistentVolume
# 删除所有持久卷(不建议)
 kubectl delete pvc --all
 kubectl delete pv --all
```