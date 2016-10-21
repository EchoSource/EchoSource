package main

import "time"

type obj struct {
	Bot 			string
	Admin 			string
	Status 			int
	BotMaster		bool
	BotCommander	string
	CmdsRun			int
}



type tasks struct {
	Task 		string
	MakeRole	string
	Message		string
	User		string
}

type statistics struct {
	ARS				int
	Channelcount	int
	Emojis			int
	Membercount		int
	Rolecount		int
	Servercount		int
}

type response struct {
	NoRole			string
	AddMaster 		string
	DelMaster 		string
	Prefix			string
	GreetMsg		string
	GreetOff		string
	ByeMsg			string
	ByeOff			string
	Autorole		string
	AutoroleOff		string
	Botrole			string
	BotroleOff		string
	SetPunish		string
	SetPunishError	string
	Give 			string
	Take			string
	NotNSFW			string
	Addnsfw 		string
	Delnsfw			string
	wasnsfw			string
	MkchanError		string
	Mkchan 			string
	IpError			string
	NoRolePerms		string
	Addrole			string
	RoleExists		string
	AddRoleError	string
	Delrole			string
	RoleNoExist		string
	Mute			string
	Unmute			string
	Rolecolor		string
	MkInvite		string
	DenyLinks		string
	AllowLinks		string
	AntiLinkKick	string
	AntiLinkBan		string
	AntiLinkWarn	string
	NoWarns			string
	ResetWarns		string
	WarnKick		string
	WarnBan			string
	Warn 			string
	WarnNotSet		string
	WarnCommander	string
	SetWarning		string
	SetWarningError	string
	Kick 			string
	Ban 			string
	Pulse			int
}


type info struct {
	Prefix 		string
	GreetMsg	string
	ByeMsg		string
	RoleSys		string
	Owner		string
	Name		string
	AntiLink	bool
	Action		string
	Log			string
	LogType		string
	Silent		bool
	Active		string
	BotAuto		string
	BotMaster	string
	AutoPerms	bool
	Warnings	int
	Password	string
	Pulse		int
}



type role struct {
	ID	string
}


type invite struct {
	MaxAge		int
	MaxUses		int
	Temporary	bool
	XkcdPass	bool
}


type commands struct {
	Greet			string
	Bye				string
	Prefix			string
	Kick			string
	Ban 			string
	Autorole		string
	SetPunish		string
	AllowLinks		string
	DenyLinks		string
	AddMaster		string
	DelMaster		string
	Invites			string
	Meme			string
	Joke			string
	Give			string
	Take			string
	Giveme			string
	Mute			string
	Unmute			string
}



type newj struct {
	autoresp	string
}



type league struct {
	id 				int
	name 			string
	profileIconId	int
	summonerLevel	int
	revisionDate	int
}






type Info struct {
	ServerID		string
    OwnerID        	string
    OwnerUser       string
    TimeStamp		int64
    Mentions 		[]*Mentions
    Users 			[]*Users
    Messages 		[]*Message
    Restricted 		[]*Channel
}


type Users struct {
	ID 			string
	Username	string
	LastSeen	time.Time
	LastStamp	int64
	TotalMsg	int
}



type Message struct {
	ID         	string
	Author		string
	Channel 	string
	Content     string
	TimeStamp	int64
}


type Channel struct {
	ID 		string
}

type Mentions struct {
	ByID			string
	Mentioned 		string
	ByUser			string
	Content 		string
	TimeStamp		int64
}


type upsong struct {
	AddedBy     string
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FullTitle   string `json:"full_title"`
	Thumbnail   string `json:"thumbnail"`
	URL         string `json:"webpage_url"`
	Duration    int    `json:"duration"`
	Remaining   int
}