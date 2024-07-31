package proto

import "encoding/binary"

const S1ProtoHeartBeatAckSize = 5

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoHeartBeatAck
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

#nullable disable
public class S1ProtoHeartBeatAck : S1ProtoBase
{
  public int resultID;
  public byte s0clientid;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteInt(this.resultID);
    buffer.WriteByte(this.s0clientid);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    uint num1 = (uint) (0 + 4);
    if (header.dataLen >= num1)
      this.resultID = buffer.ReadInt();
    uint num2 = num1 + 1U;
    if (header.dataLen >= num2)
      this.s0clientid = buffer.ReadByte();
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 5;
}

*/

func (s *S1ProtoHeartBeatAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoHeartBeatAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.ResultId))
	bytes[4] = s.S0Clientid
	return bytes
}

func (s *S1ProtoHeartBeatAck) Deserialize(bytes []byte) error {
	s.ResultId = int32(binary.LittleEndian.Uint32(bytes[0:]))
	s.S0Clientid = bytes[4]
	return nil
}
