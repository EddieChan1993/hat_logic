# 本地启动参数配置
hat_dev hat_logic 5 debug

hat_dev - 集群
hat_logic - 服务名
5 - logicId 节点

# 开服consul配置路径
`Key / Values/app_dynamic_cfg/hat_dev/gogate/logic`

# logic 存储pb生成命令
`protoc --go_out=../hat_logic/ proto/comm.proto && protoc --go_out=../hat_logic/ proto/player.proto`

# logic 请求pb生成
```
protoc --go_out=plugins=kite:../hat_logic comm.proto && protoc --go_out=plugins=kite:../hat_logic logic.proto
```

# logic api 对外rpc接口生成
```
protoc --go_out=plugins=kite:../hat_logic/ proto/api.proto
```

# 导表pb生成
```
cd /Users/duanchengwen/Re/hat_cehua/ &&
  git pull && cd /Users/duanchengwen/Re/Archive &&
  ./exporter

```

# 额外
线上部署错误日志查询路径
`/home/dhcd/back/service_deploy/re_dev-re_logic-1/log`

## 快速代码生成模版

goland的setting中搜索 live temp

设置apiMod快速生产模版代码
```
type IMod$NAME$Func interface {

}

func $NAME$Mod(gameCtx *pbreq.GameCtx) IMod$NAME$Func {
	ins, had := gameCtx.GetApiFunc(util.GetModName(&$NAME${}))
	if had {
		return ins.(IMod$NAME$Func)
	}
	return nil
}
```