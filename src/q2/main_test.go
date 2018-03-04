package main

import (
	"reflect"
	"testing"
)

func Test_IPRange2Cidr(t *testing.T) {
	if !reflect.DeepEqual(IPRange2Cidr("1.0.101.0", "1.0.150.155"),
		[]string{"1.0.101.0/24", "1.0.102.0/23", "1.0.104.0/21", "1.0.112.0/20",
			"1.0.128.0/20", "1.0.144.0/22", "1.0.148.0/23",
			"1.0.150.0/25", "1.0.150.128/28", "1.0.150.144/29",
			"1.0.150.152/30"}) {
		t.Error("Assert failed")
	}

	if !reflect.DeepEqual(IPRange2Cidr("1.0.8.0", "1.0.15.255"),
		[]string{"1.0.8.0/21"}) {
		t.Error("Assert failed")
	}

	if !reflect.DeepEqual(IPRange2Cidr("192.168.0.0", "192.168.1.0"),
		[]string{"192.168.0.0/24", "192.168.1.0/32"}) {
		t.Error("Assert failed")
	}
}
