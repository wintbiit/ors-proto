package proto

import (
	"encoding/binary"
	"math"
)

const S1ProtoLoginReqSize = 84

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoLoginReq
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

#nullable disable
public class S1ProtoLoginReq : S1ProtoBase
{
  public string account;
  public string password;
  public float version;
  public uint tid;
  public uint teamid;
  public long hash;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteString(this.account, 32);
    buffer.WriteString(this.password, 32);
    buffer.WriteFloat(this.version);
    buffer.WriteUInt(this.tid);
    buffer.WriteUInt(this.teamid);
    buffer.WriteLong(this.hash);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    uint num1 = (uint) (0 + 32);
    if (header.dataLen >= num1)
      this.account = buffer.ReadString(32);
    uint num2 = num1 + 32U;
    if (header.dataLen >= num2)
      this.password = buffer.ReadString(32);
    uint num3 = num2 + 4U;
    if (header.dataLen >= num3)
      this.version = buffer.ReadFloat();
    uint num4 = num3 + 4U;
    if (header.dataLen >= num4)
      this.tid = buffer.ReadUInt();
    uint num5 = num4 + 4U;
    if (header.dataLen >= num5)
      this.teamid = buffer.ReadUInt();
    uint num6 = num5 + 8U;
    if (header.dataLen >= num6)
      this.hash = buffer.ReadLong();
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 84;
}

*/

func (s *S1ProtoLoginReq) Serialize() []byte {
	bytes := make([]byte, S1ProtoLoginReqSize)
	copy(bytes[0:32], s.Account)
	copy(bytes[32:64], s.Password)
	binary.LittleEndian.PutUint32(bytes[64:68], math.Float32bits(s.Version))
	binary.LittleEndian.PutUint32(bytes[68:72], s.Tid)
	binary.LittleEndian.PutUint32(bytes[72:76], s.Teamid)
	binary.LittleEndian.PutUint64(bytes[76:84], uint64(s.Hash))

	return bytes
}

func (s *S1ProtoLoginReq) Deserialize(bytes []byte) error {
	s.Account = string(bytes[0:32])
	s.Password = string(bytes[32:64])
	s.Version = math.Float32frombits(binary.LittleEndian.Uint32(bytes[64:68]))
	s.Tid = binary.LittleEndian.Uint32(bytes[68:72])
	s.Teamid = binary.LittleEndian.Uint32(bytes[72:76])
	s.Hash = int64(binary.LittleEndian.Uint64(bytes[76:84]))

	return nil
}
