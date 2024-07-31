package proto

import (
	"encoding/binary"
	"math"
)

const S1ProtoMarkDetectResultSize = 97

/*
// Decompiled with JetBrains decompiler
// Type: S1ProtoMarkDetectResult
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

using System.ComponentModel;

#nullable disable
[Description("S1ProtoClientNetworkInfoNotify")]
public class S1ProtoMarkDetectResult : S1ProtoBase
{
  public float x;
  public float y;
  public float w;
  public float h;
  public float pitch;
  public float yaw;
  public float roll;
  public float[] T44_c2m = new float[16];
  public byte color;
  public ushort marker_id;
  public ushort distance;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteFloat(this.x);
    buffer.WriteFloat(this.y);
    buffer.WriteFloat(this.w);
    buffer.WriteFloat(this.h);
    buffer.WriteFloat(this.pitch);
    buffer.WriteFloat(this.yaw);
    buffer.WriteFloat(this.roll);
    for (int index = 0; index < 16; ++index)
      buffer.WriteFloat(this.T44_c2m[index]);
    buffer.WriteByte(this.color);
    buffer.WriteUShort(this.marker_id);
    buffer.WriteUShort(this.distance);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.x = buffer.ReadFloat();
    this.y = buffer.ReadFloat();
    this.w = buffer.ReadFloat();
    this.h = buffer.ReadFloat();
    this.pitch = buffer.ReadFloat();
    this.yaw = buffer.ReadFloat();
    this.roll = buffer.ReadFloat();
    for (int index = 0; index < 16; ++index)
      this.T44_c2m[index] = buffer.ReadFloat();
    this.color = buffer.ReadByte();
    this.marker_id = buffer.ReadUShort();
    this.distance = buffer.ReadUShort();
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 97;
}

*/

func (s *S1ProtoMarkDetectResult) Serialize() []byte {
	bytes := make([]byte, S1ProtoMarkDetectResultSize)
	binary.LittleEndian.PutUint32(bytes[0:], math.Float32bits(s.X))
	binary.LittleEndian.PutUint32(bytes[4:], math.Float32bits(s.Y))
	binary.LittleEndian.PutUint32(bytes[8:], math.Float32bits(s.W))
	binary.LittleEndian.PutUint32(bytes[12:], math.Float32bits(s.H))
	binary.LittleEndian.PutUint32(bytes[16:], math.Float32bits(s.Pitch))
	binary.LittleEndian.PutUint32(bytes[20:], math.Float32bits(s.Yaw))
	binary.LittleEndian.PutUint32(bytes[24:], math.Float32bits(s.Roll))
	for i := 0; i < 16; i++ {
		binary.LittleEndian.PutUint32(bytes[28+i*4:], math.Float32bits(s.T44C2M[i]))
	}
	bytes[92] = s.Color
	binary.LittleEndian.PutUint16(bytes[93:], s.MarkerId)
	binary.LittleEndian.PutUint16(bytes[95:], s.Distance)

	return bytes
}

func (s *S1ProtoMarkDetectResult) Deserialize(bytes []byte) error {
	s.X = math.Float32frombits(binary.LittleEndian.Uint32(bytes[0:]))
	s.Y = math.Float32frombits(binary.LittleEndian.Uint32(bytes[4:]))
	s.W = math.Float32frombits(binary.LittleEndian.Uint32(bytes[8:]))
	s.H = math.Float32frombits(binary.LittleEndian.Uint32(bytes[12:]))
	s.Pitch = math.Float32frombits(binary.LittleEndian.Uint32(bytes[16:]))
	s.Yaw = math.Float32frombits(binary.LittleEndian.Uint32(bytes[20:]))
	s.Roll = math.Float32frombits(binary.LittleEndian.Uint32(bytes[24:]))
	for i := 0; i < 16; i++ {
		s.T44C2M[i] = math.Float32frombits(binary.LittleEndian.Uint32(bytes[28+i*4:]))
	}
	s.Color = bytes[92]
	s.MarkerId = binary.LittleEndian.Uint16(bytes[93:])
	s.Distance = binary.LittleEndian.Uint16(bytes[95:])

	return nil
}
