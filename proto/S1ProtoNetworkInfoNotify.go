package proto

import "encoding/binary"

const S1ProtoNetworkInfoNotifySize = 34

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoNetworkInfoNotify
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

using System.ComponentModel;

#nullable disable
[Description("S1ProtoNetworkInfoNotify")]
public class S1ProtoNetworkInfoNotify : S1ProtoBase
{
  public byte robotid;
  public uint wifi_uplink_speed;
  public uint wifi_downlink_speed;
  public uint vt_uplink_speed;
  public uint vt_downlink_speed;
  public ushort wifi_uplink_PLR;
  public ushort wifi_downlink_PLR;
  public ushort vt_uplink_PLR;
  public ushort vt_downlink_PLR;
  public byte other_type;
  public ushort other_PLR;
  public ushort other_delay;
  public uint delay;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteByte(this.robotid);
    buffer.WriteUInt(this.wifi_uplink_speed);
    buffer.WriteUInt(this.wifi_downlink_speed);
    buffer.WriteUInt(this.vt_uplink_speed);
    buffer.WriteUInt(this.vt_downlink_speed);
    buffer.WriteUShort(this.wifi_uplink_PLR);
    buffer.WriteUShort(this.wifi_downlink_PLR);
    buffer.WriteUShort(this.vt_uplink_PLR);
    buffer.WriteUShort(this.vt_downlink_PLR);
    buffer.WriteByte(this.other_type);
    buffer.WriteUShort(this.other_PLR);
    buffer.WriteUShort(this.other_delay);
    buffer.WriteUInt(this.delay);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.robotid = buffer.ReadByte();
    this.wifi_uplink_speed = buffer.ReadUInt();
    this.wifi_downlink_speed = buffer.ReadUInt();
    this.vt_uplink_speed = buffer.ReadUInt();
    this.vt_downlink_speed = buffer.ReadUInt();
    this.wifi_uplink_PLR = buffer.ReadUShort();
    this.wifi_downlink_PLR = buffer.ReadUShort();
    this.vt_uplink_PLR = buffer.ReadUShort();
    this.vt_downlink_PLR = buffer.ReadUShort();
    this.other_type = buffer.ReadByte();
    this.other_PLR = buffer.ReadUShort();
    this.other_delay = buffer.ReadUShort();
    this.delay = buffer.ReadUInt();
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 34;
}

*/

func (s *S1ProtoNetworkInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoNetworkInfoNotifySize)
	bytes[0] = s.Robotid
	binary.LittleEndian.PutUint32(bytes[1:], s.WifiUplinkSpeed)
	binary.LittleEndian.PutUint32(bytes[5:], s.WifiDownlinkSpeed)
	binary.LittleEndian.PutUint32(bytes[9:], s.VtUplinkSpeed)
	binary.LittleEndian.PutUint32(bytes[13:], s.VtDownlinkSpeed)
	binary.LittleEndian.PutUint16(bytes[17:], s.WifiUplinkPlr)
	binary.LittleEndian.PutUint16(bytes[19:], s.WifiDownlinkPlr)
	binary.LittleEndian.PutUint16(bytes[21:], s.VtUplinkPlr)
	binary.LittleEndian.PutUint16(bytes[23:], s.VtDownlinkPlr)
	bytes[25] = s.OtherType
	binary.LittleEndian.PutUint16(bytes[26:], s.OtherPlr)
	binary.LittleEndian.PutUint16(bytes[28:], s.OtherDelay)
	binary.LittleEndian.PutUint32(bytes[30:], s.Delay)

	return bytes
}

func (s *S1ProtoNetworkInfoNotify) Deserialize(bytes []byte) error {
	if len(bytes) != S1ProtoNetworkInfoNotifySize {
		return InvalidDataError
	}

	s.Robotid = bytes[0]
	s.WifiUplinkSpeed = binary.LittleEndian.Uint32(bytes[1:])
	s.WifiDownlinkSpeed = binary.LittleEndian.Uint32(bytes[5:])
	s.VtUplinkSpeed = binary.LittleEndian.Uint32(bytes[9:])
	s.VtDownlinkSpeed = binary.LittleEndian.Uint32(bytes[13:])
	s.WifiUplinkPlr = binary.LittleEndian.Uint16(bytes[17:])
	s.WifiDownlinkPlr = binary.LittleEndian.Uint16(bytes[19:])
	s.VtUplinkPlr = binary.LittleEndian.Uint16(bytes[21:])
	s.VtDownlinkPlr = binary.LittleEndian.Uint16(bytes[23:])
	s.OtherType = bytes[25]
	s.OtherPlr = binary.LittleEndian.Uint16(bytes[26:])
	s.OtherDelay = binary.LittleEndian.Uint16(bytes[28:])
	s.Delay = binary.LittleEndian.Uint32(bytes[30:])

	return nil
}
