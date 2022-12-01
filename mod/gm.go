package mod

/**
gm工具
*/
import (
	"fmt"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/mod/mcst"
	config "hat_logic/pbgo/pbcfg"
	"hat_logic/pbgo/pbreq"
	"hat_logic/tool/consul/static"
	"strconv"
	"strings"
)

const maxNum = 10000000000
const minNum = 1000000

type Gm struct {
	pbreq.TModApi11
}

//========================req api========================//

func (g *Gm) ReqGm(req *pbreq.ReqGm) *pbreq.RspGm {
	res := &pbreq.RspGm{
		Type: req.Type,
	}
	if !static.StaticDebugMode() {
		//非调试模式关闭gm
		res.Status = 2
		return res
	}
	var err error
	var pbRewards *pbreq.Rewards
	var extData string
	switch req.GetType() {
	case pbreq.GMType_AddItem:
		pbRewards, err = g.addItem(req)
	}
	if err != nil {
		klog.Warn(err)
		res.Status = 1
		return res
	}
	res.Rewards = pbRewards
	res.ExtData = extData
	return res
}

//========================tool========================//

func (g *Gm) addItem(req *pbreq.ReqGm) (*pbreq.Rewards, error) {
	items := make([]*config.Reward, 0)
	if req.Data == "" {
		for _, cfgId := range config.ItemData.GetAll() {
			tmp := &config.Reward{
				Type:  mcst.GoodTypItem,
				Id:    cfgId,
				Count: maxNum,
			}
			items = append(items, tmp)
		}
	} else {
		res1 := strings.Split(req.Data, ",")
		for _, s := range res1 {
			tmp := strings.Split(s, ":")
			if len(tmp) != 2 {
				err := fmt.Errorf("格式不正确：%s", s)
				klog.Warn(g.Log(), err)
				return nil, err
			}
			id, err := strconv.Atoi(tmp[0])
			if err != nil {
				klog.Warn(g.Log(), err)
				return nil, err
			}
			nums, err := strconv.Atoi(tmp[1])
			if err != nil {
				klog.Warn(g.Log(), err)
				return nil, err
			}
			cfgItem := config.ItemData.Get(mcst.DefaultInt(id))
			if cfgItem == nil {
				err := fmt.Errorf("找不到配置：Id=%d", id)
				klog.Warn(g.Log(), err)
				return nil, err
			}
			tmpItem := &config.Reward{
				Type:  mcst.GoodTypItem,
				Id:    mcst.DefaultInt(id),
				Count: mcst.BagNum(nums),
			}
			items = append(items, tmpItem)
		}
	}
	rewards, err := GoodsMod(g.GameCtx).Rewards(items)
	if err != nil {
		klog.Warn(g.Log(), err)
		return nil, err
	}
	return rewards, nil
}

//========================mod base========================//

func (g *Gm) RspSync(req *pbreq.RspSync) {
	return
}

func NewGm() *Gm {
	ins := &Gm{}
	ins.InitModApi(ins)
	return ins
}
