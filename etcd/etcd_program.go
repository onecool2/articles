上面介绍了命令行的基本使用。补充一下，当前默认还是v2版本通过设定环境变量export ETCDCTL_API=3，设置成V3版本。

v3版本支持rpc的远程调用，这样比http的方式效率更高 
先看API定义

关于key value操作

type KV interface {
    Put(ctx context.Context, key, val string, opts ...OpOption) (*PutResponse, error)

    Get(ctx context.Context, key string, opts ...OpOption) (*GetResponse, error)

    Delete(ctx context.Context, key string, opts ...OpOption) (*DeleteResponse, error)

    Compact(ctx context.Context, rev int64, opts ...CompactOption) (*CompactResponse, error)
    Do(ctx context.Context, op Op) (OpResponse, error)

    Txn(ctx context.Context) Txn
}

关于租约

type Lease interface {

    Grant(ctx context.Context, ttl int64) (*LeaseGrantResponse, error)

    Revoke(ctx context.Context, id LeaseID) (*LeaseRevokeResponse, error)


    TimeToLive(ctx context.Context, id LeaseID, opts ...LeaseOption) (*LeaseTimeToLiveResponse, error)

    KeepAlive(ctx context.Context, id LeaseID) (<-chan *LeaseKeepAliveResponse, error)

    KeepAliveOnce(ctx context.Context, id LeaseID) (*LeaseKeepAliveResponse, error)

    Close() error
}


type Client struct {
    Cluster
    KV
    Lease
    Watcher
    Auth
    Maintenance

    conn             *grpc.ClientConn
    cfg              Config
    creds            *credentials.TransportCredentials
    balancer         *simpleBalancer
    retryWrapper     retryRpcFunc
    retryAuthWrapper retryRpcFunc

    ctx    context.Context
    cancel context.CancelFunc

    // Username is a username for authentication
    Username string
    // Password is a password for authentication
    Password string
    // tokenCred is an instance of WithPerRPCCredentials()'s argument
    tokenCred *authTokenCredential
}


package main

import (
    "github.com/coreos/etcd/clientv3"
    "time"
    "golang.org/x/net/context"
    "fmt"
)


var (
    dialTimeout    = 5 * time.Second
    requestTimeout = 10 * time.Second
    endpoints      = []string{"10.39.0.6:2379",}
)

func main() {
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   []string{"10.39.0.6:2379"},
        DialTimeout: dialTimeout,
    })
    if err != nil {
        println(err)
    }
    defer cli.Close()

    ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
    _, err = cli.Put(ctx, "/test/hello", "world")
    cancel()

    ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
    resp,err := cli.Get(ctx, "/test/hello")
    cancel()

    for _, ev := range resp.Kvs {
        fmt.Printf("%s : %s\n", ev.Key, ev.Value)
    }

    _, err = cli.Put(context.TODO(), "key", "xyz")
    ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
    _, err = cli.Txn(ctx).
        If(clientv3.Compare(clientv3.Value("key"), ">", "abc")). 
        Then(clientv3.OpPut("key", "XYZ")).                      
        Else(clientv3.OpPut("key", "ABC")).
        Commit()
    cancel()

    rch := cli.Watch(context.Background(), "/test/hello", clientv3.WithPrefix())
    for wresp := range rch {
        for _, ev := range wresp.Events {
            fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
        }
    }

    if err != nil {
        println(err)
    }
}
