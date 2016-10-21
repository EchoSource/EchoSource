package main

import (
	"fmt"
	"time"
	"os"
	"net/http"
	"io"
	"strings"
	"bufio"
)




func DownloadFile(output string, url string) bool {
chk := true
os.Remove(output)
out, err := os.Create(output)
if err != nil {
	chk = false
// fmt.Println(err)
}
defer out.Close()
resp, err := http.Get(url)
if err != nil {
	chk = false
// fmt.Println(err)
}
defer resp.Body.Close()
_, err = io.Copy(out, resp.Body)
if err != nil {
	chk = false
// fmt.Println(err)
}
 return chk
}


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








func main() {

	if _, err := os.Stat("../../updates.auf"); err == nil {
		os.Remove("../../updates.auf")
		if err != nil {
			fmt.Println("Couldn't delete updates.auf file. ")
		}
	}
	a := DownloadFile("../../updates.auf", "https://raw.githubusercontent.com/proxikal/AutoGo/master/updates.auf")
	if a == false {
		fmt.Println("Error => Couldn't download update.auf from github.")
	} else {
		fmt.Println("Downloaded updates.auf file.")
	}				


	update, err := readLines("../../updates.auf")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range update {
			pass := false
			line := strings.Split(v, "=")
			do := line[0]
			get := strings.Replace(v, line[0]+"=", "", -1)

			if do == "PACKAGE" {
				fmt.Println("Update Package Size: "+get)
			}

			if do == "MAKE" {
				err = os.Mkdir("../../"+get, 0777)
				if err != nil {
					fmt.Println(err)
				}
			} // end of make folder command.


			if do == "INSTALL" {
				thefile := strings.Replace(get, "https://raw.githubusercontent.com/proxikal/AutoGo/master/", "", -1)
				if _, err := os.Stat("../../"+thefile); err == nil {
					// path/to/whatever exists
					if thefile != "System/custom/commands.json" && thefile != "System/custom/responses.json" && thefile != "config.json" {
						err = os.Remove("../../"+thefile)
						if err != nil {
							fmt.Println(err)
						}
					}
					time.Sleep(3000 * time.Millisecond)
				}

				// check if it's commands, responses or config.ini.
				// then put them in the /updates/ folder.
				if thefile == "System/custom/commands.json" {
					pass = true
					a := DownloadFile("../updates/commands.json", get)
					if a == false {
						fmt.Println(thefile+" Error => An error has occured. Check your internet. make sure Your Autogo process is closed completely.\nAnd don't move helper from his original path /System/helper")
					} else {
						fmt.Println(thefile+" was updated.")
					}				
				}

				if thefile == "System/custom/responses.json" {
					pass = true
					a := DownloadFile("../updates/responses.json", get)
					if a == false {
						fmt.Println(thefile+" Error => An error has occured. Check your internet. make sure Your Autogo process is closed completely.\nAnd don't move helper from his original path /System/helper")
					} else {
						fmt.Println(thefile+" was updated.")
					}				
				}


				if thefile == "config.json" {
					pass = true
					a := DownloadFile("../updates/config.json", get)
					if a == false {
						fmt.Println(thefile+" Error => An error has occured. Check your internet. make sure Your Autogo process is closed completely.\nAnd don't move helper from his original path /System/helper")
					} else {
						fmt.Println(thefile+" was updated.")
					}				
				}


				if pass == false {
					// download the file.
					a := DownloadFile("../../"+thefile, get)
					if a == false {
						fmt.Println(thefile+" Error => An error has occured. Check your internet. make sure Your Autogo process is closed completely.\nAnd don't move helper from his original path /System/helper")
					} else {
						fmt.Println(thefile+" was updated.")
					}
				}
			} // end of INSTALL.
		} // loop through lines and grab updates!
		err = os.Remove("../../updates.auf")
		if err != nil {
			fmt.Println("Couldn't delete updates.auf file. You should delete this ASAP!")
		}
	} // make sure there isn't an error.
} // end of func main