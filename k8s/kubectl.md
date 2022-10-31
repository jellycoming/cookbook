```bash
# 获取节点上的所有Pod
kubectl get pods -o wide --field-selector spec.nodeName={nodeName}

# 查看Pod内Container的资源占用情况
kubectl top pod {podName} --containers

# 进入到Pod Container内部
kubectl exec -it {podName} -c {containerName} -- /bin/bash

# 删除Pod
kubectl delete pod {podName}
```