package client

import(
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/cloudwego/kitex/client"
	"log"
	"io/ioutil"
)

var Clients = make(map[string]genericclient.Client)

func InitJsonGenericClient(svcName string, filePath string) {
	//构造provider
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	contentStr := string(content)

	provider, err := generic.NewThriftContentProvider(contentStr, nil)
	if err != nil {
		panic("create provider failed")
	}
	// 构造 json 类型的泛化调用
	g, err := generic.JSONThriftGeneric(provider)
	if err != nil {
		panic(err)
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	Clients[svcName], err = genericclient.NewClient(svcName, g,
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
}
