package main


import (
  "math/rand"
  "time"
  "os"
  "bufio"
  "strings"
  "github.com/bwmarrin/discordgo"
)










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














func GetRoleID(s *discordgo.Session, guildID string, role string) string {
  var re string
  roles, err := s.GuildRoles(guildID)
  if err == nil {
    for _, v := range roles {
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












func isMemberRole(s *discordgo.Session, GuildID string, AuthorID string, role string) bool {
  var opt bool
  opt = false
// z, err := s.State.Member(GuildID, AuthorID) 
// if err != nil {
z, err := s.GuildMember(GuildID, AuthorID)
// }

if err == nil {
  var l []string
  l = z.Roles
  for r := range z.Roles {
    if strings.Contains(l[r], GetRoleID(s, GuildID, role)) {
 //     fmt.Println("Found the role!"+l[r])
      opt = true
    }
  }
} // end of err == nil
  return opt
}