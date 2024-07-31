package proto

import "encoding/binary"

const S1ProtoSvrStateAckSize = 76

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoSvrStateAck
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

#nullable disable
public class S1ProtoSvrStateAck : S1ProtoBase
{
  public e_Room_StateType state;
  public int time_left;
  public int is_pause;
  public string curr_match;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteInt((int) this.state);
    buffer.WriteInt(this.time_left);
    buffer.WriteInt(this.is_pause);
    buffer.WriteString(this.curr_match, 64);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.state = (e_Room_StateType) buffer.ReadInt();
    this.time_left = buffer.ReadInt();
    this.is_pause = buffer.ReadInt();
    this.curr_match = buffer.ReadString(64);
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 76;
}

*/

func (s *S1ProtoSvrStateAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoSvrStateAckSize)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(s.State))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(s.TimeLeft))
	binary.LittleEndian.PutUint32(bytes[8:12], uint32(s.IsPause))
	copy(bytes[12:76], s.CurrMatch)
	return bytes
}

func (s *S1ProtoSvrStateAck) Deserialize(bytes []byte) error {
	s.State = ERoomStateType(binary.LittleEndian.Uint32(bytes[0:4]))
	s.TimeLeft = int32(binary.LittleEndian.Uint32(bytes[4:8]))
	s.IsPause = int32(binary.LittleEndian.Uint32(bytes[8:12]))
	s.CurrMatch = string(bytes[12:76])
	return nil
}
