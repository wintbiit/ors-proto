package proto

import "encoding/binary"

const S1ProtoTeamInfoNotifySize = 4

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoTeamInfoNotify
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

#nullable disable
public class S1ProtoTeamInfoNotify : S1ProtoBase
{
  public int dataLen;
  public string data;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteInt(this.dataLen);
    buffer.WriteString(this.data, this.dataLen);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.dataLen = buffer.ReadInt();
    this.data = buffer.ReadString(this.dataLen);
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 4 + this.dataLen;
}

*/

func (s *S1ProtoTeamInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoTeamInfoNotifySize+s.DataLen)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.DataLen))
	copy(bytes[4:], s.Data)
	return bytes
}

func (s *S1ProtoTeamInfoNotify) Deserialize(bytes []byte) error {
	s.DataLen = int32(binary.LittleEndian.Uint32(bytes[0:]))
	s.Data = string(bytes[4:])
	return nil
}
