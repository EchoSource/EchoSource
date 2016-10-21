<?php

if(isset($_POST['server']))
{
		// $params=$_GET['params'];
      	$file=file_get_contents("servers/" . $_POST['server'] . ".json");
      	$dD = json_decode($file);
      	unset($file);
      	$dD->Name = $_POST['name'];
         $dD->Password = $_POST['password'];
      	$dD->Owner = $_POST['owner'];
      	$dD->Prefix = $_POST['prefix'];
      	$dD->RoleSys = $_POST['rolesys'];
      	$dD->Silent = $_POST['silent'];
         $warnings = $_POST['warnings'];
         settype($warnings, "integer");
      	$dD->Warnings = $warnings;
      	$dD->Action = $_POST['action'];
      	$dD->Active = $_POST['active'];
      	$dD->AntiLink = $_POST['antilink'];
      	$dD->AutoPerms = $_POST['autoperms'];
      	$dD->BotAuto = $_POST['botauto'];
      	$dD->BotMaster = $_POST['botmaster'];
      	$dD->ByeMsg = $_POST['byemsg'];
      	$dD->GreetMsg = $_POST['greetmsg'];

         if(isset($_POST['online'])) {
            $dD->Pulse = $_POST['pulse'] + 1;
         } else {
            $dD->Pulse = $_POST['pulse'];
         }

    	$newJsonString = json_encode($dD,JSON_UNESCAPED_UNICODE);
   		file_put_contents("servers/" . $_POST['server'] . ".json", $newJsonString);
    	unset($current);
      echo "You've updated your Server Database! You will automatically see changes in your server.<br><a href='http://echobot.tk/?nav=manager'>Go Back!</a>";
}


