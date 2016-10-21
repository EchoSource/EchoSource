<?php
    $backup = file_get_contents("servers/" . $_SESSION['username'] . ".json");
    $data = json_decode($backup);
    unset($backup);

    if(file_exists("response/" . $_SESSION['username'] . ".json")) {
    $resp = file_get_contents("response/" . $_SESSION['username'] . ".json");
	} else {
	$resp = file_get_contents("response.json");	
	}
    $response = json_decode($resp);
    unset($resp);

?>
<table width="50%" border="0" align="center">
<tr>
<td>


<form action="createfile.php" method="POST">
<input type="hidden" name="server" value="<?php echo $_SESSION['username']; ?>">
<input type="hidden" name="pulse" value="<?php echo $data->Pulse; ?>">
<input type="hidden" name="online" value="true">
<input type="hidden" name="name" value="<?php echo $data->Name; ?>">
<input type="hidden" name="owner" value="<?php echo $data->Owner; ?>">
<table>
<tr>
<th>&nbsp;</th>
<th>Server Database</th>
</tr>
<tr>
<td><b>Name:</b></td><td><input type="text" name="nameshow" value="<?php echo $data->Name; ?>" disabled></td></tr>
<tr class="bg">
<td><b>Owner ID:</b></td><td><input type="text" name="ownershow" value="<?php echo $data->Owner; ?>" disabled></td></tr>
<tr>
<td><b>Server Prefix:</b></td><td><input type="text" name="prefix" value="<?php echo $data->Prefix; ?>" size="5"></td></tr>
<tr class="bg">
<td><b>Autorole:</b></td><td><input type="text" name="rolesys" value="<?php echo $data->RoleSys; ?>"></td></tr>
<tr>
<td><b>Autorole [Bots]:</b></td><td><input type="text" name="botauto" value="<?php echo $data->BotAuto; ?>"></td></tr>
<tr class="bg">
<td><b>Antilinks:</b></td><td>
<select name="antilink">
<?php
if($data->AntiLink == false) {
?>
<option value="true">true</option>
<option value="false">false</option>
<?php
} else {
?>
<option value="false">false</option>
<option value="true">true</option>
<?php
}
?>
</select>
</td></tr>
<tr>
<td><b>Commander Role:</b></td><td><input type="text" name="botmaster" value="<?php echo $data->BotMaster; ?>"></td></tr>
<tr class="bg">
<td><b>Warnings:</b></td><td>
<input type="number" size="5" name="warnings" value="<?php echo $data->Warnings; ?>" size="5"></td></tr>
<tr>
<td><b>Punishment:</b></td><td>
<select name="action">
<?php
if($data->Action == "kick") {
?>
<option value="kick">kick</option>
<option value="ban">ban</option>
<option value="warn">warn</option>
<?php
}
if($data->Action == "ban") {
?>
<option value="ban">ban</option>
<option value="kick">kick</option>
<option value="warn">warn</option>
<?php
}
if($data->Action == "warn") {
?>
<option value="warn">warn</option>
<option value="kick">kick</option>
<option value="ban">ban</option>
<?php
}
?>
</select>
</td></tr>
<tr class="bg">
<td><b>Greet Message:</b></td><td><textarea cols="40" rows="8" name="greetmsg"><?php echo $data->GreetMsg; ?></textarea></td></tr>
<tr>
<td><b>Bye Message:</b></td><td><textarea cols="40" rows="8" name="byemsg"><?php echo $data->ByeMsg; ?></textarea></td></tr>
<tr class="bg">
<td><b>Password:</b></td><td><input type="text" name="password" value="<?php echo $data->Password; ?>"></td></tr>
<tr>
<td><b>Silent Autorole:</b></td><td>
<?php
if($data->Silent == false) {
  $sel = "false";
} else {
  $sel = "true";
}
?>
<select name="silent">
<option value="<?php echo $sel; ?>" selected="<?php echo $sel; ?>"><?php echo $sel; ?></option>
<option value="true">true</option>
<option value="false">false</option>
</select>
</td></tr>

<tr>
<td>&nbsp;</td><td><input type="submit" value="Update"></td></tr>
</table>
</form>


<form action="newtask.php" method="POST">
<input type="hidden" name="server" value="<?php echo $_SESSION['username']; ?>">
<input type="hidden" name="task" value="makerole">
<table width="100%">
<tr>
<th>Coming Soon:</th>
<th>Add Role</th>
</tr>
<tr>
<td>Name:</td><td><input type="text" name="action" placeholder="Role name here."></td></tr>
<tr>
<td>&nbsp;</td><td><input type="submit" value="Create Role"></td></tr>
</table>
</form>


<table width="100%">
<tr>
<th>Coming Soon:</th>
<th>PM Someone</th>
</tr>
<tr>
<td>Name:</td><td><input type="text" name-"pm_user" placeholder="Username here"></td></tr>
<tr>
<td>Message:</td><td><textarea rows="8" cols="30"></textarea></td>
</tr>
</table>



<table width="100%">
<tr>
<th>Warnings Database</th>
</tr>
<tr>
<td>
<select name="sometext" size="5" style="width: 300px;">
<option>User Data [soon]</option>
<option>User Data [soon]</option>
<option>User Data [soon]</option>
<option>User Data [soon]</option>
<option>User Data [soon]</option>
</select>
</td></tr>
<tr>
<td><input type="submit" value="Remove User">&nbsp;&nbsp;&nbsp;<input type="submit" value="Wipe List"></td>
</tr>
</table>








</td>



<td>
<form action="createresp.php" method="POST">
<input type="hidden" name="server" value="<?php echo $_SESSION['username']; ?>">
<input type="hidden" name="pulse" value="<?php echo $response->Pulse; ?>">
<input type="hidden" name="online" value="true">
<table>
<tr>
<th>&nbsp;</th>
<th>Custom Response</th>
</tr>
<tr>
<td>Addmaster:</td><td><textarea cols="50" rows="3" name="addmaster"><?php echo $response->AddMaster; ?></textarea></td></tr>
<tr class="bg">
<td>Delmaster:</td><td><textarea cols="50" rows="3" name="delmaster"><?php echo $response->DelMaster; ?></textarea></td></tr>
<tr>
<td>Prefix:</td><td><textarea cols="50" rows="3" name="prefix"><?php echo $response->Prefix; ?></textarea></td></tr>
<tr class="bg">
<td>Greet:</td><td><textarea cols="50" rows="3" name="greet"><?php echo $response->GreetMsg; ?></textarea></td></tr>
<tr>
<td>Greet Off:</td><td><textarea cols="50" rows="3" name="greetoff"><?php echo $response->GreetOff; ?></textarea></td></tr>
<tr>
<td>Bye:</td><td><textarea cols="50" rows="3" name="bye"><?php echo $response->ByeMsg; ?></textarea></td></tr>
<tr class="bg">
<td>Bye Off:</td><td><textarea cols="50" rows="3" name="byeoff"><?php echo $response->ByeOff; ?></textarea></td></tr>
<tr>
<td>Autorole:</td><td><textarea cols="50" rows="3" name="autorole"><?php echo $response->Autorole; ?></textarea></td></tr>
<tr>
<td>Autorole Off:</td><td><textarea cols="50" rows="3" name="autoroleoff"><?php echo $response->AutoroleOff; ?></textarea></td></tr>
<tr class="bg">
<td>Botrole:</td><td><textarea cols="50" rows="3" name="botrole"><?php echo $response->Botrole; ?></textarea></td></tr>
<tr>
<td>Botrole Off:</td><td><textarea cols="50" rows="3" name="botroleoff"><?php echo $response->BotroleOff; ?></textarea></td></tr>
<tr class="bg">
<td>SetPunish:</td><td><textarea cols="50" rows="3" name="setpunish"><?php echo $response->SetPunish; ?></textarea></td></tr>
<tr>
<td>SetPunish Error:</td><td><textarea cols="50" rows="3" name="setpunisherror"><?php echo $response->SetPunishError; ?></textarea></td></tr>
<tr class="bg">
<td>Give:</td><td><textarea cols="50" rows="3" name="give"><?php echo $response->Give; ?></textarea></td></tr>
<tr>
<td>Take:</td><td><textarea cols="50" rows="3" name="take"><?php echo $response->Take; ?></textarea></td></tr>
<tr class="bg">
<td>Not NSFW:</td><td><textarea cols="50" rows="3" name="notnsfw"><?php echo $response->NotNSFW; ?></textarea></td></tr>
<tr>
<td>Add NSFW</td><td><textarea cols="50" rows="3" name="addnsfw"><?php echo $response->Addnsfw; ?></textarea></td></tr>
<tr class="bg">
<td>Del NSFW</td><td><textarea cols="50" rows="3" name="delnsfw"><?php echo $response->Delnsfw; ?></textarea></td></tr>
<tr>
<td>NSFW Error1:</td><td><textarea cols="50" rows="3" name="wasnsfw"><?php echo $response->wasnsfw; ?></textarea></td></tr>
<tr class="bg">
<td>Mkchan:</td><td><textarea cols="50" rows="3" name="mkchan"><?php echo $response->Mkchan; ?></textarea></td></tr>
<tr>
<td>MkchanError</td><td><textarea cols="50" rows="3" name="mkchanerror"><?php echo $response->MkchanError; ?></textarea></td></tr>
<tr class="bg">
<td>IPError</td><td><textarea cols="50" rows="3" name="iperror"><?php echo $response->IPError; ?></textarea></td></tr>
<tr>
<td>No Role Perms:</td><td><textarea cols="50" rows="3" name="noroleperms"><?php echo $response->NoRolePerms; ?></textarea></td></tr>
<tr class="bg">
<td>Addrole:</td><td><textarea cols="50" rows="3" name="addrole"><?php echo $response->Addrole; ?></textarea></td></tr>
<tr>
<td>Role Exists:</td><td><textarea cols="50" rows="3" name="roleexists"><?php echo $response->RoleExists; ?></textarea></td></tr>
<tr class="bg">
<td>Add Role Error:</td><td><textarea cols="50" rows="3" name="addroleerror"><?php echo $response->AddRoleError; ?></textarea></td></tr>
<tr>
<td>Delrole:</td><td><textarea cols="50" rows="3" name="delrole"><?php echo $response->Delrole; ?></textarea></td></tr>
<tr class="bg">
<td>Role 404:</td><td><textarea cols="50" rows="3" name="rolenoexist"><?php echo $response->RoleNoExist; ?></textarea></td></tr>
<tr>
<td>Mute:</td><td><textarea cols="50" rows="3" name="mute"><?php echo $response->Mute; ?></textarea></td></tr>
<tr class="bg">
<td>Unmute:</td><td><textarea cols="50" rows="3" name="unmute"><?php echo $response->Unmute; ?></textarea></td></tr>
<tr>
<td>Rolecolor:</td><td><textarea cols="50" rows="3" name="rolecolor"><?php echo $response->Rolecolor; ?></textarea></td></tr>
<tr class="bg">
<td>MkInvite:</td><td><textarea cols="50" rows="3" name="mkinvite"><?php echo $response->MkInvite; ?></textarea></td></tr>
<tr>
<td>Denylinks:</td><td><textarea cols="50" rows="3" name="denylinks"><?php echo $response->DenyLinks; ?></textarea></td></tr>
<tr class="bg">
<td>Allowlinks:</td><td><textarea cols="50" rows="3" name="allowlinks"><?php echo $response->AllowLinks; ?></textarea></td></tr>
<tr>
<td>Antilink Kick:</td><td><textarea cols="50" rows="3" name="antilinkkick"><?php echo $response->AntiLinkKick; ?></textarea></td></tr>
<tr class="bg">
<td>Antilink Ban:</td><td><textarea cols="50" rows="3" name="antilinkban"><?php echo $response->AntiLinkBan; ?></textarea></td></tr>
<tr>
<td>Antilink Warn:</td><td><textarea cols="50" rows="3" name="antilinkwarn"><?php echo $response->AntiLinkWarn; ?></textarea></td></tr>
<tr class="bg">
<td>NoWarns:</td><td><textarea cols="50" rows="3" name="notnsfw"><?php echo $response->NotNSFW; ?></textarea></td></tr>
<tr>
<td>Reset Warns:</td><td><textarea cols="50" rows="3" name="resetwarns"><?php echo $response->ResetWarns; ?></textarea></td></tr>
<tr class="bg">
<td>Warn Kick:</td><td><textarea cols="50" rows="3" name="warnkick"><?php echo $response->WarnKick; ?></textarea></td></tr>
<tr>
<td>Warn Ban:</td><td><textarea cols="50" rows="3" name="warnban"><?php echo $response->WarnBan; ?></textarea></td></tr>
<tr class="bg">
<td>Warn:</td><td><textarea cols="50" rows="3" name="warn"><?php echo $response->Warn; ?></textarea></td></tr>
<tr>
<td>Warns 404:</td><td><textarea cols="50" rows="3" name="warnsnotset"><?php echo $response->WarnNotSet; ?></textarea></td></tr>
<tr class="bg">
<td>Warn Commander:</td><td><textarea cols="50" rows="3" name="warncommander"><?php echo $response->WarnCommander; ?></textarea></td></tr>
<tr>
<td>SetWarning:</td><td><textarea cols="50" rows="3" name="setwarning"><?php echo $response->SetWarning; ?></textarea></td></tr>
<tr class="bg">
<td>SetWarning 404:</td><td><textarea cols="50" rows="3" name="setwarningerror"><?php echo $response->SetWarningError; ?></textarea></td></tr>
<tr>
<td>Kick:</td><td><textarea cols="50" rows="3" name="kick"><?php echo $response->Kick; ?></textarea></td></tr>
<tr class="bg">
<td>Ban:</td><td><textarea cols="50" rows="3" name="ban"><?php echo $response->Ban; ?></textarea></td></tr>
<tr>
<td>&nbsp;</td><td><input type="submit" value="Update"></td></tr>
</table>

</form>
</td></tr>
</table>

      <div class="fix"></div>


      <div class="fix"></div>
      <!-- Text Alignment -->
     