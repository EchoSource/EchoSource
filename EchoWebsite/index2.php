<html>
<head>
<title>Echo's Revenge</title>
<meta http-equiv="content-type" content="text/html; charset=utf-8" />
<link rel="stylesheet" media="screen,projection" type="text/css" href="css/reset.css" />
<link rel="stylesheet" media="screen,projection" type="text/css" href="css/main.css" />
<link rel="stylesheet" media="screen,projection" type="text/css" href="css/2col.css" title="2col" />
<link rel="alternate stylesheet" media="screen,projection" type="text/css" href="css/1col.css" title="1col" />
<!--[if lte IE 6]><link rel="stylesheet" media="screen,projection" type="text/css" href="css/main-ie6.css" /><![endif]-->
<link rel="stylesheet" media="screen,projection" type="text/css" href="css/style.css" />
<link rel="stylesheet" media="screen,projection" type="text/css" href="css/mystyle.css" />
<link rel="shortcut icon" href="favicon.ico" type="image/x-icon" />
<link rel="icon" href="favicon.ico" type="image/x-icon" />

<script type="text/javascript" src="js/jquery.js"></script>
<script type="text/javascript" src="js/switcher.js"></script>
<script type="text/javascript" src="js/toggle.js"></script>
<script type="text/javascript" src="js/ui.core.js"></script>
<script type="text/javascript" src="js/ui.tabs.js"></script>
<script type="text/javascript">
	$(document).ready(function(){
		$(".tabs > ul").tabs();
	});
	</script>
</head>
<body>
<?php
$servers=file_get_contents('info.txt');
?>
<div id="main">
  <!-- Tray -->
  <div id="tray" class="box">
    <p class="f-left box">
      <!-- Switcher -->
      <span class="f-left" id="switcher"> <a href="javascript:void(0);" rel="1col" class="styleswitch ico-col1" title="Display one column"><img src="design/switcher-1col.gif" alt="1 Column" /></a> <a href="javascript:void(0)" rel="2col" class="styleswitch ico-col2" title="Display two columns"><img src="design/switcher-2col.gif" alt="" /></a> </span> Project: <strong>Echo</strong> </p>
    <p class="f-right">User: <strong><a href="#">Administrator</a></strong> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; <strong><a href="#" id="logout">Log out</a></strong></p>
  </div>
  <!--  /tray -->
  <hr class="noscreen" />
  <!-- Menu -->
  <div id="menu" class="box">
    <ul class="box f-right">
      <li><a href="https://www.carbonitex.net/Discord/bots" target="_new"><span><strong>More Bots &raquo;</strong></span></a></li>
    </ul>
    <ul class="box">
      <li id="menu-active"><a href="#"><span>Home</span></a></li>
      <!-- Active -->
      <li><a href="#"><span>Commands</span></a></li>
      <li><a href="#"><span>Documentation</span></a></li>
      <li><a href="#"><span>Request Feature</span></a></li>
      <li><a href="#"><span>Online Support</span></a></li>
      <li><a href="#"><span>Login</span></a></li>
      <li><a href="#"><span>Signup</span></a></li>
    </ul>
  </div>
  <!-- /header -->
  <hr class="noscreen" />
  <!-- Columns -->
  <div id="cols" class="box">
    <!-- Aside (Left Column) -->
    <div id="aside" class="box">
      <div class="padding box">
        <!-- Logo (Max. width = 200px) -->
        <p id="logo"><a href="#"><img src="tmp/logo.gif" alt="" /></a></p>
        <!-- Search -->
        <form action="#" method="get" id="search">
          <fieldset>
          <legend>Search Documents</legend>
          <p>
            <input type="text" size="17" name="" class="input-text" />
            &nbsp;
            <input type="submit" value="OK" class="input-submit-02" />
            <br />
            <a href="javascript:toggle('search-options');" class="ico-drop">Advanced search</a></p>
          <!-- Advanced search -->
          <div id="search-options" style="display:none;">
            <p>
              <label>
              <input type="checkbox" name="" checked="checked" />
              Option I.</label>
              <br />
              <label>
              <input type="checkbox" name="" />
              Option II.</label>
              <br />
              <label>
              <input type="checkbox" name="" />
              Option III.</label>
            </p>
          </div>
          <!-- /search-options -->
          </fieldset>
        </form>
        <!-- Create a new project -->
        <p id="btn-create" class="box"><a href="#"><span>Invite Echo Today!</span></a></p>
      </div>
      <!-- /padding -->
      <ul class="box">
        <li id="submenu-active"><a href="#">Home</a></li>
        <li><a href="#">Commands</a></li>
        <li><a href="#">Request Feature</a></li>
        <li><a href="#">Documentation</a>
          <!-- Active -->
          <ul>
            <li><a href="#">Base Commands</a></li>
            <li><a href="#">Auto Response System</a></li>
            <li><a href="#">Bot Commander</a></li>
            <li><a href="#">Some A.R.S Examples</a></li>
          </ul>
        </li>
        <li><a href="#">Your Server</a>
          <ul>
            <li><a href="#">Server Database</a></li>
            <li><a href="#">Auto Response Database</a></li>
            <li><a href="#">Backups</a></li>
          </ul>
        </li>
        <li><a href="#">24/7 Support</a></li>
      </ul>
    </div>
    <!-- /aside -->
    <hr class="noscreen" />
    <!-- Content (Right Column) -->
    <div id="content" class="box">
      <!-- <h1>Styles</h1> -->
      <!-- Headings -->
      <!-- <h2>Heading H2</h2>
      <h3>Heading H3</h3>
      <h4>Heading H4</h4>
      <h5>Heading H5</h5> -->
      <!-- System Messages -->

      <h3 class="tit">System Messages</h3>
      <!-- <p class="msg warning">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p> -->
      <p class="msg info">We'll be updating the website frequently. keep checking back!</b></p>
      <p class="msg done">Echo is currently on <b>(<?php echo $servers; ?>)</b> Servers.</p>
     <!-- <p class="msg error">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p> -->
      <!-- Tabs -->
      <h3 class="tit">Common Issues</h3>
      <div class="tabs box">
        <ul>
          <li><a href="#tab01"><span>Setup</span></a></li>
          <li><a href="#tab02"><span>Basic</span></a></li>
          <li><a href="#tab03"><span>Auto Response</span></a></li>
        </ul>
      </div>
      <!-- /tabs -->
      <!-- Tab01 -->
      <div id="tab01">
        <p>
          <iframe src="https://discordapp.com/widget?id=148629493676769280&theme=dark" width="350" height="500" allowtransparency="true" frameborder="0"></iframe>
        </p>
      </div>
      <!-- /tab01 -->
      <!-- Tab02 -->
      <div id="tab02">
        <p>

 <h3 class="tit">Need Help?</h3>
You can visit our server and ask any questions: <a href="https://discord.gg/0pTKzt2BDInBOrxL" target="_new">Our Discord Server</a><Br>
You can try to fix him yourself. but make sure you're typing an existing command just type `--help` in the server.<br>
Still not responding? <b>Use this as a last resort. Echo backs your database up every 2 hours <u>on the hour</u></b><br>
Private message `--help` to echo and he will walk you through grabbing your last working configuration.<br>
<Br>
<h3 class="tit">Echo's Commands</h3>

<table>
<tr>
  <th>--addmaster</th>
  <th>--delmaster</th>
  <th>--autorole</th>
</tr>
<tr>
<td>--byemsg</td><td>--getid</td><td>--rolecolor</td></tr>
<tr>
<td>--greet</td><td>--setprefix</td><td>--youtube</td></tr>
<tr>
<td>--ban</td><td>--kick</td><td>--give</td></tr>
<tr>
<td>--giveme</td><td>--take</td><td>--t</td></tr>
<tr>
<td>--giphy</td><td>--sticker</td><td>--flush</td></tr>
<tr>
<td>--coinflip</td><td>--8ball</td><td>--rolemonitor</td></tr>
<tr>
<td>--auto</td><td>--mute</td><td>--unmute</td></tr>
</table>
<Br>
You can get help on any command by typing <b>--help cmdname</b><br>

<pre><font color="#3364B7">
--rolecolor: Example: --rolecolor #FF0000 Role Name
--byemsg: Example: --byemsg {user} has left the server
--autorole Example: --autorole Role Name auto assigns role to new members.
--greet Example: --greet Welcome {user} if you need any help ask Proxy! you can use {/user} for non-mentions.
--give Example: --give @user Role Name
--take Example: --take @user Role Name
--giveme Example: --giveme Role Name
--rolemonitor Description: Alerts you when a role has been edited in your server.
--auto Description: Toggles ON\OFF Echo's Auto Response System.
--t Example: --t en-fr Hello how are you? translates text from one language to another.
</font>
</pre>


        </p>
      </div>
      <!-- /tab02 -->
      <!-- Tab03 -->
      <div id="tab03">
        <p>
<h3 class="tit">Auto Response System</h3>
The command is only for <b><U>Server Owners</U></b> no one else has access to add triggers.<br>
You need to type <b>--auto</b> to initiate the auto response for the first time!<br>
In the future you can toggle Auto Response <u>On\Off</u> by typing <b>--auto</b><br>
<B>to add a new trigger type <b>start::</b><br>
to delete a trigger type <b>delete::the trigger word</b><br>
to list all your responses type <b>list::</b><br>
to wipe your Auto Response Database type <b>wipe::</b><br>
<br><br>
<h3 class="tit">Accepted Response Key's</h3>
<ul>
  <li>{exc=USERID}</li>
  <li>{mock}</li>
  <li>{mock2}</li>
  <li>{greet}</li>
  <li>{user}</li>
  <li>{/user}</li>
  <li>{kick}</li>
  <li>{pm}</li>
  <li>{pm=USERID}</li>
  <li>{pref}</li>
  <li>{rand1}</li>
  <li>{meme1}</li>
  <li>{date1}</li>
  <li>{date2}</li>
  <li>{time}</li>
  <li>{del} <i>(THIS KEY IS CURRENTLY BUGGY)</i></li>
  <li>{cmd=} <i>More details: <a href="https://github.com/proxikal/Echo/blob/master/README.md#remake-echos-commands-with-custom-triggers">Click Here</a></i></li>
</ul>

<h3 class="tit"><b>{exc=USERID}</b>:</h3>
Excludes user(s) from the auto response. Example 1: `{exc=USERID}` or `{exc=USERID,USERID,USERID}`<br>
Echo can get a user's id by typing `--getid @User`<br>
<br>
<h3 class="tit"><b>{mock}</b>:</h3>
Echo will repeat whatever the user said.</B>
<br>
<h3 class="tit"><b>{mock2}</b>:</h3>
Echo will repeat the user's text in reverse<br>
<br>
<h3 class="tit"><b>{greet}</b>:</h3>
Echo will replace this key with your current greet message.<br>
<br>
<h3 class="tit"><b>{user}</b>:</h3>
if you use {user} it will mention the user in the response. <i>example @Username</i><br>
<br>
<h3 class="tit"><b>{/user}</b>:</h3>
If you use {/user} it will just say their name. Without the mention. <i>example Username</i><br>
<br>
<h3 class="tit"><b>{kick}</b>:</h3>
If you add {kick} into the response it will kick anyone who said your trigger!<br>
this can be used for a word filter. we will explain more below.<br>
<br>
<h3 class="tit"><b>{pm}</b>:</h3>
<b>This key needs to be at the beginning of the response!</b> if you add {pm} at the beginning of the response. it will message the user the response instead of in server.<br>
<br>
<h3 class="tit"><b>{pm=USERID}</b>:</h3>
<b>This key needs to be at the beginning of the response!</b> You can have Echo private message up to three people to "Alert" you when someone says the trigger word.<br>
To do multiple ID's you use <b>{pm=USERID,USERID,USERID}</b><i>You can get someones id by typing</i> `--getid @user`<br>
<br>
<h3 class="tit"><b>{pref}</b>:</h3>
Echo will replace this key with his prefix in your server.<br>
<br>
<h3 class="tit"><b>{rand1}</b>:</h3>
Echo will say a random joke in your response.<Br>
<br>
<h3 class="tit"><b>{meme1}</b>:</h3>
Echo will post a random meme in your response.<br>
<br>
<h3 class="tit"><b>{date1}</b>:</h3>
Displays Echo's current date like `March 10, 2001, 5:16 pm`<br>
<br>
<h3 class="tit"><b>{date2}</b>:</h3>
Displays Echo's current date like `03.10.01`<br>
<br>
<h3 class="tit"><b>{time}</b>:</h3>
Displays Echo's current time like `5:16 pm`<Br>
<br>
<h3 class="tit"><b>{del}</b>:(<b>Currently Buggy</b>}</h3>
This will delete the user's message when it triggers your response. (<b>Perfect for Word Filter</b>)<br>
<i>The api is currently bugging this feature out. we're working hard to get it fixed.</i><br>
<br>
<h3 class="tit"><b>{cmd=}</b>:</h3>
This allows Echo to use his commands in a custom trigger.<br>
Examples: `{cmd=addmaster}`, `{cmd=help}`, `{cmd=setprefix}`<br>
This will let you rename Echo's command. more info <a href="https://github.com/proxikal/Echo/master/README.md#remake-echos-commands-with-custom-triggers">Here!</a>
<br><br>
<h3 class="tit"><b>Accepted Trigger Key's</b></h3>
<ul>
  <li>{find=word}</li>
  <li>Any text</li>
</ul>
<br>
<h3 class="tit"><b>{find=word}</b>:</h3>
This will search for the word in a user's text. and than show his response.<br>
<Br>
You can type anything for his trigger. Let's teach you how to make custom @Echo commands.<br>
<br>
For echo's trigger type: @Echo what's your prefix?<Br>
for his response type: My prefix in your server is `{pref}`<br>
You have just made a custom @Echo command! You can make custom commands with any prefix.<br>
<br>
<b><font color="red">WARNING: If you add a trigger that's the same as one of his commands. it could cause your server data to get erased. and he won't respond. IF echo doesn't respond in your server. Private message him `--help` and he will walk your through
fixing your server database.</font></b>
<Br>
<h3 class="tit">Word filtering and Link filtering.</h3>
You can use Echo's auto response system for word filtering or anti-link system.<br>
From there you can decide to add {kick} in the response or not. but here's some images to show you<br>
<br>
<img src="https://github.com/proxikal/Echo/master/word_filter_howto.PNG"><br>
<img src="https://github.com/proxikal/Echo/master/word_filter_kick_howto.PNG"><br>
<img src="https://github.com/proxikal/Echo/master/antilink_howto.PNG"><br>
<img src="https://github.com/proxikal/Echo/master/antilink2_howto.PNG"><br>
<h3 class="tit">Delete an Auto Response.</h3>
<img src="https://github.com/proxikal/Echo/master/delete_example.PNG">
<h3 class="tit">Word filter with Remove offending word.</h3>
<img src="https://github.com/proxikal/Echo/master/word_filter_del_howto.PNG"><br>
<h3 class="tit">Excluding people from a trigger.</h3>
<img src="https://github.com/proxikal/Echo/master/autoresp_exclude_howto.PNG"><br>
<h3 class="tit">Remake echo's commands with custom triggers!</h3>
<img src="https://github.com/proxikal/Echo/master/customcmds_addmaster_howto.PNG"><br>
<br>
<br>
<h3 class="tit">How to make a <b>--joke</b> or <b>--meme</b></h3>
First type `start::`<br>
Now for a trigger type <u>--joke</u> or <u>--meme</u> whatever you want to add.<br>
Now in his response type <u>{rand1}</u> for joke. and <u>{meme1}</u> for a random meme!<br>
<br>
Now you have made the joke or meme command. enjoy!
        </p>
      </div>
      <!-- /tab03 -->
      <!-- 2 columns -->
      <h3 class="tit">News</h3>
      <div class="col50">
        <p class="t-justify">Content coming soon!</p>
      </div>
      <!-- /col50 -->
      <div class="col50 f-right">
        <p class="t-justify">More content Here :D</p>
      </div>
      <!-- /col50 -->
      <div class="fix"></div>


      <div class="fix"></div>
      <!-- Text Alignment -->
     


    </div>
    <!-- /content -->
  </div>
  <!-- /cols -->
  <hr class="noscreen" />
  <!-- Footer -->
  <div id="footer" class="box">
    <p class="f-left">&copy; 2009 <a href="#">Echo's Revenge</a>, All Rights Reserved &reg;</p>
  </div>
  <!-- /footer -->
</div>
<!-- /main -->
</body>
</html>
