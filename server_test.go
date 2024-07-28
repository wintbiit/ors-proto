package ors

import (
	"os"
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

func TestParseOriginal(t *testing.T) {
	f, err := os.Open("C:\\Users\\WintBit\\opt\\Game\\RM\\RoboMasterEngine 10.0.0.2.stu\\proto\\9000_S1ProtoLoginReq.bin")

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	bytes := make([]byte, 1024)
	n, err := f.Read(bytes)

	if err != nil {
		t.Fatal(err)
	}

	if n == 0 {
		t.Fatal("No data read")
	}

	bytes = bytes[:n]

	var header proto.S1ProtoHeader
	if err := header.Deserialize(bytes); err != nil {
		t.Fatal(err)
	}

	t.Logf("Header: %+v", header)

	var data proto.S1ProtoLoginReq
	if err := data.Deserialize(bytes[proto.S1ProtoHeaderSize:]); err != nil {
		t.Fatal(err)
	}

	t.Logf("Data: %+v", data)

}
