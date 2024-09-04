package main

import (
	"fmt"
	"strings"
)

type IpAddrHeap []string

func (heap IpAddrHeap) Len() int           { return len(heap) }
func (heap IpAddrHeap) Less(i, j int) bool { return compareIPAddresses(heap[i], heap[j]) < 0 }
func (heap IpAddrHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }

// Push adds an IP address string to the heap.
func (heap *IpAddrHeap) Push(x interface{}) {
	*heap = append(*heap, x.(string))
}

// Pop removes the minimum IP address string from the heap.
func (heap *IpAddrHeap) Pop() interface{} {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[0 : n-1]
	return x
}

// compareIPAddresses compares two IP addresses represented as strings.
func compareIPAddresses(ip1, ip2 string) int {
	// Split IP addresses into their octets
	octets1 := strings.Split(ip1, ".")
	octets2 := strings.Split(ip2, ".")

	for i := 0; i < 4; i++ {
		num1 := parseOctet(octets1[i])
		num2 := parseOctet(octets2[i])
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}
	return 0
}

// parseOctet parses an octet of the IP address and converts it to an integer.
func parseOctet(octet string) int {
	var num int
	fmt.Sscanf(octet, "%d", &num)
	return num
}
