package initialize

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"zg6/2112a-6/jobs/global"
)

func InitNacos() {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         global.NacosConf.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      global.NacosConf.Host,
			ContextPath: "/nacos",
			Port:        uint64(global.NacosConf.Port),
			Scheme:      "http",
		},
	}
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	content, _ := configClient.GetConfig(
		vo.ConfigParam{
			DataId: global.NacosConf.Dataid,
			Group:  global.NacosConf.Group,
		},
	)
	err = json.Unmarshal([]byte(content), &global.AppConf)
	if err != nil {
		log.Println("nacos初始化失败:" + err.Error())
		return
	}
}
