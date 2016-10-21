package main

type obj struct {
	Bot 			string
	Admin 			string
	Status 			string
	BotMaster		bool
	BotCommander	string
	CmdsRun			int
	Prefix 		string
	GreetMsg	string
	ByeMsg		string
	RoleSys		string
	Name		string
	AntiLink	bool
	Action		string
	Silent		bool
	Active		string
	HelpCmd		string
	BotAutoRole	string
}


type responses struct {
	AddMaster			string
	DelMaster			string
	Take				string
	Give				string
	Mute				string
	Unmute				string
	Greet				string
	GreetOff			string
	Bye					string
	ByeOff				string
	DenyLinks			string
	AllowLinks			string
	Prefix				string
	SetPunish			string
	AutoRole			string
	AutoRoleOff			string
	NoRole				string
	BotAutoRole			string
	AutoBotRoleOff		string
	AntiLinkKick		string
	AntiLinkBan			string
	AntiLinkWarn		string
	Kick				string
	Ban					string
	Status				string
	Giveme				string
	BotAutoRoleMsg		string
	AutoRoleMsg			string
	Rolecolor			string
}

type role struct {
	ID			string
	Name		string
	Managed		bool
	Hoist		bool
	Color		int
	Position	int
	Permissions	int
}


type invite struct {
	MaxAge		int
	MaxUses		int
	Temporary	bool
	XkcdPass	string
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
	BotAutoRole		string
	Rolecolor		string
}
