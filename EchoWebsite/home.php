      <div class="tabs box">
        <ul>
          <li><a href="#setup"><span>Setup Echo</span></a></li>
          <li><a href="#issues"><span>Having Issues?</span></a></li>
          <li><a href="#tricks"><span>A.R.S Tricks</span></a></li>
          <li><a href="#builder"><span>A.R.S Builder</span></a></li>
          <li><a href="#emojis"><span>Emojis List</span></a></li>
        </ul>
      </div>
      <!-- /tabs -->
      <!-- Tab01 -->
      <div id="setup">
        <p>
<h3 class="tit">Getting Started</h3>
After you invite Echo to your server he's ready to be used! the default prefix in your server is <b>--</b><br>
You can view a list of his commands from the <b>Top Navigation</b> or the <b>Sidebare</b>
<br>

        </p>
      </div>
      <!-- /tab01 -->
      <!-- Tab02 -->
      <div id="issues">
        <p>

 <h3 class="tit">Need Help?</h3>
You can visit our server and ask any questions: <a href="https://discord.gg/0pTKzt2BDInBOrxL" target="_new">Our Discord Server</a><Br>
Make sure to give Echo the full permissions still having issues?<br>

<b>Echo could have issues giving roles if he's not an Administrator.<br>
He will also have issues editing/giving/taking roles away from someone if they're Administrator<br>
And his role is lower than theirs on the list. read below how to fix:<br><br>

<b><u>Go into your server settings</u></b> Than click <b><u>Roles</u></b> and Drag Echo [BETA] to the top. this fixes most problems.<br>
Also be sure he has Manage Roles permissions.<br>
<br>
  <b>If the emoji is not at the beginning of your message and has a space  
  Replace with an underscore :baby girl: would be :baby_girl:</b>
<Br>
You can get help on any command by typing <b>--help cmdname</b><br>
More information in this section soon!
        </p>
      </div>
      <!-- /tab02 -->
      <!-- Tab03 -->
      <div id="tricks">
        <p>
<h3 class="tit">Auto Response System</h3>
This information will be filled soon!
        </p>
      </div>
      <!-- /tab03 -->
       <!-- Tab03 -->


<?php
if($_GET['t']!= "") {
	$trig = $_GET['t'];
	$resp = $_GET['r'];
	$pref = $_GET['p'];
} else {
	$trig = "&fuck";
	$resp = "";
	$pref = "--";
}
?>
<div id="builder">
<p>
<form class="form-container">
<div class="form-title"><center><font size="5"><b>Auto Response Builder</b></font></center></div>
<table style="border: 0px;" border="0" cellpadding-"0" cellspacing="0">
<tr>
<td style="border: 0px;">
<div class="form-title">Echo's Prefix</div>
<input class="small-field" type="text" name="prefix" id="prefix" value="<?php echo $pref; ?>" />
</td>
<td style="border: 0px;">
<div class="form-title">Trigger</div>
<input class="small-field" type="text" name="trigger" id="trigger" value="<?php echo $trig; ?>" />
</td></tr>
</table>

<div class="form-title">Response Action</div>
<select id="ars" name="ars" class="form-field" onchange="addText('resp', document.getElementById('ars').value)">
  <option selected="selected">Select a key</option>
  <option value="{user}">Mention the User</option>
  <option value="{/user}">Say users name</option>
  <option value="{pm}">Pm the User</option>
  <option value="{del}">Delete their Message</option>
  <option value="{kick}">Kick the user</option>
  <option value="{ban}">Ban the user</option>
  <option value="{role:Role Name Here}">Give someone a role</option>
  <option value="{alert:YourIDHere}">Alert Someone Via PM</option>
  <option value="{exc:Role Name,RoleName2,etc}">Exclude role(s)</option>
  <option value="{chan}">Display Channel Name</option>
  <option value="{pref}">Display Prefix</option>
  <option value="{greet}">Display Greet</option>
  <option value="{bye}">Display Bye Message</option>
  <option value="{ismaster}">Display if user is master</option>
  <option value="{listroles}">Display users roles.</option>
  <option value="{allroles}">Display Server roles.</option>
  <option value="{joined}">Display user's join date</option>
  <option value="{channels}">Display all channels in server</option>
  <option value="{meme}">Display random memes</option>
  <option value="{joke}">Display random Joke</option>
  <option value="{params}">Catch someones text</option>
  <option value="{req:Role Name,Role Name2}">Require role(s)</option>
  <option value="{ass}">Display random asses</option>
  <option value="{boobs}">Display random boobs</option>
  <option value="{warn:5}">Warn user x times</option>
  <option value="{msg:Your Message Here}">Warning Message</option>
  <option value="{getid}">Grab users id</option>
  <option value="{redirect:CHANNELID}">Redirect text to channel</option>
  <option value="{ifchan:CHANNELID}">Require channel(s)</option>
  <option value="{nsfw}">Requires channel to be NSFW</option>
  <option value="{cats}">Display random cat images</option>
  <option value="{take:Role Name Here}">Take a role from someone</option>
  <option value="{if:user==UsernameHere}">If Statement: User</option>
  <option value="{if:channel==ChannelName}">If Statement: Channel</option>
  <option value="{if:user!=UsernameHere}">If Statement: Exclude User</option>
  <option value="{if:channel!=ChannelName}">If Statement: Exclude Channel</option>
  <option value="{twitch:nl_kripp}playing: {game} viewers: {views} fps: {fps} watch: {url}">Show twitch stream info!</option>
</select><br />



<div class="form-title">Presets</div>
<select id="preset" name="preset" class="form-field" onchange="addText2('resp', document.getElementById('preset').value)">
<option selected="selected">Select a Preset</option>
  <option value="{del}{kick}I have kicked {user} for swearing.">Basic Filter with Kick</option>
  <option value="{del}{ban}I have banned {user} for swearing.">Basic Filter with Ban</option>
  <option value="{del}{warn:3}{msg:You have been Warned!}{kick}I have kicked {user} for swearing.">Basic Filter with Warning Than Kick</option>
  <option value="{del}{warn:3}{msg:You have been Warned!}{ban}I have banned {user} for swearing.">Basic Filter with Warning Than Ban</option>
  <option value="{del}{kick}{exc:Role Name}I have kicked {user} for swearing.">Basic Filter with Kick Excluding Role</option>
  <option value="{del}{ban}{exc:Role Name}I have banned {user} for swearing.">Basic Filter with Ban Excluding Role</option>
  <option value="{del}{kick}{req:Role Name}I have kicked {user} for swearing.">Basic Filter with Kick Requiring Role</option>
  <option value="{del}{ban}{req:Role Name}I have banned {user} for swearing.">Basic Filter with Ban Requiring Role</option>
  <option value="{pm}{del}{ban}{exc:Role Name}I have banned {user} for swearing.">Basic Filter with Ban Excluding Role (PM user)</option>
  <option value="{pm}{del}{ban}{req:Role Name}I have banned {user} for swearing.">Basic Filter with Kick Requiring Role (PM User)</option>
  <option value="{pm}{role:Role Name}I have given you the role `Role Name`">Give someone a role (Pm User)</option>
  <option value="{alert:YOURIDHERE}I have alerted `Your Name` that you need help.">Get an Alert PM.</option>
  <option value="{exc:Role Name1,Role Name2,Role Name3}">Exclude Multiple Roles.</option>
  <option value="{req:Role Name1,Role Name2,Role Name3}">Require Multiple Roles.</option>
  <option value="{alert:YOURID,THEIRID,ANOTHERID,ANDANOTHER}">Alert Multiple People via PM</option>
  <option value="{redirect:CHANNELID}This text wil be redirected!">Redirect text to a channel</option>
  <option value="{ifchan:CHANNELID}">Require a specific channel</option>
  <option value="{boobs}{nsfw}">Show Boobs with NSFW Restrictions</option>
  <option value="{ass}{nsfw}">Show Asses with NSFW Restrictions</option>
  <option value="{joke}{{nsfw}">Show jokes with NSFW Restrictions</option>
  <option value="{meme}{nsfw}">Show memes with NSFW Restrictions</option>
  <option value="{if:user==Proxy}">Only show response for user: Proxy</option>
  <option value="{if:channel==general}">Only show response for channel: general</option>
  <option value="{cats}{if:channel==general}">Only allow cat images in general channel</option>
  <option value="{if:user==Proxy}This response will only work if the user's name is Proxy">Show response if user is Proxy.</option>
  <option value="{if:user!=Proxy}This response will work for anyone but Proxy">Show response to everyone but a user.</option>
  <option value="{if:channel!=general}This response will work for all channels but general.">Show response in all channels but general</option>
</select><br />


<div class="form-title">Response</div>
<textarea class="form-field" id="resp" name="response" cols="25" rows="5" placeholder="{del}{warn:3}{msg:Warned for swearing!}{kick}I have kicked {user} for swearing" /><?php echo $resp; ?></textarea><br />
<div class="form-title">Results:</div>
<!-- <input class="form-field" id="results" type="text" placeholder="waiting.."> -->
<textarea class="form-field" id="results" name="results" cols="25" rows="5" placeholder="Waiting.."></textarea>
<div id="info" style="font-color: #FF0000;"></div>
<div class="submit-container">
<input class="submit-button" type="button" value="Generate" onclick="makears()" />
</div>
</form>

</p>
</div>



<div id="emojis">
<img src="https://raw.githubusercontent.com/proxikal/Echo/master/emojis1.png">  
<img src="https://raw.githubusercontent.com/proxikal/Echo/master/emojis2.png">  
<img src="https://raw.githubusercontent.com/proxikal/Echo/master/emojis3.png">  
<img src="https://raw.githubusercontent.com/proxikal/Echo/master/emojis4.png">
</div>
      <!-- 2 columns -->
      <h3 class="tit">News</h3>
      <div class="col50">
        <p class="t-justify">We're working on a system which will allow <b>Server Owners</b> To access and edit their database from this website!</p>
      </div>
      <!-- /col50 -->
      <!--
      <div class="col50 f-right">
        <p class="t-justify">More content Here :D</p>
      </div>
       /col50 -->
      <div class="fix"></div>


      <div class="fix"></div>
      <!-- Text Alignment -->
     

<div id="blanket" style="display:none"></div>
<div id="popUpDiv" style="display:none">
<form class="form-container" action="" method="POST">
<div class="form-title"><center><font size="5"><b>More data Required:</b></font></center></div>
<div class="form-title">Data</div>
<input class="form-field" type="text" name="firstname" placeholder="&fuck" /><br />
<div class="form-title">Response</div>

<textarea class="form-field" name="response" cols="25" rows="10" placeholder="{del}{warn:3}{msg:Warned for swearing!}{kick}I have kicked {user} for swearing" /></textarea><br />
<div class="submit-container">
<input class="submit-button" type="submit" value="Generate" />
</div>
</form>

<a href="#" onclick="popup('popUpDiv')" >Close Window</a>
</div>