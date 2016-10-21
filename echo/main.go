// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the New() helper function.
package main

import (
	"fmt"
	"time"
	"github.com/bwmarrin/discordgo"
//	"github.com/dustin/go-humanize"
	"flag"
	"os"
	"strconv"
 	"gotools"
 	"regexp"
 	"strings"
  	"net/url"
  	"net/http"
)


	var err error
	var startTime time.Time
	var js obj
	var cmd commands
	var limit bool
	var limitcnt int
	var start int64
	var stats statistics
	var buffer = make([][]byte, 0)
	var ConvertTime string
	var isBeep bool
	var playing map[string]string
	var mbuffer = make([][]byte, 0)
	var bot map[string]*Info = map[string]*Info{}
	var (
		dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
		spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	)





func main() {
	flag.Parse()

	start = time.Now().UnixNano()
	file, err := gto.ReadFile("config.json")
	gto.Unmarshal(file, &js)

    // Login to discord. You can use a token or email, password arguments.
	dg, err := discordgo.New(js.Bot)
	if err != nil {
		fmt.Println(err)
		return
	}
	ef, err := gto.ReadFile("info.json")
	if err == nil {
		gto.Unmarshal(ef, &stats)
	}


	db, err := gto.ReadFile("System/database.json")
	if err == nil {
		gto.Unmarshal(db, &bot)
	}



	rd, err := gto.ReadFile("playing.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	gto.Unmarshal(rd, &playing)




	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(GuildMemberAdd)
	dg.AddHandler(GuildMemberRemove)
	dg.AddHandler(onReady)
	// dg.AddHandler(GuildRoleUpdate)
	dg.AddHandler(GuildCreate)
	dg.AddHandler(GuildDelete)
	dg.AddHandler(GuildUpdate)
	dg.AddHandler(RateLimit)
	// dg.AddHandler(MessageDelete)
/*	dg.AddHandler(GuildRoleCreate)
	dg.AddHandler(GuildRoleUpdate)
*/	// dg.AddHandler(GuildRoleCreate)
	// Open the websocket and begin listening.
	dg.Open()

go task(dg)
go InfoCheck(dg)
go TwitterUpdate()
go TwitterHelp()

	// Simple way to keep program running until any key press.
	var input string
	fmt.Scanln(&input)
	return
}










/*
func MessageDelete(s *discordgo.Session, d *discordgo.MessageDelete) {

}
*/

var lastSig string
func CheckLog(s *discordgo.Session, m *discordgo.MessageCreate, in info) {
	var thisSig string
	if strings.HasPrefix(m.Content, in.Prefix) == false {
		return
	}
	if m.Author.ID == s.State.User.ID {
		return
	}
	guildID, err := gto.ServerID(s, m)
	if err != nil {
		return // couldn't find the server.
	}
	if in.Log != "" {
		_, err := s.State.Channel(in.Log)
		if err != nil {
			s.ChannelMessageSend(guildID, "The channel you have set for my command logging no longer exists.")
		}
		parse := m.Content

		thecmd := ""
		// let's parse the prefix and command out of the message
		if strings.Contains(parse, " ") {
			thecmd = gto.Split(parse, " ")[0]
		} else {
			thecmd = in.Prefix + gto.Split(parse, in.Prefix)[1]
		}

		thisSig = m.Author.ID + guildID + m.ChannelID + thecmd

		// let's look for usernames and parse them.
		if strings.Contains(parse, "<@") {
			p1 := gto.Split(parse, "<@")
			p2 := gto.Split(p1[1], ">")
			data := p2[0]
			if data != "" {
				usr, err := s.User(data)
				if err == nil {
					parse = strings.Replace(parse, "<@"+p2[0]+">", usr.Username, -1)
				}
			}
		}

		// le's look for channels
		if strings.Contains(parse, "<#") {
			p1 := gto.Split(parse, "<#")
			p2 := gto.Split(p1[1], ">")
			data := p2[0]
			if data != "" {
				usr, err := s.State.Channel(data)
				if err == nil {
					parse = strings.Replace(parse, "<#"+p2[0]+">", usr.Name, -1)
				}
			}
		}
		var conf map[string]interface{}
		rd1, err := gto.ReadFile("servers/"+guildID+"/options.json")
		if err != nil {
			return
		}
		gto.Unmarshal(rd1, &conf)
		if _, ok := conf["Log"]; ok {

		} else {
			return
		}

	// <-time.After(1000 * time.Millisecond)
		if lastSig != thisSig {
			s.ChannelMessageSend(conf["Log"].(string), "**"+m.Author.Username+"** Has triggered a Command: `"+parse+"`")
			lastSig = m.Author.ID + guildID + m.ChannelID + thecmd
		}
	}
	return
}

// CheckLog(s, m, in)


type tracker struct {
	Messages	int
	Joins		int
	Leaves		int
}


func Reminder(s *discordgo.Session, m *discordgo.MessageCreate, text string, thetime time.Duration) {
	<-time.After(thetime)

TheChan := m.ChannelID

c, err := s.State.Channel(TheChan)
if err != nil {
	c, err = s.Channel(TheChan)
	// let's try to prevent any errors from happening. hahaha....
	if err != nil {
		return
	}
}

	text = strings.Replace(text, "{user}", "<@"+m.Author.ID+">", -1)
	text = strings.Replace(text, "{everyone}", "@everyone", -1)


	if strings.Contains(text, "{user:") {
		d1 := gto.Split(text, "{user:")[1]
		d2 := gto.Split(d1, "}")[0]

		usr, err := s.User(d2)
		if err != nil {
			return
		}

		text = strings.Replace(text, "{user:"+d2+"}", "<@"+usr.ID+">", -1)
	}


	if strings.Contains(text, "{joke}") {
   		var joke []string
    	cnt := 0
    	cnt = gto.CountLines("Random.txt")
    	myrand := gto.Random(1, cnt)
    	joke, err := gto.ReadLines("Random.txt")
    	if err == nil {
    	// s.ChannelTyping(m.ChannelID)
    	// s.ChannelMessageSend(m.ChannelID, meme[myrand])
      	text = strings.Replace(text, "{joke}", joke[myrand], -1)
		}
	}

	if strings.Contains(text, "{redirect:") {
		dowork := false
		t1 := gto.Split(text, "{redirect:")[1]
		TheChan = gto.Split(t1, "}")[0]
		r, err := s.State.Guild(c.GuildID)
		if err != nil {
			return
		}
		for _, v := range r.Channels {
			if v.ID == TheChan {
				dowork = true
			}
		}
		if dowork == false {
			s.ChannelMessageSend(m.ChannelID, "This channel doesn't exist on this server.")
			return
		}

		text = strings.Replace(text, "{redirect:"+TheChan+"}", "", -1)
	}

	if strings.Contains(text, "{role:") {
		d1 := gto.Split(text, "{role:")[1]
		d2 := gto.Split(d1, "}")[0]
		therole := gto.GetRoleID(s, c.GuildID, d2)
		text = strings.Replace(text, "{role:"+d2+"}", "<@&"+therole+">", -1)
	}



	if strings.Contains(text, "{pm}") {
		text = strings.Replace(text, "{pm}", "", -1)
		k, err := s.UserChannelCreate(m.Author.ID)
		if err == nil {
			s.ChannelMessageSend(k.ID, text)
		}
	} else {
		s.ChannelMessageSend(TheChan, "**|** " + text)
	}
}





// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated user has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	chkErr := true
	sudo := false
	AllowSounds := true
	AllowReminder := true
	WordFilter := false
	WordAction := ""
	LimitRate := false
	UseRole := false
	Donator := false
	BotMaster := false
	BotCommander := ""

	var mcount map[string]int

	mc, err := gto.ReadFile("messages.json")
	if err == nil {
		gto.Unmarshal(mc, &mcount)
		if _, ok := mcount["Count"]; ok {
			mcount["Count"]++
			df, err := gto.Marshal(mcount)
			if err == nil {
				gto.WriteFile("messages.json", df, 0777)
			}
		}
	}



c, err := s.State.Channel(m.ChannelID)
if err != nil {
	c, err = s.Channel(m.ChannelID)
	// let's try to prevent any errors from happening. hahaha....
	if err != nil {
		chkErr = false // couldn't determine the Guild ID
	}
}

	if limit == true {
		limitcnt = 0
		limit = false
	//	return
	}



// ####### FUNCTIONS: Located in the funcs.go file #######
	var in info



if chkErr == true && c.GuildID != "" {
isNSFW := false


	// Load up the server information
	vfile, err := gto.ReadFile("servers/" + c.GuildID + "/main.json")
	if err != nil {
		gu, err := s.Guild(c.GuildID)
		if err == nil {
			go CreateServerFiles(s, gu.OwnerID, gu.ID, gu.Name)
		}
	} else {
	gto.Unmarshal(vfile, &in)
	}





	if _, err := os.Stat("servers/"+c.GuildID+"/options.json"); err != nil {
		cfi, err := gto.ReadFile("templates/options.json")
		if err == nil {
			gto.WriteFile("servers/"+c.GuildID+"/options.json", cfi, 0777)
		}
	}


	if _, err := os.Stat("servers/"+c.GuildID+"/tracker.json"); err != nil {
		cfi, err := gto.ReadFile("templates/options.json")
		if err == nil {
			gto.WriteFile("servers/"+c.GuildID+"/tracker.json", cfi, 0777)
		}
	}



	if _, err := os.Stat("servers/"+c.GuildID+"/namefilter.json"); err != nil {
		cfi, err := gto.ReadFile("templates/namefilter.json")
		if err == nil {
			gto.WriteFile("servers/"+c.GuildID+"/namefilter.json", cfi, 0777)
		}
	}


	if _, err := os.Stat("servers/"+c.GuildID+"/filter.json"); err != nil {
		cfi, err := gto.ReadFile("templates/namefilter.json")
		if err == nil {
			gto.WriteFile("servers/"+c.GuildID+"/filter.json", cfi, 0777)
		}
	}




	// check to see if the server is in the collection.
	// if not than add it.
	if _, ok := bot[c.GuildID]; ok {
		Task(s, m)
	} else {
		Register(s, m)
	}








	var tr map[string]int
	rtr, err := gto.ReadFile("servers/"+c.GuildID+"/tracker.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rtr, &tr)

	if _, ok := tr["Messages"]; ok  == false {
	newjs := tracker{
		Messages:	0,
		Joins:		0,
		Leaves:		0,
	}
	gg, err := gto.Marshal(newjs)
	if err != nil {
		return
	}
	gto.WriteFile("servers/"+c.GuildID+"/tracker.json", gg, 0777)
} else {
tr["Messages"]++
	gg, err := gto.Marshal(tr)
	if err != nil {
		return
	}
	gto.WriteFile("servers/"+c.GuildID+"/tracker.json", gg, 0777)
}
tr = nil




	var conf map[string]interface{}
		crd1, err := gto.ReadFile("servers/"+c.GuildID+"/options.json")
		if err != nil {
			return
		}
		gto.Unmarshal(crd1, &conf)


		if _, ok := conf["LogType"]; ok {
			in.LogType = conf["LogType"].(string)
		}

		if _, ok := conf["Log"]; ok {
			in.Log = conf["Log"].(string)
		}


		if _, ok := conf["NoReminder"]; ok {
			AllowReminder = false
		}


		if _, ok := conf["NoSounds"]; ok {
			AllowSounds = false
		}

		// check to see if they have the word filter enabled.
		// if so than the key will exist if not it won't.
		if _, ok := conf["WordFilter"]; ok {
			WordFilter = true
			WordAction = conf["WordFilter"].(string)
		}

		// Check to see if they want to use the Commander system requiring a Role
		if _, ok := conf["UseRole"]; ok {
			UseRole = true
		}


		if _, ok := conf["Restrict"]; ok {
			chh := conf["Restrict"].(string)
			_, err := s.State.Channel(chh)
			if err == nil {
				if sudo == false {
					if m.ChannelID != chh {
						return // don't work for any other channel but the channel they want
					}
				}
			}
		}







	if _, ok := conf["UseRole"]; ok {
		if in.BotMaster == "" {
			BotCommander = GetRoleID(s, c.GuildID, "Bot Commander")	// GetRoleID returns the role id.
			BotMaster = gto.MemberHasRole(s, c.GuildID, m.Author.ID, "Bot Commander")	// gto.MemberHasRole returns true or false
		} else {
			BotCommander = GetRoleID(s, c.GuildID, in.BotMaster)	// GetRoleID returns the role id.
			BotMaster = gto.MemberHasRole(s, c.GuildID, m.Author.ID, in.BotMaster)	// gto.MemberHasRole returns true or false
		}
	}


	if _, ok := conf["UseRole"]; ok == false {
		if IsManager(s, c.GuildID, m.Author.ID) == true {
			BotMaster = true
		}
	}



	if m.Author.ID == in.Owner {
		BotMaster = true
	}
	if m.Author.ID == js.Admin {
		BotMaster = true
		sudo = true
	}


	// make my Administrators Bot Masters no matter what server they are in.
	var staff map[string]string
	sfile, err := gto.ReadFile("System/Staff.json")
	if err == nil {
		gto.Unmarshal(sfile, &staff)

		for st, _ := range staff {
			if m.Author.ID == st {
				BotMaster = true
				sudo = true
			}
		}
	}



	// Detect if someone is a donator
	var donators map[string]string
	dfile, err := gto.ReadFile("System/Donators.json")
	if err == nil {
		gto.Unmarshal(dfile, &donators)

		for dn, _ := range donators {
			if m.Author.ID == dn {
				Donator = true
			}
		}
	}



if strings.HasPrefix(m.Content, in.Prefix) {
	requests("commands", m, &js)
}


	var resp response
	fresp, err := gto.ReadFile("servers/"+c.GuildID+"/response.json")
	if err != nil {
	//	s.ChannelMessageSend(m.ChannelID, "Your response file has been corrupted or never existed.\ntype `--webmaster` if you don't have a password and than login at <http://echobot.tk> do an update of your response database.")
	//	fmt.Println(err)
	// return
	}
	gto.Unmarshal(fresp, &resp)



if _, err := os.Stat("servers/"+c.GuildID+"/response.json"); err != nil {
fresp, err := gto.ReadFile("templates/response.json")
if err != nil {
//	return
	fmt.Println(err)
}
gto.Unmarshal(fresp, &resp)
}





if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "botcommander") {
	if UseRole == false {
		s.ChannelMessageSend(m.ChannelID, "Hi! By default Echo makes anyone with the permissions `Manage Server` his master. So this command is not needed. If you want to revert back to using the Bot Commander role system type `"+in.Prefix+"userole true` and this command will be available to you.")
		return
	}

	str := strings.Replace(m.Content, in.Prefix + "botcommander ", "", -1)
	verify := GetRoleID(s, c.GuildID, str)

	if verify == "" {
		if resp.NoRole == "" {
			s.ChannelMessageSend(m.ChannelID, "I couldn't find the role `"+str+"` to replace the `Bot Commander` You need to type an existing role.")
		} else {
			resp.NoRole = strings.Replace(resp.NoRole, "{data}", str, -1)
			s.ChannelMessageSend(m.ChannelID, resp.NoRole)
		}
	} else {
		var newinf map[string]interface{}
		v, err := gto.ReadFile("servers/"+c.GuildID+"/main.json")
		if err != nil {
			return
		}

		gto.Unmarshal(v, &newinf)
		newinf["BotMaster"] = str

		jm, err := gto.Marshal(newinf)
		if err != nil {
			return
		}


		gto.WriteFile("servers/"+c.GuildID+"/main.json", jm, 0777)
		s.ChannelMessageSend(m.ChannelID, "I have changed my Master role to `"+str+"`")
	}
	go CheckLog(s, m, in)
}




		if m.Content == in.Prefix + "info" {
			var mcount map[string]int
			ki, err := gto.ReadFile("messages.json")
			if err != nil {
				return
			}
			gto.Unmarshal(ki, &mcount)


			var req map[string]int
			b, err := gto.ReadFile("info.json")
			if err != nil {
				return
			}
			gto.Unmarshal(b, &req)
			active := 0
			active10m := 0
			mentions := 0
		for _, v := range bot {
			if v.TimeStamp >= time.Now().Unix() - (60 * 60 * 5) {
				active++
			}
			// track members
			for _, ac := range v.Users {
				if ac.LastStamp >= time.Now().Unix() - (60 * 60 * 5) {
					active10m++
				}
			}

			for _ = range v.Mentions {
				mentions++
			}
		}

			t := time.Unix(0, start)
			elapsed := time.Since(t)

		//	fmt.Printf("Elapsed time: %.2f hours\n", elapsed.Hours())
			upt := fmt.Sprintf("%s", elapsed)
			upti := gto.Split(upt, ".")
			uptime := upti[0]
			ts := js.CmdsRun
			i := strconv.Itoa(ts)
			servercount := stats.Servercount
			// fmt.Println(len(servercount))
			bi := strconv.Itoa(servercount)
			channelcount := strconv.Itoa(stats.Channelcount)
			membercount := strconv.Itoa(stats.Membercount)
			rolecount := strconv.Itoa(stats.Rolecount)
			msgcount := strconv.Itoa(mcount["Count"])
			emo := strconv.Itoa(req["Emojis"])
			ars := strconv.Itoa(req["ARS"])
			change, err := gto.ReadLines("System/information.txt")
			acs := strconv.Itoa(active)
			acm := strconv.Itoa(active10m)
			men := strconv.Itoa(mentions)
			newdat := ""
			if err == nil {
			//	top := 135
				for _, v := range change {
					str := strings.Replace(v, "{uptime}", uptime, -1)
					str = strings.Replace(str, "{servers}", bi, -1)
					str = strings.Replace(str, "{activeservers}", acs, -1)
					str = strings.Replace(str, "{activemembers}", acm, -1)
					str = strings.Replace(str, "{mentions}", men, -1)
					str = strings.Replace(str, "{cmdrequests}", i, -1)
					str = strings.Replace(str, "{emojirequests}", emo, -1)
					str = strings.Replace(str, "{arsrequests}", ars, -1)
					str = strings.Replace(str, "{channelcount}", channelcount, -1)
					str = strings.Replace(str, "{membercount}", membercount, -1)
					str = strings.Replace(str, "{rolecount}", rolecount, -1)
					str = strings.Replace(str, "{msgcount}", msgcount, -1)
					newdat = newdat + str + "\n"
				}
				s.ChannelTyping(m.ChannelID)
				s.ChannelMessageSend(m.ChannelID, newdat)
			}

		//	s.ChannelMessageSend(m.ChannelID, "```ruby\nEcho [v1.6.3]\nlibrary: DiscordGo\nuptime: "+uptime+"s\nservers: "+bi+"\n----------[ REQUESTS ]----------\ncommands: "+i+"\nemojis: "+emo+"\nars: "+ars+"```")
		if in.LogType == "master" {
			return
		}
			if in.LogType == "all" {
				go CheckLog(s, m, in)
			}
		}





		// have Echo grab his update and move it in the folder.
		if strings.HasPrefix(m.Content, "::moveupd") {
			if m.Author.ID != js.Admin {
				return
			}
			err = os.Rename("C:/Users/Igloo/Work/bin/test.exe", "new.exe")
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "An error occured during collecting the update.")
			} else {
				s.ChannelMessageSend(m.ChannelID, "`Completed` preparing the system for new updates.")
			}
		}




		if strings.HasPrefix(m.Content, "::cancel") {
			if m.Author.ID != js.Admin {
				return
			}
			err = os.Remove("new.exe")
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "`Error:` Update doesn't exist.")
			} else {
				s.ChannelMessageSend(m.ChannelID, "`Completed:` Update process has been cancelled.")
			}
		}









		if strings.HasPrefix(m.Content, "::grab ") {
			if m.Author.ID != js.Admin {
				return
			}
			s.ChannelMessageSend(m.ChannelID, "Searching for data..please wait")
			user := strings.Replace(m.Content, "::grab ", "", -1)
			data := ""
			g, err := s.UserGuilds()
			if err == nil {
				for _, v := range g {
					r, err := s.State.Guild(v.ID)
					if err == nil {
						if r.OwnerID == user {
							data = data + r.ID + "\n"
						}
					}
				}
				if data != "" {
					s.ChannelMessageSend(m.ChannelID, "```ruby\n"+data+"```")
				} else {
					s.ChannelMessageSend(m.ChannelID, "This user doesn't own any servers i'm on.")
				}
			}
		}




	if strings.HasPrefix(m.Content, in.Prefix + "help") {
		if m.Content == in.Prefix + "help" {
			ts := js.CmdsRun
			i := strconv.Itoa(ts)
			bi := strconv.Itoa(stats.Servercount)
			channelcount := strconv.Itoa(stats.Channelcount)
			membercount := strconv.Itoa(stats.Membercount)
			rolecount := strconv.Itoa(stats.Rolecount)
			help, err := gto.ReadLines("System/help.txt")
			newdat2 := ""
			if err == nil {
				for _, v := range help {
					str2 := strings.Replace(v, "{pref}", in.Prefix, -1)
					str2 = strings.Replace(str2, "{servers}", bi, -1)
					str2 = strings.Replace(str2, "{cmdrequests}", i, -1)
					str2 = strings.Replace(str2, "{channelcount}", channelcount, -1)
					str2 = strings.Replace(str2, "{membercount}", membercount, -1)
					str2 = strings.Replace(str2, "{rolecount}", rolecount, -1)
					newdat2 = newdat2 + str2 + "\n"
					}
					s.ChannelTyping(m.ChannelID)
					s.ChannelMessageSend(m.ChannelID, newdat2)
				}


			k, err := s.UserChannelCreate(m.Author.ID)
		if err == nil {
			s.ChannelTyping(k.ID)
			s.ChannelMessageSend(k.ID, "Commands List: <http://echobot.tk/?nav=commands>\nNeed help with A.R.S? <http://echobot.tk/?nav=arskeys>\nEcho's Wikia Page: <http://echo-bot.wikia.com/wiki/Echo_Bot_Wikia>\nCheck out the A.R.S Builder: <http://echobot.tk/?nav=home#builder>\n if you need more help feel free to ask in my server: <https://discord.gg/012s8wmCkDVgdn7yo>")
			// s.ChannelMessageSend(k.ID, "Disable all Image Embed commands in a channel: `"+in.Prefix+"noimages` to enable: `"+in.Prefix+"images`\n```ruby\nLEAGUE BANNERS:\n"+in.Prefix+"teemo\n"+in.Prefix+"vayne\n"+in.Prefix+"ekko\n"+in.Prefix+"zed\n"+in.Prefix+"anivia\n"+in.Prefix+"alistar\n"+in.Prefix+"amumu\n"+in.Prefix+"akali\n"+in.Prefix+"ahri\n"+in.Prefix+"aatrox\nFAMILY FRIENDLY:\n"+in.Prefix+"dbz\n"+in.Prefix+"cute\n"+in.Prefix+"cars\n"+in.Prefix+"trucks\n"+in.Prefix+"sky\n"+in.Prefix+"space\n"+in.Prefix+"wrecks\n:listemojis: for awesome emojis\n-------\nSOUNDS: "+in.Prefix+"911\n"+in.Prefix+"belch\n"+in.Prefix+"creepy\n"+in.Prefix+"drumroll\n"+in.Prefix+"fart\n"+in.Prefix+"fart2\n"+in.Prefix+"firelazer\n"+in.Prefix+"goofy\n"+in.Prefix+"ifarted\n"+in.Prefix+"laugh\n"+in.Prefix+"laugh2\n"+in.Prefix+"pewpew\n"+in.Prefix+"pig\n"+in.Prefix+"quack\n"+in.Prefix+"roar\n"+in.Prefix+"snore\n"+in.Prefix+"superman\n"+in.Prefix+"wetfart\n"+in.Prefix+"wookie\n"+in.Prefix+"yourmom```NSFW Commands Moved to `The #1 Porn Bot on Discord`: <http://bit.ly/25aCO2K>")
		}
	}

		if m.Content == in.Prefix + "help addmaster" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+"addmaster @User\ninfo: gives user acces to mod commands.\nEcho will try to make the Bot Commander role.\nIf non exists. make sure he keeps his original role!```")
		}

		if m.Content == in.Prefix + "help delmaster" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner\nusage: "+in.Prefix+"delmaster @User\ninfo: removes access to mod commands.```")
		}

		if m.Content == in.Prefix + "help greet" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"greet Welcome {user} if you need any help ask!\ninfo: use {user} to mention the new member.\nuse {/user} to just say their username!\nuse {pm} to message the user the greet!\njust set the message to off for turning the greet message off```")
		}

		if m.Content == in.Prefix + "help bye" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"bye {user} has left the server.\nuse {pm} to message them your byemsg\nto turn off set the bye message to off```")
		}

		if m.Content == in.Prefix + "help denylinks" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"denylinks\ninfo: enables my anti link module. by default i kick offenders. you can use setpunish command to change to ban.```")
		}

		if m.Content == in.Prefix + "help allowlinks" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"allowlinks\ninfo: turns my anti link module off.```")
		}

		if m.Content == in.Prefix + "help prefix" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"prefix #\ninfo: sets my prefix in your server to #```")
		}

		if m.Content == in.Prefix + "help autorole" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"autorole Role Name\ninfo: automatically assign a role to new members.\nsilently add roles type "+in.Prefix+"autorole -s Role Name```")
		}

		if m.Content == in.Prefix + "help invites" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Everyone\nusage: "+in.Prefix+"invites\ninfo: gives you a list of available invite codes for your channel.```")
		}

		if m.Content == in.Prefix + "help mkinvite" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"mkinvite\ninfo: creates a permenant invite code for your channel.```")
		}

		if m.Content == in.Prefix + "help give" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"give @User Role Name\ninfo: gives the user the specified role.```")
		}

		if m.Content == in.Prefix + "help take" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"take @User Role Name\ninfo: takes the user the specified role.```")
		}

		if m.Content == in.Prefix + "help giveme" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Custom\nusage: A.R.S Command\ninfo: You need to build this command in your server.\nif the A.R.S system is enabled on your server do this:\n"+in.Prefix+"auto &"+in.Prefix+"giveme {params}={role:{params}}{req:Owner}You've assumed the role **{params}**\nChange the {req:Owner} to whatever your role is.```")
		}

		if m.Content == in.Prefix + "help setpunish" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"setpunish kick/ban/warn\ninfo: you can set the anti links module punishment to either kick/ban or warn. by default its set to kick.```")
		}

		if m.Content == in.Prefix + "help mute" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"mute @User\ninfo: mutes the user. you need to make a role named muted and set the permissions to not speak and than add the role to your channels.```")
		}

		if m.Content == in.Prefix + "help unmute" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"unmute @User\ninfo: unmutes the user.```")
		}

		if m.Content == in.Prefix + "help rolecolor" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"rolecolor #000000 Role Name\ninfo: Changes the color to a role.```")
		}

		if m.Content == in.Prefix + "help auto" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"auto trigger=response\ninfo: Adds an auto response to Echo. if someone says the trigger he will reply with your response!\nhere is some keys:\n{user} will mention the user who triggered the response\n{pm}will pm the response instead of in your server.\n{kick} kicks the user that said the trigger.\n{ban} bans the user who said the trigger.\n{exc:Role Name,Role Name} excludes a role or multiple roles from a trigger.\n{role:Role Name} gives a user a role when they trigger your response.\n{alert:YOURID,ANOTHERID} pms one or more ppl when triggered!\nto get someones id type "+in.Prefix+"getid @User```View my github for all of the keys!")
		}
		if m.Content == in.Prefix + "help delauto" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"delauto trigger\ninfo: Deletes an auto response from Echo's A.R.S file.```")
		}
		if m.Content == in.Prefix + "help viewauto" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"viewauto\ninfo: Views all auto responses from echo's A.R.S file.```")
		}
		if m.Content == in.Prefix + "help botrole" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"botrole Role Name\ninfo: This is autorole but for bot accounts.\nif you put all your bots in a certain role\necho will auto add them if set.```")
		}
		if m.Content == in.Prefix + "help wipeauto" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"wipeauto\ninfo: deletes your auto response system file. you will need to type --auto to initiate the system after wiping your file.```")
		}
		if m.Content == in.Prefix + "help mkchan" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"mkchan channel-name text/voice\ninfo: creates a text or voice channel in your server.```")
		}
		if m.Content == in.Prefix + "help locateip" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Everyone\nusage: "+in.Prefix+"locateip ip/domain\ninfo: echo will display geo location information from the ip or domain name you've specified.```")
		}
		if m.Content == in.Prefix + "help getid" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Everyone\nusage: "+in.Prefix+"getid @User\ninfo: gets the users discord id.```")
		}
		if m.Content == in.Prefix + "help meme" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Custom\nusage: A.R.S Command\ninfo: You need to build this command.\nIf A.R.S is enabled on your server do this:\n"+in.Prefix+"auto "+in.Prefix+"meme={meme}```")
		}
		if m.Content == in.Prefix + "help cats" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Custom\nusage: A.R.S Command\ninfo: You need to build this command.\nIf A.R.S is enabled on your server do this:\n"+in.Prefix+"auto "+in.Prefix+"cats={cats}```")
		}
		if m.Content == in.Prefix + "help joke" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Custom\nusage: A.R.S Command\ninfo: You need to build this command.\nIf A.R.S is enabled on your server do this:\n"+in.Prefix+"auto "+in.Prefix+"joke={joke}```")
		}
		if m.Content == in.Prefix + "help nsfw" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"nsfw try/false\ninfo: enabled or disabled NSFW commands in a channel.```")
		}
		if m.Content == in.Prefix + "help grabars" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"grabars\ninfo: grabs your A.R.S file and sends it in pm.```")
		}
		if m.Content == in.Prefix + "help putars" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"putars directlink\ninfo: updates your manually edited json file to echo. you must link directly to the .json file!```")
		}
		if m.Content == in.Prefix + "help addrole" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"addrole Role Name\ninfo: creates a role with a specific name you choose!```")
		}
		if m.Content == in.Prefix + "help delrole" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"delrole Role Name\ninfo: creates a role with a specific name you choose!```")
		}
		if m.Content == in.Prefix + "help setwarning" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"setwarning 4\ninfo: sets the warning for the --warn @user command.\nby default echo kicks! you can change by typing\n--setpunish kick/warn/ban\nIMPORTANT: this changes your antilink punishment as well!```")
		}
		if m.Content == in.Prefix + "help warn" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"warn @Username\ninfo: adds a warning point to a user. and will kick/ban\nhe kicks by default. you can change this by\n--setpunish kick/ban/warn```")
		}
		if m.Content == in.Prefix + "help listwarns" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"listwarns\ninfo: lists everyone who has been warned.```")
		}
		if m.Content == in.Prefix + "help delwarn" {
			s.ChannelMessageSend(m.ChannelID, "```ruby\npermissions: Server Owner & Bot Commanders\nusage: "+in.Prefix+"delwarn @User\ninfo: resets someones warnings.```")
		}
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
return
		}
	}



GrabCommands(s,
			m,
			c.GuildID,
			BotMaster,
			sudo,
			UseRole,
			in.LogType,
			js,
			in,
			AllowReminder,
			AllowSounds,
			resp,
			BotCommander,
			Donator)



if len(m.Attachments) > 0 {
	for _, mb := range m.Attachments {
		if mb.Filename == "autoresponse.json" {
			if BotMaster == true {
				if gto.DownloadFile("servers/"+c.GuildID+"/autoresponse.json", mb.URL) == true {
					s.ChannelMessageSend(m.ChannelID, "I have detected autoresponse.json downloading file completed.")
				}
			}
		}
		if strings.HasPrefix(mb.Filename, "emoji_") {
			str := strings.TrimPrefix(mb.Filename, "emoji_")
			if strings.HasSuffix(mb.Filename, ".png") == false {
				s.ChannelMessageSend(m.ChannelID, "You can only post transparent images for emojis. `.png` format is required.")
				return
			}
			if mb.Width > 50 {
				s.ChannelMessageSend(m.ChannelID, "The emoji can't exceed `50x50` in size.")
				return
			}
			if mb.Height > 50 {
				s.ChannelMessageSend(m.ChannelID, "The emoji can't exceed `50x50` in size.")
				return
			}
			if _, err := os.Stat("images/submissions/"+m.Author.ID); err != nil {
				os.Mkdir("images/submissions/"+m.Author.ID, 0777)
			}
			str = strings.Replace(str, ":", "", -1)
			str = strings.Replace(str, ",", "", -1)
			str = strings.Replace(str, "!", "", -1)

			if gto.DownloadFile("images/submissions/"+m.Author.ID+"/"+str, mb.URL) == true {
				s.ChannelMessageSend(m.ChannelID, "`Success` Your emoji has been sent in. It could take up to 3 days to be accepted.")
			}
		}


		// Echo auto install system.
		if mb.Filename == "setup.esf" {
			fmt.Print("Detected setup.esf file")
			if BotMaster == false {
				s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this feature.")
				return
			}
			go EchoSetup(s, m, c.GuildID, m.ChannelID, mb.URL, in)
		}

		if mb.Filename == "setup.ars" {
			if BotMaster == false {
				s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this feature.")
				return
			}
			go SetupARS(s, m, c.GuildID, m.ChannelID, mb.URL)
		}


		if strings.HasSuffix(mb.Filename, ".plugin.ars") && strings.HasPrefix(mb.Filename, "install") == false {
			if m.Author.ID != "190255157647376384" {
				if BotMaster == false {
					s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this feature.")
					return
				}
			}
			go AutoScript(s, m, c.GuildID, m.ChannelID, mb.URL, BotMaster, mb.Filename, in)
		}


		if mb.Filename == "install.plugin.ars" {
			if m.Author.ID != "190255157647376384" {
				if BotMaster == false {
					s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this feature.")
					return
				}
			}
			go InstallScript(s, m, c.GuildID, m.ChannelID, mb.URL)
		}
	}
}





if strings.HasPrefix(m.Content, "::plugins") && BotMaster == true {
	ti, err := gto.ReadDir("servers/"+c.GuildID+"/scripts/", "*")
	if err == nil {
		data := ""
		for _, v := range ti {
			v = strings.Replace(v, "servers\\"+c.GuildID+"\\scripts\\", "", -1)
			v = strings.Replace(v, ".plugin.ars", "", -1)
			data = data + v + "\n"
		}
		if len(data) > 0 {
			s.ChannelMessageSend(m.ChannelID, "```xl\n"+data+"```")
		} else {
			s.ChannelMessageSend(m.ChannelID, "You don't have any plugins installed.")
		}
	}
}



if strings.HasPrefix(m.Content, "::run ") && BotMaster == true {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this feature.")
		return
	}
	plug := strings.Replace(m.Content, "::run ", "", -1)
	if _, err := os.Stat("servers/"+c.GuildID+"/scripts/"+plug+".plugin.ars"); err == nil {
		go RunScript(s, m, c.GuildID, m.ChannelID, "servers/"+c.GuildID+"/scripts/"+plug+".plugin.ars", BotMaster, in)
	}
}



if strings.HasPrefix(m.Content, "::uninstall ") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this feature.")
		return
	}
	plug := strings.Replace(m.Content, "::uninstall ", "", -1)
	err = os.Remove("servers/"+c.GuildID+"/scripts/"+plug+".plugin.ars")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "The plugin either doesn't exist or you need to check your spelling.")
	} else {
		s.ChannelMessageSend(m.ChannelID, "The plugin has been uninstalled.")
	}
}

// AutoScripts section.
/*
if _, err := os.Stat("servers/"+c.GuildID+"/scripts"); err == nil {
	// they have the scripts folder. now we need to list through
	// all the scripts and check if the trigger word has been detected.
	// if so than we need to do all the actions.

	scripts, err := gto.ReadDir("servers/"+c.GuildID+"/scripts/", "*")
	if err != nil {
		return
	}
	for _, file := range scripts {
		autoscript, err := gto.ReadLines(file)
		if err != nil {
			return
		}
		for _, fs := range autoscript {
			trig := ""
			action := ""
			ispm := false
			params := ""
			if strings.HasPrefix(fs, "trigger=") {
				trig = strings.Split(fs, "trigger=")[1]
			}

			if strings.Contains(m.Content, trig) {
				if strings.HasPrefix(fs, "action=") {
					action = strings.Split(fs, "action=")[1]
				}
				if strings.Contains(action, "{pm}") {
					ispm = true
				}
				if strings.Contains(trig, "{params}") {
			 		if strings.Contains(trig, "{params}") {
			 		//	params = strings.Replace(m.Content, trigger, "", -1)
			 			params = gto.TrimPrefix(fs, trig)
			 			trig = strings.Replace(trig, " {params}", "", -1)
			 		}
				}
				master := ""
				if BotMaster == true {
					master = "True"
				} else {
					master = "False"
				}
				newresp := action
				newresp = strings.Replace(newresp, "{user}", "<@"+m.Author.ID+">", -1)
        		newresp = strings.Replace(newresp, "{/user}", m.Author.Username, -1)
        		newresp = strings.Replace(newresp, "{chan}", "<#"+m.ChannelID+">", -1)
        		newresp = strings.Replace(newresp, "{pref}", in.Prefix, -1)
        		newresp = strings.Replace(newresp, "{greet}", in.GreetMsg, -1)
        		newresp = strings.Replace(newresp, "{bye}", in.ByeMsg, -1)
        		newresp = strings.Replace(newresp, "{ismaster}", master, -1)
				ARS(s, m, action, c.GuildID, ispm, BotMaster, params, trig)
			} // check to see if they said the trigger word.
		}
	}
}
*/


// Check to see if they want this channel limited to 5 messages per 5 seconds
if _, ok := bot[c.GuildID]; ok {
	for _, ch := range bot[c.GuildID].Restricted {
		if ch.ID == m.ChannelID {
			LimitRate = true
		}
	}
}


		if LimitRate == true && BotMaster == false {
			msgcnt := 0
			// limit := 15
			for _, v2 := range bot[c.GuildID].Messages {
				if time.Now().Unix() <= v2.TimeStamp + (3 * 2) {
					if m.Author.ID == v2.Author {
						msgcnt++
					}
				}

				if msgcnt >= 5 {
					if gto.MemberHasRole(s, c.GuildID, m.Author.ID, "muted") == false {
  						x, err := s.State.Member(c.GuildID, m.Author.ID)
            			if err != nil {
             			 	return
            			}
            			x.Roles = append(x.Roles, GetRoleID(s, c.GuildID, "muted"))
            			err = s.GuildMemberEdit(c.GuildID, m.Author.ID, x.Roles)
            			if err != nil {
            				return
            			} else {
            				s.ChannelMessageSend(m.ChannelID, "I have muted <@"+m.Author.ID+"> for flooding.")
            			}
					}
				}
			}
		}




if WordFilter == true && m.Author.ID != s.State.User.ID && strings.HasPrefix(m.Content, in.Prefix + "filter ") == false {
	var words map[string]string
	rr, err := gto.ReadFile("servers/"+c.GuildID+"/filter.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rr, &words)
	data := gto.Split(strings.ToLower(m.Content), " ")
	lm := 0
	for _, v := range data {
		if _, ok := words[v]; ok && lm == 0 {
			lm++
			// just warn the user
			if gto.ToLower(WordAction) == "warn" {
				err = s.ChannelMessageDelete(m.ChannelID, m.ID)
				if err == nil {
					s.ChannelMessageSend(m.ChannelID, "I have removed a message containing words not allowed in this server.")
				} else {
					s.ChannelMessageSend(m.ChannelID, "I don't have permissions to remove messages. Make sure i have the `Administrator` role.")
				}
			}
			// kick the user
			if gto.ToLower(WordAction) == "kick" {
				if BotMaster == false {
					err = s.GuildMemberDelete(c.GuildID, m.Author.ID)
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, "`Bad Word Detected:` I don't have permissions to kick someone in this server.")
						return
					}
					err = s.ChannelMessageDelete(m.ChannelID, m.ID)
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, "I don't have permissions to remove messages. Make sure i have the `Administrator` role.")
						return
					}
					s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> was kicked for saying a word not allowed on this server.")
				} else {
					err = s.ChannelMessageDelete(m.ChannelID, m.ID)
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, "I don't have permissions to remove messages. Make sure i have the `Administrator` role.")
						return
					}
					s.ChannelMessageSend(m.ChannelID, "I Won't kick a Bot Commander.. But I have removed your message for containing words not allowed on this server.")
				}
			}
			// ban the user
			if gto.ToLower(WordAction) == "ban" && BotMaster == false {
				if BotMaster == false {
					err = s.GuildBanCreate(c.GuildID, m.Author.ID, 10)
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, "`Bad Word Detected:` I don't have permissions to ban someone in this server.")
						return
					}
					err = s.ChannelMessageDelete(m.ChannelID, m.ID)
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, "I don't have permissions to remove messages. Make sure i have the `Administrator` role.")
						return
					}
					s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> was banned for saying a word not allowed on this server.")
				} else {
					err = s.ChannelMessageDelete(m.ChannelID, m.ID)
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, "I don't have permissions to remove messages. Make sure i have the `Administrator` role.")
						return
					}
					s.ChannelMessageSend(m.ChannelID, "I Won't ban a Bot Commander.. But I have removed your message for containing words not allowed on this server.")
				}
			}
		}
	}
}






Emojis(s, m)









cnt :=0
params := ""
if strings.HasPrefix(m.Content, in.Prefix+"auto") == false {
if m.Author.ID != s.State.User.ID && strings.HasPrefix(m.Content, in.Prefix + "delauto") == false {
	if _, err := os.Stat("servers/"+c.GuildID+"/autoresponse.json"); err == nil {
		var ars map[string]interface{}
		file, err := gto.ReadFile("servers/"+c.GuildID+"/autoresponse.json")
		if err == nil {
			master := ""
			if BotMaster == true {
				master = "True"
			} else {
				master = "False"
			}
			gto.Unmarshal(file, &ars)
		//	fmt.Println("===============")
			 for k, v := range ars {
			 	ispm := false
			 	isfind := false
			 	iscmd := false
			// 	ismod := false
			 	newresp := v.(string)
        newresp = strings.Replace(newresp, "{chan}", "<#"+m.ChannelID+">", -1)
        newresp = strings.Replace(newresp, "{pref}", in.Prefix, -1)
        newresp = strings.Replace(newresp, "{greet}", in.GreetMsg, -1)
        newresp = strings.Replace(newresp, "{bye}", in.ByeMsg, -1)
        newresp = strings.Replace(newresp, "{ismaster}", master, -1)


			 	// found the & key in the trigger
			if strings.Contains(k, "&") {
			 	trigger := strings.Replace(k, "&", "", -1)
			 	if strings.Contains(trigger, "{params}") {
			 	//	params = strings.Replace(m.Content, trigger, "", -1)
			 		params = gto.TrimPrefix(m.Content, trigger)
			 		trigger = strings.Replace(trigger, " {params}", "", -1)
			 	}







			 	if strings.HasPrefix(trigger, in.Prefix) {
			 		iscmd = true
			 		if strings.HasPrefix(gto.ToLower(m.Content), gto.ToLower(trigger)) {
			 			requests("ars", m, &js)
			 			if strings.Contains(newresp, "{pm}") {
			 				ispm = true
			 				newresp = strings.Replace(newresp, "{pm}", "", -1)
			 			}

			 			if strings.Contains(newresp, "{nsfw}") && isNSFW == false {
			 				s.ChannelMessageSend(m.ChannelID, "This channel isn't on the `NSFW` list. ask a Bot Commander to type `--nsfw true` in this channel.")
			 				return
			 			}
						cnt++
						if cnt <= 1 {
							if ispm == false {
								if BotMaster == false {
									ARS(s, m, newresp, c.GuildID, false, false, params, trigger, in)
								} else {
									ARS(s, m, newresp, c.GuildID, false, true, params, trigger, in)
								}
							} else { // ispm if statement
								if BotMaster == false {
			 						ARS(s, m, newresp, c.GuildID, true, false, params, trigger, in)
			 					} else {
			 						ARS(s, m, newresp, c.GuildID, true, true, params, trigger, in)
			 					}

							}
						}
					}
				}












			 	if strings.Contains(gto.ToLower(m.Content), gto.ToLower(trigger)) && iscmd == false && strings.HasPrefix(m.Content, in.Prefix) == false {
			 		requests("ars", m, &js)
			 		if strings.Contains(newresp, "{pm}") {
			 			ispm = true
			 			newresp = strings.Replace(newresp, "{pm}", "", -1)
			 		}
			 		if strings.Contains(newresp, "{nsfw}") && isNSFW == false {
			 			s.ChannelMessageSend(m.ChannelID, "This channel isn't on the `NSFW` list. ask a Bot Commander to type `--nsfw true` in this channel.")
			 			return
			 		}

					cnt++
					if cnt <= 1 {
						if ispm == false {
							if BotMaster == false {
							ARS(s, m, newresp, c.GuildID, false, false, params, trigger, in)
							} else {
							ARS(s, m, newresp, c.GuildID, false, true, params, trigger, in)
							}

						} else { // ispm if statement

							if BotMaster == false {
			 				ARS(s, m, newresp, c.GuildID, true, false, params, trigger, in)
			 				} else {
			 				ARS(s, m, newresp, c.GuildID, true, true, params, trigger, in)
			 				}

						}
					}
				}
			} // end of isfind block



				// basic trigger without find key.
			 	if gto.ToLower(m.Content) == gto.ToLower(k) && isfind == false {
			 		requests("ars", m, &js)
			 		if strings.Contains(newresp, "{pm}") {
			 			ispm = true
			 			newresp = strings.Replace(newresp, "{pm}", "", -1)
			 		}
			 		if strings.Contains(newresp, "{nsfw}") && isNSFW == false {
			 			s.ChannelMessageSend(m.ChannelID, "This channel isn't on the `NSFW` list. ask a Bot Commander to type `--nsfw true` in this channel.")
			 			return
			 		}

					cnt++
					if cnt <= 1 {
						if ispm == false {

							if BotMaster == false {
							ARS(s, m, newresp, c.GuildID, false, false, params, k, in)
							} else {
							ARS(s, m, newresp, c.GuildID, false, true, params, k, in)
							}

						} else {

							if BotMaster == false {
			 				ARS(s, m, newresp, c.GuildID, true, false, params, k, in)
			 				} else {
			 				ARS(s, m, newresp, c.GuildID, true, true, params, k, in)
			 				}
			 			}
					}
				}
			} // end of for loop
		}
	}
}
go CheckLog(s, m, in)
// return
}




if BotMaster == true && strings.HasPrefix(gto.ToLower(m.Content), in.Prefix + "auto ") {
	if _, err := os.Stat("servers/"+c.GuildID+"/autoresponse.json"); err != nil {
		var bm newj
		b, err := gto.Marshal(bm)
		if err == nil {
			gto.WriteFile("servers/"+c.GuildID+"/autoresponse.json", b, 0777)
		}
	}

	str := strings.Replace(m.Content, in.Prefix + "auto ", "", -1)
	tempdat := gto.Split(str, "=")
	trigger := tempdat[0]
	trigger = gto.TrimSuffix(trigger, " ")
	response := strings.Replace(str, trigger + "=", "", -1)
	var ars map[string]interface{}
	file, err := gto.ReadFile("servers/"+c.GuildID+"/autoresponse.json")
	if err == nil {

		if strings.Contains(str, "=") == true {
		gto.Unmarshal(file, &ars)
		if response != "" && trigger != "" {
			ars[trigger] = response
		}
		b, err := gto.Marshal(ars)
		if err == nil {
			response = strings.Replace(response, "`", "", -1)
			gto.WriteFile("servers/"+c.GuildID+"/autoresponse.json", b, 0777)
			s.ChannelTyping(m.ChannelID)

			s.ChannelMessageSend(m.ChannelID, "I've added `"+trigger+"` with the response:\n```ruby\n"+response+"```")
		}
		} else {
			s.ChannelTyping(m.ChannelID)

			s.ChannelMessageSend(m.ChannelID, "I'm sorry it seems you forgot to add `=` an example: ```ruby\nexample1: "+in.Prefix+"auto &help=Hey {user}!\nexample2: "+in.Prefix+"auto help=How can I help you {user}?```")
		}
	}
	go CheckLog(s, m, in)
return
}






if BotMaster == true && strings.HasPrefix(gto.ToLower(m.Content), in.Prefix+"wipeauto") {
if _, err := os.Stat("servers/" + c.GuildID + "/autoresponse.json"); err == nil {
os.Remove("servers/" + c.GuildID + "/autoresponse.json")
s.ChannelMessageSend(m.ChannelID, "I've wiped your A.R.S File.")
}
	/*
	var ars map[string]interface{}
	file, err := gto.ReadFile("servers/"+c.GuildID+"/autoresponse.json")
	if err == nil {
		gto.Unmarshal(file, ars)
		for k, _ := range ars {
			delete(ars, k)
		}
		b, err := gto.Marshal(ars)
		if err == nil {
			gto.WriteFile("servers/"+c.GuildID+"/autoresponse.json", b, 0777)
			s.ChannelMessageSend(m.ChannelID, "I've wiped your A.R.S File.")
		}
	}
*/
go CheckLog(s, m, in)
// return
}







if BotMaster == true && strings.HasPrefix(gto.ToLower(m.Content), in.Prefix + "delauto") {
	str := strings.Replace(m.Content, in.Prefix + "delauto ", "", -1)
	dcnt := 0
	var ars map[string]interface{}
	file, err := gto.ReadFile("servers/"+c.GuildID+"/autoresponse.json")
	if err == nil {
		gto.Unmarshal(file, &ars)
		for k, _ := range ars {
			if k == str {
				dcnt++
				delete(ars, k)
			}
		} // end of range ars
		b, err := gto.Marshal(ars)
		if err == nil && dcnt > 0 {
			gto.WriteFile("servers/"+c.GuildID+"/autoresponse.json", b, 0777)
			s.ChannelTyping(m.ChannelID)

			s.ChannelMessageSend(m.ChannelID, "I've deleted `"+str+"` from your A.R.S File")
		}
		if dcnt == 0 {
			s.ChannelTyping(m.ChannelID)

			s.ChannelMessageSend(m.ChannelID, "The trigger: `"+str+"` was not found in your A.R.S File.")
		}
	}
	go CheckLog(s, m, in)
	// return
}










if BotMaster == true && strings.HasPrefix(gto.ToLower(m.Content), in.Prefix + "viewauto") {
	var ars map[string]interface{}
	file, err := gto.ReadFile("servers/"+c.GuildID+"/autoresponse.json")
	if err == nil {
		gto.Unmarshal(file, &ars)
		newdata := "```ruby\n"
		newdata2 := "```ruby\n"
		newdata3 := "```ruby\n"

		for k, v := range ars {
			v = strings.Replace(v.(string), "`", "", -1)

			if len(newdata) < 1500 {
				newdata = newdata + "trigger: " + k + " response: " + v.(string) + "\n"
			} else {
				if len(newdata2) < 1500 {
					newdata2 = newdata2 + "trigger: " + k + " response: " + v.(string) + "\n"
				} else {
					newdata3 = newdata3 + "trigger: " + k + " response: " + v.(string) + "\n"
				}
			}
		}

		s.ChannelTyping(m.ChannelID)

		s.ChannelMessageSend(m.ChannelID, newdata + "```")

		if len(newdata2) > 15 {
			time.Sleep(2000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, newdata2 + "```")
		}

		if len(newdata3) > 15 {
			time.Sleep(2000 * time.Millisecond)
			s.ChannelMessageSend(m.ChannelID, newdata2 + "```")
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "You haven't created your ARS File yet. add triggers and responses by typing `--auto trigger=response`")
	} // end of err == nil
	go CheckLog(s, m, in)
// return
	} // end of viewauto



} // end of chkErr



} // ##########   END OF messageCreate











func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

TheName := m.User.Username
GreetChannel := m.GuildID
AutoRoleChannel := m.GuildID

	var namefilter map[string]string
	crd1, err := gto.ReadFile("servers/"+m.GuildID+"/namefilter.json")
	if err == nil {
		gto.Unmarshal(crd1, &namefilter)
		newnick := TheName
		for k, v := range namefilter {
			if strings.Contains(TheName, k) {
				// the user has a name that is not allowed.
				newnick = strings.Replace(newnick, k, v, -1)
			}
		}
		if newnick != m.User.Username {
			s.GuildMemberNickname(m.GuildID, m.User.ID, newnick)
			TheName = newnick
		}
	}



	var opts map[string]interface{}
	rf1, err := gto.ReadFile("servers/"+m.GuildID+"/options.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rf1, &opts)

	// check greet channel.
	if _, ok := opts["GreetChannel"]; ok {
		GreetChannel = opts["GreetChannel"].(string)
	}

	if _, ok := opts["AutoRoleChannel"]; ok {
		AutoRoleChannel = opts["AutoRoleChannel"].(string)
	}


	// AutoNick system
	if _, ok := opts["AutoNick"]; ok {
		data := opts["AutoNick"].(string)
		data = strings.Replace(data, "{/user}", m.User.Username, -1)

		if strings.Contains(data, "{random}") {
	    	var nickname []string
    		cnt := 0
    		cnt = gto.CountLines("System/nicknames.txt")
    		myrand := gto.Random(1, cnt)
    		nickname, err := gto.ReadLines("System/nicknames.txt")
    		if err == nil {
    			// s.ChannelTyping(m.ChannelID)
    			// s.ChannelMessageSend(m.ChannelID, meme[myrand])
      			data = strings.Replace(data, "{random}", nickname[myrand], -1)
    		}
		}

		if strings.Contains(data, "{joined}") {
        	joined, err := s.State.Member(m.GuildID, m.User.ID)
        	if err == nil {
            	newjo := strings.Split(joined.JoinedAt, "T")
            	thedate := newjo[0]
            	// thetime1 := gto.Split(newjo[1], ".")
            	// thetime := thetime1[0]
            	data = strings.Replace(data, "{joined}", thedate, -1)
        		// data = strings.Replace(data, "{joined}", nickname[myrand], -1)
        	}
		}

		if strings.Contains(data, "{special:rand}") {
  			chk := ""
        	reg, err := regexp.Compile("[^][.,_'0-9a-zA-Z -]")
        	if err != nil {
            	    return
        	}

        	chk = reg.ReplaceAllString(TheName, "")
        	if m.User.Username != chk {
	    		var nickname []string
    			cnt := 0
    			cnt = gto.CountLines("System/nicknames.txt")
    			myrand := gto.Random(1, cnt)
    			nickname, err := gto.ReadLines("System/nicknames.txt")
    			if err == nil {
    				// s.ChannelTyping(m.ChannelID)
    				// s.ChannelMessageSend(m.ChannelID, meme[myrand])
      				data = strings.Replace(data, "{special:rand}", nickname[myrand], -1)
      				s.ChannelMessageSend(m.GuildID, TheName + " has special characters and has been renamed to `"+nickname[myrand]+"`")
    			}
    		} else {
    			data = TheName
    		}
   		} // end of {special:rand} key


		if strings.Contains(data, "{special}") {
  			chk := ""
        	reg, err := regexp.Compile("[^][.,_'0-9a-zA-Z -]")
        	if err != nil {
            	    return
        	}
        	chk = reg.ReplaceAllString(TheName, "")
        	if m.User.Username != chk {
    			// s.ChannelTyping(m.ChannelID)
    			// s.ChannelMessageSend(m.ChannelID, meme[myrand])
    			if len(chk) > 1 {
      				data = strings.Replace(data, "{special}", chk, -1)
      				s.ChannelMessageSend(m.GuildID, TheName + " has special characters and I have removed them.")
      			} else {
	    			var nickname []string
    				cnt := 0
    				cnt = gto.CountLines("System/nicknames.txt")
    				myrand := gto.Random(1, cnt)
    				nickname, err := gto.ReadLines("System/nicknames.txt")
    				if err == nil {
    					// s.ChannelTyping(m.ChannelID)
    					// s.ChannelMessageSend(m.ChannelID, meme[myrand])
      					data = strings.Replace(data, "{special}", nickname[myrand], -1)
      					s.ChannelMessageSend(m.GuildID, TheName + " had only special characters in their name. `Giving random name`")
    				}
      			}
    		} else {
    			data = TheName
    		}
   		} // end of {special} key


   		if data != TheName {
			s.GuildMemberNickname(m.GuildID, m.User.ID, data)
		}
	} // end of AutoNick System












	// tracker system
	var tr map[string]int
	rtr, err := gto.ReadFile("servers/"+m.GuildID+"/tracker.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rtr, &tr)

	if _, ok := tr["Messages"]; ok  == false {
	newjs := tracker{
		Messages:	0,
		Joins:		0,
		Leaves:		0,
	}
	gg, err := gto.Marshal(newjs)
	if err != nil {
		return
	}
	gto.WriteFile("servers/"+m.GuildID+"/tracker.json", gg, 0777)
} else {
tr["Joins"]++
	gg, err := gto.Marshal(tr)
	if err != nil {
		return
	}
	gto.WriteFile("servers/"+m.GuildID+"/tracker.json", gg, 0777)
}
tr = nil








	// Greet System
	var in info
	vfile, err := gto.ReadFile("servers/" + m.GuildID + "/main.json")
	gto.Unmarshal(vfile, &in)

	// fmt.Println(in.RoleSys)
	roles, err := s.State.Guild(m.GuildID)

		if in.GreetMsg != "" && in.GreetMsg != "off" {
			data := strings.Replace(in.GreetMsg, "{user}", "<@"+m.User.ID+">", -1)
			data = strings.Replace(data, "{/user}", m.User.Username, -1)

			if strings.Contains(data, "{pm}") {
				data = strings.Replace(data, "{pm}", "", -1)
				pm, err := s.UserChannelCreate(m.User.ID)
				if err == nil {
					s.ChannelTyping(pm.ID)
					s.ChannelMessageSend(pm.ID, data)
				}
			} else {
				s.ChannelTyping(GreetChannel)
				s.ChannelMessageSend(GreetChannel, data)
			}
		}









// Autorole System
if in.RoleSys != "" && in.RoleSys != "off" && m.User.Bot == false {
 	for _, v := range roles.Roles {
    	if v.Name == in.RoleSys {
    		m.Roles = append(m.Roles, v.ID)
    		s.GuildMemberEdit(m.GuildID, m.User.ID, m.Roles)
    		if in.Silent == false {
				s.ChannelMessageSend(AutoRoleChannel, "I have given <@"+m.User.ID+"> the role `"+in.RoleSys+"`")
				return
			}
    	}
  	}
}





// check if they have BotAuto off and AutoRole is on. if so we need to give them the users role.
if in.BotAuto == "" || in.BotAuto == "off" && m.User.Bot == true {
	if in.RoleSys != "" && in.RoleSys != "off" {
 		for _, v := range roles.Roles {
    		if v.Name == in.RoleSys {
    			m.Roles = append(m.Roles, v.ID)
    			err = s.GuildMemberEdit(m.GuildID, m.User.ID, m.Roles)
    			if err != nil {
    				s.ChannelMessageSend(AutoRoleChannel, "`Auto Role Failed:` I don't have permissions to autorole. Visit <http://echo-bot.wikia.com/wiki/Echo_Bot_Wikia> For help.")
    				return
    			}
    			if in.Silent == false {
					s.ChannelMessageSend(AutoRoleChannel, "I have given <@"+m.User.ID+"> the role `"+in.RoleSys+"`")
					return
				}
    		}
  		}
	} // make sure they want the autorole system.
} // end of BotAuto disabled so autorole the bot to the members role.





// They have BotAuto set so give them the bot role they want.
if m.User.Bot == true && len(in.BotAuto) > 0 && in.BotAuto != "off" {
 	for _, v := range roles.Roles {
    	if v.Name == in.BotAuto {
    		m.Roles = append(m.Roles, v.ID)
    		s.GuildMemberEdit(m.GuildID, m.User.ID, m.Roles)
    		if in.Silent == false {
				s.ChannelMessageSend(AutoRoleChannel, "I have given <@"+m.User.ID+"> the role `"+in.BotAuto+"`")
				return
			}
    	}
  	}
} // if they have BotAuto Set than give the bots their role.


} // end of GuildMemberAdd













func GuildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

	var tr map[string]int
	rtr, err := gto.ReadFile("servers/"+m.GuildID+"/tracker.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rtr, &tr)

	if _, ok := tr["Messages"]; ok  == false {
	newjs := tracker{
		Messages:	0,
		Joins:		0,
		Leaves:		0,
	}
	gg, err := gto.Marshal(newjs)
	if err != nil {
		return
	}
	gto.WriteFile("servers/"+m.GuildID+"/tracker.json", gg, 0777)
} else {
tr["Leaves"]++
	gg, err := gto.Marshal(tr)
	if err != nil {
		return
	}
	gto.WriteFile("servers/"+m.GuildID+"/tracker.json", gg, 0777)
}
tr = nil


	var in info
	vfile, err := gto.ReadFile("servers/"+ m.GuildID + "/main.json")
	if err == nil {
		gto.Unmarshal(vfile, &in)
	}

	if in.ByeMsg != "" && in.ByeMsg != "off" {
		// fmt.Println(m.GuildID, m.User)
		data := strings.Replace(in.ByeMsg, "{user}", m.User.Username, -1)
		if strings.Contains(data, "{pm}") {
			data = strings.Replace(data, "{pm}", "", -1)
			pm, err := s.UserChannelCreate(m.User.ID)
			if err == nil {
				s.ChannelTyping(pm.ID)
				s.ChannelMessageSend(pm.ID, data)
			}
		} else {
			s.ChannelTyping(m.GuildID)
			s.ChannelMessageSend(m.GuildID, data)
		}
	}
}





func RateLimit(s *discordgo.Session, r *discordgo.RateLimit) {
limitcnt++
fmt.Println(r.Message)
	if limitcnt > 5 {
		limit = true
	}
}






func onReady(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Type: --help cmdname")
}









/* disabled until i make this a toggle feature.
func GuildRoleUpdate(s *discordgo.Session, m *discordgo.GuildRoleUpdate) {
s.ChannelTyping(m.GuildID)

s.ChannelMessageSend(m.GuildID, "Someone has edited the role: `"+m.Role.Name+"`")
}
*/









func GuildCreate(s *discordgo.Session, m *discordgo.GuildCreate) {
	if m.Guild.Unavailable != nil {
		return
	}
	ef, err := gto.ReadFile("info.json")
	if err == nil {
		gto.Unmarshal(ef, &stats)
	}


if _, err := os.Stat("servers/" + m.ID + "/main.json"); err != nil {
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
			BotAuto:	"",
			BotMaster:	"",
			AutoPerms:	false,
			Warnings:	0,
			Password:	"",
			Pulse:		0,
		}

		b, err := gto.Marshal(newjs)
		if err == nil {
			os.Mkdir("servers/" + m.ID, 0777)
			gto.WriteFile("servers/" + m.ID + "/main.json", b, 0777)

		}

var p string

p = ""
servercount := stats.Servercount
// fmt.Println(len(servercount))
bi := strconv.Itoa(servercount)
if isBeep == true {
	go JoinBeep()
}
go http.PostForm("https://www.carbonitex.net/discord/data/botdata.php", url.Values{"key": {p}, "servercount": {bi}})
// <-time.After(5 * time.Minute)

var inf obj
v, err := gto.ReadFile("config.json")
if err != nil {
	return
}
gto.Unmarshal(v, &inf)

	user, err := s.User(m.OwnerID)
		if err != nil {
		//	fmt.Println("Joined: " + m.Name + "\nOwner: " + m.OwnerID)
			return
		}
// fmt.Println("Joined: " + m.Name + "\nOwner: " + user.Username)

tweetit := user.Username + " Has invited Echo to " + m.Name
Tweet(tweetit, url.Values{"status": {tweetit}})

	g, err := s.State.Guild(m.ID)
	if err == nil {
		for _, v := range g.Members {
			if IsManager(s, m.ID, v.User.ID) == true {
        		k, err := s.UserChannelCreate(v.User.ID)
        		if err == nil {
        			s.ChannelMessageSend(k.ID, "Hi! Someone invited me to your server. You have `Manage Server` Permission which means you have full access to my commands. The default prefix on your server is `--` you can type `--prefix` to change your prefix to anything you want. You can view a list of commands at <http://echobot.tk> I have an Auto Response System with over 50 keys that allows you to make Echo your own! View my wikia for more help: <http://echo-bot.wikia.com/wiki/Echo_Bot_Wikia>")
        		}
        	}
		}
	}


	owner, err := s.UserChannelCreate(g.OwnerID)
	if err == nil {
        s.ChannelMessageSend(owner.ID, "Hi! Someone invited me to your server. You have `Manage Server` Permission which means you have full access to my commands. The default prefix on your server is `--` you can type `--prefix` to change your prefix to anything you want. You can view a list of commands at <http://echobot.tk> I have an Auto Response System with over 50 keys that allows you to make Echo your own! View my wikia for more help: <http://echo-bot.wikia.com/wiki/Echo_Bot_Wikia>")
	}

/*
	roles, err := s.State.Guild(g.ID)
	if err == nil {
		for _, v := range roles.Roles {
			if v.Name == "Echo [BETA]" {
				v.Position = 2
			}
		}
	}
*/

} // see if the folder exists or not.
}







func GuildDelete(s *discordgo.Session, m *discordgo.GuildDelete) {
	if isBeep == true {
		go LeaveBeep()
	}

	if _, ok := bot[m.ID]; ok {
		delete(bot, m.ID)
	}

	io, err := gto.ReadDir("servers/"+m.ID+"/", "*")
	if err == nil {
		for _, v := range io {
			os.Remove(v)
		}
	}
/*
// fmt.Println("Kicked from: " + m.Name)
os.Remove("servers/" + m.ID + "/main.json")
os.Remove("servers/" + m.ID + "/commands.json")
os.Remove("servers/" + m.ID + "/autoresponse.json")
os.Remove("servers/" + m.ID + "/nsfw.json")
os.Remove("servers/" + m.ID + "/response.json")
os.Remove("servers/" + m.ID + "/tracker.json")
os.Remove("servers/" + m.ID + "/track.json")
os.Remove("servers/" + m.ID + "/options.json")
*/
os.Remove("servers/" + m.ID)
// fmt.Println("Removed server data.")

}






func GuildUpdate(s *discordgo.Session, g *discordgo.GuildUpdate) {
	var config map[string]interface{}

	file, err := gto.ReadFile("servers/"+g.ID+"/main.json")
	if err != nil {
		return
	}
	gto.Unmarshal(file, &config)
	if g.OwnerID != config["Owner"] {
		config["Owner"] = g.OwnerID
		c, err := gto.Marshal(config)
		if err == nil {
			gto.WriteFile("servers/"+g.ID+"/main.json", c, 0777)
		}
	}
}
