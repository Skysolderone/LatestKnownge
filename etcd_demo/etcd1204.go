package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

/*实现分布式锁*/

func main() {
	config := clientv3.Config{
		Endpoint:    []string{"xxx.xxx.xxx.xxx:2379"},
		DialTimeout: 5 * time.Second,
	}
	//get client conn
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer client.Close()
	// 1. 上锁（创建租约，自动续租，拿着租约去抢占一个key ）
	// 用于申请租约
	lease := clientv3.NewLease(client)
	//申请十秒的租约
	leaseGrantResp, err := lease.Grant(context.TODO(), 10) //10s
	if err != nil {
		fmt.Println(err)
		return
	}
	//get id
	Id := leaseGrantResp.ID()
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	defer lease.Revoke(context.TODO(), Id)
	//auto 续租
	keepRespChan, err := lease.KeepAlive(ctx, Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	//process resp channel
	go func() {
		select {
		case keepResp := <-keepRespChan:
			if keepResp == nil {
				fmt.Println("expire")
				goto END
			} else {
				//
				fmt.Println("receive auto reconnect ", keepResp.ID)
			}
		}
	END:
	}()
	//if key not exists then set else get mutex failed
	kv := clientv3.NewKV(client)
	//create  transaction
	txn := kv.Txn(context.TODO())
	//if key not exists
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job7"), "=", 0)).
		Then(clientv3.OpPut("/cron/lock/job7", "", clientv3.WithLease(Id))).
		Else(clientv3.OpGet("/cron/lock/job7")) //if key exists
	//commit transaction
	txnResp, err := txn.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}
	//judge whether get mutex
	if !txnResp.Successed {
		fmt.Println("MUTEX GET FAIL:", string(txnResp.Response[0].GetResponseRange().Kvs[0].Value))
		return
	}
	//2.process  very safe
	fmt.Println("PROCESS IS SAFE")
	time.Sleep(time.Second)
	//3.free mutex
	// defer会取消续租，释放锁

}
