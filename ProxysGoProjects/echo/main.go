// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the New() helper function.
package main

import (
	"fmt"
	"time"
	"github.com/bwmarrin/discordgo"
	"strings"
	"encoding/json"
	"io/ioutil"
	"os"
	"net/http"
	"net/url"
	"strconv"
//	"github.com/zymtom/argconf"
)
	var err error
	var startTime time.Time
	var js obj
	var cmd commands









// until i can figure out how to get the roles from every member
// i have to do the masters system this way (the noob way)
/*
func isMaster(server string, user string) bool {
	mas := true

	var in info
	vfile, err := ioutil.ReadFile("servers/" + server + "/main.json")
	if err != nil {
		// mas = false
	}

	json.Unmarshal(vfile, &in)

	if _, err := os.Stat("servers/" + server + "/" + user + ".json"); err != nil {
		mas = false
	}

	if user == in.Owner {
		mas = true
	}
	return mas
}
*/





func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

















func main() {

	file, err := ioutil.ReadFile("config.json")
	json.Unmarshal(file, &js)

    // Login to discord. You can use a token or email, password arguments.
	dg, err := discordgo.New(js.Bot)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(GuildMemberAdd)
	dg.AddHandler(GuildMemberRemove)
	dg.AddHandler(onReady)
	// dg.AddHandler(GuildRoleUpdate)
	dg.AddHandler(GuildCreate)
	dg.AddHandler(GuildDelete)
	// Open the websocket and begin listening.
	dg.Open()

	// Simple way to keep program running until any key press.
	var input string
	fmt.Scanln(&input)
	return
}


















// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated user has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	chkErr := true

  c, err := s.State.Channel(m.ChannelID)
if err != nil {
c, err = s.Channel(m.ChannelID)
}


// let's try to prevent any errors from happening. hahaha....
if err != nil {
	chkErr = false // couldn't determine the Guild ID
}


// ####### FUNCTIONS: Located in the funcs.go file #######
	var in info

if chkErr == true {



	// Load up the custom commands.
	cfile, err := ioutil.ReadFile("servers/" + c.GuildID + "/commands.json")
	if err != nil {
		return
	} else {
	json.Unmarshal(cfile, &cmd)
	}



	// Load up the server information
	vfile, err := ioutil.ReadFile("servers/" + c.GuildID + "/main.json")
	if err != nil {
		return
	} else {
	json.Unmarshal(vfile, &in)
	}


js.BotCommander = GetRoleID(s, c.GuildID, "Bot Commander")	// GetRoleID returns the role id.
js.BotMaster = isMemberRole(s, c.GuildID, m.Author.ID, "Bot Commander")	// isMemberRole returns true or false

if m.Author.ID == in.Owner {
	js.BotMaster = true
}


// ############## Status Attempt #1 ##################
/*
ticker := time.NewTicker(2 * time.Minute)
quit := make(chan struct{})

    for {
       select {
        case <- ticker.C:
			myrand := random(1, 9)
			var status []string
			status, err := readLines("status.txt")
			if err == nil {
				s.UpdateStatus(0, status[myrand])
			}
        case <- quit:
            ticker.Stop()
            return
        }
    }
*/




//	fmt.Println("Prefix: " + in.Prefix + "\nGreet: " + in.GreetMsg + "\nBye: " + in.ByeMsg + "\nAutorole: " + in.RoleSys)
	// Print message to stdout.
//	fmt.Println("[" + in.Name + "] " + time.Now().Format(time.Stamp) + " - " + m.Author.Username + ": " + m.Content)







if c.GuildID == "148629493676769280" {
 // -#$-
var auto []string
auto, err := readLines("servers/148629493676769280/autoresponse.txt")
if err == nil {
	for _, ars := range auto {
	//	fmt.Println("RAW: " + ars)
		ardat := strings.Split(ars, "-#$-")
		trigger := ardat[0]
		response := ardat[1]
	//	fmt.Println("Trigger: " + trigger)
	//	fmt.Println("Response: " + response)
		// just a basic ARS trigger. Later i will code for {find=word}
		if m.Content == trigger {
			s.ChannelMessageSend(m.ChannelID, response)
		} // end of basic trigger
	}
}
} // only work in my server for now






















if strings.HasPrefix(m.Content, in.Prefix) {
	js.CmdsRun++
	newConf := obj{
		Bot:			js.Bot,
		Admin:			js.Admin,
		Status:			js.Status,
		CmdsRun:		js.CmdsRun,
		BotMaster:		js.BotMaster,
		BotCommander:	js.BotCommander,
		}
	b, err := json.Marshal(newConf)
	if err == nil {
		ioutil.WriteFile("config.json", b, 0777)
	}
}








// fmt.Println(js.BotMaster)














	if strings.HasPrefix(m.Content, in.Prefix + "help") {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)

js.BotMaster = isMemberRole(s, c.GuildID, m.Author.ID, "Bot Commander")	// isMemberRole returns true or false

if m.Author.ID == in.Owner {
	js.BotMaster = true
}


		if m.Content == in.Prefix + "help" {
			bm := "False"

			if js.BotMaster == true {
				bm = "True"
			}
			ts := js.CmdsRun
			i := strconv.Itoa(ts)
	//		fmt.Println("Converted: " + i)
	//		fmt.Println(js.CmdsRun)

		s.ChannelMessageSend(m.ChannelID, "```ruby\nEcho [BETA]\nlibrary: DiscordGO\nrequests: "+i+"\nyou: "+m.Author.Username+"\ncommander: "+bm+"\n---------\n"+in.Prefix+"help, "+in.Prefix+cmd.AddMaster+", "+in.Prefix+cmd.Greet+"\n"+in.Prefix+cmd.Bye+", "+in.Prefix+cmd.DenyLinks+", "+in.Prefix+cmd.AllowLinks+"\n"+in.Prefix+cmd.Prefix+", "+in.Prefix+cmd.Autorole+", "+in.Prefix+"invites\n"+in.Prefix+cmd.Kick+", "+in.Prefix+cmd.Ban+", "+in.Prefix+cmd.Giveme+"\n"+in.Prefix+cmd.SetPunish+", "+in.Prefix+cmd.Meme+", "+in.Prefix+cmd.Joke+"\n"+in.Prefix+cmd.Give+", "+in.Prefix+cmd.Take+", "+in.Prefix+cmd.Mute+"\n"+in.Prefix+cmd.Unmute+", "+in.Prefix+"rolecolor, "+in.Prefix+"giphy\n"+in.Prefix+"cats\n---------\n@Echo what's your prefix? displays prefix\nrun your own echo with an auto response system and more!\n: https://github.com/proxikal/AutoGo```")
		}


		if m.Content == in.Prefix + "help "+cmd.AddMaster {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+"addmaster @User\ninfo: gives user acces to mod commands.```")
		}

		if m.Content == in.Prefix + "help "+cmd.DelMaster {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+"delmaster @User\ninfo: removes access to mod commands.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Greet {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"greet Welcome {user} if you need any help ask!\ninfo: use {user} to mention the new member.\nuse {/user} to just say their username!\njust set the message to off for turning the greet message off```")
		}

		if m.Content == in.Prefix + "help "+cmd.Bye {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"bye {user} has left the server.\nto turn off set the bye message to off```")
		}

		if m.Content == in.Prefix + "help "+cmd.DenyLinks {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"denylinks\ninfo: enables my anti link module. by default i kick offenders. you can use setpunish command to change to ban.```")
		}

		if m.Content == in.Prefix + "help "+cmd.AllowLinks {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"allowlinks\ninfo: turns my anti link module off.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Prefix {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"prefix #\ninfo: sets my prefix in your server to #```")
		}

		if m.Content == in.Prefix + "help "+cmd.Autorole {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"autorole Role Name\ninfo: automatically assign a role to new members.\nsilently add roles type "+in.Prefix+"autorole -s Role Name```")
		}

		if m.Content == in.Prefix + "help "+cmd.Invites {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Everyone\nusage: "+in.Prefix+"invites\ninfo: gives you a list of available invite codes for your channel.```")
		}

		if m.Content == in.Prefix + "help mkinvite" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"mkinvite\ninfo: creates a permenant invite code for your channel.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Give {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+"give @User Role Name\ninfo: gives the user the specified role.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Take {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+"take @User Role Name\ninfo: takes the user the specified role.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Giveme {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+"giveme Role Name\ninfo: gives you the specified role.```")
		}

		if m.Content == in.Prefix + "help "+cmd.SetPunish {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"setpunish kick or ban\ninfo: you can set the anti links module punishment to either kick or ban. by default its set to kick.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Mute {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"mute @User\ninfo: mutes the user. you need to make a role named muted and set the permissions to not speak and than add the role to your channels.```")
		}

		if m.Content == in.Prefix + "help "+cmd.Unmute {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"unmute @User\ninfo: unmutes the user.```")
		}

		if m.Content == in.Prefix + "help rolecolor" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"rolecolor #000000 Role Name\ninfo: Changes the color to a role.```")
		}

	}







if strings.ToLower(m.Content) == "<@147463276840747008> what's your prefix?" {
s.ChannelMessageSend(m.ChannelID, "My prefix in your server is `"+in.Prefix+"`")
}

if strings.ToLower(m.Content) == "<@147463276840747008> whats your prefix?" {
s.ChannelMessageSend(m.ChannelID, "My prefix in your server is `"+in.Prefix+"`")
}

if strings.ToLower(m.Content) == "<@147463276840747008> whats your prefix" {
s.ChannelMessageSend(m.ChannelID, "My prefix in your server is `"+in.Prefix+"`")
}

if strings.ToLower(m.Content) == "<@147463276840747008> what's your prefix" {
s.ChannelMessageSend(m.ChannelID, "My prefix in your server is `"+in.Prefix+"`")
}

if strings.ToLower(m.Content) == "<@147463276840747008> what is your prefix" {
s.ChannelMessageSend(m.ChannelID, "My prefix in your server is `"+in.Prefix+"`")
}

if strings.ToLower(m.Content) == "<@147463276840747008> what is your prefix?" {
s.ChannelMessageSend(m.ChannelID, "My prefix in your server is `"+in.Prefix+"`")
}




if in.Owner == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + "rename") {
	str := strings.Replace(m.Content, in.Prefix + "rename ", "", -1)
	tr := strings.Split(str, " ")
	Before := tr[0]
	After := tr[1]

	if Before == cmd.Greet {
		cmd.Greet = After
	}
	if Before == cmd.Bye {
		cmd.Bye = After
	}
	if Before == cmd.Prefix {
		cmd.Prefix = After
	}
	if Before == cmd.Kick {
		cmd.Kick = After
	}
	if Before == cmd.Ban {
		cmd.Ban = After
	}
	if Before == cmd.Autorole {
		cmd.Autorole = After
	}
	if Before == cmd.SetPunish {
		cmd.SetPunish = After
	}
	if Before == cmd.AllowLinks {
		cmd.AllowLinks = After
	}
	if Before == cmd.DenyLinks {
		cmd.DenyLinks = After
	}
	if Before == cmd.AddMaster {
		cmd.AddMaster = After
	}
	if Before == cmd.DelMaster {
		cmd.DelMaster = After
	}
	if Before == cmd.Invites {
		cmd.Kick = After
	}
	if Before == cmd.Meme {
		cmd.Meme = After
	}
	if Before == cmd.Joke {
		cmd.Joke = After
	}
	if Before == cmd.Give {
		cmd.Give = After
	}
	if Before == cmd.Take {
		cmd.Take = After
	}
	if Before == cmd.Giveme {
		cmd.Giveme = After
	}
	if Before == cmd.Mute {
		cmd.Mute = After
	}
	if Before == cmd.Unmute {
		cmd.Unmute = After
	}

	newConf := commands{
		Greet: 			cmd.Greet,
		Bye:			cmd.Bye,
		Prefix:			cmd.Prefix,
		Kick:			cmd.Kick,
		Ban:			cmd.Ban,
		Autorole:		cmd.Autorole,
		SetPunish:		cmd.SetPunish,
		AllowLinks:		cmd.AllowLinks,
		DenyLinks:		cmd.DenyLinks,
		AddMaster:		cmd.AddMaster,
		DelMaster:		cmd.DelMaster,
		Invites:		cmd.Invites,
		Meme:			cmd.Meme,
		Joke:			cmd.Joke,
		Give:			cmd.Give,
		Take:			cmd.Take,
		Giveme:			cmd.Giveme,
		Mute:			cmd.Mute,
		Unmute:			cmd.Unmute,
	}
	b, err := json.Marshal(newConf)
	if err == nil {
		ioutil.WriteFile("servers/" + c.GuildID + "/commands.json", b, 0777)
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I've renamed the command `"+Before+"` to `"+After+"`")
	}


}


















if strings.HasPrefix(strings.ToLower(m.Content), in.Prefix + "cats") {
	var img map[string]interface{}

	getJson("http://random.cat/meow", &img)

	newcat := img["file"].(string)
	s.ChannelMessageSend(m.ChannelID, newcat)
}




if strings.HasPrefix(strings.ToLower(m.Content), in.Prefix + "giphy") {
	str := strings.Replace(m.Content, in.Prefix + "giphy ", "", -1)
	str = strings.Replace(str, " ", "+", -1)

	var img map[string]map[string]interface{}
	getJson("http://api.giphy.com/v1/gifs/random?api_key=dc6zaTOxFJmzC&tag="+str, &img)

	newcat := img["data"]["image_original_url"].(string)
	s.ChannelMessageSend(m.ChannelID, newcat)
}




if strings.HasPrefix(strings.ToLower(m.Content), in.Prefix + "locateip") {
	str := strings.Replace(m.Content, in.Prefix + "locateip ", "", -1)

	var ip map[string]interface{}
	getJson("http://ip-api.com/json/"+str, &ip)

	region := ip["regionName"].(string)
	zipcode := ip["zip"].(string)
	longitude := ip["lon"].(string)
	latitude := ip["lat"].(string)
	country := ip["country"].(string)
	city := ip["city"].(string)
	timezone := ip["timezone"].(string)
	isp := ip["isp"].(string)
	who := ip["as"].(string)
	theip := ip["query"].(string)

	format := "```ruby\n"+"ip: "+theip+"\ncountry: "+country+"\nregion: "+region+"\ncity: "+city+"\nlongitude: "+longitude+"\nlatitude: "+latitude+"\nzipcode: "+zipcode+"\ntimezone: "+timezone+"\nisp: "+isp+"\n"+who+"```"
	s.ChannelMessageSend(m.ChannelID, format)
}








if in.Owner == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + cmd.AddMaster) {
str := strings.Replace(m.Content, in.Prefix + cmd.AddMaster+" ", "", -1)
str = strings.Replace(str, "<@", "", -1)
str = strings.Replace(str, ">", "", -1)

z, err := s.State.Member(c.GuildID, str)
if err != nil {
return
}
		roles, err := s.GuildRoles(c.GuildID)
		if err == nil {
			for _, v := range roles {
    			if v.Name == "Bot Commander" {
    				z.Roles = append(z.Roles, v.ID)
    				s.GuildMemberEdit(c.GuildID, str, z.Roles)
					s.ChannelTyping(m.ChannelID)
					time.Sleep(1000 * time.Millisecond)
					s.ChannelMessageSend(m.ChannelID, "I've added <@" + str + "> as a `Bot Commander`")
    			}
			}
		}
	} // end of Add master command



















if in.Owner == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + "leave") {
	s.ChannelTyping(m.ChannelID)
	time.Sleep(1000 * time.Millisecond)
	s.ChannelMessageSend(m.ChannelID, "Thanks for having me!")
	s.GuildLeave(c.GuildID)
}



















if in.Owner == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + cmd.DelMaster) {
str := strings.Replace(m.Content, in.Prefix + cmd.DelMaster+" ", "", -1)
str = strings.Replace(str, "<@", "", -1)
str = strings.Replace(str, ">", "", -1)

x, err := s.State.Member(c.GuildID, str)
if err != nil {
x, err = s.GuildMember(c.GuildID, str)
}
	if err == nil {
	var mc []string
	mc = x.Roles
	for mr := range x.Roles {
		t := mc[mr]
		if strings.Contains(t, js.BotCommander) {
    		// z.Roles = append(z.Roles, t[:0])
    		x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
    		if err != nil {
    			return
    		}
    		s.GuildMemberEdit(c.GuildID, str, x.Roles)
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, "I've Removed <@" + str + "> from `Bot Commander` position.")
		}
	}
	}// end of err check
} // end of Del master command





















if in.Owner == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + cmd.Take) {
str := strings.Replace(m.Content, in.Prefix + cmd.Take + " ", "", -1)
dat := strings.Split(str, " ")
usr := dat[0]
usr = strings.Replace(usr, "<@", "", -1)
usr = strings.Replace(usr, ">", "", -1)
role := strings.Replace(str, "<@"+usr+"> ", "", -1)

 var roleID string

mroles, err := s.GuildRoles(c.GuildID)
if err == nil {
 for _, v := range mroles {
    if v.Name == role {
//    	fmt.Println("Found the role: "+role+"\nID: "+v.ID)
    	roleID = v.ID
    }
  }
  }


x, err := s.GuildMember(c.GuildID, usr)
	if err != nil {
		fmt.Println(err)
	}

	if err == nil {
		var ms []string
		ms = x.Roles
		for mr := range x.Roles {
			t := ms[mr]
			if strings.Contains(t, roleID) {
				//fmt.Println("Membert has role: "+t)
    			x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
    			s.GuildMemberEdit(c.GuildID, usr, x.Roles)
				s.ChannelTyping(m.ChannelID)
				time.Sleep(1000 * time.Millisecond)
				s.ChannelMessageSend(m.ChannelID, "I've taken <@"+usr+">'s role `"+role+"`")
    		}
		}
	}
} // end of giveme command.
























if in.Owner == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + cmd.Give) {
str := strings.Replace(m.Content, in.Prefix + cmd.Give + " ", "", -1)
dat := strings.Split(str, " ")
usr := dat[0]
usr = strings.Replace(usr, "<@", "", -1)
usr = strings.Replace(usr, ">", "", -1)
role := strings.Replace(str, "<@"+usr+"> ", "", -1)

x, err := s.State.Member(c.GuildID, usr)
if err != nil {
return
}
		roles, err := s.GuildRoles(c.GuildID)
		if err == nil {
			for _, v := range roles {
    			if v.Name == role {
    				x.Roles = append(x.Roles, v.ID)
    				s.GuildMemberEdit(c.GuildID, usr, x.Roles)
					s.ChannelTyping(m.ChannelID)
					time.Sleep(1000 * time.Millisecond)
					s.ChannelMessageSend(m.ChannelID, "I've given <@"+usr+"> the role `" + role + "`")
    			}
			}
		}
	} // end of giveme command.




























if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Mute) {
str := strings.Replace(m.Content, in.Prefix + cmd.Mute + " ", "", -1)
str = strings.Replace(str, "<@", "", -1)
str = strings.Replace(str, ">", "", -1)

// fmt.Println("Adding Master: "+str)

z, err := s.State.Member(c.GuildID, str)
if err != nil {
return
}
		roles, err := s.GuildRoles(c.GuildID)
		if err == nil {
			for _, v := range roles {
    			if strings.ToLower(v.Name) == "muted" {
    				z.Roles = append(z.Roles, v.ID)
    				s.GuildMemberEdit(c.GuildID, str, z.Roles)
					s.ChannelTyping(c.GuildID)
					time.Sleep(1000 * time.Millisecond)
					s.ChannelMessageSend(m.ChannelID, "I've muted <@" + str + "> in <#"+m.ChannelID+">")
    			}
			}
		}
	} // end of giveme command.

























if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Unmute) {
str := strings.Replace(m.Content, in.Prefix + cmd.Unmute + " ", "", -1)
usr := strings.Replace(str, "<@", "", -1)
usr = strings.Replace(usr, ">", "", -1)
	var roleID string

mroles, err := s.GuildRoles(c.GuildID)
if err == nil {

 for _, v := range mroles {
    if v.Name == "muted" {
    	roleID = v.ID
    }
  }
  }
x, err := s.State.Member(c.GuildID, usr)
if err != nil {
x, err = s.GuildMember(c.GuildID, usr)
}

if err != nil {
	fmt.Println(err)
	s.ChannelMessageSend(m.ChannelID, "You don't have the role `muted` setup in your server.")
} else {
	var ms []string
	ms = x.Roles
	for mr := range x.Roles {
		t := ms[mr]
		if t == roleID {
    		x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
    		s.GuildMemberEdit(c.GuildID, usr, x.Roles)
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, "I've unmuted <@"+usr+">")
			}
		}
	}
} // end of giveme command.















if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "rolecolor") {
str := strings.Replace(m.Content, in.Prefix + "rolecolor ", "", -1)
var roleID string
var hoist bool
var perms int

newdata := strings.Split(str, " ")
color := newdata[0]
role := strings.Replace(str, color + " ", "", -1)
color = strings.Replace(color, "#", "", -1)

// newcolor := strconv.FormatInt(h, 16)
fmt.Println(role)
  roles, err := s.GuildRoles(c.GuildID)
  if err == nil {
    for _, v := range roles {
      if v.Name == role {
       roleID = v.ID
       hoist = v.Hoist
       perms = v.Permissions
      }
    }
  } else {
  	fmt.Println("s.GuildRoles is the error")
  }

var ij int
newcode, _ := strconv.ParseInt(color, 16, 0)
d := fmt.Sprintf("%d", newcode)
fmt.Println(d)
ij, err = strconv.Atoi(d)
if err != nil {
	fmt.Println(err)
}
// if roleID != "" {
// roleID := GetRoleID(s, c.GuildID, role)
_, err = s.GuildRoleEdit(c.GuildID, roleID, role, ij, hoist, perms)
if err != nil {
  	fmt.Println("s.GuildRoles is the error")
  }
s.ChannelMessageSend(m.ChannelID, "I've changed the role color to `#"+color+"` for **"+role+"**")
// }
} // end of role color
















if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Greet) {
		str := strings.Replace(m.Content, in.Prefix + cmd.Greet+" ", "", -1)


			newjs := info{
				Prefix:		in.Prefix,
				GreetMsg:	str,
				RoleSys:	in.RoleSys,
				ByeMsg:		in.ByeMsg,
				Owner:		in.Owner,
				Name:		in.Name,
				AntiLink:	in.AntiLink,
				Action:		in.Action,
				Silent:		in.Silent,
			}
			b, err := json.Marshal(newjs)
			if err == nil {
				ioutil.WriteFile("servers/" + c.GuildID + "/main.json", b, 0777)
			}
	if str != "off" {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I have changed the Greet message: ```ruby\n"+str+"```")
	} else {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I have turned the Greet message `off` ")
	}
}




















if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Bye) {
	str := strings.Replace(m.Content, in.Prefix + cmd.Bye + " ", "", -1)
	newjs := info{
		Prefix:		in.Prefix,
		GreetMsg:	in.GreetMsg,
		RoleSys:	in.RoleSys,
		ByeMsg:		str,
		Owner:		in.Owner,
		Name:		in.Name,
		AntiLink:	in.AntiLink,
		Action:		in.Action,
		Silent:		in.Silent,
	}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("servers/" + c.GuildID + "/main.json", b, 0777)
		}
	if str != "off" {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I have changed the Bye message: ```ruby\n"+str+"```")
	} else {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I have turned the Bye message `off`")
	}
}




















	if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "mkinvite") {

		b := discordgo.Invite{
			MaxAge:		0,
			MaxUses:	0,
			Temporary:	false,
			XkcdPass:	"",
		}
	iv, err := s.ChannelInviteCreate(m.ChannelID, b)
	if err != nil {
		return
	}
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I've made the Invite Code: "+iv.Code)
		return
	}




















	if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.DenyLinks) {
		newjs := info{
			Prefix:		in.Prefix,
			GreetMsg:	in.GreetMsg,
			RoleSys:	in.RoleSys,
			ByeMsg:		in.ByeMsg,
			Owner:		in.Owner,
			Name:		in.Name,
			AntiLink:	true,
			Action:		in.Action,
			Silent:		in.Silent,
		}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("servers/" + c.GuildID + "/main.json", b, 0777)
		}

		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "Links are no longer allowed on this server.")
		return
	}























	if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.AllowLinks) {
		newjs := info{
			Prefix:		in.Prefix,
			GreetMsg:	in.GreetMsg,
			RoleSys:	in.RoleSys,
			ByeMsg:		in.ByeMsg,
			Owner:		in.Owner,
			Name:		in.Name,
			AntiLink:	false,
			Action:		in.Action,
			Silent:		in.Silent,
		}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("servers/" + c.GuildID + "/main.json", b, 0777)
		}

		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "Links are allowed on this server.")
		return
	}



















	if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Prefix) {
		str := strings.Replace(m.Content, in.Prefix + cmd.Prefix + " ", "", -1)

		newjs := info{
			Prefix:		str,
			GreetMsg:	in.GreetMsg,
			RoleSys:	in.RoleSys,
			ByeMsg:		in.ByeMsg,
			Owner:		in.Owner,
			Name:		in.Name,
			AntiLink:	in.AntiLink,
			Action:		in.Action,
			Silent:		in.Silent,
		}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("servers/" + c.GuildID + "/main.json", b, 0777)
		}

		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I have changed this servers prefix to `"+str+"`")
		return
	}





















	if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.SetPunish) {
		str := strings.Replace(m.Content, in.Prefix + cmd.SetPunish + " ", "", -1)

		check := false

		if strings.ToLower(str) == "kick" {
			check = true
		}
		if strings.ToLower(str) == "ban" {
			check = true
		}

		if check == true {
			newjs := info{
				Prefix:		in.Prefix,
				GreetMsg:	in.GreetMsg,
				RoleSys:	in.RoleSys,
				ByeMsg:		in.ByeMsg,
				Owner:		in.Owner,
				Name:		in.Name,
				AntiLink:	in.AntiLink,
				Action:		str,
				Silent:		in.Silent,
			}
			b, err := json.Marshal(newjs)
			if err == nil {
				ioutil.WriteFile("servers/" + c.GuildID + "/main.json", b, 0777)
			}
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, "I have changed the Antilink punishment to `"+str+"`")
		}
		if check == false {
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, "You need to pick a proper punishment for Anti links type `"+in.Prefix+"setpunish kick` or `"+in.Prefix+"setpunish ban`")
		}
		return
	}




















if js.BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + cmd.Autorole) {
str := strings.Replace(m.Content, in.Prefix + cmd.Autorole + " ", "", -1)
var br bool
br = false
if strings.Split(str, " ")[0] == "-s" {
br = true
str = strings.Replace(m.Content, in.Prefix + cmd.Autorole + " -s ", "", -1)
}

cnt := 0
roles, err := s.GuildRoles(c.GuildID)
if err == nil {

if str != "off" {
 for _, v := range roles {
    if v.Name == str {
    	cnt++
		newjs := info{
			Prefix:		in.Prefix,
			GreetMsg:	in.GreetMsg,
			RoleSys:	str,
			ByeMsg:		in.ByeMsg,
			Owner:		in.Owner,
			Name:		in.Name,
			AntiLink:	in.AntiLink,
			Action:		in.Action,
			Silent:		br,
			}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("servers/" + c.GuildID + "/main.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "New people will get the role `"+str+"`")
		return
	}
}
} // make sure they don't want to turn the autorole off


if str == "off" {
	cnt = 1
		newjs := info{
			Prefix:		in.Prefix,
			GreetMsg:	in.GreetMsg,
			RoleSys:	"off",
			ByeMsg:		in.ByeMsg,
			Owner:		in.Owner,
			Name:		in.Name,
			AntiLink:	in.AntiLink,
			Action:		in.Action,
			Silent:		in.Silent,
			}
		b, err := json.Marshal(newjs)
		if err == nil {
			ioutil.WriteFile("servers/" + c.GuildID + "/main.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "Autorole => `Disabled`")
}




	if cnt == 0 {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I can't find the role `"+str+"` make sure to check spelling, this is case sensitive.")
	}
}
}

























	if strings.HasPrefix(m.Content, in.Prefix + "invites") {
		o, err := s.ChannelInvites(m.ChannelID)
		if err == nil {
			data := ""
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, "Invites for: `"+in.Name+"`\n```ruby\nGrabbing Results..```")
			for _, v := range o {
			data = data + "\nissuer: "+v.Inviter.Username+"\ncode: "+v.Code
			// s.ChannelMessageDelete(m.ChannelID, theid)
			//s.ChannelMessageEdit(m.ChannelID, m.ID, "Invites for: `"+c.GuildID+"` "+v.Inviter.Username + "\n" + v.Code)
  		}
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
  		s.ChannelMessageSend(m.ChannelID, "```ruby\n"+data+"```")
		return
		}
	}












// fmt.Println("Server ID: " + c.GuildID)



















// let's see if they want advertising disabled
if in.AntiLink == true && js.BotMaster == false {
	var deny [5]string
	deny[0] = "https://"
	deny[1] = "http://"
	deny[2] = ".com"
	deny[3] = ".net"
	deny[4] = "www."

	for i := 0; i <= 4; i++ {
		if strings.Contains(strings.ToLower(strings.ToLower(m.Content)), deny[i]) && in.Owner != m.Author.ID {
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, "I have kicked <@" + m.Author.ID + "> For advertising.")
			s.GuildMemberDelete(c.GuildID, m.Author.ID)
		}
	} // end of for loop
} // end of anti link system














/* this is to change the bot's name
if strings.Contains(strings.ToLower(m.Content), "~!@#convert#@!~") && js.Admin == m.Author.ID {
	s.ChannelMessageSend(m.ChannelID, "Super Saiyan 3")
	s.UserUpdate("proxikal2@gmail.com", "Joshua(1)", "Echo", "new.png", "")
}
*/



	if strings.HasPrefix(m.Content, in.Prefix + cmd.Meme) {
		myrand := random(1, 52)
		var meme []string
		meme, err := readLines("meme.txt")
		if err == nil {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, meme[myrand])
		}
	}




	if strings.HasPrefix(m.Content, in.Prefix + cmd.Joke) {
		myrand := random(1, 81)
		var meme []string
		meme, err := readLines("random.txt")
		if err == nil {
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, meme[myrand])
		}
	}




	if strings.HasPrefix(m.Content, in.Prefix + cmd.Kick) && js.BotMaster == true {
		str := strings.Replace(m.Content, in.Prefix + cmd.Kick + " ", "", -1)
		str = strings.Replace(str, "<@", "", -1)
		str = strings.Replace(str, ">", "", -1)
		// fmt.Println("the"+str+"string")
		s.ChannelTyping(m.ID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I have kicked <@" + str + "> From the server.")
		s.GuildMemberDelete(c.GuildID, str)
	}









	if strings.HasPrefix(m.Content, in.Prefix + cmd.Ban) && js.BotMaster == true {
		str := strings.Replace(m.Content, in.Prefix + cmd.Ban + " ", "", -1)
		str = strings.Replace(str, "<@", "", -1)
		str = strings.Replace(str, ">", "", -1)
		s.ChannelTyping(m.ID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "I have banned <@" + str + "> From the server.")
		s.GuildBanCreate(c.GuildID, str, 10)
	}











	if m.Author.ID == js.Admin && strings.HasPrefix(m.Content, in.Prefix + "status") {
		str := strings.Replace(m.Content, in.Prefix + "status ", "", -1)
		s.ChannelTyping(m.ChannelID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "Attempting status change.")
		s.UpdateStatus(0, str)
		return
	}









if in.Owner == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + cmd.Giveme) {
str := strings.Replace(m.Content, in.Prefix + cmd.Giveme + " ", "", -1)
z, err := s.State.Member(c.GuildID, m.Author.ID)
if err != nil {
return
}
		roles, err := s.GuildRoles(c.GuildID)
		if err == nil {
			for _, v := range roles {
    			if v.Name == str {
    				z.Roles = append(z.Roles, v.ID)
    				s.GuildMemberEdit(c.GuildID, m.Author.ID, z.Roles)
					s.ChannelTyping(c.GuildID)
					time.Sleep(1000 * time.Millisecond)
					s.ChannelMessageSend(c.GuildID, "You have assumed the role `" + str + "`")
    			}
			}
		}
	} // end of giveme command.





// ############## Pm's??

if c.GuildID == "" {
	fmt.Println("Is PM")
	k, err := s.UserChannelCreate(m.Author.ID)
	if err == nil {
//	if strings.HasPrefix(m.Content, "--pm") {
		s.ChannelTyping(k.ID)
		time.Sleep(1000 * time.Millisecond)
		s.ChannelMessageSend(k.ID, "Hi! Pm based commands are disabled. I'm accepting some servers! You have to ask Proxy in his server: https://discord.gg/0pTKzt2BDInBOrxL")
//	}
}
}



} // end of chkErr


} // ##########   END OF messageCreate










func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	var in info
	vfile, err := ioutil.ReadFile("servers/" + m.GuildID + "/main.json")
	json.Unmarshal(vfile, &in)

	// fmt.Println(in.RoleSys)
	roles, err := s.GuildRoles(m.GuildID)

		if in.GreetMsg != "" && in.GreetMsg != "off" {
			s.ChannelTyping(m.GuildID)
			time.Sleep(1000 * time.Millisecond)
			data := strings.Replace(in.GreetMsg, "{user}", "<@"+m.User.ID+">", -1)
			data = strings.Replace(data, "{/user}", m.User.Username, -1)
			s.ChannelMessageSend(m.GuildID, data)
		}

		if err == nil {
 			for _, v := range roles {
    			if v.Name == in.RoleSys {
    				if in.RoleSys != "" && in.RoleSys != "off" {
    				m.Roles = append(m.Roles, v.ID)
    				s.GuildMemberEdit(m.GuildID, m.User.ID, m.Roles)
    				if in.Silent == false {
						s.ChannelMessageSend(m.GuildID, "I have given <@"+m.User.ID+"> The role `"+in.RoleSys+"`")
					}
				}
    		}
  		}
	}
} // end of GuildMemberAdd















func GuildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	var in info
	vfile, err := ioutil.ReadFile("servers/"+ m.GuildID + "/main.json")
	if err == nil {
		json.Unmarshal(vfile, &in)
	}

if in.ByeMsg != "" && in.ByeMsg != "off" {
// fmt.Println(m.GuildID, m.User)
s.ChannelTyping(m.GuildID)
time.Sleep(1000 * time.Millisecond)
data := strings.Replace(in.ByeMsg, "{user}", m.User.Username, -1)
s.ChannelMessageSend(m.GuildID, data)
}

}











func onReady(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Type: --help cmdname")
}









/* disabled until i make this a toggle feature.
func GuildRoleUpdate(s *discordgo.Session, m *discordgo.GuildRoleUpdate) {
s.ChannelTyping(m.GuildID)
time.Sleep(1000 * time.Millisecond)
s.ChannelMessageSend(m.GuildID, "Someone has edited the role: `"+m.Role.Name+"`")
}
*/









func GuildCreate(s *discordgo.Session, m *discordgo.GuildCreate) {

	if m.Guild.Unavailable != nil {
		return
	}

fmt.Println("Joined: " + m.Name + "\nOwner: " + m.OwnerID)
		newjs := info{
			Prefix:		"--",
			GreetMsg:	"",
			RoleSys:	"",
			ByeMsg:		"",
			Owner:		m.OwnerID,
			Name:		m.Name,
			AntiLink:	false,
			Action:		"kick",
			Silent:		false,
		}


newcmd := commands{
	Greet: 			"greet",
	Bye:			"bye",
	Prefix:			"prefix",
	Kick:			"kick",
	Ban:			"ban",
	Autorole:		"autorole",
	SetPunish:		"setpunish",
	AllowLinks:		"allowlinks",
	DenyLinks:		"denylinks",
	AddMaster:		"addmaster",
	DelMaster:		"delmaster",
	Invites:		"invites",
	Meme:			"meme",
	Joke:			"joke",
	Give:			"give",
	Take:			"take",
	Giveme:			"giveme",
	Mute:			"mute",
	Unmute:			"unmute",
}

var p string

p = "CarbonKeyNoNo"
servercount,_ := ioutil.ReadDir("servers")
// fmt.Println(len(servercount))
pi := len(servercount)
bi := strconv.Itoa(pi)
http.PostForm("https://www.carbonitex.net/discord/data/botdata.php", url.Values{"key": {p}, "servercount": {bi}})
// <-time.After(5 * time.Minute)

		k, err := json.Marshal(newcmd)
		if err != nil {
			return
		}
		b, err := json.Marshal(newjs)
		if err == nil {
			os.Mkdir("servers/" + m.ID, 0777)
			ioutil.WriteFile("servers/" + m.ID + "/main.json", b, 0777)
			ioutil.WriteFile("servers/" + m.ID + "/commands.json", k, 0777)
			// ioutil.WriteFile("servers/" + m.ID + "/masters/" + m.OwnerID, "", 0777)
		}

}





func GuildDelete(s *discordgo.Session, m *discordgo.GuildDelete) {
fmt.Println("Kicked from: " + m.Name)
os.Remove("servers/" + m.ID + "/main.json")
os.Remove("servers/" + m.ID + "/commands.json")
os.Remove("servers/" + m.ID)
fmt.Println("Removed server data.")

}
