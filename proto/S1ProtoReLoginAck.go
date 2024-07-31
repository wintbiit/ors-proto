package proto

import "encoding/binary"

const S1ProtoReLoginAckSize = 108

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoReLoginAck
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

#nullable disable
public class S1ProtoReLoginAck : S1ProtoBase
{
  public e_LoginAck_ResultType resultID;
  public ulong uid;
  public string token;
  public string wifi_name;
  public string wifi_password;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteInt((int) this.resultID);
    buffer.WriteULong(this.uid);
    buffer.WriteString(this.token, 32);
    buffer.WriteString(this.wifi_name, 32);
    buffer.WriteString(this.wifi_password, 32);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.resultID = (e_LoginAck_ResultType) buffer.ReadInt();
    this.uid = buffer.ReadULong();
    this.token = buffer.ReadString(32);
    this.wifi_name = buffer.ReadString(32);
    this.wifi_password = buffer.ReadString(32);
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 108;
}

*/

func (s *S1ProtoReLoginAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoReLoginAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.ResultId))
	binary.LittleEndian.PutUint64(bytes[4:], s.Uid)
	copy(bytes[12:], s.Token)
	copy(bytes[44:], s.WifiName)
	copy(bytes[76:], s.WifiPassword)
	return bytes
}

func (s *S1ProtoReLoginAck) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoReLoginAckSize {
		return InvalidDataError
	}

	s.ResultId = ELoginAckResultType(binary.LittleEndian.Uint32(bytes[0:]))
	s.Uid = binary.LittleEndian.Uint64(bytes[4:])
	s.Token = string(bytes[12:44])
	s.WifiName = string(bytes[44:76])
	s.WifiPassword = string(bytes[76:108])

	return nil
}
