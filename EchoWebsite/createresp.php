<?php

if(isset($_POST['server']))
{
		// $params=$_GET['params'];
      	$file=file_get_contents("response/" . $_POST['server'] . ".json");
      	$dD = json_decode($file);
      	unset($file);
      	$dD->AddMaster = $_POST['addmaster'];
         $dD->DelMaster = $_POST['delmaster'];
         $dD->Prefix = $_POST['prefix'];
         $dD->GreetMsg = $_POST['greet'];
         $dD->ByeMsg = $_POST['bye'];
         $dD->GreetOff = $_POST['greetoff'];
         $dD->ByeOff = $_POST['byeoff'];
         $dD->Autorole = $_POST['autorole'];
         $dD->AutoroleOff = $_POST['autoroleoff'];
         $dD->Botrole = $_POST['botrole'];
         $dD->BotroleOff = $_POST['botroleoff'];
         $dD->SetPunish = $_POST['setpunish'];
         $dD->SetPunishError = $_POST['setpunisherror'];
         $dD->Give = $_POST['give'];
         $dD->Take = $_POST['take'];
         $dD->NotNSFW = $_POST['notnsfw'];
         $dD->Addnsfw = $_POST['addnsfw'];
         $dD->Delnsfw = $_POST['delnsfw'];
         $dD->wasnsfw = $_POST['wasnsfw'];
         $dD->Mkchan = $_POST['mkchan'];
         $dD->MkchanError = $_POST['mkchanerror'];
         $dD->IPError = $_POST['iperror'];
         $dD->NoRolePerms = $_POST['noroleperms'];
         $dD->Addrole = $_POST['addrole'];
         $dD->RoleExists = $_POST['roleexists'];
         $dD->AddRoleError = $_POST['addroleerror'];
         $dD->Delrole = $_POST['delrole'];
         $dD->RoleNoExist = $_POST['rolenoexist'];
         $dD->Mute = $_POST['mute'];
         $dD->Unmute = $_POST['unmute'];
         $dD->Rolecolor = $_POST['rolecolor'];
         $dD->MkInvite = $_POST['mkinvite'];
         $dD->DenyLinks = $_POST['denylinks'];
         $dD->AllowLinks = $_POST['allowlinks'];
         $dD->AntiLinkKick = $_POST['antilinkkick'];
         $dD->AntiLinkBan = $_POST['antilinkban'];
         $dD->AntiLinkWarn = $_POST['antilinkwarn'];
         $dD->NoWarns = $_POST['nowarns'];
         $dD->ResetWarns = $_POST['resetwarns'];
         $dD->WarnKick = $_POST['warnkick'];
         $dD->WarnBan = $_POST['warnban'];
         $dD->Warn = $_POST['warn'];
         $dD->WarnNotSet = $_POST['warnsnotset'];
         $dD->WarnCommander = $_POST['warncommander'];
         $dD->SetWarning = $_POST['setwarning'];
         $dD->SetWarningError = $_POST['setwarningerror'];
         $dD->Kick = $_POST['kick'];
         $dD->Ban = $_POST['ban'];

         if(isset($_POST['online'])) {
            $dD->Pulse = $_POST['pulse'] + 1;
         } else {
            $dD->Pulse = $_POST['pulse'];
         }

    	$newJsonString = json_encode($dD,JSON_UNESCAPED_UNICODE);
   	file_put_contents("response/" . $_POST['server'] . ".json", $newJsonString);
    	unset($current);
      echo "You've updated your Custom Responses! You will automatically see changes in your server.<br><a href='http://echobot.tk/?nav=manager'>Go Back!</a>";
}


