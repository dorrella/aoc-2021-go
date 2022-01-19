package main

import (
	"fmt"
	"strconv"
	"strings"
)

var DECODER = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

type Parser struct {
	c    chan rune
	done bool
}

func NewParser() *Parser {
	return &Parser{
		c:    nil,
		done: false,
	}
}

func (p *Parser) ParseString(s string) {
	c := make(chan rune)
	go func() {
		for _, char := range s {
			for _, bit := range DECODER[char] {
				c <- bit
			}
		}
		p.done = true
		close(c)
	}()

	p.c = c
}

func (p *Parser) ParseBits(r []rune) {
	c := make(chan rune)
	go func() {
		for _, bit := range r {
			c <- bit
		}
		p.done = true
		close(c)
	}()

	p.c = c
}

func (p *Parser) ReadBits(num uint) []rune {
	ret := make([]rune, 0, num)
	for i := uint(0); i < num; i++ {
		bit, ok := <-p.c
		if !ok {
			panic("unexpected end")
		}
		ret = append(ret, bit)
	}

	return ret
}

func (p *Parser) ReadNumber(bits uint) uint {
	num := string(p.ReadBits(bits))
	i, err := strconv.ParseUint(num, 2, 64)
	if err != nil {
		panic(err)
	}

	return uint(i)
}

func (p *Parser) ReadHeader() Header {
	v := p.ReadNumber(3)
	i := p.ReadNumber(3)

	return Header{
		Version: v,
		Id:      i,
	}
}

func (p *Parser) ReadLiteral(h Header) (Packet, uint) {
	b := strings.Builder{}
	var num_bits uint = 0

	for true {
		chunk := p.ReadBits(5)
		num_bits += 5
		b.WriteString(string(chunk[1:5]))

		if chunk[0] != '1' {
			break
		}

	}

	num := b.String()
	i, err := strconv.ParseUint(num, 2, 64)
	if err != nil {
		panic(err)
	}

	packet := NewPacket(h, PKT_LIT)
	packet.Literal = uint(i)
	return packet, num_bits
}

func (p *Parser) ReadOp(h Header) (Packet, uint) {
	var num_bits uint
	packet := NewPacket(h, PKT_OP)

	bit_label := p.ReadBits(1)
	num_bits = 1

	if bit_label[0] == '0' {
		length := p.ReadNumber(15)
		num_bits += 15
		sub_str := p.ReadBits(length)
		num_bits += length
		new_parser := NewParser()
		new_parser.ParseBits(sub_str)

		for tmp_bits := uint(0); tmp_bits < length; {
			tmp_packet, new_bits := new_parser.ReadPacket()
			packet.Subs = append(packet.Subs, &tmp_packet)
			tmp_bits += new_bits
		}
	} else if bit_label[0] == '1' {
		packets := p.ReadNumber(11)
		num_bits += 11
		for i := 0; i < int(packets); i++ {
			tmp_packet, new_bits := p.ReadPacket()
			num_bits += new_bits
			packet.Subs = append(packet.Subs, &tmp_packet)
		}
	} else {
		panic("bad bit")
	}

	return packet, num_bits
}

func (p *Parser) ReadPacket() (Packet, uint) {
	head := p.ReadHeader()
	num_bits := uint(6)

	var packet Packet
	var new_bits uint = 0

	switch head.Id {
	case 4:
		packet, new_bits = p.ReadLiteral(head)
		num_bits += new_bits
	default:
		packet, new_bits = p.ReadOp(head)
		num_bits += new_bits
	}

	return packet, num_bits
}

func walk(p *Packet) uint {
	if p.Type == PKT_LIT {
		return p.Header.Version
	}

	sum := uint(0)
	for _, s := range p.Subs {
		sum += walk(s)
	}
	return sum + p.Header.Version

}

//i guess assume only 1 packet since its fuzzy
func (p *Parser) Part1() {
	packet, _ := p.ReadPacket()
	fmt.Println(walk(&packet))
}

func eval(p *Packet) uint {
	if p.Type == PKT_LIT {
		return p.Literal
	}

	switch p.Header.Id {
	case 0:
		//sum
		sum := uint(0)
		for _, sub := range p.Subs {
			sum += eval(sub)
		}
		return sum
	case 1:
		prod := uint(1)
		for _, sub := range p.Subs {
			prod *= eval(sub)
		}
		return prod
	case 2:
		min := uint(0)
		for i, sub := range p.Subs {
			if i == 0 {
				min = eval(sub)
				continue
			}

			if min > eval(sub) {
				min = eval(sub)
			}
		}
		return min
	case 3:
		max := uint(0)
		for i, sub := range p.Subs {
			if i == 0 {
				max = eval(sub)
				continue
			}

			if max < eval(sub) {
				max = eval(sub)
			}
		}
		return max
	case 5:
		if eval(p.Subs[0]) > eval(p.Subs[1]) {
			return 1
		}
		return 0
	case 6:
		if eval(p.Subs[0]) < eval(p.Subs[1]) {
			return 1
		}
		return 0
	case 7:
		if eval(p.Subs[0]) == eval(p.Subs[1]) {
			return 1
		}
		return 0
	default:
		panic("something")

	}
}

func (p *Parser) Part2() {
	packet, _ := p.ReadPacket()
	fmt.Println(eval(&packet))
}
