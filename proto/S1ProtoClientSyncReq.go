package proto

import "encoding/binary"

const S1ProtoClientSyncReqSize = 29

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoClientSyncReq
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

using System.Collections.Generic;

#nullable disable
public class S1ProtoClientSyncReq : S1ProtoBase
{
  public ulong uid;
  public int conn_client_state;
  public int conn_robot_state;
  public int battery;
  public int signal_quality;
  public byte battle_mode;
  public int online_module_count;
  public List<ulong> online_module = new List<ulong>();

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteULong(this.uid);
    buffer.WriteInt(this.conn_client_state);
    buffer.WriteInt(this.conn_robot_state);
    buffer.WriteInt(this.battery);
    buffer.WriteInt(this.signal_quality);
    buffer.WriteByte(this.battle_mode);
    buffer.WriteInt(this.online_module_count);
    for (int index = 0; index < this.online_module_count; ++index)
      buffer.WriteULong(this.online_module[index]);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.uid = buffer.ReadULong();
    this.conn_client_state = buffer.ReadInt();
    this.conn_robot_state = buffer.ReadInt();
    this.battery = buffer.ReadInt();
    this.signal_quality = buffer.ReadInt();
    this.battle_mode = buffer.ReadByte();
    this.online_module_count = buffer.ReadInt();
    for (int index = 0; index < this.online_module_count; ++index)
      this.online_module.Add(buffer.ReadULong());
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 29 + 8 * this.online_module_count;
}

*/

func (s *S1ProtoClientSyncReq) Serialize() []byte {
	bytes := make([]byte, 29+8*len(s.OnlineModule))
	binary.LittleEndian.PutUint64(bytes[0:], s.Uid)
	binary.LittleEndian.PutUint32(bytes[8:], uint32(s.ConnClientState))
	binary.LittleEndian.PutUint32(bytes[12:], uint32(s.ConnRobotState))
	binary.LittleEndian.PutUint32(bytes[16:], uint32(s.Battery))
	binary.LittleEndian.PutUint32(bytes[20:], uint32(s.SignalQuality))
	bytes[24] = s.BattleMode
	binary.LittleEndian.PutUint32(bytes[25:], uint32(s.OnlineModuleCount))
	for i := 0; i < len(s.OnlineModule); i++ {
		binary.LittleEndian.PutUint64(bytes[29+i*8:], s.OnlineModule[i])
	}
	return bytes
}

func (s *S1ProtoClientSyncReq) Deserialize(bytes []byte) error {
	s.Uid = binary.LittleEndian.Uint64(bytes[0:])
	s.ConnClientState = int32(binary.LittleEndian.Uint32(bytes[8:]))
	s.ConnRobotState = int32(binary.LittleEndian.Uint32(bytes[12:]))
	s.Battery = int32(binary.LittleEndian.Uint32(bytes[16:]))
	s.SignalQuality = int32(binary.LittleEndian.Uint32(bytes[20:]))
	s.BattleMode = bytes[24]
	s.OnlineModuleCount = int32(binary.LittleEndian.Uint32(bytes[25:]))
	s.OnlineModule = make([]uint64, 0, s.OnlineModuleCount)
	for i := 0; i < int(s.OnlineModuleCount); i++ {
		s.OnlineModule[i] = binary.LittleEndian.Uint64(bytes[29+i*8:])
	}
	return nil
}
