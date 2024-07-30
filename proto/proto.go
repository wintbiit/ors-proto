package proto

import (
	"reflect"
	"time"
)

type Proto interface {
	Serialize() []byte
	Deserialize(data []byte) error
}

type S1ProtoContext struct {
	Header *S1ProtoHeader
	Data   []byte

	// context.Context
	done chan struct{}
	err  error
	dl   time.Time
	m    map[interface{}]interface{}
}

func (s *S1ProtoContext) Deadline() (deadline time.Time, ok bool) {
	return s.dl, true
}

func (s *S1ProtoContext) Done() <-chan struct{} {
	return s.done
}

func (s *S1ProtoContext) Err() error {
	return s.err
}

func (s *S1ProtoContext) Value(key interface{}) interface{} {
	return s.m[key]
}

func (s *S1ProtoContext) Set(key, value interface{}) {
	s.m[key] = value
}

func (s *S1ProtoContext) Cancel() {
	close(s.done)
}

func NewS1ProtoContext(header *S1ProtoHeader, body []byte) *S1ProtoContext {
	return &S1ProtoContext{
		Header: header,
		Data:   body,
		done:   make(chan struct{}),
		m:      make(map[interface{}]interface{}),
	}
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
