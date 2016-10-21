<?php
$backup = file_get_contents('data.json');
$data = json_decode($backup);
unset($backup);
$servers=$data->servers;
$cmdsrun=$data->cmdsRun;
$ARS=$data->ARS;
$Emoji=$data->Emojis;
$chans=$data->Channelcount;
$memb=$data->Memberscount;
$role=$data->Rolecount;

if($servers == "") {
	$servers = "NaN";
} if($cmdsrun == "") {
	$cmdsrun = "NaN";
}
//echo var_dump($data).PHP_EOL;
//echo "Testing Servers: ".$servers.PHP_EOL;

// $servers=file_get_contents('info.txt');
?>
Echo is currently on (<b><?php echo $servers; ?></b>) Servers. With (<b><?php echo $cmdsrun; ?></b>) Commands ran.<br>
<!--
(<b><?php echo $ARS; ?></b>) A.R.S Requests. (<b><?php echo $Emoji; ?></b>) Emoji's used. (<b><?php echo $memb; ?></b>) Members.<br>
(<b><?php echo $chans; ?></b>) Channels. (<b><?php echo $role; ?></b>) Roles.
-->