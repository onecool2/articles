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

现在你的Istio因该已经运行在Kubernetes的集群里了。而且他因该同时创建了一个Nignx的Ingress Controller，这个东西是用来接收外面对集群内资源的访问的。我们一会儿会介绍如何设置IP。


Jaeger和Instio一起协同工作，可以实现跨服务的tracing。你可以用这个命令安装Jaege
```
kubectl create -n istio-system -f 
https://raw.githubusercontent.com/jaegertracing/jaeger-
kubernetes/master/all-in-one/jaeger-all-in-one-template.yml
```

完成之后，你可以通过Jaeger的UI访问它
![](https://cdn-images-1.medium.com/max/1600/0*Z8fQqmATKom1Au34.)

#### Instrument Code
Jaeger和Istio安装成功以后，你就可以看到跨服务的trace了。这是因为Istio注入了Envoy sidecars来处理服务间通信，而被部署的应用程序只与指定的sidecar进行通信。
从我的GitHub(https://github.com/stevenc81/jaeger-tracing-example)你可以找到一个简单的应用程序例子

main.go 
```
package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}

func getTimeHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Received getTime request")
	t := time.Now()
	ts := t.Format("Mon Jan _2 15:04:05 2006")
	fmt.Fprintf(w, "The time is %s", ts)
}

func main() {
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	injector := jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator)
	extractor := jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator)

	// Zipkin shares span ID between client and server spans; it must be enabled via the following option.
	zipkinSharedRPCSpan := jaeger.TracerOptions.ZipkinSharedRPCSpan(true)

	sender, _ := jaeger.NewUDPTransport("jaeger-agent.istio-system:5775", 0)
	tracer, closer := jaeger.NewTracer(
		"myhelloworld",
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(
			sender,
			jaeger.ReporterOptions.BufferFlushInterval(1*time.Second)),
		injector,
		extractor,
		zipkinSharedRPCSpan,
	)
	defer closer.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/gettime", getTimeHandler)
	http.ListenAndServe(
		":8080",
		nethttp.Middleware(tracer, http.DefaultServeMux))
}
```
从第28-30行开始，我们创建了一个Zipkin propagator，告诉Jaeger从OpenZipkin的request header中捕捉上下文(context)。你可能会问，这些header是怎么被放到request的开始部分的？还记得吗，当我说Istio用side care处理服务间的通信，并且应用程序只与它交互。对！，你可能已经猜到了。为了让Istio跟踪服务之间的request，当有request进入集群时，Istio的Ingress Controller将注入一组header。然后，它围绕着Envoy sidecars进行传播，并且每个都会将相关的 associated span上报给Jaeger。这有助于将span对应到每个trace。我们的应用程序代码利用这些header来收集内部服务之间的span。

下面是一个被Istio的Ingress Controller注入到OpenZipkin中header的列表
```
x-request-id
x-b3-traceid
x-b3-spanid
x-b3-parentspanid
x-b3-sampled
x-b3-flags
x-ot-span-context
```

可以用下面这个yaml文件去部署一个简单的应用程序
```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: myhelloworld
  name: myhelloworld
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: myhelloworld
    spec:
      containers:
      - image: stevenc81/jaeger-tracing-example:0.1
        imagePullPolicy: Always
        name: myhelloworld
        ports:
        - containerPort: 8080
      restartPolicy: Always
---
apiVersion: v1
kind: Service
metadata:
  name: myhelloworld
  labels:
    app: myhelloworld
spec:
  type: NodePort
  ports:
  - port: 80
    targetPort: 8080
    name: http
  selector:
    app: myhelloworld
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: myhelloworld
  annotations:
    kubernetes.io/ingress.class: "istio"
spec:
  rules:
  - http:
      paths:
      - path: /
        backend:
          serviceName: myhelloworld
          servicePort: 80
      - path: /gettime
        backend:
          serviceName: myhelloworld
          servicePort: 80
```

#### 部署
```kubectl apply -f <(istioctl kube-inject -f myhelloword.yaml)```

请注意，可以从Istio的bin目录下找到命令`istioctl`

现在是收获的时候了，当我们向Istio Ingress Controller发送请求时，它将在service以及应用程序之间进行trace。从截图中我们可以看到从不同地方报告的3个span

Ingress Controller
Envoy sidecar for application
Application code
![](https://cdn-images-1.medium.com/max/1600/0*iTwIzJGOa-w055jC.)
展开trace（看到3个span）就可以看到end-to-end的traceing
![](https://cdn-images-1.medium.com/max/1600/0*hHfzLg-i6yMtPFlk.)

### 结束语
SOA带来了一堆的待解决问题，特别是在如何observ服务之间的通信上。
Istio+Jager的集成解决方案可以从servcie到service的这个层面解决这个问题。
使用OpenZipkin prapagator配合Jaege，可以在end-to-end上tracing。


