package vendor

import (
	"testing"
)

const MacBookMac = `6c:40:08:b2:1b:a4`
const OnePlusOneMac = `C0EEFB3368FA`

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

}
