<p class="msg info">
Ok first let's look at the trigger section.<br>
You want to set this to whatever word/sentence you want Echo to respond to.<br>
</p>

You can use the <B>{params}</B> key in the trigger to catch their text!<br>
<b>--auto &--giveme {params}={role:{params}}{req:Owner}You've assumed the role {params}</b>
<br>
<p class="msg info">
The example above <b>requires</b> you to have the role <b>Owner</b> and will give you<br>
whatever role you type (<i>As long as the role exists</i>) for example: <b>--giveme Staff</b><br>
<a href="http://echobot.tk/?nav=home&p=--&t=%26--giveme {params}&r={role:{params}}{req:Owner}You've assumed the role {params}#builder">Add to A.R.S Builder!</a>
</p>
If you want to use regex you add the <b>&</b> key before the word Example Below:<br>
<b>--auto &word=The Response here!</b>
<Br><br>
<h3 class="tit">USING THE {PARAMS} KEY</h3>
<table>
<tr>
  <th>Steps</th>
  <th>You Type</th>
</tr>
<tr>
<td>Make Command</td><td>--auto &--sayhi {params}=Your Message: {params}</td></tr>
<tr class="bg">
<td>Use Command</td><td>--sayhi hey guys!<br><a href="http://echobot.tk/?nav=home&p=--&t=%26--sayhi {params}&r=Your Message: {params}#builder">Add to A.R.S Builder!</a></td></tr>
</table>
<br>
<h3 class="tit">SERVER INFO COMMAND</h3>


<table>
<tr>
  <th>Steps</th>
  <th>You Type</th>
</tr>
<tr>
<td>Make Command</td>
<td>
<pre>
--auto .server=
SERVER INFO
--------------
prefix: {pref}
greet: {greet}
bye: {bye}
roles: {allroles}
channels: {channels}
--------------
</pre>
</td></tr>
<tr class="bg">
<td>Use Command</td><td>.server<br>
<a href="http://echobot.tk/?nav=home&p=--&t=.server&r=SERVER INFO%0D--------------%0Dprefix: {pref}%0Dgreet: {greet}%0Dbye: {bye}%0Droles: {allroles}%0Dchannels: {channels}%0D--------------#builder">Add to A.R.S Builder!</a>
</td></tr>
</table>
<p class="msg info">
The example above will display your server prefix, greet message, bye message, roles & channels if someone types <b>.server</b>
</p>

<Br>
<h3 class="tit">WORD FILTER EXAMPLE:</h3>
<table>
<tr>
  <th>Steps</th>
  <th>You Type</th>
</tr>
<tr>
<td>Make Command</td>
<td>--auto &fuck={del}{pm}{kick}You've been kicked for swearing.</td></tr>
<tr class="bg">
<td>If Typed</td><td>The message will be deleted and Echo will pm the user they were kicked.<br>
<a href="http://echobot.tk/?nav=home&p=--&t=%26fuck&r={del}{pm}{kick}You've been kicked for swearing.#builder">Add to A.R.S Builder!</a>
</td></tr>
</table> 
<Br>
<h3 class="tit">GET ALERTS ON TRIGGERS:</h3>
<table>
<tr>
  <th>Steps</th>
  <th>You Type</th>
</tr>
<tr>
<td>Make Command</td>
<td>--auto &<@YOURIDHERE>={alert:YOURIDHERE}I have alerted Proxy!</td></tr>
<tr class="bg">
<td>If Typed</td><td>Echo will pm 1-3 people if you are mentioned by anyone!<br>
<a href="http://echobot.tk/?nav=home&p=--&t=%26%3C%40YOURIDHERE%3E&r={alert:YOURIDHERE}I have alerted Proxy!#builder">Add to A.R.S Builder!</a>
</td></tr>
</table> 
<br>
<h3 class="tit">CUSTOM GETID COMMAND</h3>
<table>
<tr>
  <th>Steps</th>
  <th>You Type</th>
</tr>
<tr>
<td>Make Command</td>
<td>--auto &--grabid {params}={rawid}The user's ID: {params}</td></tr>
<tr class="bg">
<td>Use Command</td><td>--grabid @User<br>
<b>{rawid}</b> is required along with <b>{params}</b> in the response to display another users ID.<br> 
<a href="http://echobot.tk/?nav=home&p=--&t=%26--grabid {params}&r={rawid}The user's ID: {params}#builder">Add to A.R.S Builder!</a></td></tr>
</table>


<h3 class="tit">MAKE THE GIPHY COMMAND</h3>
<table>
<tr>
<th>Steps</th>
<th>You Type</th>
</tr>
<tr>
<td>Make Command</td><td>--auto &--giphy {params}={giphy}</td></tr>
<tr class="bg">
<td>Use Command</td><td>--giphy Keyword Here<br>
<a href="http://echobot.tk/?nav=home&p=--&t=%26--giphy {params}&r={giphy}#builder">Add to A.R.S Builder!</a></td></tr>
</table>



<h3 class="tit">MAKE THE CATS COMMAND</h3>
<table>
<tr>
<th>Steps</th>
<th>You Type</th>
</tr>
<tr>
<td>Make Command</td><td>--auto --cats={cats}</td></tr>
<tr class="bg">
<td>Use Command</td><td>--cats<br>
<a href="http://echobot.tk/?nav=home&p=--&t=--cats&r={cats}#builder">Add to A.R.S Builder!</a></td></tr>
</table>



<h3 class="tit">MAKE THE MEME COMMAND</h3>
<table>
<tr>
<th>Steps</th>
<th>You Type</th>
</tr>
<tr>
<td>Make Command</td><td>--auto --meme={meme}</td></tr>
<tr class="bg">
<td>Use Command</td><td>--meme<br>
<a href="http://echobot.tk/?nav=home&p=--&t=--meme&r={meme}#builder">Add to A.R.S Builder!</a></td></tr>
</table>


<h3 class="tit">MAKE THE JOKE COMMAND</h3>
<table>
<tr>
<th>Steps</th>
<th>You Type</th>
</tr>
<tr>
<td>Make Command</td><td>--auto --joke={joke}</td></tr>
<tr class="bg">
<td>Use Command</td><td>--joke<br>
<a href="http://echobot.tk/?nav=home&p=--&t=--joke&r={joke}#builder">Add to A.R.S Builder!</a></td></tr>
</table>


<h3 class="tit">MAKE THE ASS COMMAND</h3>
<table>
<tr>
<th>Steps</th>
<th>You Type</th>
</tr>
<tr>
<td>Make Command</td><td>--auto --ass={ass}</td></tr>
<tr class="bg">
<td>Use Command</td><td>--ass<br>
<a href="http://echobot.tk/?nav=home&p=--&t=--ass&r={ass}#builder">Add to A.R.S Builder!</a></td></tr>
</table>


<h3 class="tit">MAKE THE BOOBS COMMAND</h3>
<table>
<tr>
<th>Steps</th>
<th>You Type</th>
</tr>
<tr>
<td>Make Command</td><td>--auto --boobs={boobs}</td></tr>
<tr class="bg">
<td>Use Command</td><td>--boobs<br>
<a href="http://echobot.tk/?nav=home&p=--&t=--boobs&r={boobs}#builder">Add to A.R.S Builder!</a></td></tr>
</table>




<!--
<table>
<tr>
  <th>Command</th>
  <th>Info</th>
  <th>Usage</th>
  <th>Permissions</th>
</tr>
<tr>
<td>help</td><td>Displays a list of Echo's commands.</td><td>--help</td><td>Everyone</td></tr>
<tr class="bg">
<td>addmaster</td><td>Gives someone access to all mod commands!</td><td>--addmaster @User</td><td>Commanders</td></tr>
<tr>
<td>greet</td><td>Says the message when someone joins<br>Use <b>{pm}</b> to pm the greet. [more]</td><td>--greet Message</td><td>Commanders</td></tr>
<tr class="bg">
<td>bye</td><td>Says the message when someone leaves</td><td>--bye Message</td><td>Commanders</td></tr>
<tr>
<td>denylinks</td><td>Turns the AntiLinks system on.</td><td>--denylinks</td><td>Commanders</td></tr>
<tr class="bg">
<td>allowlinks</td><td>Turns the AntiLinks system off</td><td>--allowlinks</td><td>Commanders</td></tr>
<tr>
<td>prefix</td><td>Changes Echo's prefix from --</td><td>--setprefix +</td><td>Commanders</td></tr>
<tr class="bg">
<td>autorole</td><td>Auto roles someone when they join.<br>For silent add <b>-s</b> before role name.</td><td>--autorole Role Name</td><td>Commanders</td></tr>
<tr>
<td>invites</td><td>Grabs a list of invites for the channel.</td><td>--invites</td><td>Everyone?</td></tr>
<tr class="bg">
<td>kick</td><td>Kicks non-commanders from the server.</td><td>--kick @User</td><td> Commanders</td></tr>
<tr>
<td>ban</td><td>Bans non-commanders from the servers.</td><td>--ban @User</td><td>Commanders</td></tr>
<tr class="bg">
<td>giveme</td><td>You need to build this in the A.R.S [more]</td><td>{role:Role Name}</td><td>Custom</td></tr>
<tr>
<td>setpunish</td><td>Set's the punishment for AntiLink and --warn</td><td>--setpunish kick\ban\warn</td><td>Commanders</td></tr>
<tr class="bg">
<td>meme</td><td>You need to build this in the A.R.S [more]</td><td>{meme}</td><td>Custom</td></tr>
<tr>
<td>joke</td><td>You need to build this in the A.R.S [more]</td><td>{joke}</td><td>Custom</td></tr>
<tr class="bg">
<td>give</td><td>Gives someone a role.</td><td>--give @User Role Name</td><td>Commanders</td></tr>
<tr>
<td>take</td><td>Takes a role away from someone.</td><td>--take @User Role Name</td><td>Commanders</td></tr>
<tr class="bg">
<td>mute</td><td>Mutes a user in your channel.</td><td>--mute @User</td><td>Commanders</td></tr>
<tr>
<td>unmute</td><td>Unmutes a user in your channel.</td><td>--unmute @User</td><td>Commanders</td></tr>
<tr class="bg">
<td>rolecolor</td><td>Changes role color from hex.</td><td>--rolecolor #000000 Role Name</td><td>Commanders</td></tr>
<tr>
<td>giphy</td><td>You need to build this in the A.R.S [more]</td><td>{giphy}</td><td>Custom</td></tr>
<tr class="bg">
<td>cats</td><td>You need to build this in the A.R.S [more]</td><td>{cats}</td><td>Custom</td></tr>
<tr>
<td>auto</td><td>Adds a trigger to your A.R.S Database.</td><td>--auto Hello=Hey {user}!</td><td>Commanders</td></tr>
<tr class="bg">
<td>delauto</td><td>Deletes an A.R.S Trigger.</td><td>--delauto triggername</td><td>Commanders</td></tr>
<tr>
<td>viewauto</td><td>Views your A.R.S Files in chat.</td><td>--viewauto</td><td>Commanders</td></tr>
<tr class="bg">
<td>wipeauto</td><td>Wipes your A.R.S File completely.</td><td>--wipeauto</td><td>Commanders</td></tr>
<tr>
<td>botrole</td><td>Auto role bots when they join.<br>for silent add <b>-s</b> before the role name.</td><td>--botrole Role Name</td><td>Commanders</td></tr>
<tr class="bg">
<td>mkchan</td><td>Creates a new channel in your server. Text or Voice.</td><td>--mkchan chan-name text</td><td>Commanders</td></tr>
<tr>
<td>locateip</td><td>Geo Location for an IP Address or Domain.</td><td>--locateip IP or .com</td><td>Everyone</td></tr>
<tr class="bg">
<td>channelid</td><td>Grabs the current channels ID.</td><td>--channelid</td><td>Everyone</td></tr>
<tr>
<td>getid</td><td>Grabs a user's ID.</td><td>--getid @User</td><td>Everyone</td></tr>
<tr class="bg">
<td>nsfw</td><td>Enables the nudity commands in current channel.<br>You have to do this for every channel.</td><td>--nsfw true</td><td>Commanders</td></tr>
<tr>
<td>grabars</td><td>Echo will send your A.R.S File to your pm via file attachment.</td><td>--grabars</td><td>Commanders</td></tr>
<tr class="bg">
<td>putars</td><td>Imports A.R.S from local to Echo! needs a direct link to your json file!</td><td>--putars directlink</td><td>Commanders</td></tr>
<tr>
<td>addrole</td><td>Creates a new role in your server.</td><td>--addrole Role Name</td><td>Commanders</td></tr>
<tr class="bg">
<td>delrole</td><td>Deletes the role from your server.</td><td>--delrole Role Name</td><td>Commanders</td></tr>
<tr>
<td>setwarning</td><td>Set's the max amounts of warns before kick\ban<br>works off of `--setpunish`</td><td>--setwarning 3</td><td>Commanders</td></tr>
<tr class="bg">
<td>warn</td><td>Adds a warning point to a user. (will kick\ban non-commanders).</td><td>--warn @User</td><td>Commanders</td></tr>
<tr>
<td>listwarns</td><td>Lists everyone who has a warning. or nothing if empty!</td><td>--listwarns</td><td>Commanders</td></tr>
<tr class="bg">
<td>delwarn</td><td>Deletes the users warnings.</td><td>--delwarn @User</td><td>Commanders</td></tr>
<tr>
<td>teemo</td><td>Creates a custom teemo banner from your text.</td><td>--temo Text=Text{br}Text</td><td>Everyone</td></tr>
<tr class="bg">
<td>vayne</td><td>Creates a custom vayne banner from your text.</td><td>--vayne Text=Text{br}Text</td><td>Everyone</td></tr>
<tr>
<td>ekko</td><td>Creates a custom Ekko banner from your text.</td><td>--ekko Text=Text{br}Text</td><td>Everyone</td></tr>
<tr class="bg">
<td>zed</td><td>Creates a custom Zed banner from your text.</td><td>--zed Text=Text{br}Text</td><td>Everyone</td></tr>
<tr>
<td>cute</td><td>Displays cute Cats \ Dogs \ Babies and more.</td><td>--cute</td><td>Everyone</td></tr>
<tr class="bg">
<td>cars</td><td>Displays random cars!</td><td>--cars</td><td>Everyone</td></tr>
<tr>
<td>trucks</td><td>Displays random trucks!</td><td>--trucks</td><td>Everyone</td></tr>
<tr class="bg">
<td>sky</td><td>Displays random sky images!</td><td>--sky</td><td>Everyone</td></tr>
<tr>
<td>space</td><td>Displays random space images!</td><td>--space</td><td>Everyone</td></tr>
<tr class="bg">
<td>lesbian</td><td>Displays random NSFW Images!</td><td>--lesbain</td><td>Everyone</td></tr>
<tr>
<td>twink</td><td>Displays random NSFW Images!</td><td>--twink</td><td>Everyone</td></tr>
<tr class="bg">
<td>upskirt</td><td>Displays random NSFW Images!</td><td>--upskirts</td><td>Everyone</td></tr>
<tr>
<td>fatgirls</td><td>Displays random NSFW Images!</td><td>--fatgirls</td><td>Everyone</td></tr>
<tr class="bg">
<td>celebs</td><td>Displays random NSFW Images!</td><td>--celebs</td><td>Everyone</td></tr>
</table>
-->

      <div class="fix"></div>


      <div class="fix"></div>
      <!-- Text Alignment -->
     
