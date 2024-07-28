package proto

type S1ProtoNetworkInfoNotify struct {
	WifiUplinkSpeed   uint32
	VtDownlinkSpeed   uint32
	WifiUplinkPlr     uint16
	VtUplinkPlr       uint16
	VtDownlinkPlr     uint16
	OtherPlr          uint16
	Delay             uint32
	Robotid           byte
	VtUplinkSpeed     uint32
	WifiDownlinkPlr   uint16
	OtherType         byte
	OtherDelay        uint16
	WifiDownlinkSpeed uint32
}

const S1ProtoNetworkInfoNotifySize = 34

func (s *S1ProtoNetworkInfoNotify) Serialize() []byte {
	bytes := make([]byte, S1ProtoNetworkInfoNotifySize)
	// TODO: Implement serialization
	return bytes
}

func (s *S1ProtoNetworkInfoNotify) Deserialize(bytes []byte) error {
	// TODO: Implement deserialization
	return nil
}
