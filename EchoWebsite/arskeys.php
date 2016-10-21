<table>
<tr>
  <th>Key</th>
  <th>Usage</th>
  <th>Info</th>
</tr>
<tr>
<td>{pm}</td><td>{pm}</td><td>Echo will private message the response.</td></tr>
<tr class="bg"><td>{user}</td><td>{user}</td><td>Echo will mention the user.</td></tr>
<tr><td>{/user}</td><td>{/user}</td><td>Echo will say the user's name</td></tr>
<tr class="bg"><td>{del}</td><td>{del}</td><td>Echo will delete the user's message.</td></tr>
<tr class="bg"><td>{kick}</td><td>{kick}</td><td>Echo will kick Non-Commanders</td></tr>
<tr><td>{ban}</td><td>{ban}</td><td>Echo will ban Non-Commanders</td></tr>
<tr class="bg"><td>{role}</td><td>{role:Role Name}</td><td>Gives a user a specific role.</td></tr>
<tr><td>{alert}</td><td>{alert:YOURID}</td><td>Alerts you if someone triggers reponse<Br>You can also do multiple alerts<Br>{alert:YourID,AnotherID}</td></tr>
<tr class="bg"><td>{exc}</td><td>{exc:Role Name}</td><td>Excludes a role from your trigger.!<br>You can exclude multiple roles!<br>{exc:Role name,Role Name}</td></tr>
<tr><td>{chan}</td><td>{chan}</td><td>Returns the current channel.</td></tr>
<tr class="bg"><td>{pref}</td><td>{pref}</td><td>Displays echo's current prefix.</td></tr>
<tr><td>{greet}</td><td>{greet}</td><td>Displays echo's current greeting.</td></tr>
<tr class="bg"><td>{bye}</td><td>{bye}</td><td>Displays echo's current bye message.</td></tr>
<tr><td>{ismaster}</td><td>{ismaster}</td><td>Displays if the user is master or not `True` or `False`</td></tr>
<tr class="bg"><td>{listroles}</td><td>{listroles}</td><td>Displays all the users roles.</td></tr>
<tr><td>{allroles}</td><td>{allroles}</td><td>Displays all roles in server.</td></tr>
<tr class="bg"><td>{joined}</td><td>{joined}</td><td>Shows the date and time a user has joined.</td></tr>
<tr><td>{channels}</td><td>{channels}</td><td>Shows all channels in the server.</td></tr>
<tr class="bg"><td>{meme}</td><td>{meme}</td><td>Shows random memes.</td></tr>
<tr><td>{joke}</td><td>{joke}</td><td>Shows random jokes.</td></tr>
<tr class="bg"><td>{params}</td><td>{params}</td><td>if you have this key in your trigger and response<br>you can catch their text!<br>`--auto &--announce {params}=**ANNOUNCEMENT:**{params}`</td></tr>
<tr><td>{req}</td><td>{req:Role Name}</td><td>Requires a user to have a role. you can also require multiple roles like `{exc}`</td></tr>
<tr class="bg"><td>{ass}</td><td>{ass}</td><td>Shows random ass.</td></tr>
<tr><td>{boobs}</td><td>{boobs}</td><td>Shows random boobs.</td></tr>
<tr class="bg"><td>{warn}</td><td>{warn:5}</td><td>Adds a warning point to a user.<Br>be default `--setpunish` is set to `kick`</td></tr>
<tr><td>{msg}</td><td>{msg:Text Here}</td><td>The warning message before he kicks\bans. Only works if`{warn}` key is found.</td></tr>
<tr class="bg"><td>{getid}</td><td>{getid}</td><td>Grabs the users id.</td></tr>
<tr><td>{redirect}</td><td>{redirect:CHANNELID}</td><td>Redirects users text to a specific channel.</td></tr>
<tr class="bg"><td>{ifchan}</td><td>{ifchan:CHANNELID}</td><td>Requires a certain channel for your response.</td></tr>
<tr><td>{nsfw}</td><td>{nsfw}</td><td>Requires your channel to be NSFW. <b>--nsfw true</b></td></tr>
<tr class="bg"><td>{cats}</td><td>{cats}</td><td>Display random cat images.</td></tr>
<tr><td>{take}</td><td>{take:Role Name}</td><td>Takes a role away from someone.</td></tr>
<tr class="bg">
<td>{if}</td><td>{if:user==Proxy}<br>{if:channel==general}<br>{if:user!=Proxy}<br>{if:channel!=general}</td><td>If statements for echo. currently supports users and channels.</td></tr>
<tr>
<td>{delauto}</td><td>{delauto}</td><td>Deletes the trigger after first use. (Good for games, prizes and keys you want to give out once!)</td></tr>
<tr class="bg">
<td>{params:flip}</td><td>{params:flip}</td><td>Reverses the text. Requires <b>{params}</b> in the trigger!</td></tr>
<tr>
<td>{params:url}</td><td>{params:url}</td><td>converts text to url encoding. Replaces space with <b>%20</b><br>Requires <b>{params}</b> in the trigger!</td></tr>
<tr class="bg">
<td>{twitch}</td><td>{twitch:username}</td><td>Replace username with your twitch name. you have access to these keys:<br>
{game} displays the game you're playing.<br>
{views} shows your current viewers.<br>
{fps} displays your current fps.<br>
{url} - shows the url.</td></tr>
<tr>
<td>{nobots}</td><td>{nobots}</td><td>Excludes bots from your trigger.</td></tr>
<tr class="bg">
<td>{bots}</td><td>{bots}</td><td>Only works if the user is a bot.</td></tr>
</table>



      <div class="fix"></div>


      <div class="fix"></div>
      <!-- Text Alignment -->
     
