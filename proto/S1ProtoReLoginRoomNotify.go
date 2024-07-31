package proto

import "encoding/binary"

const S1ProtoReLoginRoomNotifySize = 102

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoReLoginRoomNotify
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

using System.Collections.Generic;

#nullable disable
public class S1ProtoReLoginRoomNotify : S1ProtoBase
{
  public ulong self_uid;
  public uint self_tid;
  public sbyte self_teamId;
  public sbyte self_seatIndex;
  public int room_state;
  public int room_state_time_left;
  public int room_seat_islock;
  public int room_teams_count;
  public int room_teams_player_max;
  public int room_players_count;
  public List<ulong> players_uid = new List<ulong>();
  public List<uint> players_tid = new List<uint>();
  public List<sbyte> players_teamId = new List<sbyte>();
  public List<sbyte> players_seatIndex = new List<sbyte>();
  public string wifi_name;
  public string wifi_password;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteULong(this.self_uid);
    buffer.WriteUInt(this.self_tid);
    buffer.WriteSByte(this.self_teamId);
    buffer.WriteSByte(this.self_seatIndex);
    buffer.WriteInt(this.room_state);
    buffer.WriteInt(this.room_state_time_left);
    buffer.WriteInt(this.room_seat_islock);
    buffer.WriteInt(this.room_teams_count);
    buffer.WriteInt(this.room_teams_player_max);
    buffer.WriteInt(this.room_players_count);
    for (int index = 0; index < this.room_players_count; ++index)
    {
      buffer.WriteULong(this.players_uid[index]);
      buffer.WriteUInt(this.players_tid[index]);
      buffer.WriteSByte(this.players_teamId[index]);
      buffer.WriteSByte(this.players_seatIndex[index]);
    }
    buffer.WriteString(this.wifi_name, 32);
    buffer.WriteString(this.wifi_password, 32);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.self_uid = buffer.ReadULong();
    this.self_tid = buffer.ReadUInt();
    this.self_teamId = buffer.ReadSByte();
    this.self_seatIndex = buffer.ReadSByte();
    this.room_state = buffer.ReadInt();
    this.room_state_time_left = buffer.ReadInt();
    this.room_seat_islock = buffer.ReadInt();
    this.room_teams_count = buffer.ReadInt();
    this.room_teams_player_max = buffer.ReadInt();
    this.room_players_count = buffer.ReadInt();
    for (int index = 0; index < this.room_players_count; ++index)
    {
      this.players_uid.Add(buffer.ReadULong());
      this.players_tid.Add(buffer.ReadUInt());
      this.players_teamId.Add(buffer.ReadSByte());
      this.players_seatIndex.Add(buffer.ReadSByte());
    }
    this.wifi_name = buffer.ReadString(32);
    this.wifi_password = buffer.ReadString(32);
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 38 + 14 * this.room_players_count + 32 + 32;
}

*/

func (s *S1ProtoReLoginRoomNotify) Serialize() []byte {
	bytes := make([]byte, 38+(int(s.RoomPlayersCount)*14)+32+32)
	binary.LittleEndian.PutUint64(bytes[0:], s.SelfUid)
	binary.LittleEndian.PutUint32(bytes[8:], s.SelfTid)
	bytes[12] = s.SelfTeamId
	bytes[13] = s.SelfSeatIndex
	binary.LittleEndian.PutUint32(bytes[14:], uint32(s.RoomState))
	binary.LittleEndian.PutUint32(bytes[18:], uint32(s.RoomStateTimeLeft))
	binary.LittleEndian.PutUint32(bytes[22:], uint32(s.RoomSeatIslock))
	binary.LittleEndian.PutUint32(bytes[26:], uint32(s.RoomTeamsCount))
	binary.LittleEndian.PutUint32(bytes[30:], uint32(s.RoomTeamsPlayerMax))
	binary.LittleEndian.PutUint32(bytes[34:], uint32(s.RoomPlayersCount))
	for i := 0; i < int(s.RoomPlayersCount); i++ {
		binary.LittleEndian.PutUint64(bytes[38+(i*14):], s.PlayersUid[i])
		binary.LittleEndian.PutUint32(bytes[46+(i*14):], s.PlayersTid[i])
		bytes[50+(i*14)] = s.PlayersTeamId[i]
		bytes[51+(i*14)] = s.PlayersSeatIndex[i]
	}
	copy(bytes[38+(int(s.RoomPlayersCount)*14):], s.WifiName)
	copy(bytes[70+(int(s.RoomPlayersCount)*14):], s.WifiPassword)

	return bytes
}

func (s *S1ProtoReLoginRoomNotify) Deserialize(bytes []byte) error {
	s.SelfUid = binary.LittleEndian.Uint64(bytes[0:])
	s.SelfTid = binary.LittleEndian.Uint32(bytes[8:])
	s.SelfTeamId = bytes[12]
	s.SelfSeatIndex = bytes[13]
	s.RoomState = int32(binary.LittleEndian.Uint32(bytes[14:]))
	s.RoomStateTimeLeft = int32(binary.LittleEndian.Uint32(bytes[18:]))
	s.RoomSeatIslock = int32(binary.LittleEndian.Uint32(bytes[22:]))
	s.RoomTeamsCount = int32(binary.LittleEndian.Uint32(bytes[26:]))
	s.RoomTeamsPlayerMax = int32(binary.LittleEndian.Uint32(bytes[30:]))
	s.RoomPlayersCount = int32(binary.LittleEndian.Uint32(bytes[34:]))
	for i := 0; i < int(s.RoomPlayersCount); i++ {
		s.PlayersUid = append(s.PlayersUid, binary.LittleEndian.Uint64(bytes[38+(i*14):]))
		s.PlayersTid = append(s.PlayersTid, binary.LittleEndian.Uint32(bytes[46+(i*14):]))
		s.PlayersTeamId = append(s.PlayersTeamId, bytes[50+(i*14)])
		s.PlayersSeatIndex = append(s.PlayersSeatIndex, bytes[51+(i*14)])
	}
	s.WifiName = string(bytes[38+(int(s.RoomPlayersCount)*14) : 70+(int(s.RoomPlayersCount)*14)])
	s.WifiPassword = string(bytes[70+(int(s.RoomPlayersCount)*14) : 102+(int(s.RoomPlayersCount)*14)])

	return nil
}
