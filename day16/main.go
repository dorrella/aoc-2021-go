package main

import (
	"bufio"
	"os"
)

func ReadInput(file string) string {

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		return text
	}
	return ""
}
func main() {
	p := NewParser()
	p.ParseString("38006F45291200")
	p.Part1()
	p.ParseString("8A004A801A8002F478")
	p.Part1()
	p.ParseString("620080001611562C8802118E34")
	p.Part1()
	p.ParseString("C0015000016115A2E0802F182340")
	p.Part1()
	p.ParseString("A0016C880162017C3686B18A3D4780")
	p.Part1()

	p.ParseString("C200B40A82")
	p.Part2()

	p.ParseString("04005AC33890")
	p.Part2()

	p.ParseString("880086C3E88112")
	p.Part2()

	p.ParseString("CE00C43D881120")
	p.Part2()

	p.ParseString("D8005AC2A8F0")
	p.Part2()

	p.ParseString("F600BC2D8F")
	p.Part2()

	p.ParseString("9C005AC2F8F0")
	p.Part2()

	p.ParseString("9C0141080250320F1802104A08")
	p.Part2()

	i := ReadInput("input.txt")
	p.ParseString(i)
	p.Part1()

	p.ParseString(i)
	p.Part2()
	//p.ReadPacket()
	/*
		p = NewParser()
		p.ParseString("38006F45291200")
		p.ReadPacket()

		p = NewParser()
		p.ParseString("EE00D40C823060")
		p.ReadPacket()
	*/
}
