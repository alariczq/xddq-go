package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

var Endianess = binary.BigEndian

const (
	MagicNumber  = 29099
	HeaderLength = 18
)

type Header struct {
	Header   int16
	Size     int32
	MsgId    int32
	PlayerId int64
}

type Packet struct {
	Header
	Body []byte
}

func NewPacket(msgId int32, data []byte) *Packet {
	return &Packet{
		Header: Header{
			Header:   MagicNumber,
			PlayerId: 0,
			MsgId:    msgId,
			Size:     HeaderLength + int32(len(data)),
		},
		Body: data,
	}
}

func (p *Packet) String() string {
	return fmt.Sprintf("msgId: %v, size: %v, playerId: %v, bodySize: %v", p.Header.MsgId, p.Header.Size, p.Header.PlayerId, len(p.Body))
}

func (p *Packet) WithPlayerId(playerId int64) *Packet {
	p.Header.PlayerId = playerId
	return p
}

func (p *Packet) Encode() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, HeaderLength+int(p.Header.Size)))
	binary.Write(buf, Endianess, p.Header.Header)
	binary.Write(buf, Endianess, p.Header.Size)
	binary.Write(buf, Endianess, p.Header.MsgId)
	binary.Write(buf, Endianess, p.Header.PlayerId)
	buf.Write(p.Body)
	return buf.Bytes()
}

func Decode(reader io.Reader) (*Packet, error) {
	p := &Packet{}
	if err := p.Decode(reader); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Packet) Decode(reader io.Reader) error {
	binary.Read(reader, Endianess, &p.Header.Header)
	binary.Read(reader, Endianess, &p.Header.Size)
	binary.Read(reader, Endianess, &p.Header.MsgId)
	binary.Read(reader, Endianess, &p.Header.PlayerId)

	if p.Header.Size > HeaderLength {
		p.Body = make([]byte, p.Header.Size-HeaderLength)

		_, err := io.ReadFull(reader, p.Body)
		if err != nil {
			return err
		}
	}

	return nil
}
