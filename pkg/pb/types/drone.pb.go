// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openpitrix/types/drone.proto

package pbtypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DroneId struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id"`
}

func (m *DroneId) Reset()                    { *m = DroneId{} }
func (m *DroneId) String() string            { return proto.CompactTextString(m) }
func (*DroneId) ProtoMessage()               {}
func (*DroneId) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *DroneId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DroneIdList struct {
	IdList []string `protobuf:"bytes,1,rep,name=id_list,json=idList" json:"id_list"`
}

func (m *DroneIdList) Reset()                    { *m = DroneIdList{} }
func (m *DroneIdList) String() string            { return proto.CompactTextString(m) }
func (*DroneIdList) ProtoMessage()               {}
func (*DroneIdList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *DroneIdList) GetIdList() []string {
	if m != nil {
		return m.IdList
	}
	return nil
}

type DroneConfig struct {
	Id             string `protobuf:"bytes,1,opt,name=id" json:"id"`
	Host           string `protobuf:"bytes,2,opt,name=host" json:"host"`
	ListenPort     int32  `protobuf:"varint,3,opt,name=listen_port,json=listenPort" json:"listen_port"`
	CmdInfoLogPath string `protobuf:"bytes,4,opt,name=cmd_info_log_path,json=cmdInfoLogPath" json:"cmd_info_log_path"`
	ConfdSelfHost  string `protobuf:"bytes,5,opt,name=confd_self_host,json=confdSelfHost" json:"confd_self_host"`
	LogLevel       string `protobuf:"bytes,6,opt,name=log_level,json=logLevel" json:"log_level"`
}

func (m *DroneConfig) Reset()                    { *m = DroneConfig{} }
func (m *DroneConfig) String() string            { return proto.CompactTextString(m) }
func (*DroneConfig) ProtoMessage()               {}
func (*DroneConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *DroneConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DroneConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *DroneConfig) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *DroneConfig) GetCmdInfoLogPath() string {
	if m != nil {
		return m.CmdInfoLogPath
	}
	return ""
}

func (m *DroneConfig) GetConfdSelfHost() string {
	if m != nil {
		return m.ConfdSelfHost
	}
	return ""
}

func (m *DroneConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type DroneEndpoint struct {
	FrontgateId string `protobuf:"bytes,1,opt,name=frontgate_id,json=frontgateId" json:"frontgate_id"`
	DroneIp     string `protobuf:"bytes,2,opt,name=drone_ip,json=droneIp" json:"drone_ip"`
	DronePort   int32  `protobuf:"varint,3,opt,name=drone_port,json=dronePort" json:"drone_port"`
}

func (m *DroneEndpoint) Reset()                    { *m = DroneEndpoint{} }
func (m *DroneEndpoint) String() string            { return proto.CompactTextString(m) }
func (*DroneEndpoint) ProtoMessage()               {}
func (*DroneEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *DroneEndpoint) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *DroneEndpoint) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *DroneEndpoint) GetDronePort() int32 {
	if m != nil {
		return m.DronePort
	}
	return 0
}

type SetDroneConfigRequest struct {
	Endpoint *DroneEndpoint `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint"`
	Config   *DroneConfig   `protobuf:"bytes,2,opt,name=config" json:"config"`
}

func (m *SetDroneConfigRequest) Reset()                    { *m = SetDroneConfigRequest{} }
func (m *SetDroneConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*SetDroneConfigRequest) ProtoMessage()               {}
func (*SetDroneConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SetDroneConfigRequest) GetEndpoint() *DroneEndpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *SetDroneConfigRequest) GetConfig() *DroneConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*DroneId)(nil), "openpitrix.types.DroneId")
	proto.RegisterType((*DroneIdList)(nil), "openpitrix.types.DroneIdList")
	proto.RegisterType((*DroneConfig)(nil), "openpitrix.types.DroneConfig")
	proto.RegisterType((*DroneEndpoint)(nil), "openpitrix.types.DroneEndpoint")
	proto.RegisterType((*SetDroneConfigRequest)(nil), "openpitrix.types.SetDroneConfigRequest")
}

func init() { proto.RegisterFile("openpitrix/types/drone.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xba, 0xbb, 0x69, 0x33, 0x61, 0x17, 0xb0, 0x84, 0xc8, 0x02, 0xab, 0x96, 0x1c, 0xaa,
	0x72, 0x20, 0x91, 0x8a, 0x38, 0xf5, 0xc6, 0x87, 0x44, 0xa4, 0x1e, 0xaa, 0xf4, 0xc6, 0xc5, 0x6a,
	0x63, 0x27, 0xb5, 0x70, 0x3d, 0x26, 0x31, 0x08, 0x7e, 0x03, 0xff, 0x8a, 0x5f, 0x86, 0x32, 0x89,
	0x4a, 0x28, 0xda, 0x53, 0xe2, 0xf7, 0xde, 0xbc, 0x99, 0x67, 0x0f, 0xbc, 0x40, 0x2b, 0x8d, 0x55,
	0xae, 0x56, 0x3f, 0x52, 0xf7, 0xd3, 0xca, 0x26, 0x15, 0x35, 0x1a, 0x99, 0xd8, 0x1a, 0x1d, 0xb2,
	0x47, 0x7f, 0xd9, 0x84, 0xd8, 0x67, 0xff, 0xeb, 0x0b, 0x34, 0xa5, 0xe8, 0xf4, 0xf1, 0x2d, 0x8c,
	0x3f, 0xb4, 0xe5, 0x99, 0x60, 0x37, 0x30, 0x52, 0x22, 0xf2, 0x66, 0xde, 0x22, 0xc8, 0x47, 0x4a,
	0xc4, 0x73, 0x08, 0x7b, 0x6a, 0xad, 0x1a, 0xc7, 0x9e, 0xc2, 0x58, 0x09, 0xae, 0x55, 0xe3, 0x22,
	0x6f, 0x76, 0xb1, 0x08, 0x72, 0x5f, 0x11, 0x11, 0xff, 0xf6, 0x7a, 0xe1, 0x7b, 0x34, 0xa5, 0xaa,
	0xce, 0x7d, 0x18, 0x83, 0xcb, 0x03, 0x36, 0x2e, 0x1a, 0x11, 0x42, 0xff, 0x6c, 0x0a, 0x61, 0xeb,
	0x24, 0x0d, 0xb7, 0x58, 0xbb, 0xe8, 0x62, 0xe6, 0x2d, 0xae, 0x72, 0xe8, 0xa0, 0x0d, 0xd6, 0x8e,
	0xbd, 0x82, 0xc7, 0xc5, 0x51, 0x70, 0x65, 0x4a, 0xe4, 0x1a, 0x2b, 0x6e, 0x77, 0xee, 0x10, 0x5d,
	0x92, 0xc3, 0x4d, 0x71, 0x14, 0x99, 0x29, 0x71, 0x8d, 0xd5, 0x66, 0xe7, 0x0e, 0x6c, 0x0e, 0x0f,
	0x29, 0x11, 0x6f, 0xa4, 0x2e, 0x39, 0xb5, 0xba, 0x22, 0xe1, 0x35, 0xc1, 0x5b, 0xa9, 0xcb, 0x4f,
	0x6d, 0xcf, 0xe7, 0x10, 0xb4, 0x4e, 0x5a, 0x7e, 0x97, 0x3a, 0xf2, 0x49, 0x31, 0xd1, 0x58, 0xad,
	0xdb, 0x73, 0xac, 0xe1, 0x9a, 0x32, 0x7c, 0x34, 0xc2, 0xa2, 0x32, 0x8e, 0xbd, 0x84, 0x07, 0x65,
	0x8d, 0xc6, 0x55, 0x3b, 0x27, 0xf9, 0x29, 0x4f, 0x78, 0xc2, 0x32, 0xc1, 0x6e, 0x61, 0x42, 0x57,
	0xcf, 0x95, 0xed, 0xc3, 0x8d, 0xe9, 0x9c, 0x59, 0x76, 0x07, 0xd0, 0x51, 0x83, 0x78, 0x01, 0x21,
	0x6d, 0xba, 0xf8, 0x97, 0x07, 0x4f, 0xb6, 0xd2, 0x0d, 0x6e, 0x2d, 0x97, 0x5f, 0xbf, 0xc9, 0xc6,
	0xb1, 0x15, 0x4c, 0x64, 0x3f, 0x02, 0xb5, 0x0c, 0x97, 0xd3, 0xe4, 0xfc, 0x49, 0x93, 0x7f, 0x26,
	0xcd, 0x4f, 0x05, 0xec, 0x2d, 0xf8, 0x05, 0xb9, 0xd1, 0x38, 0xe1, 0xf2, 0xee, 0x9e, 0xd2, 0xbe,
	0x65, 0x2f, 0x7e, 0x97, 0x7e, 0x7e, 0x3d, 0xd0, 0x29, 0x4c, 0x07, 0x1b, 0x63, 0xbf, 0x54, 0xa9,
	0xdd, 0x77, 0x8b, 0xb3, 0xb2, 0x7b, 0xfa, 0xee, 0x7d, 0xda, 0x9d, 0x37, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa5, 0xef, 0x2c, 0x61, 0x8b, 0x02, 0x00, 0x00,
}
