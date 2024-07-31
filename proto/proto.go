package proto

import (
	"context"
	"reflect"
	"time"
)

type Proto interface {
	Serialize() []byte
	Deserialize(data []byte) error
}

const HandlerTimeout = 5 * time.Second

type S1ProtoContext struct {
	Header    *S1ProtoHeader
	ProtoName string
	Data      []byte
	ProtoData Proto

	context.Context
}

func NewS1ProtoContext(header *S1ProtoHeader, body []byte) (*S1ProtoContext, context.CancelFunc, error) {
	if header.ProtoId == 0 {
		return nil, nil, nil
	}

	protoData := CreateProtoInstance(header.ProtoId)
	if protoData == nil {
		return nil, nil, ProtoNotFoundError
	}

	err := protoData.Deserialize(body)
	if err != nil {
		return nil, nil, err
	}

	inCtx, cancel := context.WithTimeout(context.Background(), HandlerTimeout)

	ctx := &S1ProtoContext{
		Header:    header,
		ProtoName: ProtoIdMap[header.ProtoId],
		Data:      body,
		ProtoData: protoData,
		Context:   inCtx,
	}

	return ctx, cancel, nil
}

func CreateProtoInstance(protoid uint16) Proto {
	typ, ok := ProtoTypeMap[protoid]
	if !ok {
		return nil
	}

	return reflect.New(typ).Interface().(Proto)
}

const (
	S1StuMainJudgeLoginPass    = "39475983745"
	S1StuVersion               = 1.521
	S1StuMainJudgeClientId     = 71
	S1StuMainJudgeClientTId    = 10006
	S1StuMainJudgeClientTeamId = 2
)
