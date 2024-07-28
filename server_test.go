package ors

import (
	"testing"

	"github.com/wintbiit/ors-proto/proto"
)

func TestSendCommand(t *testing.T) {
	s := NewServer("127.0.0.1", proto.S1StuMainJudgeClientId, proto.S1StuMainJudgeClientTId, proto.S1StuMainJudgeClientTeamId)

	s.debug = true

	if err := s.Connect(); err != nil {
		t.Fatal(err)
	}

	defer s.Close()

	if err := s.Login(proto.S1StuMainJudgeLoginPass); err != nil {
		t.Fatal(err)
	}

	cmd := "-start"

	gmCmd := &proto.S1BattleProto2022ClientGmcommand{
		Len: byte(len(cmd)),
		Cmd: cmd,
	}

	if err := s.Send(proto.ProtoIDS1BattleProto2022ClientGMCommand, gmCmd); err != nil {
		t.Fatal(err)
	}

	select {}
}
