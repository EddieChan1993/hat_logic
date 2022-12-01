package test

import (
	"fmt"
	"log"
	"hat_logic/mod/mcst"
	config "hat_logic/pbgo/pbcfg"
	"hat_logic/rpc"
	"testing"
)

func TestSendMail(t *testing.T) {
	initServer()
	roleId := int64(85388563)
	for i := 0; i < 1; i++ {
		err := rpc.SendMail(&rpc.MailContent{
			RoleId: roleId,
			MailId: mcst.MailIdInstanceBossRank,
			Rewards: []*config.Reward{
				{Type: mcst.GoodTypItem, Id: mcst.ItemIdDIAMOND, Count: int64(1000)},
			},
			Args: fmt.Sprintf(`{"index":%d}`, i+1),
		})
		log.Println(err)
	}
}
