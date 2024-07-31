package proto

import "encoding/binary"

const S1ProtoRoomInfosNotifySize = 100

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoRoomInfosNotify
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

using System.Collections.Generic;

#nullable disable
public class S1ProtoRoomInfosNotify : S1ProtoBase
{
  public int players_count;
  public List<ulong> players_uid = new List<ulong>();
  public List<uint> players_tid = new List<uint>();
  public List<sbyte> players_teamId = new List<sbyte>();
  public List<sbyte> players_seatIndex = new List<sbyte>();
  public string wifi_name;
  public string wifi_password;
  public string curr_match;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteInt(this.players_count);
    for (int index = 0; index < this.players_count; ++index)
    {
      buffer.WriteULong(this.players_uid[index]);
      buffer.WriteUInt(this.players_tid[index]);
      buffer.WriteSByte(this.players_teamId[index]);
      buffer.WriteSByte(this.players_seatIndex[index]);
    }
    buffer.WriteString(this.wifi_name, 32);
    buffer.WriteString(this.wifi_password, 32);
    buffer.WriteString(this.curr_match, 32);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.players_count = buffer.ReadInt();
    for (int index = 0; index < this.players_count; ++index)
    {
      this.players_uid.Add(buffer.ReadULong());
      this.players_tid.Add(buffer.ReadUInt());
      this.players_teamId.Add(buffer.ReadSByte());
      this.players_seatIndex.Add(buffer.ReadSByte());
    }
    this.wifi_name = buffer.ReadString(32);
    this.wifi_password = buffer.ReadString(32);
    this.curr_match = buffer.ReadString(32);
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 4 + 14 * this.players_count + 32 + 32 + 32;
}

*/

func (s *S1ProtoRoomInfosNotify) Serialize() []byte {
	bytes := make([]byte, 4+(int(s.PlayersCount)*14)+32+32+32)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.PlayersCount))
	for i := 0; i < int(s.PlayersCount); i++ {
		binary.LittleEndian.PutUint64(bytes[4+(i*14):], s.PlayersUid[i])
		binary.LittleEndian.PutUint32(bytes[12+(i*14):], s.PlayersTid[i])
		bytes[16+(i*14)] = s.PlayersTeamId[i]
		bytes[17+(i*14)] = s.PlayersSeatIndex[i]
	}
	copy(bytes[4+(int(s.PlayersCount)*14):], s.WifiName[:32])
	copy(bytes[36+(int(s.PlayersCount)*14):], s.WifiPassword[:32])
	copy(bytes[68+(int(s.PlayersCount)*14):], s.CurrMatch[:32])

	return bytes
}

func (s *S1ProtoRoomInfosNotify) Deserialize(bytes []byte) error {
	s.PlayersCount = int32(binary.LittleEndian.Uint32(bytes[0:]))
	for i := 0; i < int(s.PlayersCount); i++ {
		s.PlayersUid = append(s.PlayersUid, binary.LittleEndian.Uint64(bytes[4+(i*14):]))
		s.PlayersTid = append(s.PlayersTid, binary.LittleEndian.Uint32(bytes[12+(i*14):]))
		s.PlayersTeamId = append(s.PlayersTeamId, bytes[16+(i*14)])
		s.PlayersSeatIndex = append(s.PlayersSeatIndex, bytes[17+(i*14)])
	}
	s.WifiName = string(bytes[4+(int(s.PlayersCount)*14) : 36+(int(s.PlayersCount)*14)])
	s.WifiPassword = string(bytes[36+(int(s.PlayersCount)*14) : 68+(int(s.PlayersCount)*14)])
	s.CurrMatch = string(bytes[68+(int(s.PlayersCount)*14) : 100+(int(s.PlayersCount)*14)])
	return nil
}
