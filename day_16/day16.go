package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var filename = "input.txt"

type packet struct {
	version, typeId, value int64
	subPackets             []packet
}

func readFile(filename string) string {
	path := filepath.Join(".", "day_16", filename)

	decod := map[rune]string{
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
		'F': "1111"}

	packet := ""

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	for _, c := range text {
		packet += decod[c]
	}

	return packet
}

func binToDec(number string) int64 {
	var result int64 = 0
	var bit uint = 0
	var n = int64(len(number) - 1)
	for (n >= 0) {
		if number[n] == '1' {
			result += (1 << (bit))
		}
		n = n - 1
		bit++
	}
	return result
}

func (p packet) calculateVersions() int64 {
	var total int64
	for _, subPacket := range p.subPackets {
		total += subPacket.calculateVersions()
	}
	return p.version + total
}

func (p packet) calculateValues() (res int64) {
	switch p.typeId {
	case 0:
		if len(p.subPackets) == 1 {
			res = p.subPackets[0].calculateValues()
		} else {
			for _, subPacket := range p.subPackets {
				res += subPacket.calculateValues()
			}
		}
	case 1:
		if len(p.subPackets) == 1 {
			res = p.subPackets[0].calculateValues()
		} else {
			res ++
			for _, subPacket := range p.subPackets {
				res *= subPacket.calculateValues()
			}
		}
	case 2:
		for i, subPacket := range p.subPackets {
			value := subPacket.calculateValues()
			if i == 0 {
				res = value
			} else {
				if res > value {
					res = value
				}
			}
		}
	case 3:
		for i, subPacket := range p.subPackets {
			value := subPacket.calculateValues()
			if i == 0 {
				res = value
			} else {
				if res < value {
					res = value
				}
			}
		}
	case 4:
		res = p.value
	case 5:
		if p.subPackets[0].calculateValues() > p.subPackets[1].calculateValues() {
			res = 1
		}
	case 6:
		if p.subPackets[0].calculateValues() < p.subPackets[1].calculateValues() {
			res = 1
		}
	case 7:
		if p.subPackets[0].calculateValues() == p.subPackets[1].calculateValues() {
			res = 1
		}
	}
	return
}

func decodeLiteral(input string) (int64, string) {
	offset := 0
	blocks := ""
	for {
		block := input[offset : offset+5]
		blocks += block[1:]
		offset += 5
		if block[0] == '0' {
			break
		}
	}
	res := binToDec(blocks)
	return res, input[offset:]
}

func decodeOperator0(input string) ([]packet, string) {
	subPacketLength := binToDec(input[:15])
	subPacket := input[15 : 15+subPacketLength]
	packets := make([]packet, 0)
	for {
		p, remainder := decodePacket(subPacket)
		packets = append(packets, p)
		subPacket = remainder
		if remainder == "" {
			break
		}
	}

	return packets, input[15+subPacketLength:]
}

func decodeOperator1(input string) ([]packet, string) {
	subPacketCount := binToDec(input[:11])
	remainder := input[11:]
	packets := make([]packet, subPacketCount)
	for i := 0; i < int(subPacketCount); i++ {
		p, rest := decodePacket(remainder)
		remainder = rest
		packets[i] = p
	}

	return packets, remainder
}

func decodePacket(input string) (packet, string) {
	p := packet{version: binToDec(input[:3]), typeId: binToDec(input[3:6])}
	if p.typeId == 4 {
		value, remainder := decodeLiteral(input[6:])
		p.value = value
		return p, remainder
	} else {
		if input[6] == '0' {
			subPackets, remainder := decodeOperator0(input[7:])
			p.subPackets = subPackets
			return p, remainder
		} else {
			subPackets, remainder := decodeOperator1(input[7:])
			p.subPackets = subPackets
			return p, remainder
		}
	}
}

func main() {
	input := readFile(filename)
	p, _ := decodePacket(input)

	fmt.Printf("First answer: %v \n", p.calculateVersions())
	fmt.Printf("Second answer: %v \n", p.calculateValues())
}