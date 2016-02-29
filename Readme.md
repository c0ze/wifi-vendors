### Wifi mac vendors

A basic lib which parses a mac address and displays associated vendor.
You need to download [oui.txt](standards-oui.ieee.org/oui.txt) from IEEE. You can place it either in the current directory or under /etc

You will get either

* a simple string identifiying the vendor
* the string `Random` when the address is in Locally Administered Space
* the string `Unknown` if it's not either in the Oui or Locally Administered Space
* the string `Malformed` if it's not a good Mac address (I should have made a good regexp for this, in this case it only checks if it's not empty string or has less than 6 chars.)
