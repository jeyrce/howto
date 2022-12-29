
### kubeadm

kubeadm 实际作用是生成 kubelet 所需要的集群启动参数，包括 service文件等。

文件写入磁盘后使用 systemctl daemon-reload 重载并 systemctl start kubelet

通过配置文件控制 kubeadm 的初始化行为: <https://kubernetes.io/zh-cn/docs/reference/setup-tools/kubeadm/kubeadm-init/>


自动转换写的不对的配置。

```text
 kubeadm config migrate --old-config kubeadm-config-v1.26.0.yaml --new-config new-kubeadm-config-v1.26.0.yaml
```

