<?php

if(isset($_POST['s']))
{

		// $params=$_GET['params'];
      	$file=file_get_contents("data.json");
      	$dD = json_decode($file);
      	unset($file);
      	$dD->servers = $_POST['s'];
      	$dD->cmdsRun = $_POST['c'];
      	$dD->ARS = $_POST['a'];
      	$dD->Channelcount = $_POST['ch'];
      	$dD->Memberscount = $_POST['m'];
      	$dD->Rolecount = $_POST['r'];
 		$dD->Emojis = $_POST['e'];
    	$newJsonString = json_encode($dD,JSON_UNESCAPED_UNICODE);
   		file_put_contents("data.json", $newJsonString);
    	unset($current);
}