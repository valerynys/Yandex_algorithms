package main

import (
	"fmt"
	"strings"
)

type Timestamp struct {
	h int
	m int
	s int
}

func (t Timestamp) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.h, t.m, t.s)
}

func getTimestamp(s string) Timestamp {
	parts := strings.Split(s, ":")
	h, m, s := 0, 0, 0
	fmt.Sscanf(parts[0], "%d", &h)
	fmt.Sscanf(parts[1], "%d", &m)
	fmt.Sscanf(parts[2], "%d", &s)
	return Timestamp{h, m, s}
}

const (
	SECS_IN_MINUTE = 60
	SECS_IN_HOUR   = SECS_IN_MINUTE * 60
	SECS_IN_DAY    = SECS_IN_HOUR * 24
)

func getSeconds(t Timestamp) int {
	ret := t.h * SECS_IN_HOUR
	ret += t.m * SECS_IN_MINUTE
	ret += t.s
	return ret
}

func getTimestampFromSeconds(secs int) Timestamp {
	h := secs / SECS_IN_HOUR
	h %= 24
	secs %= SECS_IN_HOUR
	m := secs / SECS_IN_MINUTE
	secs %= SECS_IN_MINUTE
	s := secs
	return Timestamp{h, m, s}
}

func main() {
	var A, B, C Timestamp
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)
	tm2 := 2*B + (getSeconds(C)-getSeconds(A))%SECS_IN_DAY
	tm := tm2/2 + tm2%2
	fmt.Println(getTimestampFromSeconds(tm))
}
