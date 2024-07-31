package proto

import (
	"encoding/binary"
	"math"
)

const S1BattleProto2023ClientRobotBattleInfoSize = 18

/*
// Decompiled with JetBrains decompiler
// Type: S1BattleProto2023ClientRobotBattleInfo
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

using System.ComponentModel;

#nullable disable
[Description("单个机器人战斗信息-RobotBattleInfo")]
public class S1BattleProto2023ClientRobotBattleInfo : S1ProtoBase
{
  [Description("机器人s0id号")]
  public byte Robots0id;
  [Description("是否脱战")]
  public byte BattleState;
  [Description("是否能远程补血")]
  public byte CanRemoteHeal;
  [Description("是否能远程兑换小弹丸")]
  public byte CanRemoteExchangeSmallBullet;
  [Description("是否能远程兑换大弹丸")]
  public byte CanRemoteExchangeBigBullet;
  [Description("剩余买活次数")]
  public byte RemaindBuyRevivalCount;
  [Description("当前买活所需花费")]
  public int BuyRevivalPrice;
  [Description("脱战倒计时")]
  public float BattleStateCountDown;
  [Description("是否在补血点IsRobotOnHppoint1Area")]
  public byte IsOnHpPoint;
  [Description("底盘控电方式 默认0，0:规则 1:裁判控制")]
  public byte ChassisPowerCtrl;
  [Description("云台控电方式 默认0，0:规则 1:裁判控制")]
  public byte GimbalPowerCtrl;
  [Description("枪口控电方式 默认0，0:规则 1:裁判控制")]
  public byte ShooterPowerCtrl;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteByte(this.Robots0id);
    buffer.WriteByte(this.BattleState);
    buffer.WriteByte(this.CanRemoteHeal);
    buffer.WriteByte(this.CanRemoteExchangeSmallBullet);
    buffer.WriteByte(this.CanRemoteExchangeBigBullet);
    buffer.WriteByte(this.RemaindBuyRevivalCount);
    buffer.WriteInt(this.BuyRevivalPrice);
    buffer.WriteFloat(this.BattleStateCountDown);
    buffer.WriteByte(this.IsOnHpPoint);
    buffer.WriteByte(this.ChassisPowerCtrl);
    buffer.WriteByte(this.GimbalPowerCtrl);
    buffer.WriteByte(this.ShooterPowerCtrl);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.Robots0id = buffer.ReadByte();
    this.BattleState = buffer.ReadByte();
    this.CanRemoteHeal = buffer.ReadByte();
    this.CanRemoteExchangeSmallBullet = buffer.ReadByte();
    this.CanRemoteExchangeBigBullet = buffer.ReadByte();
    this.RemaindBuyRevivalCount = buffer.ReadByte();
    this.BuyRevivalPrice = buffer.ReadInt();
    this.BattleStateCountDown = buffer.ReadFloat();
    this.IsOnHpPoint = buffer.ReadByte();
    this.ChassisPowerCtrl = buffer.ReadByte();
    this.GimbalPowerCtrl = buffer.ReadByte();
    this.ShooterPowerCtrl = buffer.ReadByte();
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 18;
}

*/

func (s *S1BattleProto2023ClientRobotBattleInfo) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2023ClientRobotBattleInfoSize)
	bytes[0] = s.Robots0Id
	bytes[1] = s.BattleState
	bytes[2] = s.CanRemoteHeal
	bytes[3] = s.CanRemoteExchangeSmallBullet
	bytes[4] = s.CanRemoteExchangeBigBullet
	bytes[5] = s.RemaindBuyRevivalCount
	binary.LittleEndian.PutUint32(bytes[6:], uint32(s.BuyRevivalPrice))
	binary.LittleEndian.PutUint32(bytes[10:], math.Float32bits(s.BattleStateCountDown))
	bytes[14] = s.IsOnHpPoint
	bytes[15] = s.ChassisPowerCtrl
	bytes[16] = s.GimbalPowerCtrl
	bytes[17] = s.ShooterPowerCtrl
	return bytes
}

func (s *S1BattleProto2023ClientRobotBattleInfo) Deserialize(bytes []byte) error {
	s.Robots0Id = bytes[0]
	s.BattleState = bytes[1]
	s.CanRemoteHeal = bytes[2]
	s.CanRemoteExchangeSmallBullet = bytes[3]
	s.CanRemoteExchangeBigBullet = bytes[4]
	s.RemaindBuyRevivalCount = bytes[5]
	s.BuyRevivalPrice = int32(binary.LittleEndian.Uint32(bytes[6:]))
	s.BattleStateCountDown = math.Float32frombits(binary.LittleEndian.Uint32(bytes[10:]))
	s.IsOnHpPoint = bytes[14]
	s.ChassisPowerCtrl = bytes[15]
	s.GimbalPowerCtrl = bytes[16]
	s.ShooterPowerCtrl = bytes[17]
	return nil
}
