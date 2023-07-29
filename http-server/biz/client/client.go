package client

import (
	"context"
	"io/ioutil"
	"sync"
	"errors"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"
	db "github.com/sherry-500/apigateway/db/sqlc"
)

var Clients = make(map[string]genericclient.Client)

var clientCache sync.Map
var idlMapCache sync.Map

func InitJsonGenericClient(svcName string, idl string, idlType string) (genericclient.Client, error){
	//构造provider
	var content []byte
	var contentStr string
	var err error

	if idlType == "PATH" {
		content, err = ioutil.ReadFile(idl)
		if err != nil {
			return nil, err
		}
		contentStr = string(content)
	} else {
		contentStr = idl

	}

	provider, err := generic.NewThriftContentProvider(contentStr, nil)
	if err != nil {
		return nil, err
	}
	// 构造 json 类型的泛化调用
	g, err := generic.JSONThriftGeneric(provider)
	if err != nil {
		return nil, err
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		return nil, err
	}

	client, err := genericclient.NewClient(svcName, g,
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetClient(svcname string)(genericclient.Client, error){
	idlmap, err := db.StoreInstance.GetIdlmapByName(context.Background(), svcname)
	if err != nil {
		return nil, err
	}

	idl := idlmap.Idl
	idlType := idlmap.Type

	if client, ok := clientCache.Load(svcname); ok {
		if cachedIdl, ok := idlMapCache.Load(svcname); ok{
			if cachedIdl != idl {
				idlMapCache.Store(svcname, idl)
			} else {
				if genericClient, ok := client.(genericclient.Client); ok {
                    return genericClient, nil
                } else {
                    return nil, errors.New("type assertion failed")
                }
			}
		}else{
			idlMapCache.Store(svcname, idl)
		}
	}

	client, err := InitJsonGenericClient(svcname, idl, idlType)
	if err != nil {
		return nil, err
	}

	clientCache.Store(svcname, client)
	idlMapCache.Store(svcname, idl)
	return client, nil
}
