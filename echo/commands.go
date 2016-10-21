package main 

import (
	"strings"
	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"gotools"
	"fmt"
	"strconv"
	"os"
	"net/url"
	"time"
	"bytes"
	"encoding/base64"
)

func GrabCommands(s *discordgo.Session, m *discordgo.MessageCreate, guildID string, BotMaster bool, sudo bool, UseRole bool, logType string, js obj, in info, AllowReminder bool, AllowSounds bool, resp response, BotCommander string, Donator bool) {



if strings.HasPrefix(m.Content, in.Prefix + "giveall ") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders have access to this feature.")
		return
	}
	role := strings.Replace(m.Content, in.Prefix + "giveall ", "", -1)
	go GiveThemAll(s, m, role, guildID, m.ChannelID)
}




if strings.HasPrefix(m.Content, in.Prefix + "takeall ") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders have access to this feature.")
		return
	}
	role := strings.Replace(m.Content, in.Prefix + "takeall ", "", -1)
	go TakeThemAll(s, m, role, guildID, m.ChannelID)
}







	if strings.HasPrefix(m.Content, in.Prefix + "stream") {
		str := strings.Replace(m.Content, in.Prefix + "stream ", "", -1)
		if str != "" && str != in.Prefix + "stream" {
    		var twi map[string]map[string]interface{}
    		gto.GetJson("https://api.twitch.tv/kraken/streams/"+str+"?client_id=j8k3gackqiwuasubw2htxd176dlbni8", &twi)

    		if twi["stream"]["game"] != nil {
    		game := twi["stream"]["game"].(string)
    		dd := fmt.Sprintf("%.2f", twi["stream"]["viewers"].(float64))
    		fps := fmt.Sprintf("%.2f", twi["stream"]["average_fps"].(float64))
    		viewers := gto.Split(dd, ".")
    		s.ChannelMessageSend(m.ChannelID, "```ruby\nuser: " + str + "\nplaying: " + game + "\nviewers: " + viewers[0] + "\nfps: " + fps + "```")
			} else {
				s.ChannelMessageSend(m.ChannelID, "The user is either offline or you typed the name wrong")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Type a twitch streamer username `"+in.Prefix+"stream imaqtpie`")
		}
	}




	if strings.HasPrefix(m.Content, in.Prefix + "tracker") {
		var track map[string]int
		gk, err := gto.ReadFile("servers/"+guildID+"/tracker.json")
		if err != nil {
			return
		}
		gto.Unmarshal(gk, &track)
		mc := strconv.Itoa(track["Messages"])
		jc := strconv.Itoa(track["Joins"])
		lc := strconv.Itoa(track["Leaves"])
		s.ChannelMessageSend(m.ChannelID, "```ruby\nmessages: " + mc + "\njoins: " + jc + "\nleaves: " + lc + "```")
	}


if gto.ToLower(m.Content) == "<@147463276840747008> what's your in.Prefix?" {
s.ChannelMessageSend(m.ChannelID, "My in.Prefix in your server is `"+in.Prefix+"`")
}

if gto.ToLower(m.Content) == "<@147463276840747008> whats your in.Prefix?" {
s.ChannelMessageSend(m.ChannelID, "My in.Prefix in your server is `"+in.Prefix+"`")
}

if gto.ToLower(m.Content) == "<@147463276840747008> whats your in.Prefix" {
s.ChannelMessageSend(m.ChannelID, "My in.Prefix in your server is `"+in.Prefix+"`")
}

if gto.ToLower(m.Content) == "<@147463276840747008> what's your in.Prefix" {
s.ChannelMessageSend(m.ChannelID, "My in.Prefix in your server is `"+in.Prefix+"`")
}

if gto.ToLower(m.Content) == "<@147463276840747008> what is your in.Prefix" {
s.ChannelMessageSend(m.ChannelID, "My in.Prefix in your server is `"+in.Prefix+"`")
}

if gto.ToLower(m.Content) == "<@147463276840747008> what is your in.Prefix?" {
s.ChannelMessageSend(m.ChannelID, "My in.Prefix in your server is `"+in.Prefix+"`")
}




if strings.HasPrefix(m.Content, in.Prefix + "prepare ") && js.Admin == m.Author.ID {
	str := strings.Replace(m.Content, in.Prefix + "prepare ", "", -1)
	go AlertUpdates(s, str)
}




	if strings.HasPrefix(m.Content, in.Prefix + "savedata") && m.Author.ID == js.Admin {
		io, err := gto.Marshal(bot)
		if err == nil {
			for k, _ := range bot[guildID].Messages {
				bot[guildID].Messages = append(bot[guildID].Messages[:k], bot[guildID].Messages[k+1:]...)
			}
			gto.WriteFile("System/database.json", io, 0777)
			s.ChannelMessageSend(m.ChannelID, "Database has been saved.")
		}
	}



	if strings.HasPrefix(m.Content, in.Prefix + "loaddata") && m.Author.ID == js.Admin {
		r1, err := gto.ReadFile("System/database.json")
		if err != nil {
			return
		}
		gto.Unmarshal(r1, &bot)
	}




	if m.Content == in.Prefix + "restart" {
		if m.Author.ID == js.Admin || sudo == true {
			s.UpdateStatus(0, "Restart4Updates!")
			gto.Print("#######################################################")
			gto.Print("## Restart has been initiated by someone (Alert)")
			gto.Print("## "+m.Author.Username+" Triggered a restart!")
			gto.Print("########################################################")
			os.Exit(1)
		}
	}






	if strings.HasPrefix(m.Content, in.Prefix + "mentions") {
		data := ""
		for _, v := range bot[guildID].Mentions {
			if v.Mentioned == m.Author.ID {
				data = data + v.ByUser + "["+v.ByID+"] Mentioned You\nMessage: " + v.Content + "\n----------\n"
			}
		}

		if data != "" {
			k, err := s.UserChannelCreate(m.Author.ID)
        	if err == nil {
				s.ChannelMessageSend(k.ID, "```ruby\n"+data+"```")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "No one has mentioned you in this server.")
		}
	}




	if strings.HasPrefix(m.Content, in.Prefix + "wipementions") {
		for k := len(bot[guildID].Mentions) - 1; k >= 0; k-- {
  			v := bot[guildID].Mentions[k]
  			if v.Mentioned == m.Author.ID {
     			bot[guildID].Mentions = append(bot[guildID].Mentions[:k], bot[guildID].Mentions[k+1:]...)
  			}
		}
		s.ChannelMessageSend(m.ChannelID, "I have wiped your mentions.")
	}



	if strings.HasPrefix(m.Content, in.Prefix + "lastseen ") {
		str := strings.Replace(m.Content, in.Prefix + "lastseen ", "", -1)
		if strings.Contains(m.Content, "<@") == false {
			return
		}

		us := strings.Split(str, "<@")[1]
		usr := strings.Split(us, ">")[0]
		usr = strings.Replace(usr, "!", "", -1)
		vj := 0
		for _, v := range bot[guildID].Users {
			if v.ID == usr {
				vj++
				lc := humanize.Time(v.LastSeen)
				s.ChannelMessageSend(m.ChannelID, "```xl\n"+fmt.Sprintf("*last seen %s*\n", lc) + fmt.Sprintf("Total Messages: %d", v.TotalMsg) + "```")
			}
		}
		if vj == 0 {
			s.ChannelMessageSend(m.ChannelID, "I haven't recorded the user in this server yet.")
		}
	}




if strings.HasPrefix(m.Content, "<@"+s.State.User.ID+"> verify") {
	// work for @Medallyon
	data := ""
	if m.Author.ID == "129036763367735297" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Administrator Detected```"
	}

	// work for proxy :D!!
	if m.Author.ID == "146046383726657536" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Owner Detected```"
	}

	// work for Paradoxum
	if m.Author.ID == "151467147367940106" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Administrator Detected```"
	}

	// work for XamTheKing
	if m.Author.ID == "107563269484490752" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Echo Support Detected```"
	}

	// work for Doorstop
	if m.Author.ID == "99312975537524736" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Administrator Detected```"
	}

	// work for EdibleDerpy
	if m.Author.ID == "106876115586383872" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Administrator Detected```"
	}

	// work for Hk
	if m.Author.ID == "104360151208706048" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Echo Support Detected```"
	}

	// work for pyyric
	/*
	if m.Author.ID == "98928980173815808" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Echo Support Detected```"
	}
	*/


	// work for Julian
	if m.Author.ID == "169596815904210944" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Echo Support Detected```"
	}


if data != "" {
	s.ChannelMessageSend(m.ChannelID, data)
} else {
	s.ChannelMessageSend(m.ChannelID, "```xl\n[✘] this user is NOT a technician```")
}

}







if strings.HasPrefix(m.Content, "<@!"+s.State.User.ID+"> verify") {
	// work for @Medallyon
	data := ""
	if m.Author.ID == "129036763367735297" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Administrator Detected```"
	}

	// work for proxy :D!!
	if m.Author.ID == "146046383726657536" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Owner Detected```"
	}



	// work for Paradoxum
	if m.Author.ID == "151467147367940106" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Administrator Detected```"
	}



	// work for XamTheKing
	if m.Author.ID == "107563269484490752" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Echo Support Detected```"
	}


	// work for Doorstop
	if m.Author.ID == "99312975537524736" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Administrator Detected```"
	}


	// work for EdibleDerpy
	if m.Author.ID == "106876115586383872" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Administrator Detected```"
	}

	// work for Hk
	if m.Author.ID == "104360151208706048" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Echo Support Detected```"
	}

	// work for moneshaq
	if m.Author.ID == "190211471009906688" {
		data = "```xl\n"+"Technician: " + m.Author.Username + "\nTech ID: " +m.Author.ID + "\n[✓] Echo Support Detected```"
	}

if data != "" {
	s.ChannelMessageSend(m.ChannelID, data)
} else {
	s.ChannelMessageSend(m.ChannelID, "```xl\n[✘] this user is NOT a technician```")
}

}







if strings.HasPrefix(m.Content, in.Prefix + "ratelimit true") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this command.")
		return
	}
	AddChannelLimit(s, m)
	s.ChannelMessageSend(m.ChannelID, "People are now limited to 1 message every 2 seconds.\nMake sure you have a role named `muted` setup to not allow `Send Messages` in your channels.")
}





if strings.HasPrefix(m.Content, in.Prefix + "ratelimit false") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this command.")
		return
	}
	ci := 0

		for k := len(bot[guildID].Restricted) - 1; k >= 0; k-- {
  			v := bot[guildID].Restricted[k]
  			if v.ID == m.ChannelID {
  				ci++
     			bot[guildID].Restricted = append(bot[guildID].Restricted[:k], bot[guildID].Restricted[k+1:]...)
  			}
		}
	if ci > 0 {
		s.ChannelMessageSend(m.ChannelID, "I have disabled rate limiting in this channel.")
	}
}



	if strings.HasPrefix(m.Content, in.Prefix + "countm") && m.Author.ID == js.Admin {
		count := 0
		for _, v := range bot {
			for _ = range v.Messages {
				count++
			}
		}
		s.ChannelMessageSend(m.ChannelID, "There are " + strconv.Itoa(count) + " Messages in memory at this moment.")
	}










// USERNAME FILTER OPTIONS ####################################################


if strings.HasPrefix(m.Content, in.Prefix + "namefilter ") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only a Bot Commander can use this command.")
		return
	}
	work := true
	str := strings.Replace(m.Content, in.Prefix + "namefilter ", "", -1)
	if strings.Contains(str, ":") == false {
		s.ChannelMessageSend(m.ChannelID, "You need to type the word to remove and a word to replace with.\nexamples: `"+in.Prefix+"namefilter badword:newword`\n`"+in.Prefix+"namefilter del:word`\n`"+in.Prefix+"namefilter list:all`")
		return
	}

	if len(strings.Split(str, ":")) > 1 {

	} else {
		s.ChannelMessageSend(m.ChannelID, "I see you're having issues.\nexamples: `"+in.Prefix+"namefilter badword:newword`\n`"+in.Prefix+"namefilter del:word`\n`"+in.Prefix+"namefilter list:all`")
		return
	}

	opt := strings.Split(str, ":")[0]
	tar := strings.Split(str, ":")[1]

	var filter map[string]string
	rr, err := gto.ReadFile("servers/"+guildID+"/namefilter.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rr, &filter)




	// they want to view the list
	if opt == "list" {
		work = false
		data := ""
		for k, v := range filter {
			data = data + "replace: " + k + " with: " + v + "\n"
		}
		if len(data) > 0 {
			s.ChannelMessageSend(m.ChannelID, "```xl\n"+data+"```")
			return
		} else {
			s.ChannelMessageSend(m.ChannelID, "You don't have any rules in your namefilter database.")
		}
	}


	// they want to delete a rule
	if opt == "del" {
		work = false
		cd := 0
		for k, _ := range filter {
			if k == tar {
				cd++
				delete(filter, k)
			}
		}
		if cd > 0 {
			io, err := gto.Marshal(filter)
			if err == nil {
				gto.WriteFile("servers/"+guildID+"/namefilter.json", io, 0777)
				s.ChannelMessageSend(m.ChannelID, "I have deleted the rule from your name filter.")
				return
			}
		}
	}





	if work == true {
		filter[opt] = tar
		io, err := gto.Marshal(filter)
		if err == nil {
			gto.WriteFile("servers/"+guildID+"/namefilter.json", io, 0777)
			s.ChannelMessageSend(m.ChannelID, "I have added the rule to your name filter.")
			return
		}		
	}

}





// USERNAME FILTER OPTIONS END #################################################















// ############### Word Filter Options ##########################################


if strings.HasPrefix(m.Content, in.Prefix + "filter ") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only a Bot Commander can use this command.")
		return
	}
	str := strings.Replace(m.Content, in.Prefix + "filter ", "", -1)
	if strings.Contains(str, ":") == false {
		s.ChannelMessageSend(m.ChannelID, "You need to type the word to remove and a word to replace with.\nexamples: `"+in.Prefix+"filter add:badword`\n`"+in.Prefix+"filter del:word`\n`"+in.Prefix+"filter list:all`")
		return
	}

	if len(strings.Split(str, ":")) > 1 {

	} else {
		s.ChannelMessageSend(m.ChannelID, "I see you're having issues.\nexamples: `"+in.Prefix+"filter add:badword`\n`"+in.Prefix+"filter del:word`\n`"+in.Prefix+"filter list:all`")
		return
	}

	opt := strings.Split(str, ":")[0]
	tar := strings.ToLower(strings.Split(str, ":")[1])

	var filter map[string]string
	rr, err := gto.ReadFile("servers/"+guildID+"/filter.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rr, &filter)




	// they want to view the list
	if opt == "list" {
		data := ""
		for k, _ := range filter {
			data = data + "word: " + k + "\n"
		}
		if len(data) > 0 {
		    k, err := s.UserChannelCreate(m.Author.ID)
    		if err == nil {
				s.ChannelMessageSend(k.ID, "```xl\n"+data+"```")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "You don't have any rules in your filter database.")
		}
	}


	// they want to delete a rule
	if opt == "del" {
		cd := 0
		for k, _ := range filter {
			if k == tar {
				cd++
				delete(filter, k)
			}
		}
		if cd > 0 {
			io, err := gto.Marshal(filter)
			if err == nil {
				gto.WriteFile("servers/"+guildID+"/filter.json", io, 0777)
				s.ChannelMessageSend(m.ChannelID, "I have deleted the rule from your word filter.")
				return
			}
		}
	}





	if opt == "add" {
		filter[tar] = tar
		io, err := gto.Marshal(filter)
		if err == nil {
			gto.WriteFile("servers/"+guildID+"/filter.json", io, 0777)
			s.ChannelMessageSend(m.ChannelID, "I have added the rule to your filter.")
			return
		}		
	}

}
// Word filter end ##########################################################








	if strings.HasPrefix(m.Content, in.Prefix + "userole true") {
		if BotMaster == false {
			s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this command.")
			return
		}
		var opts map[string]interface{}
		r1, err := gto.ReadFile("servers/"+guildID+"/options.json")
		if err == nil {
			gto.Unmarshal(r1, &opts)
			opts["UseRole"] = true
			io, err := gto.Marshal(opts)
			if err == nil {
				gto.WriteFile("servers/"+guildID+"/options.json", io, 0777)
				s.ChannelMessageSend(m.ChannelID, "Alright I now require users to have the Bot Commander role. By default the role is `Bot Commander` you can change this role by typing `"+in.Prefix+"botcommander Role Name`")
			}
		}
	}





	if strings.HasPrefix(m.Content, in.Prefix + "userole false") {
		if BotMaster == false {
			s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this command.")
			return
		}
		var opts map[string]interface{}
		r1, err := gto.ReadFile("servers/"+guildID+"/options.json")
		if err == nil {
			gto.Unmarshal(r1, &opts)
			
			for k, _ := range opts {
				if k == "UseRole" {
					delete(opts, k)
				}
			}
			io, err := gto.Marshal(opts)
			if err == nil {
				gto.WriteFile("servers/"+guildID+"/options.json", io, 0777)
				s.ChannelMessageSend(m.ChannelID, "Alright anyone who has `Manage Server` permission will be considered my master.")
			}
		}
	}






if strings.HasPrefix(m.Content, in.Prefix + "stopfilter") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this command.")
		return
	}
	var opts map[string]string
	rr, err := gto.ReadFile("servers/"+guildID+"/options.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rr, &opts)
	for k, _ := range opts {
		if k == "WordFilter" {
			delete(opts, k)
		}
	}
	io, err := gto.Marshal(opts)
	if err == nil {
		s.ChannelMessageSend(m.ChannelID, "Word filter has been disabled.")
		gto.WriteFile("servers/"+guildID+"/options.json", io, 0777)
	}
}




if strings.HasPrefix(m.Content, in.Prefix + "startfilter ") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this command.")
		return
	}
	str := gto.Replace(m.Content, in.Prefix + "startfilter ", "", -1)
	str = strings.ToLower(str)
	if str != "kick" && str != "ban" && str != "warn" {
		s.ChannelMessageSend(m.ChannelID, "You need to pick a punishment. warn, kick or ban. an example: `"+in.Prefix+"startfilter warn`")
		return
	}
	
	var opts map[string]string
	rr, err := gto.ReadFile("servers/"+guildID+"/options.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rr, &opts)
	opts["WordFilter"] = str
	io, err := gto.Marshal(opts)
	if err == nil {
		s.ChannelMessageSend(m.ChannelID, "Word filter has been enabled.")
		gto.WriteFile("servers/"+guildID+"/options.json", io, 0777)
	}
}






if strings.HasPrefix(m.Content, in.Prefix + "greetchannel ") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this command.")
		return
	}
	str := gto.Replace(m.Content, in.Prefix + "greetchannel ", "", -1)
	if strings.Contains(str, "<#") == false {
		s.ChannelMessageSend(m.ChannelID, "You need to mention the channel. `"+in.Prefix+"greetchannel #channel-name`")
		return
	}
	str = strings.Replace(str, "<#", "", -1)
	str = strings.Replace(str, ">", "", -1)

	_, err := s.State.Channel(str)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "This channel doesn't exist in this server.")
		return
	}
	var opts map[string]string
	rr, err := gto.ReadFile("servers/"+guildID+"/options.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rr, &opts)
	if str == guildID {
		for k, _ := range opts {
			if k == "GreetChannel" {
				delete(opts, k)
			}
		}
	} else {
		opts["GreetChannel"] = str
	}
	io, err := gto.Marshal(opts)
	if err == nil {
		s.ChannelMessageSend(m.ChannelID, "I have set your greet channel to <#"+str+">")
		gto.WriteFile("servers/"+guildID+"/options.json", io, 0777)
	}
}





if strings.HasPrefix(m.Content, in.Prefix + "autochannel ") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this command.")
		return
	}
	str := gto.Replace(m.Content, in.Prefix + "autochannel ", "", -1)
	if strings.Contains(str, "<#") == false {
		s.ChannelMessageSend(m.ChannelID, "You need to mention the channel. `"+in.Prefix+"autochannel #channel-name`")
		return
	}
	str = strings.Replace(str, "<#", "", -1)
	str = strings.Replace(str, ">", "", -1)

	_, err := s.State.Channel(str)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "This channel doesn't exist in this server.")
		return
	}
	var opts map[string]string
	rr, err := gto.ReadFile("servers/"+guildID+"/options.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rr, &opts)
	if str == guildID {
		for k, _ := range opts {
			if k == "AutoRoleChannel" {
				delete(opts, k)
			}
		}
	} else {
		opts["AutoRoleChannel"] = str
	}
	io, err := gto.Marshal(opts)
	if err == nil {
		s.ChannelMessageSend(m.ChannelID, "I have set your auto role channel to <#"+str+">")
		gto.WriteFile("servers/"+guildID+"/options.json", io, 0777)
	}
}








if strings.HasPrefix(m.Content, in.Prefix + "8ball ") {
	q := []string{
			"It is certain",
			"It is decidedly so",
			"Without a doubt",
			"Yes, definitely",
			"You may rely on it",
			"As I see it, yes",
			"Most likely",
			"Outlook good",
			"Yes",
			"Signs point to yes",
			"Reply hazy try again",
			"Ask again later",
			"Better not tell you now",
			"Cannot predict now",
			"Concentrate and ask again",
			"Don't count on it",
			"My reply is no",
			"My sources say no",
			"Outlook not so good",
			"Very doubtful",
		}
		i := gto.Random(0, 19)
		s.ChannelMessageSend(m.ChannelID, q[i])
}









if strings.HasPrefix(m.Content, in.Prefix + "tweet ") {
	tweetit := strings.Replace(m.Content, in.Prefix + "tweet ", "", -1)
	tweetit = gto.ToLower(tweetit)

	tweetit = strings.Replace(tweetit, "nigger", "", -1)
	tweetit = strings.Replace(tweetit, "faggot", "", -1)
	tweetit = strings.Replace(tweetit, "n1gger", "", -1)
	tweetit = strings.Replace(tweetit, "pussy", "", -1)
	tweetit = strings.Replace(tweetit, "whore", "", -1)
	tweetit = strings.Replace(tweetit, "n1gg3r", "", -1)
	tweetit = strings.Replace(tweetit, "niger", "", -1)
	tweetit = strings.Replace(tweetit, "echo sucks", "", -1)
	tweetit = strings.Replace(tweetit, "bitch", "", -1)
	tweetit = strings.Replace(tweetit, "clitoris", "", -1)
	tweetit = strings.Replace(tweetit, "asshole", "", -1)
	tweetit = strings.Replace(tweetit, "ass ", "", -1)
	tweetit = strings.Replace(tweetit, "tits", "", -1)
	tweetit = strings.Replace(tweetit, "titties", "", -1)


	banlist, err := gto.ReadLines("System/tweetbans.txt")
	if err == nil {
		//	top := 135
		for _, v := range banlist {
			if m.Author.ID == v {
				s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> You've been banned from my twitter system. For more information ask Proxy in his server.")
				return
			}
		}

	if len(tweetit) > 140 {
		s.ChannelMessageSend(m.ChannelID, "Tweet is longer than 140 characters in length.")
		return
	}

	tweetit = tweetit + "\n\nBy #" + strings.Replace(m.Author.Username, " ", "", -1) + "\nID: " + m.Author.ID
	Tweet(tweetit, url.Values{"status": {tweetit}})
	s.ChannelMessageSend(m.ChannelID, "You're message has been sent to my twitter: <https://twitter.com/EchoTheBot>")
}
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}



	if strings.HasPrefix(m.Content, in.Prefix + "monitor ") {
		if m.Author.ID != js.Admin {
			return
		}
		str := strings.Replace(m.Content, in.Prefix + "monitor ", "", -1)
		if gto.ToLower(str) == "true" {
			isBeep = true
			s.ChannelMessageSend(m.ChannelID, "Alright I've enabled the monitor system. Lower your sound\n```go\n3 beeps = invited to server\n2 beeps = kicked from server\n1 beep = playing music```")
		} else {
			isBeep = false
			s.ChannelMessageSend(m.ChannelID, "I have turned the monitor system off.")
		}
	}







if strings.HasPrefix(m.Content, in.Prefix + "autonick ") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this command.")
		return
	}
	str := strings.Replace(m.Content, in.Prefix + "autonick ", "", -1)
	if str == "" {
		s.ChannelMessageSend(m.ChannelID, "You need to pick the nickname an example: `--autonick [GR]-{/user}`")
		return
	}
	var opts map[string]interface{}
	rf1, err := gto.ReadFile("servers/"+guildID+"/options.json")
	if err != nil {
		return
	}
	gto.Unmarshal(rf1, &opts)

	if str == "off" {
		for k,_ := range opts {
			if k == "AutoNick" {
				delete(opts, k)
			}
		}
	} else {
		opts["AutoNick"] = str
	}

	io, err := gto.Marshal(opts)
	if err == nil {
		gto.WriteFile("servers/"+guildID+"/options.json", io, 0777)
		s.ChannelMessageSend(m.ChannelID, "I have set the AutoNick System.")
	}
}














if strings.HasPrefix(m.Content, in.Prefix + "nick ") {
	if BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders have access to this command.")
		return
	}
	if strings.Contains(m.Content, "<@") {
		usr := gto.Split(m.Content, "<@")[1]
		user := gto.Split(usr, ">")[0]
		
		nickname := strings.Replace(m.Content, in.Prefix + "nick <@!"+user+"> ", "", -1)
		nickname = strings.Replace(nickname, in.Prefix + "nick <@"+user+"> ", "", -1)

		user = strings.Replace(user, "!", "", -1)
		err = s.GuildMemberNickname(guildID, user, nickname)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "`Failed:` Check my permissions also *I cannot edit the server owners nickname*")
		} else {
			s.ChannelMessageSend(m.ChannelID, "I have changed <@"+user+">'s nickname.")
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "You need to mention the user. `"+in.Prefix+"nick @User new nickname`")
	}
}










	if strings.HasPrefix(m.Content, in.Prefix + "stopsounds") {
		if BotMaster == true {
			var options map[string]interface{}
			// let's load the file options.json
			ops, err := gto.ReadFile("servers/"+guildID+"/options.json")
			if err != nil {
				return // something went wrong and they don't have options.json
			}
			gto.Unmarshal(ops, &options)
			options["NoSounds"] = true
			gi, err := gto.Marshal(options)
			if err == nil {
				gto.WriteFile("servers/"+guildID+"/options.json", gi, 0777)
				s.ChannelMessageSend(m.ChannelID, "I have disabled all sounds on this server.")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders have access to this command.")
		}
		go CheckLog(s, m, in)
	}






	if strings.HasPrefix(m.Content, in.Prefix + "startsounds") {
		if BotMaster == true {
			var options map[string]interface{}
			// let's load the file options.json
			ops, err := gto.ReadFile("servers/"+guildID+"/options.json")
			if err != nil {
				return // something went wrong and they don't have options.json
			}
			gto.Unmarshal(ops, &options)

			if _, ok := options["NoSounds"]; ok {
				for k, _ := range options {
					if k == "NoSounds" {
						delete(options, k)
					}
				}
			}
			gi, err := gto.Marshal(options)
			if err == nil {
				gto.WriteFile("servers/"+guildID+"/options.json", gi, 0777)
				s.ChannelMessageSend(m.ChannelID, "Sounds are now enabled on this server.")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders have access to this command.")
		}
		go CheckLog(s, m, in)
	}







	if strings.HasPrefix(m.Content, in.Prefix + "restrict ") {
		if BotMaster == true {
			channel := strings.Replace(m.Content, in.Prefix + "restrict ", "", -1)
			if channel == "" {
				s.ChannelMessageSend(m.ChannelID, "An example of this command would be `"+in.Prefix+"restrict #channel`")
				return // stop what you are doing they messed up.
			}
			if strings.Contains(channel, "<#") == false {
				s.ChannelMessageSend(m.ChannelID, "You need to mention the channel using **#**channelname")
				return
			}

			channel = strings.Replace(channel, "<#", "", -1)
			channel = strings.Replace(channel, ">", "", -1)
			var options map[string]interface{}
			// let's load the file options.json
			ops, err := gto.ReadFile("servers/"+guildID+"/options.json")
			if err != nil {
				return // something went wrong and they don't have options.json
			}
			gto.Unmarshal(ops, &options)

			options["Restrict"] = channel
			gi, err := gto.Marshal(options)
			if err == nil {
				gto.WriteFile("servers/"+guildID+"/options.json", gi, 0777)
				s.ChannelMessageSend(m.ChannelID, "I have restricted all my commands to the channel <#"+channel+">")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders have access to this command.")
		}
	//	go CheckLog(s, m, in)
	}








	if strings.HasPrefix(m.Content, in.Prefix + "denyreminder") {
		if BotMaster == true {
			var options map[string]interface{}
			// let's load the file options.json
			ops, err := gto.ReadFile("servers/"+guildID+"/options.json")
			if err != nil {
				return // something went wrong and they don't have options.json
			}
			gto.Unmarshal(ops, &options)
			options["NoReminder"] = true
			gi, err := gto.Marshal(options)
			if err == nil {
				gto.WriteFile("servers/"+guildID+"/options.json", gi, 0777)
				s.ChannelMessageSend(m.ChannelID, "I have restricted the `remind` command to Commanders only.")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders have access to this command.")
		}
	}


	if strings.HasPrefix(m.Content, in.Prefix + "allowreminder") {
		if BotMaster == true {
			var options map[string]interface{}
			// let's load the file options.json
			ops, err := gto.ReadFile("servers/"+guildID+"/options.json")
			if err != nil {
				return // something went wrong and they don't have options.json
			}
			gto.Unmarshal(ops, &options)

			if _, ok := options["NoReminder"]; ok {
				for k, _ := range options {
					if k == "NoReminder" {
						delete(options, k)
					}
				}
			}
			gi, err := gto.Marshal(options)
			if err == nil {
				gto.WriteFile("servers/"+guildID+"/options.json", gi, 0777)
				s.ChannelMessageSend(m.ChannelID, "Everyone can now use the reminder system in this server.")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders have access to this command.")
		}
	//	go CheckLog(s, m, in)
	}








	if strings.HasPrefix(m.Content, in.Prefix + "unrestrict") {
		if BotMaster == true {
			var options map[string]interface{}
			// let's load the file options.json
			ops, err := gto.ReadFile("servers/"+guildID+"/options.json")
			if err != nil {
				return // something went wrong and they don't have options.json
			}
			gto.Unmarshal(ops, &options)

			if _, ok := options["Restrict"]; ok {
				for k, _ := range options {
					if k == "Restrict" {
						delete(options, k)
					}
				}
			}
			gi, err := gto.Marshal(options)
			if err == nil {
				gto.WriteFile("servers/"+guildID+"/options.json", gi, 0777)
				s.ChannelMessageSend(m.ChannelID, "I have removed the restriction.")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders have access to this command.")
		}
		go CheckLog(s, m, in)
	}










	if strings.HasPrefix(m.Content, in.Prefix + "remind ") {
		if AllowReminder == false && BotMaster == false {
			s.ChannelMessageSend(m.ChannelID, "The reminder system has been restricted in this server.")
			return
		}

		ti := strings.Replace(m.Content, in.Prefix + "remind ", "", -1)
		if strings.Contains(ti, " ") == false {
			s.ChannelMessageSend(m.ChannelID, "I see you're having issues. you can view my website for an indepth example but for now `"+in.Prefix+"remind 2h15s {user} I have reminded you!`")
			return
		}

		tim := gto.Split(ti, " ")[0] 
		text := strings.Replace(m.Content, in.Prefix + "remind " + tim + " ", "", -1)
		if strings.Contains(text, "{everyone}") {
			if BotMaster == false {
				s.ChannelMessageSend(m.ChannelID, "Only my Bot Commanders have access to mention everyone.")
				return
			}
		}
		if strings.Contains(gto.ToLower(text), "{role:everyone}") {
			if BotMaster == false {
				s.ChannelMessageSend(m.ChannelID, "Only my Bot Commanders have access to mention everyone.")
				return
			}
		}


		io, err := time.ParseDuration(tim)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "You need to format the time like `2h15m` or `35s`, `15m`, `24h15m` etc..")
		} else {
			s.ChannelMessageSend(m.ChannelID, "I have set your reminder. for `"+tim+"`")
			go Reminder(s, m, text, io)
		}
	}




if strings.HasPrefix(m.Content, in.Prefix + "playcount") {
	if m.Author.ID == js.Admin {
		cnt := 0
		for _ = range playing {
			cnt++
		}
		bo, err := gto.String(cnt)
		if err != nil {
			return
		}
		s.ChannelMessageSend(m.ChannelID, "I'm currently playing on `"+bo+"` servers.")
	}
}




if strings.HasPrefix(m.Content, in.Prefix + "play ") {
	if sudo == false && Donator == false {
		return
	}
	str := strings.Replace(m.Content, in.Prefix + "play ", "", -1)
	s.ChannelMessageSend(m.ChannelID, "You need to be a Commander or have the role `Controller` to use this command.")

if isConverting(m) == true {
	s.ChannelMessageSend(m.ChannelID, "You're currently processing the max songs at a time. Please wait for your current tasks to complete.")
	return
}


	if CountQueue(guildID) <= 5 {
		go Convert(s, m, str)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Hi, my music feature is currently being tested for everyone. During this process i need to limit each server `queue` to 3 songs. You can play from your queue `"+in.Prefix+"qplay` or wipe your queue `"+in.Prefix+"qclear`")
		return
	}
}



if strings.HasPrefix(m.Content, in.Prefix + "qclear") {
	if sudo == false || m.Author.ID != js.Admin || Donator == false {
		return
	}
	if ClearQueue(guildID) == true {
		s.ChannelMessageSend(m.ChannelID, "I have cleared your queue.")
	} else {
		s.ChannelMessageSend(m.ChannelID, "You do not have any songs in your queue.")
	}
}








if strings.HasPrefix(m.Content, in.Prefix + "mstop") {
	if sudo == false || m.Author.ID != js.Admin || Donator == false {
		return
	}
	if gto.MemberHasRole(s, guildID, m.Author.ID, "Controller") == false && BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "You need to be a Commander or have the role `Controller` to use this command.")
		return
	}

	for k, v := range playing {
		if k == guildID {
			delete(playing, k)
			os.Remove("music/"+v+".dca")
		}
	}
	mbuffer = nil
	if _, ok := s.VoiceConnections[guildID]; ok {
		s.VoiceConnections[guildID].Disconnect()
	}
	s.ChannelMessageSend(m.ChannelID, "I have stopped playing music.")
}










	if strings.HasPrefix(m.Content, in.Prefix + "qplay") {
		if sudo == false || m.Author.ID != js.Admin || Donator == false {
			return
		}
	if gto.MemberHasRole(s, guildID, m.Author.ID, "Controller") == false && BotMaster == false {
		s.ChannelMessageSend(m.ChannelID, "You need to be a Commander or have the role `Controller` to use this command.")
		return
	}

		playit := true
		for k1, _ := range playing {
			//fmt.Println(k1)
			if k1 == guildID {
				playit = false
			}
		}

		// if playit = false than we need to add the song to the queue.
		var queue map[string]string
		fq, err := gto.ReadFile("queue/"+guildID+".json")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "This server doesn't have any songs in queue.")
			return
		}
		gto.Unmarshal(fq, &queue)
		if playit == false {
			s.ChannelMessageSend(m.ChannelID, "You can't start the queue when it's already playing. Having an error? type `"+in.Prefix+"mstop`and try again!")
		}

		if playit == true {
		// Find the guild for that channel.
		g, err := s.State.Guild(guildID)
		if err != nil {
			// Could not find guild.
			return
		}

		ck := 0
		thesong := ""
		for k, _ := range queue {
			ck++
			if ck == 1 {
				thesong = k
				delete(queue, k)
				b, err := gto.Marshal(queue)
				if err == nil {
					gto.WriteFile("queue/"+guildID+".json", b, 0777)
				}
			}
		}


		if ck > 0 {
		// Look for the message sender in that guilds current voice states.
		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {
				err := loadMusic("music/" + thesong + ".dca")
				if err != nil {
					// s.ChannelMessageSend(m.ChannelID, "I couldn't find `"+str+"` Check your spelling. type `!msearch keyword` to find media!")
					fmt.Println("Error loading sound: ", err)
					for k, _ := range queue {
						if k == thesong {
							delete(queue, k)
							b, err := gto.Marshal(queue)
							if err == nil {
								gto.WriteFile("queue/"+guildID+".json", b, 0777)
							}
						}
					}
				//	return
				}
				err = playMusic(s, m, g.ID, vs.ChannelID, thesong)
				if err != nil {
					for k, _ := range queue {
						if k == thesong {
							delete(queue, k)
							b, err := gto.Marshal(queue)
							if err == nil {
								gto.WriteFile("queue/"+guildID+".json", b, 0777)
							}
						}
					}
					// s.ChannelMessageSend(m.ChannelID, "I couldn't find `"+str+"` Check your spelling. type `!msearch keyword` to find media!")
				}
			}
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "This server doesn't have any songs in the queue.")
	}
}
} // end of !mplay




if strings.HasPrefix(m.Content, in.Prefix + "mstop") {
	if _, ok := s.VoiceConnections[guildID]; ok {
		s.VoiceConnections[guildID].Disconnect()
		s.ChannelMessageSend(m.ChannelID, "I have left the voice session.")
	} else {
		s.ChannelMessageSend(m.ChannelID, "I wasn't in a voice session to begin with.")
	}
}




	if strings.HasPrefix(m.Content, in.Prefix + "911") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/911.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	}





	if strings.HasPrefix(m.Content, in.Prefix + "belch") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/belch.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	}








	if strings.HasPrefix(m.Content, in.Prefix + "creepy") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/creepy.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	}






	if strings.HasPrefix(m.Content, in.Prefix + "fart") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/fart.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay






	if strings.HasPrefix(m.Content, in.Prefix + "sstop") {
		if BotMaster == true {
			if _, ok := s.VoiceConnections[guildID]; ok {
				s.VoiceConnections[guildID].Disconnect()
			}
			s.ChannelMessageSend(m.ChannelID, "I have stopped playing sounds.")
		} else {
			s.ChannelMessageSend(m.ChannelID, "Only the Owner of this server or people in the `Bot Commander` role can use this feature.")
		}
		go CheckLog(s, m, in)
		return
	}







	if strings.HasPrefix(m.Content, in.Prefix + "fart2") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/fart2.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "firelazer") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/firelazer.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "goofy") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/goofy.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "ifarted") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/ifarted.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "laugh") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/laugh.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "laugh2") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/laugh2.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay


	if strings.HasPrefix(m.Content, in.Prefix + "pewpew") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/pewpew.dca")
				if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "pig") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/pig.dca")
						if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay




	if strings.HasPrefix(m.Content, in.Prefix + "quack") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/quack.dca")
						if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay


	if strings.HasPrefix(m.Content, in.Prefix + "roar") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/roar.dca")
						if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "snore") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/snore.dca")
						if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "drumroll") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/drumroll.dca")
						if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "superman") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/superman.dca")
						if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "wetfart") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/wetfart.dca")
						if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay


	if strings.HasPrefix(m.Content, in.Prefix + "wookie") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/wookie.dca")
						if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay



	if strings.HasPrefix(m.Content, in.Prefix + "yourmom") {
		if AllowSounds == false {
			s.ChannelMessageSend(m.ChannelID, "Sounds have been disabled on this server.")
			return
		}
		go PlaySystem(s, m, "dca/yourmom.dca")
						if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
			return
		}
	} // end of !mplay









if strings.HasPrefix(m.Content, in.Prefix + "logtype ") {
	var work bool
	work = false
	if BotMaster == true {
		channel := strings.Replace(m.Content, in.Prefix + "logtype ", "", -1)
		if channel == "all" {
			work = true
		}
		if channel == "master" {
			work = true
		}

		if work == false {
			s.ChannelMessageSend(m.ChannelID, "An example of this command would be `"+in.Prefix+"logtype all` or `"+in.Prefix+"logtype master`")
			return
		}


		var myconf map[string]interface{}
		rd1, err := gto.ReadFile("servers/"+guildID+"/options.json")
		if err == nil {
			gto.Unmarshal(rd1, &myconf)
		} else {
			return
		}
		myconf["LogType"] = channel
		io, err := gto.Marshal(myconf)
		if err == nil {
			gto.WriteFile("servers/"+guildID+"/options.json", io, 0777)
			s.ChannelMessageSend(m.ChannelID, "Command logging has been enabled.")
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "Only a Bot master can do this.")
	}
	go CheckLog(s, m, in)
return
}










if strings.HasPrefix(m.Content, in.Prefix + "logstart ") {
	if BotMaster == true {
		channel := strings.Replace(m.Content, in.Prefix + "logstart ", "", -1)
		if channel == "" || strings.Contains(m.Content, "<#") == false {
			s.ChannelMessageSend(m.ChannelID, "You need to enter a channel for broadcasting. `"+in.Prefix+"startlog #channel`")
			return
		}
		channel = strings.Replace(channel, "<#", "", -1)
		channel = strings.Replace(channel, ">", "", -1)

		var myconf map[string]interface{}
		rd1, err := gto.ReadFile("servers/"+guildID+"/options.json")
		if err == nil {
			gto.Unmarshal(rd1, &myconf)
		} else {
			return
		}
		myconf["Log"] = channel
		io, err := gto.Marshal(myconf)
		if err == nil {
			gto.WriteFile("servers/"+guildID+"/options.json", io, 0777)
			s.ChannelMessageSend(m.ChannelID, "Command logging has been enabled.")
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "Only a Bot master can do this.")
	}
	go CheckLog(s, m, in)
}








if strings.HasPrefix(m.Content, in.Prefix + "logstop") {
	if BotMaster == true {
		var myconf map[string]interface{}
		rd1, err := gto.ReadFile("servers/"+guildID+"/options.json")
		if err == nil {
			gto.Unmarshal(rd1, &myconf)
		} else {
			return
		}
		for k, _ := range myconf {
			if k == "Log" {
			delete(myconf, k)
			}
			if k == "LogType" {
				delete(myconf, k)
			}
		}

		io, err := gto.Marshal(myconf)
		if err == nil {
			gto.WriteFile("servers/"+guildID+"/options.json", io, 0777)
			s.ChannelMessageSend(m.ChannelID, "Command logging has been disabled.")
			go CheckLog(s, m, in)
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "Only a Bot master can do this.")
		go CheckLog(s, m, in)
	}
}






if strings.HasPrefix(m.Content, in.Prefix + "coopvsai") {
	str := strings.Replace(m.Content, in.Prefix + "coopvsai ", "", -1)
	r:= gto.Split(str, " ")
	region := r[0]
	user := strings.Replace(str, region + " ", "", -1)
	if len(r) > 1 {
		theid := GetSummonerID(region, user)
		if theid > 0 {
			id := fmt.Sprintf("%d", theid)
			idd, err := strconv.Atoi(id)
			i64 := int64(idd)
			if err != nil {
				return
			}
   			s.ChannelMessageSend(m.ChannelID,  StatSummariesBySummoner(region, i64, "CoopVsAI"))
   		}
	}
					if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}	
}



if strings.HasPrefix(m.Content, in.Prefix + "unranked") {
	str := strings.Replace(m.Content, in.Prefix + "unranked ", "", -1)
	r:= gto.Split(str, " ")
	region := r[0]
	user := strings.Replace(str, region + " ", "", -1)
	if len(r) > 1 {
		theid := GetSummonerID(region, user)
		if theid > 0 {
			id := fmt.Sprintf("%d", theid)
			idd, err := strconv.Atoi(id)
			i64 := int64(idd)
			if err != nil {
				return
			}
   			s.ChannelMessageSend(m.ChannelID,  StatSummariesBySummoner(region, i64, "Unranked"))
   		}
	}
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}



if strings.HasPrefix(gto.ToLower(m.Content), in.Prefix + "changelog") {
	s.ChannelMessageSend(m.ChannelID, "This command is deprecated. Add a channel to your server named `#echo-updates` for live updates and change logs.")
}










if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "clear") {
	var list []string
	str := strings.Replace(m.Content, in.Prefix + "clear ", "", -1)
	do := 0
	if strings.Contains(str, "@") {
		if strings.Contains(str, " ") == false {
			s.ChannelMessageSend(m.ChannelID, "You need to type an amount. `"+in.Prefix+"clear @User 10`")
			return
		}
		dat := gto.Split(str, " ")
		user := dat[0]
		num, err := gto.Integer(dat[1])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Incorrect syntax. `"+in.Prefix+"clear @User 10`")
			return
		}

		if num > 100 {
			s.ChannelMessageSend(m.ChannelID, "I will only wipe 100 messages at a time.")
			return
		}
		user = strings.Replace(user, "<@", "", -1)
		user = strings.Replace(user, "!", "", -1)
		user = strings.Replace(user, ">", "", -1)
		msgs, err := s.ChannelMessages(m.ChannelID, 100, "", "")
		for _, v := range msgs {
			if v.Author.ID == user && do <= num {
				do++
				list = append(list, v.ID)
				// s.ChannelMessageDelete(m.ChannelID, v.ID)
				// time.Sleep(1000 * time.Millisecond)
			}
		}
		if do > 0 {
			s.ChannelMessagesBulkDelete(m.ChannelID, list)
		}

	} else {
		num, err := gto.Integer(str)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Incorrect syntax. `"+in.Prefix+"clear @User 10`")
			return
		}
		if num > 100 {
			s.ChannelMessageSend(m.ChannelID, "I will only wipe 100 messages at a time.")
			return
		}
		msgs, err := s.ChannelMessages(m.ChannelID, num, "", "")
		for _, v := range msgs {
			if do <= num {
				do++
				list = append(list, v.ID)
				// s.ChannelMessageDelete(m.ChannelID, v.ID)
				// time.Sleep(1000 * time.Millisecond)
			}
		}		
		if do > 0 {
			s.ChannelMessagesBulkDelete(m.ChannelID, list)
		}
		// they just want to clear messages not a users messages.
	}
	go CheckLog(s, m, in)
return
}












if strings.HasPrefix(m.Content, in.Prefix + "rip") {
	str := strings.Replace(m.Content, in.Prefix + "rip ", "", -1)
	if str == in.Prefix + "rip" {
		s.ChannelMessageSend(m.ChannelID, "Example for my custom rip: `"+in.Prefix+"rip TheirName=The Message You want to type!`")
		return		
	}
	if strings.Contains(m.Content, "=") == false {
		s.ChannelMessageSend(m.ChannelID, "Example for my custom rip: `"+in.Prefix+"rip TheirName=The Message You want to type!`")
		return			
	}
	dat := gto.Split(str, "=")
	if len(dat) < 1 {
		s.ChannelMessageSend(m.ChannelID, "Example for my custom rip: `"+in.Prefix+"rip TheirName=The Message You want to type!`")
		return
	}

	name := dat[0]
	text := dat[1]
	if len(name) > 19 {
		s.ChannelMessageSend(m.ChannelID, "Keep the name under 20 characters.")
		return
	}

	if len(text) > 100 {
		s.ChannelMessageSend(m.ChannelID, "Keep the message under 100 characters.")
		return
	}

/*
	var imgT []string
	cnd := 0
	dat1 := ""
	thedat := gto.Split(text, " ")
	if strings.Contains(text, " ") {
		for _, v := range thedat {
			cnd++
			dat1 = dat1 + v
			if cnd == 3 {
				append(imgT, v)
				dat1 = ""
			}
		}
	}
*/

WriteFile("rip", m.Author.ID, 50, 136, "fonts/ariblk.ttf", 20, name)
testing := text + " "
t := gto.Split(testing, "{br}")
cnt := 0
data := ""
top := 160
for _, v := range t {
	cnt++
	data = data + v + " "
	WriteFile(m.Author.ID, m.Author.ID, 40, top, "fonts/cour.ttf", 10, data)	
	top = top + 12
	data = ""
	cnt = 0
}
time.Sleep(2000 * time.Millisecond)
mk, err := gto.ReadFile("images/custom/"+m.Author.ID+".png")
if err != nil {
	return
}
tc := bytes.NewReader(mk)
_, err = s.ChannelFileSend(m.ChannelID, "EchoRipSystem.png", tc)
if err != nil {
	s.ChannelMessageSend(m.ChannelID, "I don't have permissions to post images in this channel.")
}

					if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}

}












if strings.HasPrefix(m.Content, in.Prefix + "donate") {
	s.ChannelMessageSend(m.ChannelID, "Hey user! We could always use help for server fees! if you feel generous <https://www.paypal.me/EchoBot/5> Thank you!")
					if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}




if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "grabars") {
	mk, err := gto.ReadFile("servers/"+guildID+"/autoresponse.json")
	if err == nil {
		s.ChannelTyping(m.ChannelID)
        tc := bytes.NewReader(mk)
        pm, err := s.UserChannelCreate(m.Author.ID)
        if err == nil {
        	s.ChannelMessageSend(m.ChannelID, "Your A.R.S File has been sent to your DM!")
        	s.ChannelFileSend(pm.ID, "autoresponse.json", tc)
        } else {
        	s.ChannelMessageSend(m.ChannelID, "You have me blocked. Or you have PM's disabled. I wasn't able to send your A.R.S file.")
        }
	}	
	go CheckLog(s, m, in)
return
}



if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "putars ") {
	/*
	str := strings.Replace(m.Content, in.Prefix + "putars ", "", -1)
	if str != "" {
		if strings.Contains(str, ".json") {
			if gto.DownloadFile("servers/"+guildID+"/autoresponse.json", str) == true {
				s.ChannelMessageSend(m.ChannelID, "I have updated your `A.R.S` File. to view changes type `--viewauto`")
			} else {
				s.ChannelMessageSend(m.ChannelID, "There was an error updating your `A.R.S` File. make sure the link directly links to the .json file!")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "You need to link directly to your .json file. example: <http://yourserver.com/autoresponse.json> no short urls allowed.")
		}
	}
	go CheckLog(s, m, in)
return
*/
s.ChannelMessageSend(m.ChannelID, "This command is deprecated. Drag your `autoresponse.json` file in any channel I can see to update your A.R.S.")
}



if strings.HasPrefix(m.Content, ":listemojis:") && m.Author.ID != s.State.User.ID {
s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> You can view the guide here: <http://echobot.tk>")
}







if strings.HasPrefix(m.Content, in.Prefix + "sky") {
go PostImage("skies", s, m)
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}



if strings.HasPrefix(m.Content, in.Prefix + "space") {
go PostImage("space", s, m)
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}



if strings.HasPrefix(m.Content, in.Prefix + "dbz") {
go PostImage("dbz", s, m)
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}



if strings.HasPrefix(m.Content, in.Prefix + "cute") {
go PostImage("cute", s, m)
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}



if strings.HasPrefix(m.Content, in.Prefix + "cars") {
go PostImage("cars", s, m)
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}



if strings.HasPrefix(m.Content, in.Prefix + "wrecks") {
go PostImage("wrecks", s, m)
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}



if strings.HasPrefix(m.Content, in.Prefix + "trucks") {
go PostImage("trucks", s, m)
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}








if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "nsfw") {
	s.ChannelMessageSend(m.ChannelID, "The NSFW Commands have been moved to my other bot. Invite #NSFW Today! Click Here <http://bit.ly/25aCO2K>")
}











if BotMaster == true && strings.HasPrefix(gto.ToLower(m.Content), in.Prefix + "mkchan") {
	chtype := "text"
	str := strings.Replace(gto.ToLower(m.Content), in.Prefix + "mkchan ", "", -1)
	dat := gto.Split(str, " ")
	chname := strings.Replace(dat[0], " ", "-", -1)
	chdat := strings.Replace(str, chname, "", -1)
	if strings.Contains(chdat, " ") {
		chtype = dat[1]
	}
	if chname != "" && chtype != "" {
		_, err = s.GuildChannelCreate(guildID, chname, chtype)
		if err == nil {
			newdata := strings.Replace(resp.Mkchan, "{channel}", "#"+chname+"", -1)
			newdata = strings.Replace(newdata, "{data}", chtype, -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
		} else {
			s.ChannelMessageSend(m.ChannelID, resp.MkchanError)
		}
	}
	go CheckLog(s, m, in)
return
} 





if strings.HasPrefix(gto.ToLower(m.Content), in.Prefix + "locateip") {
	str := strings.Replace(m.Content, in.Prefix + "locateip ", "", -1)
	str = strings.Replace(str, ",", ".", -1)
	if strings.Contains(str, ".") {
	var ip map[string]interface{}
	gto.GetJson("http://ip-api.com/json/"+str, &ip)

if ip["status"].(string) != "fail" {
	region := ip["regionName"].(string)
	zipcode := ip["zip"].(string)
	country := ip["country"].(string)
	city := ip["city"].(string)
	timezone := ip["timezone"].(string)
	isp := ip["isp"].(string)
	who := ip["as"].(string)
	theip := ip["query"].(string)

	format := "```ruby\n"+"ip: "+theip+"\ncountry: "+country+"\nregion: "+region+"\ncity: "+city+"\nzipcode: "+zipcode+"\ntimezone: "+timezone+"\nisp: "+isp+"\n"+who+"```"
	s.ChannelMessageSend(m.ChannelID, format)
} else {
	s.ChannelMessageSend(m.ChannelID, resp.IpError)	
}
} else {
	s.ChannelMessageSend(m.ChannelID, resp.IpError)
}
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}











if strings.HasPrefix(m.Content, in.Prefix + "myserver") {
	pm, err := s.UserChannelCreate(m.Author.ID)
	if err == nil {
		s.ChannelMessageSend(pm.ID, "Join our server and ask for help! We're always around.\n<https://discord.gg/012s8wmCkDVgdn7yo>")
	} else {
		s.ChannelMessageSend(m.ChannelID, "You've Disabled your PM's or you have blocked me.\nI cannot send you an invite link at this time.")
	}
				if in.LogType == "master" {
			return
		}
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}
















if m.Author.ID == js.Admin && strings.HasPrefix(m.Content, in.Prefix + "avatar") {
      img, err := gto.ReadFile("test.png")
        if err != nil {
                fmt.Println(err)
        }

        base64 := base64.StdEncoding.EncodeToString(img)

        avatar := fmt.Sprintf("data:image/png;base64,%s", string(base64))

        _, err = s.UserUpdate("", "", "Echo", avatar, "")
        if err != nil {
                fmt.Println(err)
        }
}











if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "emojis off") {
	therole := GetRoleID(s, guildID, "Echo [BETA]")
	if therole != "" {
	_, err = s.GuildRoleEdit(guildID, therole, "Echo [BETA]", 0, false, 280067127)
	if err != nil {
		return
	}
	s.ChannelMessageSend(m.ChannelID, "I have disabled the emoji system.")
	}
	go CheckLog(s, m, in)
return
}


if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "emojis on") {
	therole := GetRoleID(s, guildID, "Echo [BETA]")
	if therole != "" {
	_, err = s.GuildRoleEdit(guildID, therole, "Echo [BETA]", 0, false, 334625847)
	if err != nil {
		return
	}
	s.ChannelMessageSend(m.ChannelID, "I have enabled the emoji system.")
	}
	go CheckLog(s, m, in)
return
}











if in.Owner == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + "addmaster") {
	if UseRole == false {
		s.ChannelMessageSend(m.ChannelID, "Hi! By default Echo makes anyone with the permissions `Manage Server` his master. So this command is not needed. If you want to revert back to using the Bot Commander role system type `"+in.Prefix+"userole true` and this command will be available to you.")
		return
	}

	master := ""
	if in.BotMaster != "" {
		master = in.BotMaster
	} else {
		master = "Bot Commander"
	}

/*
if BotCommander == "" {
	com, err := s.GuildRoleCreate(guildID)
	if err != nil {
		return
	}
	_, err = s.GuildRoleEdit(guildID, com.ID, master, 0, false, 36785153)
	if err != nil {
		return
	}
}
*/
str := strings.Replace(m.Content, in.Prefix+"addmaster ", "", -1)
str = strings.Replace(str, "<@", "", -1)
str = strings.Replace(str, ">", "", -1)
str = strings.Replace(str, "!", "", -1)

z, err := s.State.Member(guildID, str) 
if err != nil {
z, err = s.GuildMember(guildID, str)
	if err != nil {
		return
	}
}

		roles, err := s.GuildRoles(guildID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, resp.NoRolePerms)
			return
		}
		if err == nil {
			for _, v := range roles {
    			if v.Name == master {
    				z.Roles = append(z.Roles, v.ID)
    				s.GuildMemberEdit(guildID, str, z.Roles)
					s.ChannelTyping(m.ChannelID)
					newdata := strings.Replace(resp.AddMaster, "{user}", "<@"+str+">", -1)
					s.ChannelMessageSend(m.ChannelID, newdata)
    			}
			}
		}
		go CheckLog(s, m, in)
return
	} // end of Add master command




















if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "leave") {
	go CheckLog(s, m, in)
return
	s.ChannelMessageSend(m.ChannelID, "Thanks for having me!")
	s.GuildLeave(guildID)
}



















if in.Owner == m.Author.ID && strings.HasPrefix(m.Content, in.Prefix + "delmaster") {
	if UseRole == false {
		s.ChannelMessageSend(m.ChannelID, "Hi! By default Echo makes anyone with the permissions `Manage Server` his master. So this command is not needed. If you want to revert back to using the Bot Commander role system type `"+in.Prefix+"userole true` and this command will be available to you.")
		return
	}
	var mrcom string
	if in.BotMaster != "" {
		mrcom = GetRoleID(s, guildID, in.BotMaster)
	} else {
		mrcom = GetRoleID(s, guildID, "Bot Commander")
	}
str := strings.Replace(m.Content, in.Prefix + "delmaster ", "", -1)
str = strings.Replace(str, "<@", "", -1)
str = strings.Replace(str, ">", "", -1)
str = strings.Replace(str, "!", "", -1)

x, err := s.State.Member(guildID, str) 
if err != nil {
x, err = s.GuildMember(guildID, str)
}
	if err == nil {
	var mc []string
	mc = x.Roles
	for mr := range x.Roles {
		t := mc[mr]
		if t == mrcom {
    		// z.Roles = append(z.Roles, t[:0])
    		x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
    		if err != nil {
    			return
    		}
    		s.GuildMemberEdit(guildID, str, x.Roles)
			s.ChannelTyping(m.ChannelID)
			newdata := strings.Replace(resp.DelMaster, "{user}", "<@"+str+">", -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
		}
	}
	}// end of err check
	go CheckLog(s, m, in)
return
} // end of Del master command


















	// Take a role away from someone.
	// Let's make this command Masters only (Server Owner & Bot Commanders)
	// this command is basically the same as -give we just need to tweak a few things here and there.
	if strings.HasPrefix(m.Content, in.Prefix + "take ") {
		if BotMaster == true {
			if gto.TakeRoleCommand(s, m, in.Prefix, "take") == true {
				// it worked!
				s.ChannelMessageSend(m.ChannelID, "I have taken the role `"+gto.Role+"` from <@"+gto.User+">")
			} else {
				// an error has occured. (usually permission errors)
			//	s.ChannelMessageSend(m.ChannelID, "An error occured. Check my permissions *(Drag my role to the top to manage Admins roles)*")
			}
		} else { // the user is not my master.
			s.ChannelMessageSend(m.ChannelID, "You're not a Bot Commander.")
		}
		go CheckLog(s, m, in)
return
	} // end of -take @User Role Name command.


















	if strings.HasPrefix(m.Content, in.Prefix + "give ") {
		if BotMaster == true { // Check if the user is a master or not.
			if gto.GiveRoleCommand(s, m, in.Prefix, "give")  == true { // we're going to send the prefix & the command name we want (you can change this)
				// for example if you want this command to be called GiveIt you do gto.GiveRole(s, m, prefix, "GiveIt")
				// just be sure to change the strings.HasPrefix() line to account for the new command name.
				// it worked!
				// s.ChannelMessageSend(m.ChannelID, "I have given the user <@"+user+"> the role `"+role+"`")
				s.ChannelMessageSend(m.ChannelID, "I have given <@"+gto.User+"> the role `"+gto.Role+"`")
			} else {
				// this usually means permission problems. Let's let them know
				// Also discord changed the role system around with the Administrator entry. This requires Gotools rank to be higher
				// than theirs in order to edit them.
				s.ChannelMessageSend(m.ChannelID, "Something went wrong. *(Usually Permissions)* Check to make sure I have the appropriate permissions")
			}
		} else { // the user is not my master.
			s.ChannelMessageSend(m.ChannelID, "You're not a Bot Commander.")
		}
		go CheckLog(s, m, in)
return
	}










if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix+"addrole") {
str := strings.Replace(m.Content, in.Prefix+"addrole ", "", -1)

	if str != in.Prefix + "addrole" {
		if GetRoleID(s, guildID, str) == "" {
			com, err := s.GuildRoleCreate(guildID)
			if err != nil {
				return
			}
			_, err = s.GuildRoleEdit(guildID, com.ID, str, 0, false, 11656193)
			if err != nil {
				return
			}
				newdata := strings.Replace(resp.Addrole, "{data}", str, -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
		} else {
			s.ChannelMessageSend(m.ChannelID, resp.RoleExists)
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, resp.AddRoleError)
	}
	go CheckLog(s, m, in)
return
}










if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix+"delrole") {
str := strings.Replace(m.Content, in.Prefix+"delrole ", "", -1)
roleid := GetRoleID(s, guildID, str)
	if str != in.Prefix + "delrole" {
		if roleid != "" {
			err = s.GuildRoleDelete(guildID, roleid)
			if err != nil {
				return
			}
			newdata := strings.Replace(resp.Delrole, "{data}", str, -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
		} else {
			s.ChannelMessageSend(m.ChannelID, resp.RoleNoExist)
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, resp.AddRoleError)
	}
	go CheckLog(s, m, in)
return
}













if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "mute") {
	if GetRoleID(s, guildID, "muted") == "" {
		com, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return
		}
		_, err = s.GuildRoleEdit(guildID, com.ID, "muted", 0, false, 0)
		if err != nil {
			return
		}
	}

	str := strings.Replace(m.Content, in.Prefix + "mute ", "", -1)
	if strings.Contains(str, "!") {
		str = strings.Replace(str, "<@!", "", -1)
	} else {
		str = strings.Replace(str, "<@", "", -1)
	}
	str = strings.Replace(str, ">", "", -1)

	// fmt.Println("Adding Master: "+str)

	z, err := s.State.Member(guildID, str) 
	if err != nil {
		return
	}
	roles, err := s.State.Guild(guildID)
	if err == nil {
		for _, v := range roles.Roles {
    		if v.Name == "muted" {
    			z.Roles = append(z.Roles, v.ID)
    			s.GuildMemberEdit(guildID, str, z.Roles)
				newdata := strings.Replace(resp.Mute, "{user}", "<@" + str + ">", -1)
				newdata = strings.Replace(newdata, "{data}", "<#"+m.ChannelID+">", -1)
      			msgs, err := s.ChannelMessages(m.ChannelID, 100, "", "")
      			if err == nil {
      				var list []string
      				do := 0
        			for _, v := range msgs {
          				if str == v.Author.ID {
            				do++
            				list = append(list, v.ID)
          				}
        			}
        			if do > 0 {
         				 s.ChannelMessagesBulkDelete(m.ChannelID, list)
        			}
      			}
      			s.ChannelMessageSend(m.ChannelID, newdata)
    		}
		}
	}
	go CheckLog(s, m, in)
	return
} // end of giveme command.













if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "unmute") {
str := strings.Replace(m.Content, in.Prefix + "unmute ", "", -1)
usr := strings.Replace(str, "<@", "", -1)
usr = strings.Replace(usr, "!", "", -1)
usr = strings.Replace(usr, ">", "", -1)
	var roleID string

mroles, err := s.State.Guild(guildID)
if err == nil {

 for _, v := range mroles.Roles {
    if v.Name == "muted" {
    	roleID = v.ID
    }
  }
  }
x, err := s.State.Member(guildID, usr) 
if err != nil {
x, err = s.GuildMember(guildID, usr)
}

if err != nil {
	s.ChannelMessageSend(m.ChannelID, "You don't have the role `muted` setup in your server.")
} else {
	var ms []string
	ms = x.Roles
	for mr := range x.Roles {
		t := ms[mr]
		if t == roleID {
    		x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
    		s.GuildMemberEdit(guildID, usr, x.Roles)
			s.ChannelTyping(m.ChannelID)
			newdata := strings.Replace(resp.Unmute, "{user}", "<@"+usr+">", -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
			}
		}
	}
	go CheckLog(s, m, in)
return
} // end of giveme command.















if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "rolecolor") {
str := strings.Replace(m.Content, in.Prefix + "rolecolor ", "", -1)
var roleID string
var hoist bool
var perms int

newdata := gto.Split(str, " ")
color := newdata[0]
role := strings.Replace(str, color + " ", "", -1)
color = strings.Replace(color, "#", "", -1)

// newcolor := strconv.FormatInt(h, 16)
// fmt.Println(role)
  roles, err := s.State.Guild(guildID)
  if err == nil {
    for _, v := range roles.Roles {
      if v.Name == role {
       roleID = v.ID
       hoist = v.Hoist
       perms = v.Permissions
      }
    }
  } else {
  //	fmt.Println("s.GuildRoles is the error")

  }

var ij int
newcode, _ := strconv.ParseInt(color, 16, 0)
d := fmt.Sprintf("%d", newcode)
// fmt.Println(d)
ij, err = strconv.Atoi(d)
if err != nil {
//	fmt.Println(err)
}
// if roleID != "" {
// roleID := GetRoleID(s, guildID, role)
_, err = s.GuildRoleEdit(guildID, roleID, role, ij, hoist, perms)
if err != nil {
  	s.ChannelMessageSend(m.ChannelID, resp.NoRole)
} else {
	newdata := strings.Replace(resp.Rolecolor, "{data}", "#"+color, -1)
	s.ChannelMessageSend(m.ChannelID, newdata)
}
// }
go CheckLog(s, m, in)
return
} // end of role color
















	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "greet ") {
		str := strings.Replace(m.Content, in.Prefix + "greet ", "", -1)
		in.GreetMsg = str
		b, err := gto.Marshal(in)
		if err == nil {
			gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
		}
		if str != "off" {
			s.ChannelTyping(m.ChannelID)
			newdata := strings.Replace(resp.GreetMsg, "{data}", str, -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
		} else {
			s.ChannelTyping(m.ChannelID)
			s.ChannelMessageSend(m.ChannelID, resp.GreetOff)		
		}
		go CheckLog(s, m, in)
return
	}









	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "bye ") {
		str := strings.Replace(m.Content, in.Prefix + "bye ", "", -1)
		in.ByeMsg = str
		b, err := gto.Marshal(in)
		if err == nil {
			gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
		}
		if str != "off" {
			s.ChannelTyping(m.ChannelID)
			newdata := strings.Replace(resp.ByeMsg, "{data}", str, -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
		} else {
			s.ChannelTyping(m.ChannelID)
		
			s.ChannelMessageSend(m.ChannelID, resp.ByeOff)	
		}
		go CheckLog(s, m, in)
return
	}




















	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "mkinvite") {

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
		newdata := strings.Replace(resp.MkInvite, "{data}", iv.Code, -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
		go CheckLog(s, m, in)
return
	}




















	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "denylinks") {
		in.AntiLink = true
		b, err := gto.Marshal(in)
		if err == nil {
			gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		s.ChannelMessageSend(m.ChannelID, resp.DenyLinks)
		go CheckLog(s, m, in)
return
	}












	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "allowlinks") {
		in.AntiLink = false
		b, err := gto.Marshal(in)
		if err == nil {
			gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		s.ChannelMessageSend(m.ChannelID, resp.AllowLinks)
		go CheckLog(s, m, in)
return
	}



















	if strings.HasPrefix(m.Content, in.Prefix + "prefix ") {
		if BotMaster == false {
			s.ChannelMessageSend(m.ChannelID, "Only Bot Commanders can use this command.")
			return
		}
		str := strings.Replace(m.Content, in.Prefix + "prefix ", "", -1)
		if str == "" {
			s.ChannelMessageSend(m.ChannelID, "You need to set a prefix. `"+in.Prefix+"prefix +`")
			return
		}
		in.Prefix = str
		b, err := gto.Marshal(in)
		if err == nil {
			gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		newdata := strings.Replace(resp.Prefix, "{data}", str, -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
		go CheckLog(s, m, in)
return
	}





















	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "setpunish ") {
		str := strings.Replace(m.Content, in.Prefix + "setpunish ", "", -1)
		check := false
		if gto.ToLower(str) == "kick" {
			check = true
		}
		if gto.ToLower(str) == "ban" {
			check = true
		}
		if gto.ToLower(str) == "warn" {
			check = true
		}

		if check == true {
			in.Action = str
			b, err := gto.Marshal(in)
			if err == nil {
				gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
			}
			s.ChannelTyping(m.ChannelID)
			newdata := strings.Replace(resp.SetPunish, "{data}", str, -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
		}
		if check == false {
			s.ChannelTyping(m.ChannelID)
			s.ChannelMessageSend(m.ChannelID, resp.SetPunishError)
		}
		go CheckLog(s, m, in)
return
	}

















	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "autorole ") {
		str := strings.Replace(m.Content, in.Prefix + "autorole ", "", -1)
		var br bool
		br = false
		if gto.Split(str, " ")[0] == "-s" {
			br = true
			str = strings.Replace(m.Content, in.Prefix + "autorole -s ", "", -1)
		}

		cnt := 0
		roles, err := s.State.Guild(guildID)
		if err == nil {
			if str != "off" {
 			for _, v := range roles.Roles {
    			if v.Name == str {
    				cnt++
    				in.RoleSys = str
    				in.Silent = br
					b, err := gto.Marshal(in)
					if err == nil {
						gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
					}
					s.ChannelTyping(m.ChannelID)
					newdata := strings.Replace(resp.Autorole, "{data}", str, -1)
					s.ChannelMessageSend(m.ChannelID, newdata)
					}
				}
			} // make sure they don't want to turn the autorole off


			if str == "off" {
				cnt = 1
				in.RoleSys = "off"
				b, err := gto.Marshal(in)
				if err == nil {
					gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
				}
				s.ChannelTyping(m.ChannelID)
				s.ChannelMessageSend(m.ChannelID, resp.AutoroleOff)
			}
			if cnt == 0 {
				s.ChannelTyping(m.ChannelID)
				s.ChannelMessageSend(m.ChannelID, resp.NoRole)
			}
		}
		go CheckLog(s, m, in)
	}






















if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "botrole") {
str := strings.Replace(m.Content, in.Prefix + "botrole ", "", -1)
var br bool
br = false
if gto.Split(str, " ")[0] == "-s" {
br = true
str = strings.Replace(m.Content, in.Prefix + "botrole -s ", "", -1)
}

cnt := 0
roles, err := s.State.Guild(guildID)
if err == nil {
if str != "off" {
 for _, v := range roles.Roles {
    if v.Name == str {
    	cnt++
    	in.BotAuto = str
    	in.Silent = br
		b, err := gto.Marshal(in)
		if err == nil {
			gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		newdata := strings.Replace(resp.Botrole, "{data}", str, -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
	}
}
} // make sure they don't want to turn the autorole off


if str == "off" {
	cnt = 1
	in.BotAuto = "off"
		b, err := gto.Marshal(in)
		if err == nil {
			gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
		}
		s.ChannelTyping(m.ChannelID)
		s.ChannelMessageSend(m.ChannelID, resp.BotroleOff)
}




	if cnt == 0 {
		s.ChannelTyping(m.ChannelID)
		
		s.ChannelMessageSend(m.ChannelID, resp.NoRole)
	}
}
go CheckLog(s, m, in)
return
}














	if strings.HasPrefix(m.Content, in.Prefix + "invites") {
		o, err := s.ChannelInvites(m.ChannelID)
		if err == nil {
			data := ""
			s.ChannelTyping(m.ChannelID)
			
			s.ChannelMessageSend(m.ChannelID, "Invites for: `"+in.Name+"`\n```ruby\nGrabbing Results..```")
			for _, v := range o {
			data = data + "\nissuer: "+v.Inviter.Username+"\ncode: "+v.Code
			// s.ChannelMessageDelete(m.ChannelID, theid)
			//s.ChannelMessageEdit(m.ChannelID, m.ID, "Invites for: `"+guildID+"` "+v.Inviter.Username + "\n" + v.Code)
  		}
			s.ChannelTyping(m.ChannelID)
			
  		s.ChannelMessageSend(m.ChannelID, "```ruby\n"+data+"```")
		return
		}
					if in.LogType == "master" {
			return
		}
		if in.LogType == "all" {
			go CheckLog(s, m, in)
return
		}		
	}





















// let's see if they want advertising disabled
if in.AntiLink == true && BotMaster == false && m.Author.ID != s.State.User.ID {
	var deny [5]string
	deny[0] = "https://"
	deny[1] = "http://"
	deny[2] = ".com"
	deny[3] = ".net"
	deny[4] = "www."

	for i := 0; i <= 4; i++ {
		if strings.Contains(gto.ToLower(gto.ToLower(m.Content)), deny[i]) && in.Owner != m.Author.ID {
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			s.ChannelTyping(m.ChannelID)
			
			if in.Action == "kick" {
				newdata := strings.Replace(resp.AntiLinkKick, "{user}", "<@" + m.Author.ID + ">", -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
			s.GuildMemberDelete(guildID, m.Author.ID)
			}

			if in.Action == "ban" {
				newdata := strings.Replace(resp.AntiLinkBan, "{user}", "<@" + m.Author.ID + ">", -1)
				s.ChannelMessageSend(m.ChannelID, newdata)
				s.GuildBanCreate(guildID, m.Author.ID, 10)
			}

			if in.Action == "warn" {
				newdata := strings.Replace(resp.AntiLinkWarn, "{user}", "<@" + m.Author.ID + ">", -1)
				s.ChannelMessageSend(m.ChannelID, newdata)
			}
		}
	} // end of for loop
} // end of anti link system













	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix+"listwarns") {
		var lst map[string]interface{}
		g, err := gto.ReadFile("servers/"+guildID+"/warnings.json")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, resp.NoWarns)
			return
		}
		gto.Unmarshal(g, &lst)
		data := ""
		for k,v := range lst {
			if k != "Ignore" {
			user, err := s.User(k)
			if err != nil {
				return
			}
			dg := v.(string)
			// dg := fmt.Sprintf("%d", v)
			data = data + "user: " + user.Username + ": warnings: " + dg + "\n"
			dg = ""
			}
		}
		if len(data) > 0 {
			s.ChannelMessageSend(m.ChannelID, "```ruby\n"+data+"```")
		} else {
			s.ChannelMessageSend(m.ChannelID, resp.NoWarns)
		}
		go CheckLog(s, m, in)
return
	}




	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "delwarn") {
		str := strings.Replace(m.Content, in.Prefix + "delwarn ", "", -1)
		str = strings.Replace(str, "<@", "", -1)
		str = strings.Replace(str, "!", "", -1)
		str = strings.Replace(str, ">", "", -1)
		user, err := s.User(str)
		if err != nil {
			return
		}
		var lst map[string]interface{}
		g, err := gto.ReadFile("servers/"+guildID+"/warnings.json")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, resp.NoWarns)
			return
		}
		gto.Unmarshal(g, &lst)

			delete(lst, user.ID)
			dv, err := gto.Marshal(lst)
			if err != nil {
				return
			}
			gto.WriteFile("servers/"+guildID+"/warnings.json", dv, 0777)
			newdata := strings.Replace(resp.ResetWarns, "{user}", "<@" + str + ">", -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
			go CheckLog(s, m, in)
return
	}





	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "wipewarns") {
		err = os.Remove("servers/"+guildID+"/warnings.json")
		if err == nil {
			s.ChannelMessageSend(m.ChannelID, "I have wiped your warnings list.")
		} else {
			s.ChannelMessageSend(m.ChannelID, "Your warnings file doesn't exist.")
			return
		}
		go CheckLog(s, m, in)
return
	}









	if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "warn") { 
		str := strings.Replace(m.Content, in.Prefix + "warn ", "", -1)
		str = strings.Replace(str, "<@", "", -1)
		str = strings.Replace(str, "!", "", -1)
		str = strings.Replace(str, ">", "", -1)
		master := gto.MemberHasRole(s, guildID, str, BotCommander)
		if str == in.Owner {
			master = true
		}


	if master == false {
		if in.Warnings > 0 {
			if CheckWarn(guildID, str, in.Warnings) == true {
			s.ChannelTyping(m.ChannelID)
			if in.Action == "kick" {
			newdata := strings.Replace(resp.WarnKick, "{user}", "<@" + str + ">", -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
			s.GuildMemberDelete(guildID, str)
			}

			if in.Action == "ban" {
				newdata := strings.Replace(resp.WarnBan, "{user}", "<@" + str + ">", -1)
				s.ChannelMessageSend(m.ChannelID, newdata)
				s.GuildBanCreate(guildID, str, 10)
			}

			if in.Action == "warn" {
				newdata := strings.Replace(resp.Warn, "{user}", "<@" + str + ">", -1)
				s.ChannelMessageSend(m.ChannelID, newdata)
			}
			} else {
				newdata := strings.Replace(resp.Warn, "{user}", "<@" + str + ">", -1)
				s.ChannelMessageSend(m.ChannelID, newdata)
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, resp.SetWarningError)
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, resp.WarnCommander)
	}

	go CheckLog(s, m, in)
return
	}








if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "setwarning ") {
		str := strings.Replace(m.Content, in.Prefix + "setwarning ", "", -1)
		check := true

		vr,err := strconv.Atoi(str)
		if err != nil {
			check = false
		}
		if check == true {
			in.Warnings = vr
			b, err := gto.Marshal(in)
			if err == nil {
				gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
			}
			s.ChannelTyping(m.ChannelID)
			newdata := strings.Replace(resp.SetWarning, "{data}", str, -1)
			s.ChannelMessageSend(m.ChannelID, newdata)
		}
		if check == false {
			s.ChannelTyping(m.ChannelID)
			s.ChannelMessageSend(m.ChannelID, resp.SetWarningError)
		}
		go CheckLog(s, m, in)
return
	}







	if strings.HasPrefix(m.Content, in.Prefix + "kick ") && BotMaster == true {
		str := strings.Replace(m.Content, in.Prefix + "kick ", "", -1)
		str = strings.Replace(str, "<@", "", -1)
		str = strings.Replace(str, "!", "", -1)
		str = strings.Replace(str, ">", "", -1)
		// fmt.Println("the"+str+"string")
		s.ChannelTyping(m.ID)
		newdata := strings.Replace(resp.Kick, "{user}", "<@" + str + ">", -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
		s.GuildMemberDelete(guildID, str)
		go CheckLog(s, m, in)
return
	}






	if strings.HasPrefix(m.Content, in.Prefix + "ban ") && BotMaster == true {
		str := strings.Replace(m.Content, in.Prefix + "ban ", "", -1)
		str = strings.Replace(str, "<@", "", -1)
		str = strings.Replace(str, "!", "", -1)
		str = strings.Replace(str, ">", "", -1)
		s.ChannelTyping(m.ID)
		newdata := strings.Replace(resp.Kick, "{user}", "<@" + str + ">", -1)
		s.ChannelMessageSend(m.ChannelID, newdata)
		s.GuildBanCreate(guildID, str, 10)
		go CheckLog(s, m, in)
return
	}




if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "noimages") {
    s.ChannelPermissionSet(m.ChannelID, GetRoleID(s, guildID, "Echo [BETA]"), "text", 268516369, 49152)
    s.ChannelMessageSend(m.ChannelID, "Any image based command and Emoji's has been disabled in this channel.\nYou can re-enable by typing `"+in.Prefix+"images`")
    go CheckLog(s, m, in)
return
}



if BotMaster == true && strings.HasPrefix(m.Content, in.Prefix + "images") {
    s.ChannelPermissionSet(m.ChannelID, GetRoleID(s, guildID, "Echo [BETA]"), "text", 268565521, 0)
    s.ChannelMessageSend(m.ChannelID, "Image based command and Emoji's have been Enabled in this channel.\nYou can Disable by typing `"+in.Prefix+"noimages`")
    go CheckLog(s, m, in)
return
}








	if m.Author.ID == js.Admin && strings.HasPrefix(m.Content, in.Prefix + "status") {
		str := strings.Replace(m.Content, in.Prefix + "status ", "", -1)
		s.ChannelTyping(m.ChannelID)
		
		s.ChannelMessageSend(m.ChannelID, "Attempting status change.")
		s.UpdateStatus(0, str)
		return
	}











if strings.HasPrefix(gto.ToLower(m.Content), in.Prefix + "channelid") {
			s.ChannelTyping(m.ChannelID)
			time.Sleep(1000 * time.Millisecond)
	s.ChannelMessageSend(m.ChannelID, "<#"+m.ChannelID+"> ID: `"+m.ChannelID+"`")
		if in.LogType == "master" {
			return
		}	
		if in.LogType == "all" {
			go CheckLog(s, m, in)
return
		}
}



if strings.HasPrefix(gto.ToLower(m.Content), in.Prefix + "getid") {
	str := strings.Replace(m.Content, in.Prefix + "getid ", "", -1)
	theid := strings.Replace(str, "<@", "", -1)
	theid = strings.Replace(theid, ">", "", -1)
	s.ChannelTyping(m.ChannelID)
	time.Sleep(1000 * time.Millisecond)
	s.ChannelMessageSend(m.ChannelID, "<@"+theid+">'s ID: `"+theid+"`")
	if in.LogType == "master" {
		return
	}	
	if in.LogType == "all" {
		go CheckLog(s, m, in)
return
	}
}



} // GrabCommands function end.