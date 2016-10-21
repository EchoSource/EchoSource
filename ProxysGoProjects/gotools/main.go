package gotools

import (
	"io"
	"os"
	"net/http"
	"encoding/json"
	"ioutil"
	"time"
	"fmt"
	"runtime"
	"bufio"
	"math/rand"

)




func OpenURL(url string) {

switch runtime.GOOS {
case "linux":
    err = exec.Command("xdg-open", url).Start()
case "windows", "darwin":
    err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
default:
    err = fmt.Errorf("can't open url. unsupported platform.")
}

}




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



func ReadJSON(path string, key string) (string, error) {
	newval := ""
	var rjson map[string]interface{}

	// ################## COMMANDS.JSON UPDATE ########################
	file, err := ioutil.ReadFile(path)
	if err != nil {
	//	fmt.Println("gotools error =>")
	//	fmt.Println(err)
		return err
	} else {
		json.Unmarshal(file, &rjson)
		for k1, v1 := range rjson {
			if k1 == key {
			 	newval = v1
			}
		} // end of oldcommand for loop
	} // check if error is nil
	return newval, err
}



func WriteJSON(path string, key string, val string) error {
	newval := ""
	var rjson map[string]interface{}

	// ################## COMMANDS.JSON UPDATE ########################
	file, err := ioutil.ReadFile(path)
	if err != nil {
	//	fmt.Println("gotools error =>")
	//	fmt.Println(err)
		return err
	} else {
		json.Unmarshal(file, &rjson)
		rjson[kew] = val
	b, err := json.MarshalIndent(rjson, "", "   ")
	if err == nil {
		ioutil.WriteFile(path, b, 0777)
		// it works
	}	else {
		return err
	}

	} // check if error is nil
return err
}





// displays random integer
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
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







// readLines reads a whole file into memory
// and returns a slice of its lines.
func countLines(path string) int {
  counter := 0

  file, err := os.Open(path)
  if err != nil {

  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    counter++
  }
  return counter
}


