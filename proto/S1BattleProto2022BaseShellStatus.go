package proto

const S1BattleProto2022BaseShellStatusSize = 12

/*
// Decompiled with JetBrains decompiler
// Type: S1BattleProto2022BaseShellStatus
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

using System.ComponentModel;

#nullable disable
[Description("基地外壳状态(上行)-BaseShellStatus")]
public class S1BattleProto2022BaseShellStatus : S1ProtoBase
{
  [Description("0:闭合状态  1:正在展开 2.正在闭合 3.展开状态")]
  public byte BaseStatus;
  [Description("一号电机状态; 0:正常  1：不正常")]
  public byte BaseMotoOneStatus;
  [Description("0:堵转  1：不堵转")]
  public byte BaseMotoOneBlockStatus;
  [Description("在线 1.不在线")]
  public byte BaseMotoOneOnlineStatus;
  [Description("二号电机状态; 0:正常  1：不正常")]
  public byte BaseMotoTwoStatus;
  [Description("0:堵转  1：不堵转")]
  public byte BaseMotoTwoBlockStatus;
  [Description("在线 1.不在线")]
  public byte BaseMotoTwoOnlineStatus;
  [Description("三号电机状态； 0:正常  1：不正常")]
  public byte BaseMotoThreeStatus;
  [Description("0:堵转  1：不堵转")]
  public byte BaseMotoThreeBlockStatus;
  [Description("在线 1.不在线")]
  public byte BaseMotoThreeOnlineStatus;
  public byte Reserved0;
  public byte Reserved2;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteByte(this.BaseStatus);
    buffer.WriteByte(this.BaseMotoOneStatus);
    buffer.WriteByte(this.BaseMotoOneBlockStatus);
    buffer.WriteByte(this.BaseMotoOneOnlineStatus);
    buffer.WriteByte(this.BaseMotoTwoStatus);
    buffer.WriteByte(this.BaseMotoTwoBlockStatus);
    buffer.WriteByte(this.BaseMotoTwoOnlineStatus);
    buffer.WriteByte(this.BaseMotoThreeStatus);
    buffer.WriteByte(this.BaseMotoThreeBlockStatus);
    buffer.WriteByte(this.BaseMotoThreeOnlineStatus);
    buffer.WriteByte(this.Reserved0);
    buffer.WriteByte(this.Reserved2);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.BaseStatus = buffer.ReadByte();
    this.BaseMotoOneStatus = buffer.ReadByte();
    this.BaseMotoOneBlockStatus = buffer.ReadByte();
    this.BaseMotoOneOnlineStatus = buffer.ReadByte();
    this.BaseMotoTwoStatus = buffer.ReadByte();
    this.BaseMotoTwoBlockStatus = buffer.ReadByte();
    this.BaseMotoTwoOnlineStatus = buffer.ReadByte();
    this.BaseMotoThreeStatus = buffer.ReadByte();
    this.BaseMotoThreeBlockStatus = buffer.ReadByte();
    this.BaseMotoThreeOnlineStatus = buffer.ReadByte();
    this.Reserved0 = buffer.ReadByte();
    this.Reserved2 = buffer.ReadByte();
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 12;
}

*/

func (s *S1BattleProto2022BaseShellStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022BaseShellStatusSize)
	bytes[0] = s.BaseStatus
	bytes[1] = s.BaseMotoOneStatus
	bytes[2] = s.BaseMotoOneBlockStatus
	bytes[3] = s.BaseMotoOneOnlineStatus
	bytes[4] = s.BaseMotoTwoStatus
	bytes[5] = s.BaseMotoTwoBlockStatus
	bytes[6] = s.BaseMotoTwoOnlineStatus
	bytes[7] = s.BaseMotoThreeStatus
	bytes[8] = s.BaseMotoThreeBlockStatus
	bytes[9] = s.BaseMotoThreeOnlineStatus
	bytes[10] = s.Reserved0
	bytes[11] = s.Reserved2

	return bytes
}

func (s *S1BattleProto2022BaseShellStatus) Deserialize(bytes []byte) error {
	if len(bytes) < S1BattleProto2022BaseShellStatusSize {
		return InvalidDataError
	}

	s.BaseStatus = bytes[0]
	s.BaseMotoOneStatus = bytes[1]
	s.BaseMotoOneBlockStatus = bytes[2]
	s.BaseMotoOneOnlineStatus = bytes[3]
	s.BaseMotoTwoStatus = bytes[4]
	s.BaseMotoTwoBlockStatus = bytes[5]
	s.BaseMotoTwoOnlineStatus = bytes[6]
	s.BaseMotoThreeStatus = bytes[7]
	s.BaseMotoThreeBlockStatus = bytes[8]
	s.BaseMotoThreeOnlineStatus = bytes[9]
	s.Reserved0 = bytes[10]
	s.Reserved2 = bytes[11]

	return nil
}
