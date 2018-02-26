When Istio Meets Jaeger — An Example of End-to-end Distributed Tracing

![](https://cdn-images-1.medium.com/max/1600/1*6MjgQZk-pWVtF88PENpNhA.png)

Kubernetes很NB！因为他能帮助很多的工程师团队去实现SOA（面向服务的架构体系）。在过去很长的一段时间里，我们都是围绕monolith mindset的概念构来构建我们的应用程序。也就是说，我们会在一个很牛X的计算机上运行一个应用的所有的组件。像帐户管理，结算，报告生成等这些工作，都是在一个机器上用共享资源的方式运行的。这种模式一直很ok，直到SOA出现了。它通过将应用程序拆分成一个个相对小的组件，并让它们之间使用REST或gRPC进行通信。我们其实仅仅希望这样做会比以前容易点，但后来我们发现，其实等待我们的是一堆新的挑战。跨服务的访问如何通信？如何去observe两个微型服务之间的通信（如日志或tracing）？本文演示如何在Kubernetes集群内部设置OpenTracing，以便在服务之间进行end-to-end的去跟踪，和在一个服务内部使用正确的工具进行跟踪。


### 创建Kubernetes

首先，我们需要有一个可用的Kubernetes集群。我在AWS上使用kops(https://github.com/kubernetes/kops)，因为它提供了一系列的K8S群集自动化操作命令，如upgrade，scaling up/down和多个instance group。除了方便的集群操作之外，kops团队还随着Kubernetes版本升级而升级，以支持Kubernetes的最新版本。我觉得这东西很酷，很有用！

按照这个(https://github.com/kubernetes/kops/blob/master/docs/aws.md)步骤可以安装和使用kops。

#### 创建集群

```
kops create cluster \ 
--name steven.buffer-k8s.com \ 
--cloud aws \ 
--master-size t2.medium \ 
--master-zones=us-east-1b \ 
--node-size m5.large \ 
--zones=us-east-1a,us-east-1b,us-east-1c,us-east-1d,us-east-1e,us-east-1f \ 
--node-count=2 \ 
--kubernetes-version=1.8.6 \ 
--vpc=vpc-1234567a \ 
--network-cidr=10.0.0.0/16 \ 
--networking=flannel \ 
--authorization=RBAC \ 
--ssh-public-key="~/.ssh/kube_aws_rsa.pub" \ 
--yes
```
这个命令会告诉AWS在`us-east-1`上创建一个VPC`1234456a`，会有1个master节点和2个worker（minion）节点，它的CIDR(无类别域间路由，Classless Inter-Domain Routing)是`10.0.0.0/16`。大约10分钟后，你的集群就可以使用了。当然在此过程中，你也可以用`watch kubectl get nodes`去查看整个创建的过程。

一旦Kubernetes集群安装完成，你就可以在上面去安装Istio(https://istio.io/)了。他是一个Service mesh，可以去管理在一个集群内部的服务之间的流量。由于它的这种特性，使得Istio能够很轻松的实现在两个服务之间进行trace。

### 安装Istio
从GitHub repo(https://github.com/istio/istio/releases/tag/0.4.0)下载Istio

在你下载的Istio目录中可以找到istio.yaml文件，运行：
`kubectl apply -f install/kubernetes/istio.yaml`






