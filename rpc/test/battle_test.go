package test

import (
	"hat_logic/pbgo/pbreq"
	"hat_logic/rpc"
	"testing"
)

func TestBattleVerify(t *testing.T) {
	rpc.VerifyBattle(&pbreq.ReqInitBattle{}, []*pbreq.BattleReq{})
}
