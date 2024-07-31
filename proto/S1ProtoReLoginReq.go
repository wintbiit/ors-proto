package proto

import (
	"encoding/binary"
	"math"
)

const S1ProtoReLoginReqSize = 72

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoReLoginReq
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

#nullable disable
public class S1ProtoReLoginReq : S1ProtoBase
{
  public string account;
  public string password;
  public float version;
  public uint tid;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteString(this.account, 32);
    buffer.WriteString(this.password, 32);
    buffer.WriteFloat(this.version);
    buffer.WriteUInt(this.tid);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.account = buffer.ReadString(32);
    this.password = buffer.ReadString(32);
    this.version = buffer.ReadFloat();
    this.tid = buffer.ReadUInt();
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 72;
}

*/

func (s *S1ProtoReLoginReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoReLoginReqSize)
	copy(bytes[0:], s.Account)
	copy(bytes[32:], s.Password)
	binary.LittleEndian.PutUint32(bytes[64:], math.Float32bits(s.Version))
	binary.LittleEndian.PutUint32(bytes[68:], s.Tid)
	return bytes
}

func (s *S1ProtoReLoginReq) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoReLoginReqSize {
		return InvalidDataError
	}

	s.Account = string(bytes[0:32])
	s.Password = string(bytes[32:64])
	s.Version = math.Float32frombits(binary.LittleEndian.Uint32(bytes[64:]))
	s.Tid = binary.LittleEndian.Uint32(bytes[68:])
	return nil
}
