package vendor

import (
	"bufio"
	//	"fmt"
	"github.com/c0ze/golang-utils"
	"log"
	"os"
	"regexp"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

var randomMac = regexp.MustCompile(`^([A-F0-9][2,6,A,E][A-F0-9]*)$`)
var VendorMap map[string]string

func Init() {
	VendorMap = make(map[string]string)
	var lines []string
	var err error
	lines, err = readLines("/etc/fup/oui.txt")
	if err != nil {
		lines, err = readLines("oui.txt")
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
	}

	for _, line := range lines {
		if len(line) > 1 {
			if string(line[0]) != " " && string(line[0]) != "" && string(line[0]) != "\t" && string(line[2]) != "-" {
				words := strings.Fields(line)
				if len(words) > 3 {
					//					VendorMap[words[0]] = fmt.Sprintf("%v\n", strings.Join(words[3:], " "))
					VendorMap[words[0]] = strings.Join(words[3:], " ")
				}
			}
		}
	}
}

func Lookup(mac string) string {
	vendor := ""

	sanitizedMac := utils.StripColon(mac)

	if len(sanitizedMac) > 5 {
		vendor = VendorMap[sanitizedMac[0:6]]
		if vendor == "" {
			if randomMac.MatchString(utils.StripColon(mac)) {
				vendor = "Random"
			} else {
				vendor = "Unknown"
			}
		}
	} else {
		vendor = "Malformed"
	}

	return vendor
}
