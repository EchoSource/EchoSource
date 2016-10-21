package main


import (

  "time"
  "os"
  "fmt"
  "strings"
  "strconv"
  "github.com/TrevorSStone/goriot"
  "github.com/bwmarrin/discordgo"
  "gotools"
  "bytes"
  "os/exec"
  "encoding/binary"
  "io"
  "bufio"
  "net/http"
  "net/url"
)





var (
  personalkey = os.Getenv("bf4d4d68-9b08-4591-855d-6763a2ec5c62")
)



func TestSetup() {
  goriot.SetSmallRateLimit(10, 10*time.Second)
  goriot.SetLongRateLimit(500, 10*time.Minute)
}








func EchoSetup(s *discordgo.Session, m *discordgo.MessageCreate, guildID string, channelID string, URL string, in info) {
      if gto.DownloadFile("temp/"+guildID+".txt", URL) == true {
        setup, err := gto.ReadLines("temp/"+guildID+".txt")
        if err != nil {
          fmt.Print(err)
          return
        }
        cnn := 0
        for _, se := range setup {
          if strings.HasPrefix(se, "//") == false {
            if strings.HasPrefix(se, "Prefix=") {
              cnn++
              in.Prefix = strings.Split(se, "Prefix=")[1]
              str := strings.Split(se, "Prefix=")[1]
              s.ChannelMessageSend(channelID, "`Prefix Status:` Preparing to change the prefix to `"+str+"`")
              time.Sleep(10 * time.Second)
          }
            if strings.HasPrefix(se, "Autorole=") {
              cnn++
              str := strings.Split(se, "Autorole=")[1]
              in.RoleSys = str
              s.ChannelMessageSend(channelID, "`Autorole Status:` Preparing to change the Autorole to `"+str+"`")
              time.Sleep(10 * time.Second)
            }
            if strings.HasPrefix(se, "Botrole=") {
              cnn++
              str := strings.Split(se, "Botrole=")[1]
              in.BotAuto = str
              s.ChannelMessageSend(channelID, "`Botrole Status:` Preparing to change the Bot Autorole to `"+str+"`")
              time.Sleep(10 * time.Second)
            }
            if strings.HasPrefix(se, "Greet=") {
              cnn++
              str := strings.Split(se, "Greet=")[1]
              in.GreetMsg = str
              s.ChannelMessageSend(channelID, "`Greet Status:` Preparing to change the Greet to `"+str+"`")
              time.Sleep(10 * time.Second)
            }
            if strings.HasPrefix(se, "Bye=") {
              cnn++
              str := strings.Split(se, "Bye=")[1]
              in.ByeMsg = str
              s.ChannelMessageSend(channelID, "`Bye Status:` Preparing to change the Bye message to `"+str+"`")
              time.Sleep(10 * time.Second)
            }
            if strings.HasPrefix(se, "Antilink=") {
              cnn++
              str := strings.Split(se, "Antilink=")[1]
              if strings.ToLower(str) == "true" {
                in.AntiLink = true
              }
              if strings.ToLower(str) == "false" {
                in.AntiLink = false
              }
              s.ChannelMessageSend(channelID, "`Antilink Status:` Preparing to change the Antilinks to `"+str+"`")
              time.Sleep(10 * time.Second)
            }
            if strings.HasPrefix(se, "Punishment=") {
              cnn++
              str := strings.Split(se, "Punishment=")[1]
              in.Action = str
              s.ChannelMessageSend(channelID, "`Punishment Status:` Preparing to change the punishment to `"+str+"`")
              time.Sleep(3 * time.Second)
            }
          } // make sure it isn't a comment line
        }
        if cnn > 0 {
          b, err := gto.Marshal(in)
          if err == nil {
            gto.WriteFile("servers/" + guildID + "/main.json", b, 0777)
          }
          cnn = 0
          os.Remove("temp/"+guildID+".txt")
          s.ChannelMessageSend(channelID, "`Session Complete` I have imported your `setup.esf` file.")
        }
      } else {
        s.ChannelMessageSend(channelID, "There were issues downloading the file.")
      }
}







func SetupARS(s *discordgo.Session, m *discordgo.MessageCreate, guildID string, channelID string, URL string) {
  if gto.DownloadFile("temp/"+guildID+"-ars.txt", URL) == true {
    setup, err := gto.ReadLines("temp/"+guildID+"-ars.txt")
    if err != nil {
      fmt.Print(err)
      return
    }
    var ars map[string]interface{}
    file, err := gto.ReadFile("servers/"+guildID+"/autoresponse.json")
    if err != nil {
      return
    }
    gto.Unmarshal(file, &ars)
    cnn := 0

    for _, se := range setup {
      if strings.Contains(se, "=") {
        trig := strings.Split(se, "=")[0]
        response := strings.Replace(se, trig + "=", "", -1)
        trigger := trig
        if response != "" && trigger != "" {
          ars[trigger] = response
          time.Sleep(3 * time.Second)
          response = strings.Replace(response, "`", "", -1)
          s.ChannelMessageSend(channelID, "Preparing `"+trigger+"` with the response:\n`"+response+"`")
          cnn++
        }
      }
    }
    if cnn > 0 {
      b, err := gto.Marshal(ars)
      if err == nil {
        gto.WriteFile("servers/"+guildID+"/autoresponse.json", b, 0777)
        os.Remove("temp/"+guildID+"-ars.txt")
        s.ChannelMessageSend(channelID, "`Success` Added `"+strconv.Itoa(cnn)+"` triggers to your A.R.S File.")
      }
    }
  }
}




func AutoScript(s *discordgo.Session, m *discordgo.MessageCreate, guildID string, channelID string, URL string, BotMaster bool, name string, in info) {
  if gto.DownloadFile("temp/"+guildID+name+".txt", URL) == true {
    setup, err := gto.ReadLines("temp/"+guildID+name+".txt")
    if err != nil {
      fmt.Print(err)
      return
    }
    pname := ""
    pauthor := ""
    pversion := ""
    do := 0
    for _, se := range setup {
      if strings.HasPrefix(strings.ToLower(se), "[name=") == false && do == 0 {
        s.ChannelMessageSend(channelID, "`Syntax error:` You need to define the plugin name on line `1` using `[name=Plugin Name Here]`")
        return
      }

      if strings.HasPrefix(strings.ToLower(se), "[author=") == false && do == 1 {
        s.ChannelMessageSend(channelID, "`Syntax error:` You need to define the plugin author on line `2` using `[author=Your Name Here]`")
        return
      }

      if strings.HasPrefix(strings.ToLower(se), "[version=") == false && do == 2 {
        s.ChannelMessageSend(channelID, "`Syntax error:` You need to define the plugin version on line `3` using `[version=1.0.0]`")
        return
      }

      if strings.HasPrefix(strings.ToLower(se), "[name=") {
        pname = strings.Split(strings.ToLower(se), "[name=")[1]
        pname = strings.TrimSuffix(pname, "]")
      }
      if strings.HasPrefix(strings.ToLower(se), "[author=") {
        pauthor = strings.Split(strings.ToLower(se), "[author=")[1]
        pauthor = strings.TrimSuffix(pauthor, "]")
      }
      if strings.HasPrefix(strings.ToLower(se), "[version=") {
        pversion = strings.Split(strings.ToLower(se), "[version=")[1]
        pversion = strings.TrimSuffix(pversion, "]")
      }

      if strings.HasPrefix(se, "//") == false && strings.HasPrefix(strings.ToLower(se), "[name=") == false && strings.HasPrefix(strings.ToLower(se), "[author=") == false && strings.HasPrefix(strings.ToLower(se), "[version=") == false && se != "" {
        ispm := false
        master := ""
        if BotMaster == true {
          master = "True"
        } else {
          master = "False"
        }
        newresp := se
        newresp = strings.Replace(newresp, "  ", "\n", -1)
        newresp = strings.Replace(newresp, "{plugin:name}", pname, -1)
        newresp = strings.Replace(newresp, "{plugin:author}", pauthor, -1)
        newresp = strings.Replace(newresp, "{plugin:version}", pversion, -1)
        newresp = strings.Replace(newresp, "{chan}", "<#"+m.ChannelID+">", -1)
        newresp = strings.Replace(newresp, "{pref}", in.Prefix, -1)
        newresp = strings.Replace(newresp, "{greet}", in.GreetMsg, -1)
        newresp = strings.Replace(newresp, "{bye}", in.ByeMsg, -1)
        newresp = strings.Replace(newresp, "{ismaster}", master, -1)

        if strings.Contains(newresp, "{pm}") {
          ispm = true
          newresp = strings.Replace(newresp, "{pm}", "", -1)
        }
        if strings.Contains(newresp, "{sleep:") {
          newt := strings.Split(newresp, "{sleep:")[1]
          newtime := strings.Split(newt, "}")[0]
          io, err := time.ParseDuration(newtime)
          if err == nil {
            time.Sleep(io)
            newresp = strings.Replace(newresp, "{sleep:"+newtime+"}", "", -1)
          }
        } else {
          time.Sleep(3 * time.Second)
        }
        ARS(s, m, newresp, guildID, ispm, BotMaster, "", "", in)
      }
      do++
    }
    os.Remove("temp/"+guildID+name+".txt")
  }
}






func RunScript(s *discordgo.Session, m *discordgo.MessageCreate, guildID string, channelID string, path string, BotMaster bool, in info) {
    setup, err := gto.ReadLines(path)
    if err != nil {
      fmt.Print(err)
      return
    }
    pname := ""
    pauthor := ""
    pversion := ""
    do := 0
    for _, se := range setup {
      if strings.HasPrefix(strings.ToLower(se), "[name=") == false && do == 0 {
        s.ChannelMessageSend(channelID, "`Syntax error:` You need to define the plugin name on line `1` using `[name=Plugin Name Here]`")
        return
      }

      if strings.HasPrefix(strings.ToLower(se), "[author=") == false && do == 1 {
        s.ChannelMessageSend(channelID, "`Syntax error:` You need to define the plugin author on line `2` using `[author=Your Name Here]`")
        return
      }

      if strings.HasPrefix(strings.ToLower(se), "[version=") == false && do == 2 {
        s.ChannelMessageSend(channelID, "`Syntax error:` You need to define the plugin version on line `3` using `[version=1.0.0]`")
        return
      }

      if strings.HasPrefix(strings.ToLower(se), "[name=") {
        pname = strings.Split(strings.ToLower(se), "[name=")[1]
        pname = strings.TrimSuffix(pname, "]")
      }
      if strings.HasPrefix(strings.ToLower(se), "[author=") {
        pauthor = strings.Split(strings.ToLower(se), "[author=")[1]
        pauthor = strings.TrimSuffix(pauthor, "]")
      }
      if strings.HasPrefix(strings.ToLower(se), "[version=") {
        pversion = strings.Split(strings.ToLower(se), "[version=")[1]
        pversion = strings.TrimSuffix(pversion, "]")
      }

      if strings.HasPrefix(se, "//") == false && strings.HasPrefix(strings.ToLower(se), "[name=") == false && strings.HasPrefix(strings.ToLower(se), "[author=") == false && strings.HasPrefix(strings.ToLower(se), "[version=") == false && se != "" {
        ispm := false
        master := ""
        if BotMaster == true {
          master = "True"
        } else {
          master = "False"
        }
        newresp := se
        newresp = strings.Replace(newresp, "  ", "\n", -1)
        newresp = strings.Replace(newresp, "{plugin:name}", pname, -1)
        newresp = strings.Replace(newresp, "{plugin:author}", pauthor, -1)
        newresp = strings.Replace(newresp, "{plugin:version}", pversion, -1)
        newresp = strings.Replace(newresp, "{chan}", "<#"+m.ChannelID+">", -1)
        newresp = strings.Replace(newresp, "{pref}", in.Prefix, -1)
        newresp = strings.Replace(newresp, "{greet}", in.GreetMsg, -1)
        newresp = strings.Replace(newresp, "{bye}", in.ByeMsg, -1)
        newresp = strings.Replace(newresp, "{ismaster}", master, -1)

        /*
        trigger := strings.Replace(k, "&", "", -1)
        if strings.Contains(trigger, "{params}") {
        //  params = strings.Replace(m.Content, trigger, "", -1)
          params = gto.TrimPrefix(m.Content, trigger)
          trigger = strings.Replace(trigger, " {params}", "", -1)
        }
        */

        if strings.Contains(newresp, "{pm}") {
          ispm = true
          newresp = strings.Replace(newresp, "{pm}", "", -1)
        }
        if strings.Contains(newresp, "{sleep:") {
          newt := strings.Split(newresp, "{sleep:")[1]
          newtime := strings.Split(newt, "}")[0]
          io, err := time.ParseDuration(newtime)
          if err == nil {
            time.Sleep(io)
            newresp = strings.Replace(newresp, "{sleep:"+newtime+"}", "", -1)
          }
        } else {
          time.Sleep(3 * time.Second)
        }
        ARS(s, m, newresp, guildID, ispm, BotMaster, "", "", in)
      }
      do++
    }
}






func InstallScript(s *discordgo.Session, m *discordgo.MessageCreate, guildID string, channelID string, URL string) {
  if _, err := os.Stat("servers/"+guildID+"/scripts"); err != nil {
    os.Mkdir("servers/"+guildID+"/scripts", 0777)
  }
  if gto.DownloadFile("servers/"+guildID+"/scripts/install.plugin.ars", URL) == true {
    setup, err := gto.ReadLines("servers/"+guildID+"/scripts/install.plugin.ars")
    if err != nil {
      fmt.Print(err)
      return
    }
    newname := ""
    for _, pl := range setup {
      if strings.HasPrefix(pl, "[name=") {
        plug := strings.Split(pl, "[name=")[1]
        plugin := strings.Split(plug, "]")[0]
        newname = strings.Replace(plugin, " ", "", -1)
        newname = gto.CleanPath(newname)
        newname = strings.ToLower(newname)
      }
    }
    if newname != "" {
      err = os.Rename("servers/"+guildID+"/scripts/install.plugin.ars", "servers/"+guildID+"/scripts/"+newname+".plugin.ars")
      if err != nil {
        os.Remove("servers/"+guildID+"/scripts/install.plugin.ars")
        s.ChannelMessageSend(channelID, "An error occured during renaming. Make sure your `Plugin Name` is without special chars.")
        return
      } else {
        if m.Author.ID != "190255157647376384" {
          s.ChannelMessageSend(channelID, "I have installed the plugin `"+newname+"` you can view your plugins by typing `::plugins` and also run them by typing `::run plugin-name`")
        } else {
          s.ChannelMessageSend(channelID, "Installed `"+newname+"` plugin from `Echo Package Manager` type `::plugins` to view and `::run plugin-name` to run.")
        }
      }
    }
  }
}







func GiveThemAll(s *discordgo.Session, m *discordgo.MessageCreate, role string, guildID string, channelID string) {
  s.ChannelMessageSend(channelID, "Please wait while I calculate the wait time.")

  r, err := s.State.Guild(guildID)
  if err != nil {
    return
  }

  // Get total member count. Calculate the wait time.
  tc := 0
  for _ = range r.Members {
    tc++
  }

  // Search for the role. Make sure it exists.
  check := false
  roleID := ""
  for _, ro := range r.Roles {
    if ro.Name == role {
      check = true
      roleID = ro.ID
    }
  }


  if check == false {
  //  s.ChannelMessageSend(m.ChannelID, "The role doesn't exist in this server. Check your spelling.")
    return
  }


  if tc > 0 {
    tcc := strconv.Itoa(tc)
    wait := tc * 3 / 60
    wt := strconv.Itoa(wait)
    s.ChannelMessageSend(channelID, "Alright I have calculated `"+tcc+"` members in this server. ```go\nWait time: " + wt + " minutes.```")
  }
  count := 0
  ocount := 0
  mcnt := 0
  for _, m := range r.Members {
    mcnt++
    if gto.MemberHasRole(s, guildID, m.User.ID, roleID) == false {
      time.Sleep(3 * time.Second)
      x, err := s.State.Member(guildID, m.User.ID)
      if err != nil {
        return
      }

      x.Roles = append(x.Roles, roleID)
      err = s.GuildMemberEdit(guildID, m.User.ID, x.Roles)
      if err == nil {
        count++
      }
    } else {
      ocount++
    } // make sure they don't have the role
  }


  if count > 0 {
    cnt := strconv.Itoa(count)
    ocnt := strconv.Itoa(ocount)
    if ocount > 0 {
      s.ChannelMessageSend(channelID, "I have given the role to `"+cnt+"` members. I have passed over `"+ocnt+"` members that already had the role.")
    } else {
      s.ChannelMessageSend(channelID, "I have given the role to `"+cnt+"` members.")
    }
  } else {
    s.ChannelMessageSend(channelID, "Something has happened. an Error occured.")
  }
}







func TakeThemAll(s *discordgo.Session, m *discordgo.MessageCreate, role string, guildID string, channelID string) {

  r, err := s.State.Guild(guildID)
  if err != nil {
    return
  }

  // Get total member count. Calculate the wait time.
  tc := 0
  for _ = range r.Members {
    tc++
  }

  check := false
  roleID := ""

  for _, ro := range r.Roles {
    if ro.Name == role {
      check = true
      roleID = ro.ID
    }
  }


  if check == false {
    s.ChannelMessageSend(m.ChannelID, "The role doesn't exist in this server. Check your spelling.")
    return
  }

  if tc > 0 {
    tcc := strconv.Itoa(tc)
    wait := tc * 3 / 60
    wt := strconv.Itoa(wait)
    s.ChannelMessageSend(channelID, "Alright I have calculated `"+tcc+"` members in this server. ```go\nWait time: " + wt + " minutes.```")
  }

  count := 0
  ocount := 0
  for _, m := range r.Members {
    time.Sleep(3 * time.Second)
        x, err := s.State.Member(guildID, m.User.ID)
        if err != nil {
          return
        }

        for k := len(x.Roles) - 1; k >= 0; k-- {
          if x.Roles[k] == roleID {
              x.Roles = append(x.Roles[:k], x.Roles[k+1:]...)
              err = s.GuildMemberEdit(guildID, m.User.ID, x.Roles)
              if err == nil {
                count++
              }
          }
      }
  }

  if count > 0 {
    cnt := strconv.Itoa(count)
    ocnt := strconv.Itoa(ocount)
    if ocount > 0 {
      s.ChannelMessageSend(m.ChannelID, "I have taken the role from `"+cnt+"` members. I have passed over `"+ocnt+"` members that didn't have the role to begin with.")
    } else {
      s.ChannelMessageSend(m.ChannelID, "I have taken the role from `"+cnt+"` members.")
    }
  } else {
    s.ChannelMessageSend(m.ChannelID, "Something has happened. an Error occured.")
  }
}



type blank struct {
  Ignore    string
}





func LeaveBeep() {
  // beeps 2 times with a low frequency.
  <-time.After(1000 * time.Millisecond)
    exec.Command("beep", "-c", "2", "-f", "600", "-v", "1").Run()
    return
}


func PlayBeep() {
  // beeps 1 time with a high frequency
  <-time.After(1000 * time.Millisecond)
    exec.Command("beep", "-c", "1", "-f", "400", "-v", "1").Run()
    return
}


func JoinBeep() {
  // beeps 2 times with a low frequency.
    <-time.After(1000 * time.Millisecond)
    exec.Command("beep", "-c", "3", "-f", "800", "-v", "1").Run()
    return
}







func AlertUpdates(s *discordgo.Session, message string) {
  var updates map[string]string
  rd1, err := gto.ReadFile("updates.json")
  if err == nil {
    gto.Unmarshal(rd1, &updates)
    for _, v := range updates {
      s.ChannelMessageSend(v, "```AutoIt\n"+message+"```")
      time.Sleep(1000 * time.Millisecond)
    }
  }
}









func AddQueue(thesong string, guildID string, requester string) error {
  var queue map[string]string

  file, err := gto.ReadFile("queue/"+guildID+".json")
  if err != nil {
    temp, err := gto.ReadFile("queue_template.json")
    if err == nil {
      gto.WriteFile("queue/"+guildID+".json", temp, 0777)
      newfile, err := gto.ReadFile("queue/"+guildID+".json")
      if err != nil {
        return err
      }
      gto.Unmarshal(newfile, &queue)
    }
    // the queue file doesn't exist. so make it!
  }
  gto.Unmarshal(file, &queue)
  queue[thesong] = requester
  i, err := gto.Marshal(queue)
  if err != nil {
    return err
  } else {
    gto.WriteFile("queue/"+guildID+".json", i, 0777)
  }
  return nil
}




func DeleteQueue(thesong string, guildID string) error {
  var queue map[string]string
  file, err := gto.ReadFile("queue/"+guildID+".json")
  if err != nil {
    // the queue file doesn't exist.
    return err
  }

  gto.Unmarshal(file, &queue)
  delete(queue, thesong)

  i, err := gto.Marshal(queue)
  if err != nil {
    return err
  } else {
    gto.WriteFile("queue/"+guildID+".json", i, 0777)
  }
  return nil
}







func QueueExists(thesong string, guildID string) bool {
  var queue map[string]string
  file, err := gto.ReadFile("queue/"+guildID+".json")
  if err != nil {
    // the queue file doesn't exist.
    return false
  }

  gto.Unmarshal(file, &queue)
  cnt := 0
  for k1,_ := range queue {
    if k1 == thesong {
      cnt++
      return true
    }
  }
  if cnt == 0 {
    return false
  }
  return false
}






func ClearQueue(guildID string) bool {
  var queue map[string]string
  file, err := gto.ReadFile("queue/"+guildID+".json")
  if err != nil {
    temp, err := gto.ReadFile("queue_template.json")
    if err == nil {
      gto.WriteFile("queue/"+guildID+".json", temp, 0777)
      newfile, err := gto.ReadFile("queue/"+guildID+".json")
      if err != nil {
        return false
      }
      gto.Unmarshal(newfile, &queue)
    }
  }
  gto.Unmarshal(file, &queue)
  se := 0
  for k1, _ := range queue {
    se++
    delete(queue, k1)
  }

  io, err := gto.Marshal(queue)
  if err != nil {
    return false
  }
  gto.WriteFile("queue/"+guildID+".json", io, 0777)
  return true
}






func isConverting(m *discordgo.MessageCreate) bool {

  io, err := gto.ReadDir(".", "*")
  if err != nil {
    return false
  }
  chk := 0
  for _, v := range io {
    if strings.Contains(v, m.Author.ID) {
      chk++
    }
  }

  if chk >= 4 {
    return true
  } else {
    return false
  }

  // check for MP3
  if _, err := os.Stat(m.Author.ID + ".mp3"); err == nil {
    return true
  } else {
    return false
  }
  // check for webm
  if _, err := os.Stat(m.Author.ID + ".webm"); err == nil {
    return true
  } else {
    return false
  }

  // check for m4a
  if _, err := os.Stat(m.Author.ID + ".m4a"); err == nil {
    return true
  } else {
    return false
  }

  // check for dca
  if _, err := os.Stat(m.Author.ID + ".dca"); err == nil {
    return true
  } else {
    return false
  }
  return false
}






func CountQueue(guildID string) int {
  var queue map[string]string
  file, err := gto.ReadFile("queue/"+guildID+".json")
  if err != nil {
    temp, err := gto.ReadFile("queue_template.json")
    if err == nil {
      gto.WriteFile("queue/"+guildID+".json", temp, 0777)
      newfile, err := gto.ReadFile("queue/"+guildID+".json")
      if err != nil {
        return 0
      }
      gto.Unmarshal(newfile, &queue)
    }
  }
  gto.Unmarshal(file, &queue)
  se := 0
  for _ = range queue {
    se++
  }
  return se
}






func isPlaying(guildID string) bool {
  if _, ok := playing[guildID]; ok {
    return true
  } else {
    return false
  }

  return false
}



func SetPlaying(guildID string, thesong string) error {
  playing[guildID] = thesong
  return nil
}



/*
func SaveData() {
  save:
  <-time.After(5 * time.Minute)
  io, err := gto.Marshal(bot)
  if err == nil {
    for _, v := range bot {
      for k, _ := range v.Messages {
        v.Messages = append(v.Messages[:k], v.Messages[k+1:]...)
      }
    }
    gto.WriteFile("System/database.json", io, 0777)
  }
  goto save
}
*/



func Convert(s *discordgo.Session, m *discordgo.MessageCreate, song string) {

  var Title string
  Title = ""
  var Duration int
  Duration = 0
  code := strconv.FormatInt(time.Now().Unix(), 10)

  // var ss map[string]interface{}

  vd := upsong{
    AddedBy:  "",
    ID:     "",
    Title:    "",
    Description:  "",
    FullTitle:    "",
    Thumbnail:    "",
    URL:    "",
    Duration: 0,
    Remaining:  0,
  }

  startAvg := time.Now().UnixNano()


  guildID, err := gto.ServerID(s, m)
  if err != nil {
    // cant find the server.
    s.ChannelMessageSend(m.ChannelID, "You need to be in a server to use this command.")
    return
  }

  // youtube-dl https://www.youtube.com/watch?v=1Vq_WDf6NT0 -v -f bestaudio -o Convert/%(title)s.%(ext)s --write-info-json
  cmd := exec.Command("youtube-dl", "--dump-json", "--default-search", "ytsearch", "--no-playlist", "--no-check-certificate", song)
  output, err := cmd.StdoutPipe()
  if err != nil {
    // log.Println(err)
    gto.Println(err)
    return
  }

  err = cmd.Start()
  if err != nil {
    gto.Println(err)
  //  log.Println(err)
    return
  }
  defer func() {
    go cmd.Wait()
  }()

  scanner := bufio.NewScanner(output)

  for scanner.Scan() {
    // s := song{}
    err = gto.Unmarshal(scanner.Bytes(), &vd)
    if err != nil {
      //log.Println(err)
      gto.Println(err)
      continue
    }

  }

  Title = vd.Title
  Duration = vd.Duration

  gto.Print(Duration)

    // DEBUGGING PURPOSE: REMOVE WHEN LIVE.
    if ConvertTime == "" {
      ConvertTime = "Not Available"
    }

    s.ChannelMessageSend(m.ChannelID, "Preparing: `"+Title+"`\nAverage Wait: `"+ConvertTime+"`")
    // song doesn't exist in my library so add it.

    // saves the song as it's original titl (causes issues cleaning)
    // exec.Command("youtube-dl", song, "-f", "bestaudio", "-o", "%(title)s.%(ext)s").Run()
    exec.Command("youtube-dl", song, "--no-playlist", "--default-search", "ytsearch", "--no-check-certificate", "-f", "bestaudio", "-o", m.Author.ID + code + ".%(ext)s").Run()


    newdata := ""
      newname := ""
      newname = gto.CleanPath(Title)

      ext := ".webm"
      if _, err := os.Stat(m.Author.ID+code+".webm"); err == nil {
      //  os.Rename(ss.Title+".webm", newname + ".webm")
        ext = ".webm"
      }
      if _, err := os.Stat(m.Author.ID+code+".m4a"); err == nil {
        //os.Rename(ss.Title+".m4a", newname + ".m4a")
        ext = ".m4a"
      }
      if _, err := os.Stat(m.Author.ID+code+".mp3"); err == nil {
        //os.Rename(ss.Title+".m4a", newname + ".m4a")
        ext = ".mp3"
      }
      newdata = "dca-rs.exe -i \"" + m.Author.ID + code + ext + "\" --raw > \"" + m.Author.ID + code + ".dca" + "\""
      Title = newname



      doit := []byte(newdata)
      gto.WriteFile(m.Author.ID+code+".bat", doit, 0777)
      exec.Command(m.Author.ID+code+".bat").Run()


      os.Remove(m.Author.ID + code + ".webm")
      os.Remove(m.Author.ID + code + ".m4a")
      os.Remove(m.Author.ID + code + ".mp3")

      t := time.Unix(0, startAvg)
      elapsed := time.Since(t)

    //  fmt.Printf("Elapsed time: %.2f hours\n", elapsed.Hours())
      upt := fmt.Sprintf("%s", elapsed)
      upti := gto.Split(upt, ".")
      uptime := upti[0]
      ConvertTime = uptime + "s"

      file, err := os.Open(m.Author.ID + code + ".dca")
      if err != nil {
        s.ChannelMessageSend(m.ChannelID, "Something happened to your song. Most likely corruption during the conversion process.")
        file.Close()
        os.Remove(m.Author.ID + code + ".bat")
        Title = ""
        return
      }
      fi, err := file.Stat()
      if err != nil {
        s.ChannelMessageSend(m.ChannelID, "Sorry I had problems converting the file. Try again or with a different version of the song.")
        file.Close()
        os.Remove(m.Author.ID + code + ".bat")
        Title = ""
        return
      }

      file.Close()
      if fi.Size() > 0 {

        // the download worked now let's do some work.
        os.Rename(m.Author.ID + code + ".dca", "music/"+m.Author.ID+" "+Title+".dca")
        os.Remove(m.Author.ID + code + ".dca")
        // s.ChannelMessageSend(m.ChannelID, ":notes: `"+ss.Title+"`")
        os.Remove(m.Author.ID + code + ".bat")
        go MusicSystem(s, m, guildID, m.Author.ID+" "+Title, Title)

      } else {
        os.Remove(m.Author.ID + code + ".dca")
        s.ChannelMessageSend(m.ChannelID, "There was an error trying to grab your file. Try again later.")
        os.Remove(m.Author.ID + code + ".bat")

      } // make sure the file worked.
    // Now the file has been created we need to convert it to .dca and than move it to the music folder.
      Title = ""
}






// loadSound attempts to load an encoded sound file from disk.
func loadMusic(path string) error {
  mbuffer = nil
  file, err := os.Open(path)

  if err != nil {
    gto.Println("Error opening dca file :", err)
    return err
  }

  var opuslen int16

  for {
    // Read opus frame length from dca file.
    err = binary.Read(file, binary.LittleEndian, &opuslen)

    // If this is the end of the file, just return.
    if err == io.EOF || err == io.ErrUnexpectedEOF {
      return nil
    }

    if err != nil {
      gto.Println("Error reading from dca file :", err)
      return err
    }

    // Read encoded pcm from dca file.
    InBuf := make([]byte, opuslen)
    err = binary.Read(file, binary.LittleEndian, &InBuf)

    // Should not be any end of file errors
    if err != nil {
      gto.Println("Error reading from dca file :", err)
      return err
    }

    // Append encoded pcm data to the buffer.
    mbuffer = append(mbuffer, InBuf)
  }
}





func MusicSystem(s *discordgo.Session, m *discordgo.MessageCreate, guildID string, song string, raw string) error {
  // let's try to have the file Autoplay after they add it.
  g, err := s.State.Guild(guildID)
  if err != nil {
    // Could not find guild.
    return err
  }

  if isPlaying(guildID) == false {
        SetPlaying(guildID, song) // set's the playing.json file for the server.
        s.ChannelMessageSend(m.ChannelID, "I'm now playing: `"+raw+"`")

    for _, vs := range g.VoiceStates {
      if vs.UserID == m.Author.ID {
        err = loadMusic("music/" + song + ".dca")
        if err != nil {
          s.ChannelMessageSend(m.ChannelID, "`LOADING: 404` - File doesn't exist or corrupted data.")
          return err
        }
        err = playMusic(s, m, guildID, vs.ChannelID, song)
        if err != nil {
          s.ChannelMessageSend(m.ChannelID, "`SOUND: 404` - Possibly corrupted data.")
          return err
        }
      }
    }
  } else {
    if QueueExists(song, guildID) == false {
      AddQueue(song, guildID, m.Author.ID)
      s.ChannelMessageSend(m.ChannelID, "The song has been added to your Queue.")
    } else {
      s.ChannelMessageSend(m.ChannelID, "The song is already in your queue.")
    }
  }
  return nil
}







// playSound plays the current buffer to the provided channel.
func playMusic(s *discordgo.Session, m *discordgo.MessageCreate, guildID, channelID string, thesong string) (err error) {
  // Join the provided voice channel.
  vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
  if err != nil {
    return err
  }

  vc.LogLevel = discordgo.LogInformational
  if isBeep == true {
    go PlayBeep()
  }
  // Sleep for a specified amount of time before playing the sound
  time.Sleep(250 * time.Millisecond)

  // Start speaking.
  vc.Speaking(true)

  playing[guildID] = thesong


  // Send the buffer data.
  for _, buff := range mbuffer {
    vc.OpusSend <- buff
  }

  // Stop speaking
  vc.Speaking(false)


  // delete the json key for the guild. because he's done playing.
  for k1, _ := range playing {
    if k1 == guildID {
      delete(playing, k1)
    }
  }
    // vc.OpusSend = nil
  time.Sleep(1000 * time.Millisecond)
  os.Remove("music/"+thesong+".dca")


  var queue map[string]string

  qu, err := gto.ReadFile("queue/"+guildID+".json")
  if err != nil {
    return err
  }
  gto.Unmarshal(qu, &queue)


  var next string
  var by string
  var cd int
  cd = 0
  que := ""
  cn := 0
  for k1,_ := range queue {
    cn++
    que = que + k1 + "\n"
  }


  for k, v := range queue {
    cd++
    if cd == 1 {
      next = k
      by = v
      by = strings.Replace(by, "a", "", -1)
      by = strings.Replace(by, "s", "", -1)
      delete(queue, k)
      bd, err := gto.Marshal(queue)
      if err != nil {
        return err
      }
      gto.WriteFile("queue/"+guildID+".json", bd, 0777)
    }
  }

  // Disconnect from the provided voice channel.
  if cd < 1 {
    vc.Disconnect()
    s.ChannelMessageSend(m.ChannelID, "`Echo 103.1 FM`\nMusic session complete.")

  }
    mbuffer = nil
    // vc.OpusSend = nil
  time.Sleep(1000 * time.Millisecond)
  onext := ""
  if cd > 0 {
    if strings.Contains(next, " ") {
      net := gto.Split(next, " ")[0]
      onext = strings.Replace(next, net, "", -1)
    }
    s.ChannelMessageSend(m.ChannelID, "`Echo 103.1 FM`\nNow playing: `"+onext+"`")
    // they have songs queued up. let's play the first one and than delete it from the queue.
    err = loadMusic("music/" + next + ".dca")
    if err != nil {
      return nil
    }
    err = playMusic(s, m, guildID, channelID, next)
    if err != nil {
      return nil
    }
  }


  return nil
}



// END OF YOUTUBE STUFF





func task(s *discordgo.Session) {
  d:
  <-time.After(2 * time.Minute)
  ef, err := gto.ReadFile("info.json")
  if err == nil {
    gto.Unmarshal(ef, &stats)
  }
    cnt := 0
    cnt = gto.CountLines("System/status.txt")
    myrand := gto.Random(1, cnt)
    st, err := gto.ReadLines("System/status.txt")
    if err == nil {
   //   fmt.Println("2 Minute State Change =>"+st[myrand])
      if _, err := os.Stat("new.exe"); err == nil {
        s.UpdateStatus(0, "Restarting for Updates!")
      fmt.Println("Found new.exe")
      os.Exit(1)
    } else {
      s.UpdateStatus(0, st[myrand])
      var inf obj
      v, err := gto.ReadFile("config.json")
      if err != nil {
        return
      }
      gto.Unmarshal(v, &inf)
      cmds := strconv.Itoa(inf.CmdsRun)
      servercount := stats.Servercount
      // fmt.Println(len(servercount))
      bi := strconv.Itoa(servercount)
      membercnt := strconv.Itoa(stats.Membercount)
      rolecnt := strconv.Itoa(stats.Rolecount)
      channelcnt := strconv.Itoa(stats.Channelcount)
      ars := strconv.Itoa(stats.ARS)
      emoji := strconv.Itoa(stats.Emojis)
      go http.PostForm("http://echobot.tk/info.php", url.Values{"s": {bi}, "c": {cmds}, "a": {ars}, "m": {membercnt}, "ch": {channelcnt}, "r": {rolecnt}, "e": {emoji}})
    }
    }
  io, err := gto.Marshal(bot)
  if err == nil {
    for _, v := range bot {
      v.Messages = nil
    }
    gto.WriteFile("System/database.json", io, 0777)
  }
  for _, gg := range bot {
    for k := len(gg.Mentions) - 1; k >= 0; k-- {
      v := gg.Mentions[k].TimeStamp
      if v <= time.Now().Unix() - (60 * 60 * 24) {
          gg.Mentions = append(gg.Mentions[:k], gg.Mentions[k+1:]...)
        }
    }
  }
    goto d
}



func Tweet(message string, url url.Values) {
  gto.ConsumerKey("")
    gto.ConsumerSecret("")
  api := gto.TwitterApi("", "")
  // api.PostTweet(str, url.Values{"status": {str}})
  api.PostTweet(message, url)
}







// loadSound attempts to load an encoded sound file from disk.
func loadSound(path string) error {
  buffer = nil
  file, err := os.Open(path)

  if err != nil {
    fmt.Println("Error opening dca file :", err)
    return err
  }

  var opuslen int16

  for {
    // Read opus frame length from dca file.
    err = binary.Read(file, binary.LittleEndian, &opuslen)

    // If this is the end of the file, just return.
    if err == io.EOF || err == io.ErrUnexpectedEOF {
      return nil
    }

    if err != nil {
      fmt.Println("Error reading from dca file :", err)
      return err
    }

    // Read encoded pcm from dca file.
    InBuf := make([]byte, opuslen)
    err = binary.Read(file, binary.LittleEndian, &InBuf)

    // Should not be any end of file errors
    if err != nil {
      fmt.Println("Error reading from dca file :", err)
      return err
    }

    // Append encoded pcm data to the buffer.
    buffer = append(buffer, InBuf)
  }
}

// playSound plays the current buffer to the provided channel.
func playSound(s *discordgo.Session, guildID, channelID string) (err error) {
  // Join the provided voice channel.
  vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
  if err != nil {
    return err
  }

  // Sleep for a specified amount of time before playing the sound
  time.Sleep(250 * time.Millisecond)

  // Start speaking.
  vc.Speaking(true)

  // Send the buffer data.
  for _, buff := range buffer {
    vc.OpusSend <- buff
  }

  // Stop speaking
  vc.Speaking(false)

  // Sleep for a specificed amount of time before ending.
  time.Sleep(250 * time.Millisecond)

  // Disconnect from the provided voice channel.
  vc.Disconnect()
  buffer = nil
  vc.OpusSend = nil

  return nil
}








func requests(opt string, m *discordgo.MessageCreate, js *obj) {
  if opt == "commands" {
    js.CmdsRun++
    newConf := obj{
      Bot:      js.Bot,
      Admin:      js.Admin,
      Status:     js.Status,
      CmdsRun:    js.CmdsRun,
      BotMaster:    false,
      BotCommander: "",
    }
    b, err := gto.Marshal(newConf)
    if err == nil {
      gto.WriteFile("config.json", b, 0777)
    }
  }

  if opt == "emojis" {
    var sys map[string]int
    b, err := gto.ReadFile("info.json")
    if err == nil {
      gto.Unmarshal(b, &sys)
      if sys["Emojis"] > 0 {
        sys["Emojis"]++
      } else {
        sys["Emojis"] = stats.Emojis
      }

      d, err := gto.Marshal(sys)
      if err == nil {
        gto.WriteFile("info.json", d, 0777)
      }
    }
  }


  if opt == "ars" {
    /*
    var sys map[string]int
    b, err := gto.ReadFile("info.json")
    if err == nil {
      gto.Unmarshal(b, &sys)

      if sys["ARS"] > 0 {
        sys["ARS"]++
      } else {
        sys["ARS"] = stats.ARS
      }
      d, err := gto.Marshal(sys)
      if err == nil {
        gto.WriteFile("info.json", d, 0777)
      }
    }
    */
  }
}








func InfoCheck(s *discordgo.Session) {
  doit:
  <-time.After(15 * time.Minute)
  var newj map[string]interface{}
  servercount := 0
  rolecount := 0
  channelcount := 0
  membercount := 0

  var updates map[string]string
  file, err := gto.ReadFile("updates.json")
  if err != nil {
  // the queue file doesn't exist.
    return
  }
  gto.Unmarshal(file, &updates)

  g, err := s.UserGuilds()
  if err == nil {
    for _, v := range g {
      servercount++
      r, err := s.State.Guild(v.ID)
      if err == nil {
        for _ = range r.Roles {
          rolecount++
        }

        for _, v2 := range r.Channels {
          channelcount++
          if v2.Name == "echo-updates" {
            updates[v.ID] = v2.ID
          }
        }

        for _ = range r.Members {
          membercount++
        }
      }
    }

    i, err := gto.Marshal(updates)
    if err != nil {
      return
    } else {
      gto.WriteFile("updates.json", i, 0777)
    }


    file, err := gto.ReadFile("info.json")
    if err == nil {
      gto.Unmarshal(file, &newj)
      if _, ok := newj["Servercount"]; ok {
        newj["Servercount"] = servercount
      }
      if _, ok := newj["Rolecount"]; ok {
        newj["Rolecount"] = rolecount
      }
      if _, ok := newj["Channelcount"]; ok {
        newj["Channelcount"] = channelcount
      }
      if _, ok := newj["Membercount"]; ok {
        newj["Membercount"] = membercount
      }

      t, err := gto.Marshal(newj)
      if err == nil {
        gto.WriteFile("info.json", t, 0777)
      }
    }
  }
  ef, err := gto.ReadFile("info.json")
  if err == nil {
    gto.Unmarshal(ef, &stats)
  }
  /*
  fmt.Print("Servers: " + strconv.Itoa(servercount) + "\n")
  fmt.Print("Roles: " + strconv.Itoa(rolecount) + "\n")
  fmt.Print("Channels: " + strconv.Itoa(channelcount) + "\n")
  fmt.Print("Members: " + strconv.Itoa(membercount) + "\n")
  */
  goto doit
}




















func TwitterUpdate() {
  do:
  <-time.After(240 * time.Minute)
  ef, err := gto.ReadFile("info.json")
  if err == nil {
    gto.Unmarshal(ef, &stats)
  }
  servers := strconv.Itoa(stats.Servercount)
  members := strconv.Itoa(stats.Membercount)
  channels := strconv.Itoa(stats.Channelcount)
  roles := strconv.Itoa(stats.Rolecount)
  ARS := strconv.Itoa(stats.ARS)
  emojis := strconv.Itoa(stats.Emojis)
  tweetit := "Let's see how I'm doing:\nServers: " + servers + "\nMembers: " + members + "\nChannels: " + channels + "\nRoles: " + roles + "\nARS Requests: " + ARS + "\nEmoji's Requested: " + emojis + "\n\n#EchoStatus"
  Tweet(tweetit, url.Values{"status": {tweetit}})
  goto do
}






func TwitterHelp() {
  doo:
  tweetit := ""
  <-time.After(30 * time.Minute)
    var status []string
    cnt := 0
    cnt = gto.CountLines("System/auto_tweet.txt")
    myrand := gto.Random(1, cnt)
    status, err := gto.ReadLines("System/auto_tweet.txt")
    if err != nil {
      fmt.Println("Twitter Auto Tweet Error:")
      fmt.Println(err)
      return
    }

    tweetit = status[myrand]
    Tweet(tweetit, url.Values{"status": {tweetit}})


    goto doo
}




func ResetPlaying() {
  for k, v := range playing {
    delete(playing, k)
    os.Remove("music/"+v+".dca")
  }
}




func IsManager(s *discordgo.Session, GuildID string, AuthorID string) bool {
z, err := s.State.Member(GuildID, AuthorID)
if err != nil {
z, err = s.GuildMember(GuildID, AuthorID)
}
// }

if err == nil {
  var l []string
  l = z.Roles
  for r := range z.Roles {
    if CheckManager(s, GuildID, l[r]) == true {
      return true
    }
  }
} // end of err == nil
  return false
}






func CreateServerFiles(s *discordgo.Session, owner string, GuildID string, name string) {
if _, err := os.Stat("servers/"+GuildID+"/main.json"); err == nil {
  return
}

  var bn info
  os.Mkdir("servers/"+GuildID, 0777)
  // gto.CopyFile("templates/main.json", "servers/"+GuildID+"/main.json")
  rd1, err := gto.ReadFile("templates/main.json")
  if err == nil {
    gto.WriteFile("servers/"+GuildID+"/main.json", rd1, 0777)
  }

  time.Sleep(1000 * time.Millisecond)
  bu, err := gto.ReadFile("servers/"+GuildID+"/main.json")
  if err == nil {
    gto.Unmarshal(bu, &bn)
    bn.Prefix = "--"
    bn.GreetMsg = ""
    bn.RoleSys =  ""
    bn.ByeMsg = ""
    bn.Owner = owner
    bn.Name = name
    bn.AntiLink = false
    bn.Action = "kick"
    bn.Silent = false
    bn.BotAuto =  ""
    bn.BotMaster = ""
    bn.AutoPerms = false
    bn.Warnings = 0
    bn.Password = ""
    bn.Pulse = 0
    jf, err := gto.Marshal(bn)
    if err != nil {
    //  return
    }
    gto.WriteFile("servers/"+GuildID+"/main.json", jf, 0777)
    gto.Print("FIX [Joined: " + name + "\nOwner: " + owner + "]")
  }
}




func Register(s *discordgo.Session, m *discordgo.MessageCreate) {
  // grab the server id c.GuildID
  c, err := s.State.Channel(m.ChannelID)
  if err != nil {
    c, err = s.Channel(m.ChannelID)
    // let's try to prevent any errors from happening. hahaha....
    if err != nil {
      return
    }
  }


  // grab the server information.
  g, err := s.State.Guild(c.GuildID)
  if err != nil {
    // Could not find guild.
    return
  }

  // check and make sure the server already exists in my collection.
  if _, ok := bot[c.GuildID]; ok {
  //  fmt.Print("The server doesn't exist in my records yet.")
    return
  }


  // collect server owner's name
  owner, err := s.User(g.OwnerID)
  if err != nil {
    // fmt.Print("Error grabbing the server owner")
    return
  }


// Create a new Info pointer.
  info := &Info{
    ServerID: g.ID,
    OwnerID: g.OwnerID,
    OwnerUser: owner.Username,
    TimeStamp:  time.Now().Unix(),
  }

  // Add our Info object to the bot map.
  bot[c.GuildID] = info
}






func Task(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID == s.State.User.ID {
    return
  }
  // grab the server id c.GuildID
  c, err := s.State.Channel(m.ChannelID)
  if err != nil {
    c, err = s.Channel(m.ChannelID)
    // let's try to prevent any errors from happening. hahaha....
    if err != nil {
      return
    }
  }



  if _, ok := bot[c.GuildID]; ok == false {
    return
  }


  // let's handle mentions and store the data.
  mention := "" // initiate the variable: mention
  if strings.Contains(m.Content, "<@") {
    me := strings.Split(m.Content, "<@")[1]
    mention = strings.Split(me, ">")[0]
    mention = strings.Replace(mention, "!", "", -1)
    clean := strings.Replace(m.Content, "<@"+mention+">", "", -1)
    clean = strings.Replace(clean, "<@!"+mention+">", "", -1)

    men := &Mentions{
      ByID:   m.Author.ID,
      Mentioned:  mention,
      ByUser:   m.Author.Username,
      Content:  clean,
      TimeStamp:  time.Now().Unix(),
    }
    bot[c.GuildID].Mentions = append(bot[c.GuildID].Mentions, men)
  }


  work := true
  for _, v := range bot[c.GuildID].Users {
    if v.ID == m.Author.ID {
      v.LastSeen = time.Now()
      v.TotalMsg++
      work = false
    }
  }

  if work == true {
    usr := &Users{
      ID:   m.Author.ID,
      Username: m.Author.Username,
      LastSeen: time.Now(),
      LastStamp:  time.Now().Unix(),
      TotalMsg: 0,
    }
    bot[c.GuildID].Users = append(bot[c.GuildID].Users, usr)
  }


  // Create a new Message pointer.
  msg := &Message{
    ID:     m.ID,
    Author:   m.Author.ID,
    Channel:  m.ChannelID,
    Content:  m.Content,
    TimeStamp:  time.Now().Unix(),
  }

  // Add this Message to our Info object.
  bot[c.GuildID].Messages = append(bot[c.GuildID].Messages, msg)
}







func AddChannelLimit(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID == s.State.User.ID {
    return
  }
  // grab the server id c.GuildID
  c, err := s.State.Channel(m.ChannelID)
  if err != nil {
    c, err = s.Channel(m.ChannelID)
    // let's try to prevent any errors from happening. hahaha....
    if err != nil {
      return
    }
  }

  if _, ok := bot[c.GuildID]; ok == false {
    fmt.Print("The server doesn't exist in my records yet.")
    return
  }

  ch := &Channel{
    ID:   m.ChannelID,
  }
  bot[c.GuildID].Restricted = append(bot[c.GuildID].Restricted, ch)

}



func CheckManager(s *discordgo.Session, GuildID string, therole string) bool {
  // perms = discordgo.PermissionManageServer + discordgo.PermissionManageRoles + discordgo.PermissionManageChannels
    x, err := s.State.Guild(GuildID)
    if err != nil {
        return false
    } else {
      for _, v := range x.Roles {
        if v.ID == therole {
          if (v.Permissions & discordgo.PermissionManageServer) > 0 {
            return true
          }
        }
      }
      return false
    }
}




func TimeDelete(s *discordgo.Session, channelID, msgID string, thet time.Duration) {
  <-time.After(thet)
  s.ChannelMessageDelete(channelID, msgID)
}






func Emojis(s *discordgo.Session, m *discordgo.MessageCreate) {
// ######################## NSFW COMMANDS HERE!!!!! ############################
  look := gto.Split(m.Content, " ")
  if len(look) > 1 {
  vk := 0
    for _, v := range look {
      v = gto.Replace(v, "_", " ", -1)
      if gto.HasPrefix(v, ":") && gto.HasSuffix(v, ":") {

        content := gto.Replace(v, ":", "", -1)
        kappa, _ := gto.ReadDir("images/emotes", "*")
        for _, f := range kappa {
          emoji := gto.ToLower(f)
          emoji = gto.Replace(emoji, ".jpg", "", -1)
          emoji = gto.Replace(emoji, ".png", "", -1)
          emoji = gto.Replace(emoji, "images\\emotes\\", "", -1)
          if gto.ToLower(content) == emoji && vk == 0 {
            vk++
            mk, err := gto.ReadFile(f)
            if err == nil {
              requests("emojis", m, &js)
              tc := bytes.NewReader(mk)
              _, err = s.ChannelFileSend(m.ChannelID, "EchoEmojiSystem.png", tc)
              if err != nil {
                s.ChannelMessageSend(m.ChannelID, "I don't have permissions to post images in this channel.")
              }
            }
          }
        }
      }
    }
  }


if gto.HasPrefix(m.Content, ":") && gto.HasSuffix(m.Content, ":") && m.Author.ID != s.State.User.ID {
  content := gto.Replace(m.Content, ":", "", -1)
    kappa, _ := gto.ReadDir("images/emotes", "*")
    for _, f := range kappa {
        emoji := gto.ToLower(f)
        emoji = gto.Replace(emoji, ".jpg", "", -1)
        emoji = gto.Replace(emoji, ".png", "", -1)
        emoji = gto.Replace(emoji, "images\\emotes\\", "", -1)

        if gto.ToLower(content) == emoji {

              mk, err := gto.ReadFile(f)
              if err == nil {
                requests("emojis", m, &js)
                //  s.ChannelMessageDelete(m.ChannelID, m.ID)
                tc := bytes.NewReader(mk)
                _, err = s.ChannelFileSend(m.ChannelID, "EchoEmojiSystem.png", tc)
                if err != nil {
                  s.ChannelMessageSend(m.ChannelID, "I don't have permissions to post images in this channel.")
                }
            }
          }
      }
  }
}











func PlaySystem(s *discordgo.Session, m *discordgo.MessageCreate, sound string) {
    cnt := 0
    // Find the channel that the message came from.
    c, err := s.State.Channel(m.ChannelID)
    if err != nil {
      // Could not find channel.
      return
    }

    // Find the guild for that channel.
    g, err := s.State.Guild(c.GuildID)
    if err != nil {
      // Could not find guild.
      return
    }

    // Look for the message sender in that guilds current voice states.
    for _, vs := range g.VoiceStates {
      if vs.UserID == m.Author.ID {
        cnt++
        err := loadSound(sound)
        if err != nil {
          fmt.Println("Error loading sound: ", err)
          fmt.Println("Please copy $GOPATH/src/github.com/bwmarrin/examples/airhorn/airhorn.dca to this directory.")
          return
        }
        err = playSound(s, g.ID, vs.ChannelID)
        if err != nil {
          fmt.Println("Error playing sound:", err)
        }

        return
      }
    }
    if cnt == 0 {
      s.ChannelMessageSend(m.ChannelID, "You need to be in a voice channel to use this command.")
    }
}













func CheckWarn(GuildID string, user string, max int) bool {
  opt := false
  var sys map[string]interface{}
  var newsys blank
  if _, err := os.Stat("servers/" + GuildID + "/warnings.json"); err != nil {
    b, err := gto.Marshal(newsys)
    if err == nil {
      gto.WriteFile("servers/"+GuildID+"/warnings.json", b, 0777)
    }
  }


// the warnings file exists now let's loop through the ids and grab their warnings.
// if they don't exist, add their key with 1 value.
global := 0
newwarn := 0
  file, err := gto.ReadFile("servers/"+GuildID+"/warnings.json")
  if err == nil {
    gto.Unmarshal(file, &sys)
    if sys[user] == nil {
      newwarn = 0
    } else {
      newwar := sys[user].(string)
      newwarn, err = strconv.Atoi(newwar)
      if err != nil {
        return false
      }
    }

    if err == nil {
      newwarn++
      backto := strconv.Itoa(newwarn)
      sys[user] = backto
      global = newwarn

      if global >= max {
        opt = true
        for k, _ := range sys {
          if k == user {
            delete(sys, k)
          }
        }
      } else {
        opt = false
      }

      mk, err := gto.Marshal(sys)
      if err == nil {
        gto.WriteFile("servers/"+GuildID+"/warnings.json", mk, 0777)
      }
    }
  }
return opt
} // end of CheckWarn Function.






func GetSummonerID(region string, user string) int {
  data := 0
  goriot.SetAPIKey("bf4d4d68-9b08-4591-855d-6763a2ec5c62")
  sum, err := goriot.SummonerByName(region, user)
  if err != nil {
//    t.Error(err.Error())
    fmt.Println(err)
    return 0
  }
  fmt.Println(sum)
  for _, v := range sum {
    if v.Name == user {
      userid := fmt.Sprintf("%d", v.ID)
    //  icon := fmt.Sprintf("%d", v.ProfileIconID)
     //  level := strconv.Itoa(v.SummonerLevel)
    //  fmt.Println("User ID:" + userid)
      data, err = strconv.Atoi(userid)
      if err != nil {
        return 0
      }
    }
  }
//  theid := strconv.Itoa(sum.ID)
//  fmt.Println("User:"+sum.Name + "\nID:"+theid)
  return data
}


func StatSummariesBySummoner(region string, id int64, sumType string) string {
  data := ""
  goriot.SetAPIKey("bf4d4d68-9b08-4591-855d-6763a2ec5c62")
  sum, err := goriot.StatSummariesBySummoner(region, id, "SEASON3")
  if err != nil {
    fmt.Println(err)
    return "Error Retrieving summoner stats."
  }

  for _, v := range sum {
    if v.PlayerStatSummaryType == sumType {
      AverageAssists := fmt.Sprintf("%d", v.AggregatedStats.AverageAssists)
      AverageChampionsKilled := fmt.Sprintf("%d", v.AggregatedStats.AverageChampionsKilled)
      AverageCombatPlayerScore := fmt.Sprintf("%d", v.AggregatedStats.AverageCombatPlayerScore)
      AverageNodeCapture := fmt.Sprintf("%d", v.AggregatedStats.AverageNodeCapture)
      AverageNodeCaptureAssist := fmt.Sprintf("%d", v.AggregatedStats.AverageNodeCaptureAssist)
      AverageNumDeaths  := fmt.Sprintf("%d", v.AggregatedStats.AverageNumDeaths)
      AverageTotalPlayerScore := fmt.Sprintf("%d", v.AggregatedStats.AverageTotalPlayerScore)
      TotalChampionKills := fmt.Sprintf("%d", v.AggregatedStats.TotalChampionKills)
      TotalMinionKills := fmt.Sprintf("%d", v.AggregatedStats.TotalMinionKills)
      TotalNeutralMinionsKilled := fmt.Sprintf("%d", v.AggregatedStats.TotalNeutralMinionsKilled)
      TotalNodeCapture := fmt.Sprintf("%d", v.AggregatedStats.TotalNodeCapture)
      TotalNodeNeutralize := fmt.Sprintf("%d", v.AggregatedStats.TotalNodeNeutralize)
      TotalTurretsKilled := fmt.Sprintf("%d", v.AggregatedStats.TotalTurretsKilled)
      MaxChampionKills := fmt.Sprintf("%d", v.AggregatedStats.MaxChampionsKilled)
      MaxAssists := fmt.Sprintf("%d", v.AggregatedStats.MaxAssists)

      theid := fmt.Sprintf("%d", id)
      wins := fmt.Sprintf("%d", sum[0].Wins)

      data = "```ruby\n"+"Summoner ID: " + theid + "\n"
      data = data + "Wins: " + wins + "```"

      data = data + "```ruby\n"+"Total Kills: " + TotalChampionKills + "\n"
      data = data + "Average Champions Killed: " + AverageChampionsKilled + "\n"
      data = data + "Total Champion Kills: " + TotalChampionKills + "\n"
      data = data + "Max Champion Kills: " + MaxChampionKills + "\n"
      data = data + "Average Deaths: " + AverageNumDeaths + "\n"
      data = data + "Average Assists: " + AverageAssists + "\n"
      data = data + "Max Assists: " + MaxAssists + "\n"
      data = data + "```"

      data = data + "```ruby\n"+"Average Combat Score: " + AverageCombatPlayerScore + "\n"
      data = data + "Average Player Score: " + AverageTotalPlayerScore + "\n"
      data = data + "Average Node Capture: " + AverageNodeCapture + "\n"
      data = data + "Average Node Assists: " + AverageNodeCaptureAssist + "\n"
      data = data + "Total Minion Kills: " + TotalMinionKills + "\n"
      data = data + "Total Neutral Minion Kills: " + TotalNeutralMinionsKilled + "\n"
      data = data + "Total Nodes Captured: " + TotalNodeCapture + "\n"
      data = data + "Total Nodes Neutralized: " + TotalNodeNeutralize + "\n"
      data = data + "Total Turrets Killed: " + TotalTurretsKilled + "```"



    }
  }
//  fmt.Println(sum[0].AggregatedStats.playerStatSummaryType)
//  fmt.Println(sum[0].AggregatedStats.AverageChampionsKilled)
  return data
}




func ARSWait(s *discordgo.Session, m *discordgo.MessageCreate, newresp string, GuildID string, ispm bool, BotMaster bool, params string, trigger string, thet time.Duration, in info) {
  <-time.After(thet)
  ARS(s, m, newresp, GuildID, ispm, BotMaster, params, trigger, in)
}




func ARS(s *discordgo.Session, m *discordgo.MessageCreate, newresp string, GuildID string, ispm bool, BotMaster bool, params string, trigger string, in info) {
therole := ""
theid := ""
work := true
cd := 0
dont := 0
theuser := m.Author.ID
p1 := ""
p2 := ""
p3 := ""
p4 := ""


if ispm == false {
theid = m.ChannelID
} else {
    k, err := s.UserChannelCreate(m.Author.ID)
    if err == nil {
      theid = k.ID
    }
}




if theid != "" {


  if gto.Contains(newresp, "{params}") {
    params = gto.Replace(params, trigger + " ", "", -1)
    newresp = gto.Replace(newresp, "{params}", params, -1)
  }


if gto.Contains(newresp, "{p1}") {
    if gto.Contains(params, "//") {
      if len(gto.Split(params, "//")) > 1 {
        p1 = gto.Split(gto.Split(params, " ")[1], "//")[0]
        p1 = gto.TrimPrefix(p1, " ")
      }
    }
    newresp = gto.Replace(newresp, "{p1}", p1, -1)
}


if gto.Contains(newresp, "{p2}") {
    if gto.Contains(params, "//") {
      if len(gto.Split(params, "//")) > 2 {
        p2 = gto.Split(params, "//")[1]
        p2 = gto.TrimPrefix(p2, " ")
      }
    }
     newresp = gto.Replace(newresp, "{p2}", p2, -1)
}


if gto.Contains(newresp, "{p3}") {
    if gto.Contains(params, "//") {
      if len(gto.Split(params, "//")) > 3 {
        p3 = gto.Split(params, "//")[2]
        p3 = gto.TrimPrefix(p3, " ")
      }
    }
     newresp = gto.Replace(newresp, "{p3}", p3, -1)
}



if gto.Contains(newresp, "{p4}") {
    if gto.Contains(params, "//") {
      if len(gto.Split(params, "//")) > 4 {
        p4 = gto.Split(params, "//")[3]
        p4 = gto.TrimPrefix(p4, " ")
      }
    }
     newresp = gto.Replace(newresp, "{p4}", p4, -1)
}





if gto.Contains(newresp, "{sleep:") {
  ti := gto.Split(newresp, "{sleep:")[1]
  tim := gto.Split(ti, "}")[0]
      io, err := time.ParseDuration(tim)
    if err != nil {
      gto.SendMessage(s, m.ChannelID, "You need to format the time correctly. {sleep:15s} {sleep:15m} {sleep:1h}")
    }
    newresp = gto.Replace(newresp, "{sleep:"+tim+"}", "", -1)
  go ARSWait(s, m, newresp, GuildID, ispm, BotMaster, params, trigger, io, in)
  return
}





if gto.Contains(m.Content, "<@") {
  t := gto.Split(m.Content, "<@")[1]
  theuser = gto.Split(t, ">")[0]
  theuser = gto.Replace(theuser, "!", "", -1)
}


if gto.Contains(newresp, "{user}") {
  newresp = strings.Replace(newresp, "{user}", "<@"+theuser+">", -1)
 }


if gto.Contains(newresp, "{/user}") {
  usr, err := s.State.Member(GuildID, theuser)
  if err == nil {
    newresp = strings.Replace(newresp, "{/user}", usr.User.Username, -1)
  } else {
    newresp = strings.Replace(newresp, "{/user}", "{NaN}", -1)
  }
}



if gto.Contains(newresp, "{plugin:") {
  plug := strings.Split(newresp, "{plugin:")[1]
  plugin := strings.Split(plug, "}")[0]
  if _, err := os.Stat("servers/"+GuildID+"/scripts/"+plugin+".plugin.ars"); err == nil {
    go RunScript(s, m, GuildID, m.ChannelID, "servers/"+GuildID+"/scripts/"+plugin+".plugin.ars", BotMaster, in)
    newresp = strings.Replace(newresp, "{plugin:"+plugin+"}", "", -1)
  }
}



if gto.Contains(newresp, "{sky}") {
  go PostImage("skies", s, m)
  newresp = strings.Replace(newresp, "{sky}", "", -1)
}


if gto.Contains(newresp, "{space}") {
  go PostImage("space", s, m)
  newresp = strings.Replace(newresp, "{space}", "", -1)
}


if gto.Contains(newresp, "{dbz}") {
  go PostImage("dbz", s, m)
  newresp = strings.Replace(newresp, "{dbz}", "", -1)
}


if gto.Contains(newresp, "{cute}") {
  go PostImage("cute", s, m)
  newresp = strings.Replace(newresp, "{cute}", "", -1)
}


if gto.Contains(newresp, "{cars}") {
  go PostImage("cars", s, m)
  newresp = strings.Replace(newresp, "{cars}", "", -1)
}


if gto.Contains(newresp, "{wrecks}") {
  go PostImage("wrecks", s, m)
  newresp = strings.Replace(newresp, "{wrecks}", "", -1)
}


if gto.Contains(newresp, "{trucks}") {
  go PostImage("trucks", s, m)
  newresp = strings.Replace(newresp, "{trucks}", "", -1)
}





if gto.Contains(newresp, "{play:") {
  dat := gto.Split(newresp, "{play:")[1]
  data := gto.Split(dat, "}")[0]
  go PlaySystem(s, m, "dca/"+data+".dca")
  newresp = gto.Replace(newresp, "{play:"+data+"}", "", -1)
}



if gto.Contains(newresp, "{nickname:") {
  nic := gto.Split(newresp, "{nickname:")[1]
  nick := gto.Split(nic, "}")[0]
  gto.Print("User changed nicknames to: " + nick)
  err = s.GuildMemberNickname(GuildID, theuser, nick)
  if err != nil {
    newresp = gto.Replace(newresp, "{nickname:"+nick+"}", "[Nickname Failed][Can't edit server owners nickname or no permissions.]", -1)
  } else {
    newresp = gto.Replace(newresp, "{nickname:"+nick+"}", "", -1)
  }
}










if gto.Contains(newresp, "{tclear:") {
  do := 0
  var list []string
  cl := gto.Split(newresp, "{tclear:")[1]
  clear := gto.Split(cl, "}")[0]
  if clear != "" {
    // convert to integer
    num, err := gto.Integer(clear)
    if err == nil {
      msgs, err := s.ChannelMessages(m.ChannelID, 100, "", "")
      if err == nil {
        for _, v := range msgs {
          if do <= num && m.Author.ID == v.Author.ID {
            do++
            list = append(list, v.ID)
          }
        }
        if do > 0 {
          s.ChannelMessagesBulkDelete(m.ChannelID, list)
        }
      }
    }
  }
  newresp = gto.Replace(newresp, "{tclear:"+clear+"}", "", -1)
}










if gto.Contains(newresp, "{clear:") {
  do := 0
  var list []string
  cl := gto.Split(newresp, "{clear:")[1]
  clear := gto.Split(cl, "}")[0]
  if clear != "" {
    // convert to integer
    num, err := gto.Integer(clear)
    if err == nil {
      msgs, err := s.ChannelMessages(m.ChannelID, 100, "", "")
      if err == nil {
        for _, v := range msgs {
          if do <= num {
            do++
            list = append(list, v.ID)
          }
        }
        if do > 0 {
          s.ChannelMessagesBulkDelete(m.ChannelID, list)
        }
      }
    }
  }
  newresp = gto.Replace(newresp, "{clear:"+clear+"}", "", -1)
}











if gto.Contains(newresp, "{find:") {
  w := gto.Split(newresp, "{find:")[1]
  words := gto.Split(w, "}")[0]
  newresp = gto.Replace(newresp, "{find:"+words+"}", "", -1)

  if gto.Contains(words, ",") == false {
    if gto.Contains(gto.ToLower(m.Content), gto.ToLower(words)) == false {
      return
    }
  } else {
    cnd := 0
    // they want multiple words to be found.
    data := gto.Split(words, ",")
    for _, v := range data {
      if gto.Contains(gto.ToLower(m.Content), gto.ToLower(v)) {
        cnd++
      }
    }
    if cnd == 0 {
      return
    }
  }
}





if gto.Contains(newresp, "{ignore:") {
  iggy := gto.Split(newresp, "{ignore:")[1]
  ignore := gto.Split(iggy, "}")[0]
  newresp = gto.Replace(newresp, "{ignore:"+ignore+"}", "", -1)
  if gto.Contains(m.Content, ignore) {
    return
  }
}



if gto.Contains(newresp, "{nobots}") {
  usr, err := s.User(m.Author.ID)
  if err != nil {
    return
  }
  if usr.Bot == true {
    return
  }
}



if gto.Contains(newresp, "{bots}") {
  usr, err := s.User(m.Author.ID)
  if err != nil {
    return
  }
  if usr.Bot == false {
    return
  }
}




if gto.Contains(newresp, "{twitch:") {
    var twi map[string]map[string]interface{}
    data := gto.Split(newresp, "{twitch:")
    usr := gto.Split(data[1], "}")
    user := usr[0]
    if user != "" {
        gto.GetJson("https://api.twitch.tv/kraken/streams/"+user+"?client_id=j8k3gackqiwuasubw2htxd176dlbni8", &twi)
        if twi["stream"]["game"] != nil {
        game := twi["stream"]["game"].(string)
        dd := fmt.Sprintf("%.2f", twi["stream"]["viewers"].(float64))
        fps := fmt.Sprintf("%.2f", twi["stream"]["average_fps"].(float64))
        viewers := gto.Split(dd, ".")
        newresp = gto.Replace(newresp, "{twitch:"+user+"}", "", -1)
        newresp = gto.Replace(newresp, "{game}", game, -1)
        newresp = gto.Replace(newresp, "{views}", viewers[0], -1)
        newresp = gto.Replace(newresp, "{fps}", fps, -1)
        newresp = gto.Replace(newresp, "{url}", "<https://www.twitch.tv/"+user+">", -1)
      } else {
        newresp = gto.Replace(newresp, "{twitch:"+user+"}", "user-offline", -1)
        newresp = gto.Replace(newresp, "{game}", "user-offline", -1)
        newresp = gto.Replace(newresp, "{views}", "user-offline", -1)
        newresp = gto.Replace(newresp, "{fps}", "user-offline", -1)
        newresp = gto.Replace(newresp, "{url}", "<https://www.twitch.tv/"+user+">", -1)
      }
    }
}








	// Let's try to make a custom if statement code.
	if gto.Contains(newresp, "{if:") {
		opti := gto.Split(newresp, "{if:")
		option := gto.Split(opti[1], "}")





    // ####### IF COMPARISON FOR == EQUAL TO.
    if gto.Contains(option[0], "==") {
		comp := gto.Split(option[0], "==")
		beg := comp[0]
		end := comp[1]

		if beg == "user" {
			if m.Author.Username != end {
				cd++
        return
			}
		}



		if beg == "channel" {
			channel, err := s.State.Channel(m.ChannelID)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Channel Name doesn't exist")
				return
			}
			if channel.Name != end {
				cd++
        return
			}
		}


    if gto.Contains(end, "&") {
      channel, err := s.State.Channel(m.ChannelID)
      if err != nil {
        s.ChannelMessageSend(m.ChannelID, "Channel Name doesn't exist")
        return
      }
      if gto.Contains(channel.Name, end) == false {
        cd++
      }
    }
  }
// #############################################


    if gto.Contains(option[0], "!=") {
    comp := gto.Split(option[0], "!=")
    beg := comp[0]
    end := comp[1]

    if beg == "user" {
      if m.Author.Username == end {
        cd++
      }
    }

    if beg == "channel" {
      channel, err := s.State.Channel(m.ChannelID)
      if err != nil {
        s.ChannelMessageSend(m.ChannelID, "Channel Name doesn't exist")
        return
      }
      if channel.Name == end {
        cd++
      }
    }

      if gto.Contains(end, "&") {
      channel, err := s.State.Channel(m.ChannelID)
      if err != nil {
        s.ChannelMessageSend(m.ChannelID, "Channel Name doesn't exist")
        return
      }
      if gto.Contains(channel.Name, end) == true {
        cd++
      }
    }
  }
	newresp = gto.Replace(newresp, "{if:"+option[0]+"}", "", -1)
	}





if gto.Contains(newresp, "{delauto}") {
  dcnt := 0
  var ars map[string]interface{}
  file, err := gto.ReadFile("servers/"+GuildID+"/autoresponse.json")
  if err == nil {
    gto.Unmarshal(file, &ars)
    for k, _ := range ars {
      if k == "&" + trigger {
        dcnt++
        delete(ars, k)
      }
      if k == trigger {
        dcnt++
        delete(ars, k)
      }
    } // end of range ars
    b, err := gto.Marshal(ars)
    if err == nil && dcnt > 0 {
      gto.WriteFile("servers/"+GuildID+"/autoresponse.json", b, 0777)
     // s.ChannelMessageSend(m.ChannelID, "I've deleted `"+str+"` from your A.R.S File")
    }
    if dcnt == 0 {
     // s.ChannelMessageSend(m.ChannelID, "The trigger: `"+str+"` was not found in your A.R.S File.")
    }
  }
  newresp = gto.Replace(newresp, "{delauto}", "", -1)
}








	if gto.Contains(newresp, "{nsfw}") {
		newresp = gto.Replace(newresp, "{nsfw}", "", -1)
	}

    if gto.Contains(newresp, "{ifchan:") {
      data := gto.Split(newresp, "{ifchan:")
      data = gto.Split(data[1], "}")
      chans := data[0]

      if gto.Contains(chans, ",") {
      	dh := 0
      	channels := gto.Split(chans, ",")
      	for _, v := range channels {
      		if m.ChannelID == v {
      			dh++
      			newresp = gto.Replace(newresp, "{ifchan:"+chans+"}", "", -1)
      		}
      	}
      	if dh == 0 {
      		work = false
      		cd++
          return
      	}

      } else {
      	if m.ChannelID == chans {
      		newresp = gto.Replace(newresp, "{ifchan:"+chans+"}", "", -1)
      	} else {
      		cd++
      		work = false
          return
      	}
      }
    }



        if gto.Contains(newresp, "{getid}") {
          newresp = gto.Replace(newresp, "{getid}", theuser, -1)
        }


        if gto.Contains(newresp, "{joined}") {
          joined, err := s.State.Member(GuildID, theuser)
          if err != nil {
            return
          } else {
            if joined.JoinedAt != "" {
            newjo := gto.Split(joined.JoinedAt, "T")
              thedate := newjo[0]
              thetime1 := gto.Split(newjo[1], ".")
              thetime := thetime1[0]
              newresp = gto.Replace(newresp, "{joined}", thedate + " " + thetime, -1)
            } else {
              newresp = gto.Replace(newresp, "{joined}", "NaN", -1)
            }
          }
        }


        if gto.Contains(newresp, "{cats}") {
            var img map[string]interface{}
            gto.GetJson("http://random.cat/meow", &img)
            if _, ok := img["file"]; ok {
              newcat := img["file"].(string)
              newresp = gto.Replace(newresp, "{cats}", newcat, -1)
            } else {
              newresp = gto.Replace(newresp, "{cats}", "API is currently down.", -1)
            }
        }


      if gto.Contains(newresp, "{exc:") {
          str := gto.Split(newresp, "{exc:")
          estr := gto.Split(str[1], "}")
          role := estr[0]

          if gto.Contains(role, ",") == false {
            newresp = gto.Replace(newresp, "{exc:"+role+"}", "", -1)
            if gto.MemberHasRole(s, GuildID, m.Author.ID, role) == false {
              therole = GetRoleID(s, GuildID, role)
            } else {
              work = false
              return
            }
          } else {
            // loop through roles
            rolez := gto.Split(role, ",")
            for _, v := range rolez {
              if gto.MemberHasRole(s, GuildID, m.Author.ID, v) == true {
                cd++
                work = false
              }
            } // end of for loop
            if cd > 0 {
              return
            }
          }
          newresp = gto.Replace(newresp, "{exc:"+role+"}", "", -1)
      }






      if gto.Contains(newresp, "{req:") {
          str := gto.Split(newresp, "{req:")
          estr := gto.Split(str[1], "}")
          role := estr[0]

          if gto.Contains(role, ",") == false {
            newresp = gto.Replace(newresp, "{req:"+role+"}", "", -1)
            if gto.MemberHasRole(s, GuildID, m.Author.ID, role) == false {
              work = false
              return
            } else {
              work = true
            }
          } else {
            // loop through roles
            rolez := gto.Split(role, ",")
            for _, v := range rolez {
              if gto.MemberHasRole(s, GuildID, m.Author.ID, v) == false {
                cd++
                work = false
              }
            } // end of for loop
            if cd > 0 {
              return
            }
          }
          newresp = gto.Replace(newresp, "{req:"+role+"}", "", -1)
      }


  if gto.Contains(newresp, "{del}") == true {
    newresp = gto.Replace(newresp, "{del}", "", -1)
    s.ChannelMessageDelete(m.ChannelID, m.ID)
    //  time.Sleep(2000 * time.Millisecond)
  }

if gto.Contains(newresp, "{del:") {
  ti := gto.Split(newresp, "{del:")[1]
  tim := gto.Split(ti, "}")[0]
    io, err := time.ParseDuration(tim)
    if err != nil {
      gto.SendMessage(s, m.ChannelID, "Error completing `{del:"+tim+"} the time format was incorrect.")
    } else {
      go TimeDelete(s, m.ChannelID, m.ID, io)
    }
    newresp = gto.Replace(newresp, "{del:"+tim+"}", "", -1)
}




  if gto.Contains(newresp, "{del*}") == true {
    newresp = gto.Replace(newresp, "{del*}", "", -1)
    if BotMaster == false {
        s.ChannelMessageDelete(m.ChannelID, m.ID)
     }
    //  time.Sleep(2000 * time.Millisecond)
  }




        if gto.Contains(newresp, "{warn:") {
          x := gto.Split(newresp, "{warn:")
          x2 := gto.Split(x[1], "}")
          warns, err := strconv.Atoi(x2[0])
          if err == nil {
            newresp = gto.Replace(newresp, "{warn:"+x2[0]+"}", "", -1)
            // strconv.Atoi(warns)
            if CheckWarn(GuildID, theuser, warns) == true {
              // kick the user.


              if gto.Contains(newresp, "{kick}") {
                dont++
                if BotMaster == false {
                  s.GuildMemberDelete(GuildID, theuser)
              //     s.ChannelMessageSend(theid, newresp)
                  newresp = gto.Replace(newresp, "{kick}", "", -1)
                  if gto.Contains(newresp, "{msg:") {
                    l := gto.Split(newresp, "{msg:")
                    l2 := gto.Split(l[1], "}")
                   // msg := l2[0]
                    newresp = gto.Replace(newresp, "{msg:"+l2[0]+"}", "", -1)
                  }
                //  fmt.Println("I have Kicked "+m.Author.ID)
                //  time.Sleep(2000 * time.Millisecond)
                } else {
                  return
                //  fmt.Println("User is a bot commander, don't show response.")
                }
              }




              if gto.Contains(newresp, "{ban}") {
                dont++
                if BotMaster == false {
                  s.GuildBanCreate(GuildID, theuser, 10)
              //     s.ChannelMessageSend(theid, newresp)
                  newresp = gto.Replace(newresp, "{ban}", "", -1)
                  if gto.Contains(newresp, "{msg:") {
                    l := gto.Split(newresp, "{msg:")
                    l2 := gto.Split(l[1], "}")
                  //  msg := l2[0]
                    newresp = gto.Replace(newresp, "{msg:"+l2[0]+"}", "", -1)
                  }
               //   fmt.Println("I have Banned "+m.Author.ID)
                //  time.Sleep(2000 * time.Millisecond)
                } else {
                  return
               //   fmt.Println("User is a bot commander, don't show response.")
                }
              }
            } else {
              if BotMaster == false {
                  if gto.Contains(newresp, "{msg:") {
                    l := gto.Split(newresp, "{msg:")
                    l2 := gto.Split(l[1], "}")
                    msg := l2[0]
                    newresp = gto.Replace(newresp, "{msg:"+l2[0]+"}", msg, -1)
                    s.ChannelMessageSend(theid, msg)
                  }
              }
             return
            }
          }
        }





          if gto.Contains(newresp, "{listroles}") {
            newdat := ""
            var l []string
            x, err := s.State.Member(GuildID, theuser)
            l = x.Roles
            if err != nil {
              return
            } else {
                for r := range x.Roles {
                  newdat = newdat + GetRoleName(s, GuildID, l[r]) + ", "
                }
            }
            if newdat != "" {
              newdat = gto.TrimSuffix(newdat, ", ")
              newresp = gto.Replace(newresp, "{listroles}", newdat, -1)
            } else {
              newresp = gto.Replace(newresp, "{listroles}", "None", -1)
            }
          }






  if gto.Contains(newresp, "{channels}") {
    newdat := ""
    x, err := s.State.Guild(GuildID)
    if err == nil {
      for _, v := range x.Channels {
        newdat = newdat + v.Name + "("+v.Type+"), "
      }
        if newdat != "" {
          newdat = gto.TrimSuffix(newdat, ", ")
          newresp = gto.Replace(newresp, "{channels}", newdat, -1)
        }
    }
  }





          if gto.Contains(newresp, "{allroles}") {
            newdat := ""
            x, err := s.State.Guild(GuildID)
            if err != nil {
              return
            } else {
                for _, v := range x.Roles {
                  if v.Name != "@everyone" {
                  newdat = newdat + v.Name + ", "
                  }
                }
            }
            if newdat != "" {
              newdat = gto.TrimSuffix(newdat, ", ")
              newresp = gto.Replace(newresp, "{allroles}", newdat, -1)
            }
          }





              if gto.Contains(newresp, "{kick}") == true && dont == 0 {
                if BotMaster == false {
                  newresp = gto.Replace(newresp, "{kick}", "", -1)
                  s.GuildMemberDelete(GuildID, theuser)
                   s.ChannelMessageSend(theid, newresp)
                   cd++
                //  time.Sleep(2000 * time.Millisecond)
                } else {
                  work = false
                }
              }




if work == true {

/// ALERT ONE OR MULTIPLE PPL VIA PM.
      if gto.Contains(newresp, "{alert:") {
          str := gto.Split(newresp, "{alert:")
          estr := gto.Split(str[1], "}")
          ppl := estr[0]

          if gto.Contains(ppl, ",") == false {
            newresp = gto.Replace(newresp, "{alert:"+ppl+"}", "", -1)
             pk, err := s.UserChannelCreate(ppl)
             if err == nil {
              s.ChannelMessageSend(pk.ID, "<@"+m.Author.ID+"> Has triggered your alert!")
             }
          } else {
            newresp = gto.Replace(newresp, "{alert:"+ppl+"}", "", -1)
            dat := gto.Split(ppl, ",")
            for _, v := range dat {
              pk, err := s.UserChannelCreate(v)
              if err == nil {
                s.ChannelMessageSend(pk.ID, "<@"+m.Author.ID+"> Has triggered your alert!")
                time.Sleep(1000 * time.Millisecond)
              }
            }
          } // check if it was one or multiple ppl to alert.
      }
}




            if gto.Contains(newresp, "{params:url}") {
              params = gto.Replace(params, trigger + " ", "", -1)
              params = gto.Replace(params, " ", "%20", -1)
              newresp = gto.Replace(newresp, "{params:url}", params, -1)
            }



            if gto.Contains(newresp, "{params:flip}") {
              params = gto.Replace(params, trigger + " ", "", -1)

              input := params
              // Get Unicode code points.
              n := 0
              rune := make([]rune, len(input))
              for _, r := range input {
                  rune[n] = r
                  n++
              }
              rune = rune[0:n]
              // Reverse
              for i := 0; i < n/2; i++ {
                  rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
              }
              // Convert back to UTF-8.
              output := string(rune)
              newresp = gto.Replace(newresp, "{params:flip}", output, -1)
            }



        if gto.Contains(newresp, "{rawid}") {
          newresp = gto.Replace(newresp, "{rawid}", "", -1)
          newresp = gto.Replace(newresp, "<@", "", -1)
          newresp = gto.Replace(newresp, ">", "", -1)
        }


              if gto.Contains(newresp, "{ban}") == true && dont == 0 {
                if BotMaster == false {
                  newresp = gto.Replace(newresp, "{ban}", "", -1)
                  s.GuildBanCreate(GuildID, theuser, 10)
                   s.ChannelMessageSend(theid, newresp)
                   cd++
                //  time.Sleep(2000 * time.Millisecond)
                } else {
                  work = false
                }
              }


  if gto.Contains(newresp, "{meme}") {
    var meme []string
    cnt := 0
    cnt = gto.CountLines("meme.txt")
    myrand := gto.Random(1, cnt)
    meme, err := gto.ReadLines("meme.txt")
    if err == nil {
   // s.ChannelTyping(m.ChannelID)
    // s.ChannelMessageSend(m.ChannelID, meme[myrand])
      newresp = gto.Replace(newresp, "{meme}", meme[myrand], -1)
    }
  }


  if gto.Contains(newresp, "{joke}") {
    var joke []string
    cnt := 0
    cnt = gto.CountLines("Random.txt")
    myrand := gto.Random(1, cnt)
    joke, err := gto.ReadLines("Random.txt")
    if err == nil {
   // s.ChannelTyping(m.ChannelID)
    // s.ChannelMessageSend(m.ChannelID, meme[myrand])
      newresp = gto.Replace(newresp, "{joke}", joke[myrand], -1)
    }
  }




        if gto.Contains(newresp, "{role:") {
          str := gto.Split(newresp, "{role:")
          estr := gto.Split(str[1], "}")
          role := estr[0]
          newresp = gto.Replace(newresp, "{role:"+role+"}", "", -1)
          if gto.MemberHasRole(s, GuildID, theuser, role) == false {
            therole = GetRoleID(s, GuildID, role)
          } else {
            work = false
          }

          if therole != "" && work == true {
            x, err := s.State.Member(GuildID, theuser)
            if err != nil {
              return
            }
            x.Roles = append(x.Roles, therole)
            err = s.GuildMemberEdit(GuildID, theuser, x.Roles)
            if err == nil {
            // s.ChannelMessageSend(theid, newresp)
            } else {
              fmt.Println(err)
              s.ChannelMessageSend(theid, "Discord prevented me from saving the role: "+role)
            }

            //  time.Sleep(2000 * time.Millisecond)
          }
        }


        if gto.Contains(newresp, "{take:") {
          str := gto.Split(newresp, "{take:")
          estr := gto.Split(str[1], "}")
          role := estr[0]

          if gto.Contains(role, ",") == false {
          newresp = gto.Replace(newresp, "{take:"+role+"}", "", -1)
          if gto.MemberHasRole(s, GuildID, theuser, role) == true {
            therole = GetRoleID(s, GuildID, role)
          } else {
            work = false
            cd++
          }

            if therole != "" && work == true {
              x, err := s.State.Member(GuildID, theuser)
            if err != nil {
              return
            }
            if err == nil {
              var ms []string
              ms = x.Roles
              for mr := range x.Roles {
                t := ms[mr]
                if gto.Contains(t, therole) {
                  //fmt.Println("Membert has role: "+t)
                  x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
                  s.GuildMemberEdit(GuildID, theuser, x.Roles)
                  // s.ChannelMessageSend(m.ChannelID, newdata)
                }
              }
            }
          }
        } else { // end of the single row
          // they want to take away more roles.

          gto.Print("They want multiple roles taken away")
          data := gto.Split(role, ",")
          newresp = gto.Replace(newresp, "{take:"+role+"}", "", -1)
            x, err := s.State.Member(GuildID, theuser)
            if err != nil {
              return
            }
            if err == nil {
              var ms []string
              ms = x.Roles
            for mr := range x.Roles {
              gto.Print(ms[mr])
              for _, r1 := range data {
                gto.Print(r1)
                  if ms[mr] == r1 {
                    //fmt.Println("Membert has role: "+t)
                    x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
                    s.GuildMemberEdit(GuildID, theuser, x.Roles)
                    // s.ChannelMessageSend(m.ChannelID, newdata)
                  }
                }
              }
            }
          }
        }




        if gto.Contains(newresp, "{ass}") {
          res := gto.Random(300, 3420)
          re := strconv.Itoa(res)
          newresp = gto.Replace(newresp, "{ass}", "http://media.obutts.ru/butts_preview/0"+re+".jpg", -1)
        }


        if gto.Contains(newresp, "{boobs}") {
          res := gto.Random(8999, 9999)
          re := strconv.Itoa(res)
          newresp = gto.Replace(newresp, "{boobs}", "http://media.oboobs.ru/boobs_preview/0"+re+".jpg", -1)
        }



        if params != "" && gto.Contains(newresp, "{giphy}") {
          params = gto.Replace(params, trigger + " ", "", -1)

        //  fmt.Println(params)
          var img map[string]map[string]interface{}
          params = gto.Replace(params, " ", "+", -1)
          gto.GetJson("http://api.giphy.com/v1/gifs/random?api_key=dc6zaTOxFJmzC&tag="+params, &img)

          if img["data"]["image_original_url"] != nil {
            newcat := img["data"]["image_original_url"].(string)
            // s.ChannelMessageSend(m.ChannelID, newcat)
            newresp = gto.Replace(newresp, "{giphy}", newcat, -1)
          } else {
           // s.ChannelMessageSend(m.ChannelID, "Couldn't find any results with your keyword. `"+in.Prefix+"giphy keyword`")
            newresp = gto.Replace(newresp, "{giphy}", "No results found with your keyword..", -1)
          }
        }





if gto.Contains(newresp, "{replace:") {
  if gto.Contains(newresp, "{with:") {
    re := gto.Split(newresp, "{replace:")[1]
    rep := gto.Split(re, "}")[0]
    newresp = gto.Replace(newresp, "{replace:"+rep+"}", "", -1)
    wi := gto.Split(newresp, "{with:")[1]
    wit := gto.Split(wi, "}")[0]
    newresp = gto.Replace(newresp, "{with:"+wit+"}", "", -1)


    if gto.Contains(newresp, "{log:") {
    og := gto.Split(newresp, "{log:")[1]
    chanid := gto.Split(og, "}")[0]
    channel, err := s.State.Channel(chanid)
    if err == nil {
      newresp = gto.Replace(newresp, "{log:"+chanid+"}", "", -1)
      newresp = gto.Replace(newresp, "& =", "", -1)
      newresp = gto.Replace(newresp, trigger + "=", "", -1)
      gto.SendMessage(s, channel.ID, newresp)
    }
  }



    newresp = m.Content
    if gto.Contains(rep, ",") == false {
      newresp = gto.Replace(gto.ToLower(m.Content), gto.ToLower(rep), gto.ToLower(wit), -1)
    } else {
      data := gto.Split(rep, ",")
      for _, v := range data {
        newresp = gto.Replace(gto.ToLower(newresp), v, wit, -1)
        newresp = gto.Replace(newresp, gto.ToLower(v), gto.ToLower(wit), -1)
      }
    }
  }
}





if gto.Contains(newresp, "{trigger:") {
  tri := gto.Split(newresp, "{trigger:")[1]
  trig := gto.Split(tri, "}")[0]
  gto.Print(trig)
  newresp = gto.Replace(newresp, "{trigger:"+trig+"}", "", -1)
  gto.Print(newresp)
  gto.SendMessage(s, m.ChannelID, newresp)
  ARS(s, m, newresp, GuildID, ispm, BotMaster, params, trig, in)
  work = false
}






  if gto.Contains(newresp, "{log:") && gto.Contains(newresp, "replace") == false {
    og := gto.Split(newresp, "{log:")[1]
    chanid := gto.Split(og, "}")[0]
      channel, err := s.State.Channel(chanid)
      if err != nil {
        return
      }
      newresp := gto.Replace(newresp, "{log:"+chanid+"}", "", -1)
      gto.SendMessage(s, channel.ID, newresp)
      return
  }



if gto.Contains(newresp, "{pm:") {
  pmto := gto.Split(newresp, "{pm:")[1]
  pm := gto.Split(pmto, "}")[0]
  newresp = gto.Replace(newresp, "{pm:"+pm+"}", "", -1)

  if gto.Contains(pm, ",") == false {
    k, err := s.UserChannelCreate(pm)
    if err == nil {
      gto.SendMessage(s, k.ID, newresp)
    }
  } else {
    // they want to pm multiple ppl...here we go
    // limit the amount of ppl to 5
    lmt := 0
    dat := gto.Split(pm, ",")
    for _, v := range dat {
      lmt++
      if lmt <= 5 {
        k, err := s.UserChannelCreate(v)
        if err == nil {
          gto.SendMessage(s, k.ID, newresp)
        }
      }
    }
  }
}




        if gto.Contains(newresp, "{redirect:") && work == true && cd == 0 {
          str := gto.Split(newresp, "{redirect:")
          str = gto.Split(str[1], "}")
          chanid := str[0]
          newresp = gto.Replace(newresp, "{redirect:"+chanid+"}", "", -1)

          if gto.Contains(newresp, "{msg:") {
            dat := gto.Split(newresp, "{msg:")
            data := gto.Split(dat[1], "}")
            msg := data[0]
            newresp = gto.Replace(newresp, "{msg:"+msg+"}", "", -1)
            s.ChannelMessageSend(m.ChannelID, msg)
          }

          if gto.Contains(newresp, "{params}") {
            s.ChannelMessageSend(chanid, newresp)
          } else {
          s.ChannelMessageSend(chanid, newresp + " " + gto.TrimPrefix(m.Content, trigger))
          }
          cd++
          return
        }



              if cd == 0 && work == true {
                  s.ChannelMessageSend(theid, newresp)
              }
  } // check if theid is != nil
}


















func GetRoleID(s *discordgo.Session, guildID string, role string) string {
  var re string
  roles, err := s.State.Guild(guildID)
  if err == nil {
    for _, v := range roles.Roles {
      if v.Name == role {
        re = v.ID
      }
    }
  }
  return re
}










/* caused a memory error :( will work into it more later.
func ServerID(s *discordgo.Session, ChannelID string) string {
  c, err := s.State.Channel(ChannelID)
if err != nil {
 // channel not found
}
  return c.GuildID
}
*/








func GetRoleName(s *discordgo.Session, guildID string, role string) string {
  var re string
  roles, err := s.State.Guild(guildID)
  if err == nil {
    for _, v := range roles.Roles {
      if v.ID == role {
        re = v.Name
      }
    }
  }
  return re
}
