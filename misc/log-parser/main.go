/* Code to read a file, transform and save in another file.
https://www.loggly.com/ultimate-guide/parsing-apache-logs/

https://github.com/kiritbasu/Fake-Apache-Log-Generator

TODO

- [ ] Read file log
- [ ] Parse logs
- [ ] Print each part

*/

package main

import (
	"log"
	"fmt"
	"strings"
)

func main() {
	log.Println("Starting app")

	log_string := `153.48.39.181 - rempel2286 [23/Apr/2021:17:01:12 -0300] "GET /seize/e-markets/e-tailers/dynamic HTTP/1.1" 301 16465`

	fmt.Println(log_string)
	splited := strings.Split(log_string, " ")

	fmt.Println("IP", splited[0])
	fmt.Println("User", splited[2])
}