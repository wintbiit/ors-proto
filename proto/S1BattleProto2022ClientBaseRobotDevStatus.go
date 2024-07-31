package proto

import "encoding/binary"

const S1BattleProto2022ClientBaseRobotDevStatusSize = 44

/*
// Decompiled with JetBrains decompiler
// Type: S1BattleProto2022ClientBaseRobotDevStatus
// Assembly: Assembly-CSharp, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
// MVID: 7B09A0F1-4657-4604-BC1E-E6B74CAD8CE9
// Assembly location: C:\Users\WintBit\opt\Game\RM\RoboMasterEngine 10.0.0.2.stu\RoboMasterEngine_Data\Managed\Assembly-CSharp.dll

using System.ComponentModel;

#nullable disable
[Description("基地信息同步客户端-BaseRobotDevStatus")]
public class S1BattleProto2022ClientBaseRobotDevStatus : S1ProtoBase
{
  [Description("基地设备外壳当前状态")]
  public byte ShellStatus;
  [Description("基地设备是否在线  1：在线  2.离线 ")]
  public byte Online;
  [Description("飞镖靶实时位置")]
  public short DartTargetPosition;
  [Description("飞镖靶是否到达点位")]
  public byte DartTargetIsInplace;
  [Description("基地护甲一号传感器 0 正常 1 异常")]
  public byte BaseShellOneSensorStatus;
  [Description("基地护甲2号传感器 0 正常 1 异常")]
  public byte BaseShellTwoSensorStatus;
  [Description("基地护甲3号传感器 0 正常 1 异常")]
  public byte BaseShellThreeSensorStatus;
  [Description("飞镖靶传感器 0 正常 1 异常")]
  public byte BaseDartboardSensorStatus;
  [Description("基地护甲1号电机堵转状态 0 正常 1 异常")]
  public byte BaseShellMotoOneBlockStatus;
  [Description("基地护甲1号电机在线状态 0 在线 1 离线")]
  public byte BaseShellMotoOneOnlineStatus;
  [Description("基地护甲2号电机堵转状态  0 正常 1 堵转")]
  public byte BaseShellMotoTwoBlockStatus;
  [Description("基地护甲2号电机在线状态 0 在线 1 离线")]
  public byte BaseShellMotoTwoOnlineStatus;
  [Description("基地护甲3号电机堵转状态 0 正常 1 异常")]
  public byte BaseShellMotoThreeBlockStatus;
  [Description("基地护甲3号电机在线状态 0 在线 1 离线")]
  public byte BaseShellMotoThreeOnlineStatus;
  [Description("飞镖靶电机堵转状态 0 正常 1 异常")]
  public byte BaseDartboardMotoBlockStatus;
  [Description("飞镖靶电机在线状态 0 在线 1 离线")]
  public byte BaseDartboardMotoOnlineStatus;
  [Description("基地护甲自检结果状态 0 正常 1 失败")]
  public byte BaseShellSelfcheckStatus;
  [Description("飞镖靶自检结果状态 0 正常 1 失败")]
  public byte BaseDartboardSelfcheckStatus;
  [Description("飞镖靶移动超时 0 正常 1 超时")]
  public byte BaseDartboardMoveTimeout;
  [Description("外壳开启超时  0 正常 1 超时")]
  public byte BaseShellOpenTimeout;
  [Description("外壳关闭超时  0 正常 1 超时")]
  public byte BaseShellCloseTimeout;
  [Description("飞镖靶红外模块离线 0 在线 1为离线")]
  public byte DartBoardIROnline;
  [Description("飞镖检测靶自检后的异常码 0 正常 1 接收板红外灯珠损坏 2 飞镖引导灯损坏 3 红外灯珠和引导灯均损坏")]
  public byte DartBoard_BrokenErr;
  [Description("飞镖检测靶自检后的 左接收板灯珠损坏情况，0-31bit代表对应灯珠ID")]
  public uint IRledStatusLeft;
  [Description("飞镖检测靶自检后的 右接收板灯珠损坏情况，0-31bit代表对应灯珠ID")]
  public uint IRledStatusRight;
  [Description("基地护甲1号电机过热  0 正常 1 异常")]
  public byte BaseShellMotoOneOverHeat;
  [Description("基地护甲2号电机过热 0 正常 1 异常")]
  public byte BaseShellMotoTwoOverHeat;
  [Description("基地护甲3号电机过热 0 正常 1 异常")]
  public byte BaseShellMotoThreeOverHeat;
  [Description("基地靶子电机过热 0 正常 1 异常")]
  public byte BaseDartBoardMotoOverHeat;
  [Description("基地护甲1号电机电流过大 0 正常 1 异常")]
  public byte BaseShellMotoOneOverCurrent;
  [Description("基地护甲2号电机电流过大 0 正常 1 异常")]
  public byte BaseShellMotoTwoOverCurrent;
  [Description("基地护甲3号电机电流过大 0 正常 1 异常")]
  public byte BaseShellMotoThreeOverCurrent;
  [Description("基地飞镖靶电机电流过大 0 正常 1 异常")]
  public byte BaseDartBoardMotoOverCurrent;
  [Description("基地护甲1号电机差速过大 0 正常 1 异常")]
  public byte BaseShellMotoOneOverSpeed;
  [Description("基地护甲2号电机差速过大 0 正常 1 异常")]
  public byte BaseShellMotoTwoOverSpeed;
  [Description("基地护甲3号电机差速过大 0 正常 1 异常")]
  public byte BaseShellMotoThreeOverSpeed;
  [Description("基地飞镖靶电机差速过大 0 正常 1 异常")]
  public byte BaseDartBoardMotoOverSpeed;

  public override void Pack(ByteBuffer buffer)
  {
    buffer.WriteByte(this.ShellStatus);
    buffer.WriteByte(this.Online);
    buffer.WriteShort(this.DartTargetPosition);
    buffer.WriteByte(this.DartTargetIsInplace);
    buffer.WriteByte(this.BaseShellOneSensorStatus);
    buffer.WriteByte(this.BaseShellTwoSensorStatus);
    buffer.WriteByte(this.BaseShellThreeSensorStatus);
    buffer.WriteByte(this.BaseDartboardSensorStatus);
    buffer.WriteByte(this.BaseShellMotoOneBlockStatus);
    buffer.WriteByte(this.BaseShellMotoOneOnlineStatus);
    buffer.WriteByte(this.BaseShellMotoTwoBlockStatus);
    buffer.WriteByte(this.BaseShellMotoTwoOnlineStatus);
    buffer.WriteByte(this.BaseShellMotoThreeBlockStatus);
    buffer.WriteByte(this.BaseShellMotoThreeOnlineStatus);
    buffer.WriteByte(this.BaseDartboardMotoBlockStatus);
    buffer.WriteByte(this.BaseDartboardMotoOnlineStatus);
    buffer.WriteByte(this.BaseShellSelfcheckStatus);
    buffer.WriteByte(this.BaseDartboardSelfcheckStatus);
    buffer.WriteByte(this.BaseDartboardMoveTimeout);
    buffer.WriteByte(this.BaseShellOpenTimeout);
    buffer.WriteByte(this.BaseShellCloseTimeout);
    buffer.WriteByte(this.DartBoardIROnline);
    buffer.WriteByte(this.DartBoard_BrokenErr);
    buffer.WriteUInt(this.IRledStatusLeft);
    buffer.WriteUInt(this.IRledStatusRight);
    buffer.WriteByte(this.BaseShellMotoOneOverHeat);
    buffer.WriteByte(this.BaseShellMotoTwoOverHeat);
    buffer.WriteByte(this.BaseShellMotoThreeOverHeat);
    buffer.WriteByte(this.BaseDartBoardMotoOverHeat);
    buffer.WriteByte(this.BaseShellMotoOneOverCurrent);
    buffer.WriteByte(this.BaseShellMotoTwoOverCurrent);
    buffer.WriteByte(this.BaseShellMotoThreeOverCurrent);
    buffer.WriteByte(this.BaseDartBoardMotoOverCurrent);
    buffer.WriteByte(this.BaseShellMotoOneOverSpeed);
    buffer.WriteByte(this.BaseShellMotoTwoOverSpeed);
    buffer.WriteByte(this.BaseShellMotoThreeOverSpeed);
    buffer.WriteByte(this.BaseDartBoardMotoOverSpeed);
  }

  public override S1ProtoBase UnPack(S1ProtoHeader header, ByteBuffer buffer)
  {
    this.ShellStatus = buffer.ReadByte();
    this.Online = buffer.ReadByte();
    this.DartTargetPosition = buffer.ReadShort();
    this.DartTargetIsInplace = buffer.ReadByte();
    this.BaseShellOneSensorStatus = buffer.ReadByte();
    this.BaseShellTwoSensorStatus = buffer.ReadByte();
    this.BaseShellThreeSensorStatus = buffer.ReadByte();
    this.BaseDartboardSensorStatus = buffer.ReadByte();
    this.BaseShellMotoOneBlockStatus = buffer.ReadByte();
    this.BaseShellMotoOneOnlineStatus = buffer.ReadByte();
    this.BaseShellMotoTwoBlockStatus = buffer.ReadByte();
    this.BaseShellMotoTwoOnlineStatus = buffer.ReadByte();
    this.BaseShellMotoThreeBlockStatus = buffer.ReadByte();
    this.BaseShellMotoThreeOnlineStatus = buffer.ReadByte();
    this.BaseDartboardMotoBlockStatus = buffer.ReadByte();
    this.BaseDartboardMotoOnlineStatus = buffer.ReadByte();
    this.BaseShellSelfcheckStatus = buffer.ReadByte();
    this.BaseDartboardSelfcheckStatus = buffer.ReadByte();
    this.BaseDartboardMoveTimeout = buffer.ReadByte();
    this.BaseShellOpenTimeout = buffer.ReadByte();
    this.BaseShellCloseTimeout = buffer.ReadByte();
    this.DartBoardIROnline = buffer.ReadByte();
    this.DartBoard_BrokenErr = buffer.ReadByte();
    this.IRledStatusLeft = buffer.ReadUInt();
    this.IRledStatusRight = buffer.ReadUInt();
    this.BaseShellMotoOneOverHeat = buffer.ReadByte();
    this.BaseShellMotoTwoOverHeat = buffer.ReadByte();
    this.BaseShellMotoThreeOverHeat = buffer.ReadByte();
    this.BaseDartBoardMotoOverHeat = buffer.ReadByte();
    this.BaseShellMotoOneOverCurrent = buffer.ReadByte();
    this.BaseShellMotoTwoOverCurrent = buffer.ReadByte();
    this.BaseShellMotoThreeOverCurrent = buffer.ReadByte();
    this.BaseDartBoardMotoOverCurrent = buffer.ReadByte();
    this.BaseShellMotoOneOverSpeed = buffer.ReadByte();
    this.BaseShellMotoTwoOverSpeed = buffer.ReadByte();
    this.BaseShellMotoThreeOverSpeed = buffer.ReadByte();
    this.BaseDartBoardMotoOverSpeed = buffer.ReadByte();
    return (S1ProtoBase) this;
  }

  public override int GetSize() => 44;
}

*/

func (s *S1BattleProto2022ClientBaseRobotDevStatus) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022ClientBaseRobotDevStatusSize)
	bytes[0] = s.ShellStatus
	bytes[1] = s.Online
	binary.LittleEndian.PutUint16(bytes[2:], uint16(s.DartTargetPosition))
	bytes[4] = s.DartTargetIsInplace
	bytes[5] = s.BaseShellOneSensorStatus
	bytes[6] = s.BaseShellTwoSensorStatus
	bytes[7] = s.BaseShellThreeSensorStatus
	bytes[8] = s.BaseDartboardSensorStatus
	bytes[9] = s.BaseShellMotoOneBlockStatus
	bytes[10] = s.BaseShellMotoOneOnlineStatus
	bytes[11] = s.BaseShellMotoTwoBlockStatus
	bytes[12] = s.BaseShellMotoTwoOnlineStatus
	bytes[13] = s.BaseShellMotoThreeBlockStatus
	bytes[14] = s.BaseShellMotoThreeOnlineStatus
	bytes[15] = s.BaseDartboardMotoBlockStatus
	bytes[16] = s.BaseDartboardMotoOnlineStatus
	bytes[17] = s.BaseShellSelfcheckStatus
	bytes[18] = s.BaseDartboardSelfcheckStatus
	bytes[19] = s.BaseDartboardMoveTimeout
	bytes[20] = s.BaseShellOpenTimeout
	bytes[21] = s.BaseShellCloseTimeout
	bytes[22] = s.DartBoardIronline
	bytes[23] = s.DartBoardBrokenErr
	binary.LittleEndian.PutUint32(bytes[24:], s.IrledStatusLeft)
	binary.LittleEndian.PutUint32(bytes[28:], s.IrledStatusRight)
	bytes[32] = s.BaseShellMotoOneOverHeat
	bytes[33] = s.BaseShellMotoTwoOverHeat
	bytes[34] = s.BaseShellMotoThreeOverHeat
	bytes[35] = s.BaseDartBoardMotoOverHeat
	bytes[36] = s.BaseShellMotoOneOverCurrent
	bytes[37] = s.BaseShellMotoTwoOverCurrent
	bytes[38] = s.BaseShellMotoThreeOverCurrent
	bytes[39] = s.BaseDartBoardMotoOverCurrent
	bytes[40] = s.BaseShellMotoOneOverSpeed
	bytes[41] = s.BaseShellMotoTwoOverSpeed
	bytes[42] = s.BaseShellMotoThreeOverSpeed
	bytes[43] = s.BaseDartBoardMotoOverSpeed

	return bytes
}

func (s *S1BattleProto2022ClientBaseRobotDevStatus) Deserialize(bytes []byte) error {
	if len(bytes) < S1BattleProto2022ClientBaseRobotDevStatusSize {
		return InvalidDataError
	}

	s.ShellStatus = bytes[0]
	s.Online = bytes[1]
	s.DartTargetPosition = int16(binary.LittleEndian.Uint16(bytes[2:]))
	s.DartTargetIsInplace = bytes[4]
	s.BaseShellOneSensorStatus = bytes[5]
	s.BaseShellTwoSensorStatus = bytes[6]
	s.BaseShellThreeSensorStatus = bytes[7]
	s.BaseDartboardSensorStatus = bytes[8]
	s.BaseShellMotoOneBlockStatus = bytes[9]
	s.BaseShellMotoOneOnlineStatus = bytes[10]
	s.BaseShellMotoTwoBlockStatus = bytes[11]
	s.BaseShellMotoTwoOnlineStatus = bytes[12]
	s.BaseShellMotoThreeBlockStatus = bytes[13]
	s.BaseShellMotoThreeOnlineStatus = bytes[14]
	s.BaseDartboardMotoBlockStatus = bytes[15]
	s.BaseDartboardMotoOnlineStatus = bytes[16]
	s.BaseShellSelfcheckStatus = bytes[17]
	s.BaseDartboardSelfcheckStatus = bytes[18]
	s.BaseDartboardMoveTimeout = bytes[19]
	s.BaseShellOpenTimeout = bytes[20]
	s.BaseShellCloseTimeout = bytes[21]
	s.DartBoardIronline = bytes[22]
	s.DartBoardBrokenErr = bytes[23]
	s.IrledStatusLeft = binary.LittleEndian.Uint32(bytes[24:])
	s.IrledStatusRight = binary.LittleEndian.Uint32(bytes[28:])
	s.BaseShellMotoOneOverHeat = bytes[32]
	s.BaseShellMotoTwoOverHeat = bytes[33]
	s.BaseShellMotoThreeOverHeat = bytes[34]
	s.BaseDartBoardMotoOverHeat = bytes[35]
	s.BaseShellMotoOneOverCurrent = bytes[36]
	s.BaseShellMotoTwoOverCurrent = bytes[37]
	s.BaseShellMotoThreeOverCurrent = bytes[38]
	s.BaseDartBoardMotoOverCurrent = bytes[39]
	s.BaseShellMotoOneOverSpeed = bytes[40]
	s.BaseShellMotoTwoOverSpeed = bytes[41]
	s.BaseShellMotoThreeOverSpeed = bytes[42]
	s.BaseDartBoardMotoOverSpeed = bytes[43]

	return nil
}
