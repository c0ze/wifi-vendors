package vendor

import (
	"testing"
)

const MacBookMac = `6c:40:08:b2:1b:a4`
const OnePlusOneMac = `C0EEFB3368FA`
const RandomMac1 = `C2EEFB3368FA`
const RandomMac2 = `6e:40:08:b2:1b:a4`

func TestLookup(t *testing.T) {
	Init()
	MacBookVendorString := Lookup(MacBookMac)
	if MacBookVendorString != "Apple, Inc." {
		t.Errorf("Error looking up MacBook Mac. %v", MacBookVendorString)
	}

	OnePlusVendorString := Lookup(OnePlusOneMac)
	if OnePlusVendorString != "OnePlus Tech (Shenzhen) Ltd" {
		t.Errorf("Error looking up OnePlus Mac. %v", OnePlusVendorString)
	}

	if Lookup(RandomMac1) != "Random" {
		t.Errorf("Error looking up Random Mac. 1")
	}

	if Lookup(RandomMac2) != "Random" {
		t.Errorf("Error looking up Random Mac. 2")
	}

	if Lookup("") != "Malformed" {
		t.Errorf("Error looking up empty string %v", Lookup(""))
	}

	if Lookup("dummy") != "Malformed" {
		t.Errorf("Error looking up dummy string")
	}
}

func TestLookupNoInit(t *testing.T) {
	VendorMap = make(map[string]string)
	MacBookVendorString := Lookup(MacBookMac)
	if MacBookVendorString != "Unknown" {
		t.Errorf("Error looking up MacBook Mac. %v", MacBookVendorString)
	}

	OnePlusVendorString := Lookup(OnePlusOneMac)
	if OnePlusVendorString != "Unknown" {
		t.Errorf("Error looking up OnePlus Mac. %v", OnePlusVendorString)
	}

	if Lookup(RandomMac1) != "Random" {
		t.Errorf("Error looking up Random Mac. 1")
	}

	if Lookup(RandomMac2) != "Random" {
		t.Errorf("Error looking up Random Mac. 2")
	}

	if Lookup("") != "Malformed" {
		t.Errorf("Error looking up empty string %v", Lookup(""))
	}

	if Lookup("dummy") != "Malformed" {
		t.Errorf("Error looking up dummy string")
	}
}
