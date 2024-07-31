package proto

import (
	"encoding/binary"
	"math"
)

const S1ProtoMarkDetectResultAckSize = 25

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoMarkDetectResultAck
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

using System.ComponentModel;

#nullable disable
[Description("S1ProtoClientNetworkInfoNotify")]
public class S1ProtoMarkDetectResultAck : S1ProtoBase
{
  public int result;
  public float x;
  public float y;
  public float w;
  public float h;
  public byte color;
  public ushort marker_id;
  public ushort distance;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteInt(this.result);
    buffer.WriteFloat(this.x);
    buffer.WriteFloat(this.y);
    buffer.WriteFloat(this.w);
    buffer.WriteFloat(this.h);
    buffer.WriteByte(this.color);
    buffer.WriteUShort(this.marker_id);
    buffer.WriteUShort(this.distance);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.result = buffer.ReadInt();
    this.x = buffer.ReadFloat();
    this.y = buffer.ReadFloat();
    this.w = buffer.ReadFloat();
    this.h = buffer.ReadFloat();
    this.color = buffer.ReadByte();
    this.marker_id = buffer.ReadUShort();
    this.distance = buffer.ReadUShort();
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 25;
}

*/

func (s *S1ProtoMarkDetectResultAck) Serialize() []byte {
	bytes := make([]byte, S1ProtoMarkDetectResultAckSize)
	binary.LittleEndian.PutUint32(bytes[0:], uint32(s.Result))
	binary.LittleEndian.PutUint32(bytes[4:], math.Float32bits(s.X))
	binary.LittleEndian.PutUint32(bytes[8:], math.Float32bits(s.Y))
	binary.LittleEndian.PutUint32(bytes[12:], math.Float32bits(s.W))
	binary.LittleEndian.PutUint32(bytes[16:], math.Float32bits(s.H))
	bytes[20] = s.Color
	binary.LittleEndian.PutUint16(bytes[21:], s.MarkerId)
	binary.LittleEndian.PutUint16(bytes[23:], s.Distance)

	return bytes
}

func (s *S1ProtoMarkDetectResultAck) Deserialize(bytes []byte) error {
	s.Result = int32(binary.LittleEndian.Uint32(bytes[0:]))
	s.X = math.Float32frombits(binary.LittleEndian.Uint32(bytes[4:]))
	s.Y = math.Float32frombits(binary.LittleEndian.Uint32(bytes[8:]))
	s.W = math.Float32frombits(binary.LittleEndian.Uint32(bytes[12:]))
	s.H = math.Float32frombits(binary.LittleEndian.Uint32(bytes[16:]))
	s.Color = bytes[20]
	s.MarkerId = binary.LittleEndian.Uint16(bytes[21:])
	s.Distance = binary.LittleEndian.Uint16(bytes[23:])

	return nil
}
