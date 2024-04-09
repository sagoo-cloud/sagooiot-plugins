package base

import (
	"encoding/binary"
	"fmt"
)

const (
	//头的字节长度
	HeaderLen = 10
	//魔数
	magicNum = "CYRC"
	//协议版本号
	protocolVer = 0x10
	//保留字段
	reserve = 0
)

type Packet struct {
	Len      uint16 // 负载长度
	Protocol byte   // 协议版本号
	Op       byte   // 包类型
	Checksum byte   // 校验位
	Reserve  byte   // Reserve
	Payload  []byte // 负载内容
}

func (p *Packet) Encode() ([]byte, error) {
	buf := make([]byte, HeaderLen+len(p.Payload))
	copy(buf, magicNum)
	p.Len = uint16(len(p.Payload))
	binary.BigEndian.PutUint16(buf[4:6], p.Len)
	buf[6] = protocolVer
	buf[7] = p.Op
	// 设置checksum
	for _, v := range p.Payload {
		p.Checksum += v
	}
	buf[8] = p.Checksum
	// 固定保留
	buf[9] = reserve
	copy(buf[HeaderLen:], p.Payload)
	return buf, nil
}

func Decode(buf []byte) (Packet, error) {
	var p Packet
	p.Len = binary.BigEndian.Uint16(buf[4:6])
	p.Protocol = buf[6]
	p.Op = buf[7]
	p.Checksum = buf[8]
	p.Reserve = buf[9]
	if len(buf) < int(10+p.Len) {
		return p, fmt.Errorf("invalid data length,p:%v buf:%x", p, buf)
	}
	if p.Len > 0 {
		p.Payload = buf[10 : 10+p.Len]
	}
	return p, nil
}
