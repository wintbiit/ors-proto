package proto

type S1BattleProto2022IoctrState struct {
	Dp3Adc1In6        byte
	IsPwmDp2Io4       byte
	IsPwmDp2Io7       byte
	ModuleType        byte
	Dp1Gpio1          byte
	Dp1Gpio3          byte
	Dp1Gpio6          byte
	Dp3Adc1In0        byte
	IsPwmDp2Io5       byte
	Dp2OutVars        []uint16
	Dp2PwmFrq         uint16
	Dp3Adc1In1        byte
	IsPwmDp2Io2       byte
	Dp2OutValsListLen int32
	Dp3AdcVals        []uint16
	Error             byte
	ModuleId          byte
	Dp1Gpio7          byte
	Dp3Adc1In4        byte
	Dp3AdcValsListLen int32
	IsPwmDp2Io3       byte
	Dp1Gpio0          byte
	Dp1Gpio4          byte
	Dp3Adc1In7        byte
	IsPwmDp2Io1       byte
	IsPwmDp2Io6       byte
	Dp1Gpio5          byte
	Dp3Adc1In3        byte
	IsPwmDp2Io0       byte
	Dp3Adc1In2        byte
	Dp3Adc1In5        byte
	Dp1Gpio2          byte
}

const S1BattleProto2022IoctrStateSize = 69

func (s *S1BattleProto2022IoctrState) Serialize() []byte {
	bytes := make([]byte, S1BattleProto2022IoctrStateSize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1BattleProto2022IoctrState) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
