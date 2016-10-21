package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os/exec"
	"io/ioutil"
	"encoding/json"
	"runtime"
	"time"
	"strconv"
	"os"
	"net/http"
	"io"
	"strings"
	"bufio"
)

	var err error
	var startTime time.Time
	var js obj
	var cmd commands
	var resp responses

type addnew struct {

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








var (
	Invite     	string
	MakeApp   	bool
	MakeRoles	bool
	Initiate	bool
	UpdateFiles	bool
	UpdateFiles2 bool
	UpdateFiles3 bool
	AlertChannel	string
	AlertHour		string
	UpdateChk		bool
	Role		string
	Color 		string
	Silent		bool
	RoleRename	string
	RoletoName	string
	Status 		string
	NewUpdateChk	bool
)

func init() {

	flag.StringVar(&Invite, "i", "", "Requires the client id of your bot application. helper.exe -i CLIENT ID")
	flag.BoolVar(&MakeApp, "m", false, "Opens your browser to the Discord Developer site.")
	flag.BoolVar(&MakeRoles, "b", false, "Creates the Bot Commander & muted role.")
	flag.BoolVar(&Initiate, "o", false, "Listens for the PM if a Guild owner can't be found.")
	flag.BoolVar(&UpdateFiles, "upd1", false, "Updates commands.json")
	flag.BoolVar(&UpdateFiles2, "upd2", false, "Updates response.json")
	flag.BoolVar(&UpdateFiles3, "upd3", false, "Updates config.ini")
	flag.StringVar(&AlertChannel, "a", "", "Alerts a specified channel of whatever text you have in notice.txt")
	flag.StringVar(&AlertHour, "h", "0", "Chooses how long to wait before Autogo sends notice text to desired channel.")
	flag.BoolVar(&UpdateChk, "chkupd", false, "Looks for update.")
	flag.BoolVar(&Silent, "s", false, "Helper won't alert you when he completes a task.")
	flag.BoolVar(&NewUpdateChk, "up", false, "Auto updates required files from update.auf file.")

	// rolecolor
	flag.StringVar(&Role, "role", "", "helper.exe -role \"Role Name\" -color #00FFFF")
	flag.StringVar(&Color, "color", "", "helper.exe -role \"Role Name\" -color #00FFFF")

	//role rename
	flag.StringVar(&RoleRename, "r1", "", "Role to rename: helper.exe -r1 \"Original Name\" -r2 \"New Name\"")
	flag.StringVar(&RoletoName, "r2", "", "the new name")

	// Name change
	flag.StringVar(&Status, "status", "", "Changes AutoGo's Status.")

	flag.Parse()
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

	file, err := ioutil.ReadFile("../../config.json")
	json.Unmarshal(file, &js)

    // Login to discord. You can use a token or email, password arguments.
	dg, err := discordgo.New(js.Bot)
	if err != nil {
		fmt.Println(err)
		return
	}


	dg.Open()

	// Simple way to keep program running until any key press.









if Status != "" {
	fmt.Println("Triggered Status change =>")
//	dg.UserUpdate(dg.State.User.Email, js.Bot, Name, dg.State.User.Avatar, "")
//	dg.UserUpdate(dg.State.User.Email, js.Bot, Name, dg.State.User.Avatar, "")
	dg.UpdateStatus(0, Status)
}













if Role != "" {
	jolt := 0
	fmt.Println("Triggered Helper: Role Color")
	id := ""
        guilds, err := dg.UserGuilds()
        if err != nil {
                fmt.Println(err)
                return
        }
        for _, v := range guilds {
        	id = v.ID
          //  fmt.Printf("%d : %#v\n", k, v)
        }
var roleID string
var hoist bool
var perms int

role := Role
color := strings.Replace(Color, "#", "", -1)


// newcolor := strconv.FormatInt(h, 16)
fmt.Println(role)
  roles, err := dg.GuildRoles(id)
  if err == nil {
    for _, v := range roles {
      if v.Name == role {
       roleID = v.ID
       hoist = v.Hoist
       perms = v.Permissions
      }
    }
  }

var ij int
newcode, _ := strconv.ParseInt(color, 16, 0)
d := fmt.Sprintf("%d", newcode)
fmt.Println(d)
ij, err = strconv.Atoi(d)
if err != nil {
	fmt.Println(err)
}
_, err = dg.GuildRoleEdit(id, roleID, role, ij, hoist, perms)
if err != nil {
if Silent == false {
//  	fmt.Println("s.GuildRoles is the error")
// 	dg.ChannelMessageSend(id, "**Helper:** I've changed the role color to `#"+color+"` on the role `"+role+"`")
k, err := dg.UserChannelCreate(js.Admin)
if err == nil {
	jolt++
	dg.ChannelTyping(k.ID)
	time.Sleep(1000 * time.Millisecond)
	dg.ChannelMessageSend(k.ID, "**Helper:** There was an error!\n ```fix\nMake sure I have permissions to edit roles\nCheck the role name"+role+"\nmake sure you formatted the color properly. example: #000FFF```")
}

} // check if silent == false
}




if Silent == false && jolt == 0 {
// dg.ChannelMessageSend(id, "**Helper:** I've changed the role color to `#"+color+"` on the role `"+role+"`")
k, err := dg.UserChannelCreate(js.Admin)
if err == nil {
	dg.ChannelTyping(k.ID)
	time.Sleep(1000 * time.Millisecond)
	dg.ChannelMessageSend(k.ID, "**Helper:** I've changed the role color to `#"+color+"` on the role `"+role+"`")
}
} // check if silent == false.


} // end of rolecolor parameter
























if RoleRename != "" {
	jolt := 0
	fmt.Println("Triggered Helper: Rename Role")
	id := ""
        guilds, err := dg.UserGuilds()
        if err != nil {
                fmt.Println(err)
                return
        }
        for _, v := range guilds {
        	id = v.ID
          //  fmt.Printf("%d : %#v\n", k, v)
        }
var roleID string
var hoist bool
var perms int
var color int

role := RoleRename
to := RoletoName


// newcolor := strconv.FormatInt(h, 16)
fmt.Println(role)
  roles, err := dg.GuildRoles(id)
  if err == nil {
    for _, v := range roles {
      if v.Name == role {
       roleID = v.ID
       hoist = v.Hoist
       perms = v.Permissions
       color = v.Color
      }
    }
  }


_, err = dg.GuildRoleEdit(id, roleID, to, color, hoist, perms)
if err != nil {
if Silent == false {
//  	fmt.Println("s.GuildRoles is the error")
// 	dg.ChannelMessageSend(id, "**Helper:** I've changed the role color to `#"+color+"` on the role `"+role+"`")
k, err := dg.UserChannelCreate(js.Admin)
if err == nil {
	jolt++
	dg.ChannelTyping(k.ID)
	time.Sleep(1000 * time.Millisecond)
	dg.ChannelMessageSend(k.ID, "**Helper:** There was an error!\n ```fix\nMake sure I have permissions to edit roles\nCheck the role name"+role+"```")
}

} // check if silent == false
}




if Silent == false && jolt == 0 {
// dg.ChannelMessageSend(id, "**Helper:** I've changed the role color to `#"+color+"` on the role `"+role+"`")
k, err := dg.UserChannelCreate(js.Admin)
if err == nil {
	dg.ChannelTyping(k.ID)
	time.Sleep(1000 * time.Millisecond)
	dg.ChannelMessageSend(k.ID, "**Helper:** I've changed the role `"+role+"` to `"+to+"`")
}
} // check if silent == false.


} // end of role rename parameter
















/*
if NewUpdateChk {
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
					if thefile != "System/custom/commands.json" && thefile != "System/custom/responses.json" && thefile != "config.json" && thefile != "System/helper/helper.exe" && thefile != "System/helper/helper.exe" {
						if thefile != "updates.auf" {
							err = os.Remove("../../"+thefile)
							if err != nil {
								fmt.Println(err)
							}
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


				if thefile == "System/custom/config.json" {
					pass = true
					a := DownloadFile("../updates/responses.json", get)
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
	} // make sure there isn't an error.
} // end of NewUpdateChk
*/













if UpdateChk {
	fmt.Println("Checking for updates...")
	time.Sleep(1000 * time.Millisecond)
	a := DownloadFile("../../autogo.exe", "https://raw.githubusercontent.com/proxikal/AutoGo/master/autogo.exe")
	if a == false {
		fmt.Println("autogo.exe Error => An error has occured. Check your internet. make sure Your Autogo process is closed completely.\nAnd don't move helper from his original path /System/helper")
	} else {
		fmt.Println("autogo.exe was updated.")
	}
	b := DownloadFile("../../autogo_32bit.exe", "https://raw.githubusercontent.com/proxikal/AutoGo/master/autogo_32bit.exe")
	if b == false {
		fmt.Println("autogo_32bit.exe Error => An error has occured. Check your internet, make sure Your Autogo process is closed completely.\nAnd don't move helper from his original path /System/helper")
	} else {
		fmt.Println("autogo_32bit.exe was updated.")
	}

	c := DownloadFile("../updates/commands.json", "https://raw.githubusercontent.com/proxikal/AutoGo/master/System/custom/commands.json")
	if c == false {
		fmt.Println("commands.json Error => An error has occured. Check your internet, make sure commands.json exists in System/custom/\nAnd don't move helper from his original path /System/helper")
	} else {
		fmt.Println("commands.json placed in update queue.")
	}

	d := DownloadFile("../updates/config.json", "https://raw.githubusercontent.com/proxikal/AutoGo/master/config.json")
	if d == false {
		fmt.Println("config.json Error => An error has occured. Check your internet, make sure config.json file exists in autogo's folder.\nAnd don't move helper from his original path /System/helper")
	} else {
		fmt.Println("config.json placed in update queue.")
	}

	e := DownloadFile("../updates/responses.json", "https://raw.githubusercontent.com/proxikal/AutoGo/master/System/custom/responses.json")
	if e == false {
		fmt.Println("responses.json Error => An error has occured. Check your internet, make sure responses.json file exists in System/custom\nAnd don't move helper from his original path /System/helper")
	} else {
		fmt.Println("responses.json placed in update queue.")
	}
	time.Sleep(5000 * time.Millisecond)
	/*
	cc := exec.Command("clear")
	cc.Stdout = os.Stdout
	cc.Run()
	*/
	// fmt.Println("You need to run StartHelper.bat Type 1 and hit enter to update json files.")

	time.Sleep(10000 * time.Millisecond)
} // end of update check













if Invite != "" {
switch runtime.GOOS {
case "linux":
    err = exec.Command("xdg-open", "https://discordapp.com/oauth2/authorize?&client_id="+Invite+"&scope=bot&permissions=36961343").Start()
case "windows", "darwin":
    err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://discordapp.com/oauth2/authorize?&client_id="+Invite+"&scope=bot&permissions=36961343").Start()
default:
    err = fmt.Errorf("can't open url. unsupported platform. you need to obtain your Client ID in your application.\nfrom there: https://discordapp.com/oauth2/authorize?&client_id=CLIENTIDHERE&scope=bot&permissions=36961343")
}
} // end of Invite



if MakeApp {
switch runtime.GOOS {
case "linux":
    err = exec.Command("xdg-open", "https://discordapp.com/developers/applications/me").Start()
case "windows", "darwin":
    err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://discordapp.com/developers/applications/me").Start()
default:
    err = fmt.Errorf("can't open url. unsupported platform. you need to visit https://discordapp.com/developers/applications/me")
}
} // end of Invite









if AlertChannel != "" {
	id := ""

        guilds, err := dg.UserGuilds()
        if err != nil {
                fmt.Println(err)
                return
        }
        for _, v := range guilds {
        	id = v.ID
          //  fmt.Printf("%d : %#v\n", k, v)
        }


th, err := ioutil.ReadFile("../custom/notice.txt")
if err != nil {
	fmt.Println("There was an error read below:")
	fmt.Println(err)
} else {

	if AlertHour == "0" {
	if AlertChannel == "general" {
		dg.ChannelMessageSend(id, string(th))
	} else {
		dg.ChannelMessageSend(AlertChannel, string(th))
	}
	} else {
	G:
	if AlertChannel == "general" {
		dg.ChannelMessageSend(id, string(th))
	} else {
		dg.ChannelMessageSend(AlertChannel, string(th))
	}
	newsel, err := strconv.Atoi(AlertHour)
	if err == nil {
	if newsel == 1 {
	time.Sleep(1 * time.Hour)
	}
	if newsel == 2 {
	time.Sleep(2 * time.Hour)
	}
	if newsel == 3 {
	time.Sleep(3 * time.Hour)
	}
	if newsel == 4 {
	time.Sleep(4 * time.Hour)
	}
	if newsel == 5 {
	time.Sleep(5 * time.Hour)
	}
	if newsel == 6 {
	time.Sleep(1 * time.Hour)
	}
	if newsel == 7 {
	time.Sleep(1 * time.Hour)
	}
	if newsel == 8 {
	time.Sleep(1 * time.Hour)
	}
	if newsel == 9 {
	time.Sleep(1 * time.Hour)
	}
	if newsel == 10 {
	time.Sleep(1 * time.Hour)
	}
	if newsel == 24 {
	time.Sleep(24 * time.Hour)
	}
	goto G
	}

	}
}

} // end of alert channel















if MakeRoles {
isBot := 0
isMute := 0
id := ""

        guilds, err := dg.UserGuilds()
        if err != nil {
                fmt.Println(err)
                return
        }
        for _, v := range guilds {
        	id = v.ID
          //  fmt.Printf("%d : %#v\n", k, v)
        }


        roles, err := dg.GuildRoles(id)
        if err != nil {
                fmt.Println(err)
                return
        }

        for _, v := range roles {
        	if v.Name == "Bot Commander" {
        		isBot++
        	
        	}
        	if v.Name == "muted" {
        		isMute++
        	
        	}
           //     fmt.Printf("%d : %#v\n", k, v)
        }




    if isBot == 0 {
    	botrole, err := dg.GuildRoleCreate(id)
    	if err != nil {
    		fmt.Println(err)
    	} else {
    	_, err = dg.GuildRoleEdit(id, botrole.ID, "Bot Commander", 0, false, 0)
    	if err != nil {
    		fmt.Println(err)
    	} else {
    		fmt.Println("Bot Commander role created.")
    	}

    	}
    }

time.Sleep(1000 * time.Millisecond)

    if isMute == 0 {
    	botrole, err := dg.GuildRoleCreate(id)
    	if err != nil {
    		fmt.Println(err)
    	} else {
    	_, err = dg.GuildRoleEdit(id, botrole.ID, "muted", 0, false, 0)
    	if err != nil {
    		fmt.Println(err)
    	} else {
    		fmt.Println("muted role created.")
    	}

    	}
    }

time.Sleep(10000 * time.Millisecond)
} // end of makeroles



if Initiate {
OwnerID := ""

  guilds, err := dg.UserGuilds()
  if err != nil {
   fmt.Println(err)
    return
  }
  for _, v := range guilds {
    OwnerID = v.OwnerID
    //  fmt.Printf("%d : %#v\n", k, v)
  }

if OwnerID != "" {
newConf := obj{
		Bot:			js.Bot,
		Admin:			OwnerID,
		Status:			js.Status,
		BotMaster:		js.BotMaster,
		BotCommander:	js.BotCommander,
		CmdsRun:		js.CmdsRun,
		Prefix:			js.Prefix,
		GreetMsg:		js.GreetMsg,
		ByeMsg:			js.ByeMsg,
		RoleSys:		js.RoleSys,
		Name:			js.Name,
		AntiLink:		js.AntiLink,
		Action:			js.Action,
		Silent:			js.Silent,
		HelpCmd:		js.HelpCmd,
		BotAutoRole:	js.BotAutoRole,
		}
	b, err := json.MarshalIndent(newConf, "", "   ")
	if err == nil {
		ioutil.WriteFile("config.json", b, 0777)
		fmt.Println("My owner has been found! Complete.")
	}
	if err != nil {
		fmt.Println("Something went wrong. config file was not found or broken.")
		fmt.Println(err)
	}
} else {
	fmt.Println("I'm sorry I couldn't find the Serve Owner. Private message me to gain ownership!")
}
} // end of initiate












if UpdateFiles {
	work := true
	cn := 0

	var newcommand map[string]interface{}
	var oldcommand map[string]interface{}


// ################## COMMANDS.JSON UPDATE ########################
	file, err := ioutil.ReadFile("../custom/commands.json")
	if err != nil {
		work = false
		fmt.Println("Can't find your original commands.json file.")
	}
	file1, err := ioutil.ReadFile("../updates/commands.json")
	if err == nil && work == true{
	json.Unmarshal(file, &oldcommand)	
	json.Unmarshal(file1, &newcommand)


for k, _ := range oldcommand {
	cn++
	for k1, _ := range newcommand {
		if k == k1 {
			delete(newcommand, k1)
		}
	} // end of oldcommand for loop
} // end of newcommand for loop


for k2, v2 := range newcommand {
 oldcommand[k2] = v2
}


	b, err := json.MarshalIndent(oldcommand, "", "   ")
	if err == nil {
		ioutil.WriteFile("../custom/commands.json", b, 0777)
		fmt.Println("[Worked] I have updated your commands.json")
		os.Remove("../updates/commands.json")
	}	else {
		fmt.Println("[Failed] Something has happend read below:")
		fmt.Println(err)
	}
	} else {
		fmt.Println("[Skip] commands.json update file not found in updates folder.")
	}

} // end of updatefiles for commands.json













if UpdateFiles2 {
	work := true
	cn := 0
	var newresponse map[string]interface{}
	var oldresponse map[string]interface{}

// ################### RESPONSES.JSON UPDATE ########################
work = true

	file2, err := ioutil.ReadFile("../custom/responses.json")
	if err != nil {
		work = false
		fmt.Println("Can't find your original responses.json file.")
	}
	file3, err := ioutil.ReadFile("../updates/responses.json")
	if err == nil && work == true{
	json.Unmarshal(file2, &oldresponse)	
	json.Unmarshal(file3, &newresponse)


for k, _ := range oldresponse {
	cn++
	for k1, _ := range newresponse {
		if k == k1 {
			delete(newresponse, k1)
		}
	} // end of oldcommand for loop
} // end of newcommand for loop

for k2, v2 := range newresponse {
oldresponse[k2] = v2
}

	b2, err := json.MarshalIndent(oldresponse, "", "   ")
	if err == nil {
		ioutil.WriteFile("../custom/responses.json", b2, 0777)
		fmt.Println("[Worked] I have updated your responses.json")
		os.Remove("../updates/responses.json")
	}	else {
		fmt.Println("[Failed] Something has happend read below:")
		fmt.Println(err)
	}
	} else {
		fmt.Println("[Skip] responses.json update file not found in /updates/ folder.")
	}
work = true

} // end of updatefiles2 for responses.json















if UpdateFiles3 {
	work := true
	cn := 0
	var newconfig map[string]interface{}
	var oldconfig map[string]interface{}

	// ################ CONFIG.JSON #######################

	file4, err := ioutil.ReadFile("../../config.json")
	if err != nil {
		work = false
		fmt.Println("Can't find your original config.json file.")
	}
	file5, err := ioutil.ReadFile("../updates/config.json")
	if err == nil && work == true{
	json.Unmarshal(file4, &oldconfig)	
	json.Unmarshal(file5, &newconfig)


for k, _ := range oldconfig {
	cn++
	for k1, _ := range newconfig {
		if k == k1 && k != "Version" {
			delete(newconfig, k1)
		}
	} // end of oldcommand for loop
} // end of newcommand for loop


for k2, v2 := range newconfig {
oldconfig[k2] = v2
}


	b3, err := json.MarshalIndent(oldconfig, "", "   ")
	if err == nil {
		ioutil.WriteFile("../../config.json", b3, 0777)
		fmt.Println("[Worked] I have updated your config.json")
		os.Remove("../updates/config.json")
	}	else {
		fmt.Println("[Failed] Something has happend read below:")
		fmt.Println(err)
	}
	} else {
		fmt.Println("[Skip] config.json update file not found in /updates/ folder.")
	}

} // end of updatefiles3 config.json



} // end of main func