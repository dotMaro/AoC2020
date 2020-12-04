package main

import (
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day4/input.txt")
	present, valid := validPassports(input)
	utils.Print("Task 1. There are %d present passports", present)
	utils.Print("Task 2. There are %d valid passports", valid)
}

// validPassports returns how many present and valid passports there are.
func validPassports(input string) (int, int) {
	presentPassports := 0
	validPassports := 0
	var pass passport
	for _, line := range utils.SplitLine(input) {
		if len(line) == 0 {
			// Moving to new passport
			// Validate old pass
			if pass.present() {
				presentPassports++
			}
			if pass.valid() {
				validPassports++
			}
			// Reset it and continue
			pass = passport{}
			continue
		}

		for _, s := range strings.Split(line, " ") {
			pass.parseField(s)
		}
	}
	if pass.present() {
		presentPassports++
	}
	if pass.valid() {
		validPassports++
	}
	return presentPassports, validPassports
}

type passport struct {
	byr int    // Birth Year
	iyr int    // Issue Year
	eyr int    // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid string // Passport ID
	cid string // Country ID
}

// parseField and set it.
func (p *passport) parseField(s string) {
	switch s[:3] {
	case "byr":
		i, err := strconv.Atoi(s[4:])
		if err != nil {
			return
		}
		p.byr = i
	case "iyr":
		i, err := strconv.Atoi(s[4:])
		if err != nil {
			return
		}
		p.iyr = i
	case "eyr":
		i, err := strconv.Atoi(s[4:])
		if err != nil {
			return
		}
		p.eyr = i
	case "hgt":
		p.hgt = s[4:]
	case "hcl":
		p.hcl = s[4:]
	case "ecl":
		p.ecl = s[4:]
	case "pid":
		p.pid = s[4:]
	case "cid":
		p.cid = s[4:]
	default:
		panic(s)
	}
}

// presents returns true if all fields except cid (country ID) are present.
func (p *passport) present() bool {
	return p.byr != 0 &&
		p.iyr != 0 &&
		p.eyr != 0 &&
		len(p.hgt) != 0 &&
		len(p.hcl) != 0 &&
		len(p.ecl) != 0 &&
		len(p.pid) != 0
}

// valid returns true if all fields are deemed valid according to their
// respective policy.
func (p *passport) valid() bool {
	if !p.present() {
		return false
	}

	// byr, iyr, eyr
	if !(p.byr >= 1920 && p.byr <= 2002 &&
		p.iyr >= 2010 && p.iyr <= 2020 &&
		p.eyr >= 2020 && p.eyr <= 2030) {
		return false
	}

	// hgt
	if strings.HasSuffix(p.hgt, "cm") {
		height, err := strconv.Atoi(p.hgt[:len(p.hgt)-2])
		if err != nil {
			return false
		}
		if height < 150 || height > 193 {
			return false
		}
	} else if strings.HasSuffix(p.hgt, "in") {
		height, err := strconv.Atoi(p.hgt[:len(p.hgt)-2])
		if err != nil {
			return false
		}
		if height < 59 || height > 76 {
			return false
		}
	} else {
		// Invalid suffix
		return false
	}

	// hcl
	if len(p.hcl) != 7 || p.hcl[0] != '#' {
		return false
	}
	for _, c := range p.hcl[1:] {
		if !(c >= 'a' && c <= 'f') && !(c >= '0' && c <= '9') {
			return false
		}
	}

	// ecl
	if p.ecl != "amb" && p.ecl != "blu" && p.ecl != "brn" &&
		p.ecl != "gry" && p.ecl != "grn" && p.ecl != "hzl" && p.ecl != "oth" {
		return false
	}

	// pid
	if len(p.pid) != 9 {
		return false
	}

	return true
}
