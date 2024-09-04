package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	PATH  = "./root/devops/"
	REGEX = `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`
	hq    = &IpAddrHeap{}
)

func readIpAddressFromFolder(path string) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, dirEntry := range dirEntries {
		path := filepath.Join(path, dirEntry.Name())
		if dirEntry.IsDir() {
			readIpAddressFromFolder(path)
		} else {
			readFile(path, hq)
		}
	}
}

// by default golang is pass by value, we need to pass by reference ensure
// the modification can persists
func readFile(path string, ipaddrs *IpAddrHeap) {
	// fmt.Printf("%d size before parse %v\n", ipaddrs.Len(), path)
	file, err := os.Open(path)
	// IO handler must be close when functions completes defer allow this done nicely

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	regex := regexp.MustCompile(REGEX)

	for scanner.Scan() {
		for _, ipaddr := range parseIpFromFile(scanner.Text(), regex) {
			heap.Push(ipaddrs, ipaddr)
		}
	}
	// debug string found out the heap was re init after the readingDir
	// fmt.Printf("%d size after parse %v\n", ipaddrs.Len(), path)
	// fmt.Println("***************************")
	defer file.Close()
}

func parseIpFromFile(text string, regex *regexp.Regexp) []string {
	// since regex is defined as constant we can ignore the error
	return regex.FindAllString(text, -1)
}

func main() {
	heap.Init(hq)
	readIpAddressFromFolder(PATH)
	for hq.Len() != 0 {
		fmt.Printf("%v\n", heap.Pop(hq))
	}

}
