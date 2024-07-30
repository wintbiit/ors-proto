package ors

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/wintbiit/ors-proto/proto"
)

func TestSendCommand(t *testing.T) {
	s := NewServer("192.168.1.2", 72, proto.S1StuMainJudgeClientTId, proto.S1StuMainJudgeClientTeamId)

	s.debug = true

	if err := s.Connect(); err != nil {
		t.Fatal(err)
	}

	defer s.Close()

	time.Sleep(1 * time.Second)

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

func TestValidateGeneratedProtoData(t *testing.T) {
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

	var s = new(Server)
	s.ClientId = proto.S1StuMainJudgeClientId
	s.ClientTId = proto.S1StuMainJudgeClientTId
	s.ClientTeamId = proto.S1StuMainJudgeClientTeamId
	var genBytes = s.packProtoData(proto.ProtoIDS1ProtoLoginReq, &proto.S1ProtoLoginReq{
		Account:  fmt.Sprintf("s0unit_client_%d_%d", s.ClientId, s.ClientId),
		Password: proto.S1StuMainJudgeLoginPass,
		Version:  proto.S1StuVersion,
		Tid:      s.ClientTId,
		Teamid:   s.ClientTeamId,
		Hash:     s.Hash,
	})

	t.Logf("Generated bytes: \t%v", genBytes)
	t.Logf("Original bytes: \t%v", bytes)

	if len(genBytes) != len(bytes) {
		t.Fatalf("Generated bytes length %d does not match original bytes length %d", len(genBytes), len(bytes))
	}

	match := true
	for i := 0; i < len(genBytes); i++ {
		if genBytes[i] != bytes[i] {
			t.Errorf("Generated bytes do not match original bytes at index %d", i)
			match = false
		}
	}

	if match {
		t.Log("Generated bytes match original bytes")
	}
}
