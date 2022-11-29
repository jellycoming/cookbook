```bash
# 获取节点上的所有Pod
kubectl get pods -o wide --field-selector spec.nodeName={nodeName}

# 查看Pod内Container的资源占用情况
kubectl top pod {podName} --containers

# 进入到Pod Container内部
kubectl exec -it {podName} -c {containerName} -- /bin/bash

# 删除Pod
kubectl delete pod {podName}

# 设置节点不可调度/驱逐节点上的Pod/设置节点可调度
kubectl cordon {nodeName}
kubectl drain {nodeName} --delete-local-data --ignore-daemonsets --force
kubectl uncordon {nodeName}

# 删除节点
kubectl delete node {nodeName}

# 加入节点
# kubeadm token create --print-join-command
# ?
```