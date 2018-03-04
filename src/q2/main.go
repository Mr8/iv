package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var MASK = []uint32{0x00000000, 0x80000000, 0xC0000000, 0xE0000000,
	0xF0000000, 0xF8000000, 0xFC000000, 0xFE000000,
	0xFF000000, 0xFF800000, 0xFFC00000, 0xFFE00000,
	0xFFF00000, 0xFFF80000, 0xFFFC0000, 0xFFFE0000,
	0xFFFF0000, 0xFFFF8000, 0xFFFFC000, 0xFFFFE000,
	0xFFFFF000, 0xFFFFF800, 0xFFFFFC00, 0xFFFFFE00,
	0xFFFFFF00, 0xFFFFFF80, 0xFFFFFFC0, 0xFFFFFFE0,
	0xFFFFFFF0, 0xFFFFFFF8, 0xFFFFFFFC, 0xFFFFFFFE,
	0xFFFFFFFF}

func IP2Int(sip string) uint32 {
	splitedIP := strings.Split(sip, ".")
	var ip [4]uint32
	for i := 0; i < 4; i++ {
		_ip, _ := strconv.ParseInt(splitedIP[i], 10, 32)
		ip[i] = uint32(_ip)
	}

	return (ip[0] << 24) + (ip[1] << 16) + (ip[2] << 8) + ip[3]
}

func IntToIP(iIP uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(iIP>>24), byte(iIP>>16), byte(iIP>>8), byte(iIP))
}

func IPMaxDiffBit(start uint32, end uint32) int {
	return 32 - int(math.Floor(math.Log(float64(end-start+1))/math.Log(2)))

}

func Max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func FindMask(ip uint32) int {
	maxMask := 32
	for i := maxMask; i > 0; i-- {
		if ip&MASK[maxMask-1] != ip {
			break
		}
		maxMask = i
	}
	return maxMask
}

func IPRange2Cidr(startIP string, endIP string) []string {
	istart := IP2Int(startIP)
	iend := IP2Int(endIP)

	var cidrs []string
	for iend >= istart {

		// find max mask match
		mask := FindMask(istart)

		// find max different between end IP
		maxDiff := IPMaxDiffBit(istart, iend)

		// set mask as max(mask, maxDiff)
		mask = Max(mask, maxDiff)

		cidrs = append(cidrs, fmt.Sprintf("%s/%d", IntToIP(istart), mask))

		// increase one Mask range IP
		istart += uint32(math.Pow(2, float64(32-mask)))
	}
	return cidrs
}

func main() {
	fmt.Println(IPRange2Cidr("1.0.101.0", "1.0.150.155"))
}
