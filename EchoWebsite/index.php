<?php
session_start();
error_reporting(E_ALL);
$logout = false;
if($_GET['logout'] != "") {
unset($_SESSION['username']);
$logout=true;
}

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
?>
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
<script type="text/javascript" src="css-pop.js"></script>

<script>
function addText(event) {
    var targ = event.target || event.srcElement;
    document.getElementById("alltext").value += targ.textContent || targ.innerText;
}
</script>


<script type="text/javascript">
  $(document).ready(function(){
    $(".tabs > ul").tabs();
  });
  </script>

 <meta name="description" content="over 180 Emojis Auto Response system that has endless possibilities! Lots of SFW\NSFW Commands And Custom League of Legends Banners & Custom RIP Images!">
<meta name="author" content="Proxy">

<style>
#blanket {
background-color:#111;
opacity: 0.85;
*background:none;
position:absolute;
z-index: 9001;
top:0px;
left:0px;
width:100%;
}

#popUpDiv {
position:absolute;
width:400px;
height:400px;
z-index: 9002;
}
.form-container {
   border: 1px solid #030847;
   background: #006c7a;
   background: -webkit-gradient(linear, left top, left bottom, from(#208bbd), to(#006c7a));
   background: -webkit-linear-gradient(top, #208bbd, #006c7a);
   background: -moz-linear-gradient(top, #208bbd, #006c7a);
   background: -ms-linear-gradient(top, #208bbd, #006c7a);
   background: -o-linear-gradient(top, #208bbd, #006c7a);
   background-image: -ms-linear-gradient(top, #208bbd 0%, #006c7a 100%);
   -webkit-border-radius: 8px;
   -moz-border-radius: 8px;
   border-radius: 8px;
   -webkit-box-shadow: rgba(000,000,000,0.9) 0 1px 2px, inset rgba(255,255,255,0.4) 0 0px 0;
   -moz-box-shadow: rgba(000,000,000,0.9) 0 1px 2px, inset rgba(255,255,255,0.4) 0 0px 0;
   box-shadow: rgba(000,000,000,0.9) 0 1px 2px, inset rgba(255,255,255,0.4) 0 0px 0;
   font-family: 'Helvetica Neue',Helvetica,sans-serif;
   text-decoration: none;
   vertical-align: middle;
   min-width:500px;
   padding:20px;
   width:500px;
   }
.form-field {
   border: 1px solid #030645;
   background: #9db2d1;
   -webkit-border-radius: 4px;
   -moz-border-radius: 4px;
   border-radius: 4px;
   color: #000000;
   -webkit-box-shadow: rgba(255,255,255,0.4) 0 0px 0, inset rgba(000,000,000,0.7) 0 0px 0px;
   -moz-box-shadow: rgba(255,255,255,0.4) 0 0px 0, inset rgba(000,000,000,0.7) 0 0px 0px;
   box-shadow: rgba(255,255,255,0.4) 0 0px 0, inset rgba(000,000,000,0.7) 0 0px 0px;
   padding:8px;
   margin-bottom:20px;
   width:480px;
   }

::-webkit-input-placeholder { /* Chrome/Opera/Safari */
  color: black;
}
::-moz-placeholder { /* Firefox 19+ */
  color: black;
}
:-ms-input-placeholder { /* IE 10+ */
  color: black;
}
:-moz-placeholder { /* Firefox 18- */
  color: black;
}

.small-field {
   border: 1px solid #030645;
   background: #9db2d1;
   -webkit-border-radius: 4px;
   -moz-border-radius: 4px;
   border-radius: 4px;
   color: #000000;
   -webkit-box-shadow: rgba(255,255,255,0.4) 0 0px 0, inset rgba(000,000,000,0.7) 0 0px 0px;
   -moz-box-shadow: rgba(255,255,255,0.4) 0 0px 0, inset rgba(000,000,000,0.7) 0 0px 0px;
   box-shadow: rgba(255,255,255,0.4) 0 0px 0, inset rgba(000,000,000,0.7) 0 0px 0px;
   padding:8px;
   margin-bottom:20px;
   width:200px;
   }
.form-field:focus {
   background: #0c7ba3;
   color: #ffffff;
   }
.form-container h2 {
   text-shadow: #000000 0 1px 0;
   font-size:18px;
   margin: 0 0 10px 0;
   font-weight:bold;
   text-align:center;
    }
.form-title {
   margin-bottom:10px;
   color: #ffffff;
   text-shadow: #000000 0 1px 0;
   }
.submit-container {
   margin:8px 0;
   text-align:right;
   }
.submit-button {
   border: 1px solid #000000;
   background: #9db2d1;
   background: -webkit-gradient(linear, left top, left bottom, from(#9db2d1), to(#9db2d1));
   background: -webkit-linear-gradient(top, #9db2d1, #9db2d1);
   background: -moz-linear-gradient(top, #9db2d1, #9db2d1);
   background: -ms-linear-gradient(top, #9db2d1, #9db2d1);
   background: -o-linear-gradient(top, #9db2d1, #9db2d1);
   background-image: -ms-linear-gradient(top, #9db2d1 0%, #9db2d1 100%);
   -webkit-border-radius: 4px;
   -moz-border-radius: 4px;
   border-radius: 4px;
   -webkit-box-shadow: rgba(255,255,255,0.4) 0 1px 0, inset rgba(255,255,255,0.4) 0 1px 0;
   -moz-box-shadow: rgba(255,255,255,0.4) 0 1px 0, inset rgba(255,255,255,0.4) 0 1px 0;
   box-shadow: rgba(255,255,255,0.4) 0 1px 0, inset rgba(255,255,255,0.4) 0 1px 0;
   text-shadow: #32697a 0 1px 0;
   color: #000000;
   font-family: helvetica, serif;
   padding: 8.5px 18px;
   font-size: 14px;
   text-decoration: none;
   vertical-align: middle;
   }
.submit-button:hover {
   border: 1px solid #000000;
   text-shadow: #167080 0 1px 0;
   background: #0c7ba3;
   background: -webkit-gradient(linear, left top, left bottom, from(#0c7ba3), to(#0c7ba3));
   background: -webkit-linear-gradient(top, #0c7ba3, #0c7ba3);
   background: -moz-linear-gradient(top, #0c7ba3, #0c7ba3);
   background: -ms-linear-gradient(top, #0c7ba3, #0c7ba3);
   background: -o-linear-gradient(top, #0c7ba3, #0c7ba3);
   background-image: -ms-linear-gradient(top, #0c7ba3 0%, #0c7ba3 100%);
   color: #fff;
   }
.submit-button:active {
   text-shadow: #000000 0 1px 0;
   border: 1px solid #ff0000;
   background: #8a2c2c;
   background: -webkit-gradient(linear, left top, left bottom, from(#c41c1c), to(#0c7ba3));
   background: -webkit-linear-gradient(top, #c41c1c, #8a2c2c);
   background: -moz-linear-gradient(top, #c41c1c, #8a2c2c);
   background: -ms-linear-gradient(top, #c41c1c, #8a2c2c);
   background: -o-linear-gradient(top, #c41c1c, #8a2c2c);
   background-image: -ms-linear-gradient(top, #c41c1c 0%, #8a2c2c 100%);
   color: #fff;
   }
   </style>

<script>
function addText(elId,text) {
    document.getElementById(elId).value += text;

    var d = document.getElementById('ars').value;
    if (d == "{user}") {
    	document.getElementById("info").innerHTML = "adding the <b>{user}</b> key will cause echo to mention the user.";
    }
    if (d == "{/user}") {
    	document.getElementById("info").innerHTML = "adding the <b>{/user}</b> key will cause echo to say their name.";
    }
    if (d == "{kick}") {
    	document.getElementById("info").innerHTML = "adding the <b>{kick}</b> key will cause echo to kick the user.";
    }
    if (d == "{ban}") {
    	document.getElementById("info").innerHTML = "adding the <b>{ban}</b> key will cause echo to ban the user.";
    }
    if (d == "{role:Role Name Here}") {
    	document.getElementById("info").innerHTML = "adding the <b>{role:Role Name}</b> key will cause echo to give the role to the user.";
    }
}

function addText2(elId,text) {
    document.getElementById(elId).value = text;
}

function makears() {
	var trigger = document.getElementById('trigger').value;
	var prefix = document.getElementById('prefix').value;
	var response = document.getElementById('resp').value

	document.getElementById('results').value = prefix + "auto " + trigger + "=" + response;
}




function auto_load(){
        $.ajax({
          url: "data.php",
          cache: false,
          success: function(data){
             $("#auto_load_div").html(data);
          } 
        });
}
 
$(document).ready(function(){
 
auto_load(); //Call auto_load() function when DOM is Ready
 
});
setInterval(auto_load,30000);
</script>

</head>
<body>
<div id="main">
  <!-- Tray -->
  <div id="tray" class="box">
    <p class="f-left box">
      <!-- Switcher -->
      <span class="f-left" id="switcher"> <a href="javascript:void(0);" rel="1col" class="styleswitch ico-col1" title="Display one column"><img src="design/switcher-1col.gif" alt="1 Column" /></a> <a href="javascript:void(0)" rel="2col" class="styleswitch ico-col2" title="Display two columns"><img src="design/switcher-2col.gif" alt="" /></a> </span> Hosted by: <strong><a href="http://northstarofficial.com/" target="_new">NorthstarOfficial</a></strong> </p>
    <?php if($_SESSION['username'] != "") {?>
      <p class="f-right">User: <strong><a href="#"><?php echo $_SESSION['username']; ?></a></strong> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; <strong><a href="?logout=true" id="logout">Log out</a></strong></p>
    <?php } else { ?>
      <p class="f-right">User: <strong><a href="#">Guest</a></strong> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; <strong><a href="?nav=login" id="logout">Login</a></strong></p>
    <?php } ?>

  </div>
  <!--  /tray -->
  <hr class="noscreen" />
  <!-- Menu -->
  <div id="menu" class="box">
    <ul class="box f-right">
      <li><a href="https://www.carbonitex.net/Discord/bots" target="_new"><span><strong>More Bots &raquo;</strong></span></a></li>
    </ul>
    <ul class="box">
      <li><a href="http://echobot.tk"><span>Home</span></a></li>
      <!-- Active -->
      <li><a href="?nav=commands"><span>Commands</span></a></li>
      <li><a href="?nav=arskeys"><span>A.R.S Keys</span></a></li>
      <li><a href="?nav=arsexamples"><span>A.R.S Examples</span></a></li>
      <li><a href="?nav=home#builder"><span>A.R.S Builder</span></a></li>
      <li><a href="?nav=home#emojis"><span>Emojis List</span></a></li>
      <li><a href="?nav=nsfw"><span>NSFW Bot</span></a></li>
      <li><a href="?nav=music"><span>Music!</span></a></li>
      <?php
      if($_SESSION['username'] != "") {?>
        <li><a href="?nav=manager"><span>Server Manager</span></a></li>
     <?php } ?>
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
        <!-- Create a new project -->
        <p id="btn-create" class="box"><a href="https://discordapp.com/oauth2/authorize?&client_id=161096989986127872&scope=bot&permissions=271842359" target="_new"><span>Invite Echo Today!</span></a></p>
      </div>
      <center>
      <font face="verdana" size="2">
      <u><?php echo $ARS; ?></u> A.R.S Requests.<br>
      <u><?php echo $Emoji; ?></u> Emoji's used.<br>
      <u><?php echo $memb; ?></u> Members.<br>
      <u><?php echo $chans; ?></u> Channels.<br>
      <u><?php echo $role; ?></u> Roles.
      </font>
      </center>
      <!-- /padding -->
      <ul class="box">
      <li><hr></li>
        <li><a href="http://echobot.tk/?nav=home#builder">A.R.S Builder</a></li>
        <li><a href="#">Documentation</a>
          <!-- Active -->
          <ul>
            <li><a href="?nav=home#setup">Setting up Echo</a></li>
            <li><a href="?nav=home#issues">Permissions & Issues</a></li>
            <li><a href="?nav=commands">Basic Commands</a></li>
            <li><a href="?nav=arsexamples">Some A.R.S Examples</a></li>
          </ul>
        </li>
        <?php if($_SESSION['username'] != "") {?>
        <li><a href="#">Your Server</a>
          <ul>
            <li><a href="?nav=manager">Server Manager</a></li>
            <li><a href="#">Auto Response Database [SOON]</a></li>
            <li><a href="#">Backups [SOON]</a></li>
          </ul>
        </li>
        <?php } ?>
        <li><a href="#">24/7 Support</a></li>
      </ul>
                <iframe src="https://discordapp.com/widget?id=148629493676769280&theme=dark" width="225" height="500" allowtransparency="true" frameborder="0"></iframe>
    </div>
    <!-- /aside -->
    <hr class="noscreen" />
    <!-- Content (Right Column) -->
    <div id="content" class="box">
     <?php if($logout == true) { ?>
     <p class="msg warning">You have logged out.</p>
     <?php } ?>
     <!--
     <p class="msg info">Looking for a Graphics designer for the Custom League of Legends banners. Your work would be displayed all over! private message me in discord!</p>
    -->
  <!--    <p class="msg info"><b>[NEW]</b> - Edit Server Database from the website! AND customize echo's response to all of his commands!<br>First time: type <b>--webmaster</b> in your server for the username and password!</p>-->
  <p>Check out EdibleDerpy's <a href="https://discordapp.com/oauth2/authorize?&response_type=code&client_id=173283811243589634&scope=bot&permissions=506588223&redirect_uri=http://Hexacircle.ml/bot/success" target="_new3">Hexacircle Bot</a> Today!</p>

      <p class="msg done" id="auto_load_div"></p>
     <!-- <p class="msg error">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p> -->
      <!-- Tabs -->
     <!-- <p class="msg error"></b></p> -->
     <p class="msg done"><b>NEW:</b> Music bot soon! <a href="http://echobot.tk/?nav=music">Read More Here!</a></p>
      <?php
      if($_GET['nav'] != "") {
      	include $_GET['nav'] . ".php";
      }
      else
      {
      	include "home.php";
      }
      ?>
      <div style="position:relative; left: 15px;">
<!--      <script type="text/javascript" src="//www5.yourshoutbox.com/shoutbox/start.php?key=962460785"></script> -->
      </div>
      </div>
      <!-- /tab03 -->
      <!-- 2 columns -->
      <!-- /col50 -->
<!--
      <div class="col50 f-right">
        <p class="t-justify">More content Here :D</p>
      </div> -->
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
    <p class="f-left">&copy; 2016 <a href="http://echobot.tk">Echo's Revenge</a>, All Rights Reserved &reg; Thanks to <a href="http://northstarofficial.com" target="_new">NorthstarOfficial.com</a> for hosting us!</p>
  </div>
  <!-- /footer -->
</div>

<!-- <a href="#" onclick="popup('popUpDiv')">Click to Open CSS Pop Up</a> -->

<!-- Start of StatCounter Code for Default Guide -->
<script type="text/javascript">
var sc_project=10909223;
var sc_invisible=0;
var sc_security="7b85d947";
var scJsHost = (("https:" == document.location.protocol) ?
"https://secure." : "http://www.");
document.write("<sc"+"ript type='text/javascript' src='" + scJsHost+
"statcounter.com/counter/counter.js'></"+"script>");
</script>
<noscript><div class="statcounter"><a title="shopify analytics ecommerce
tracking" href="http://statcounter.com/shopify/" target="_blank"><img
class="statcounter" src="http://c.statcounter.com/10909223/0/7b85d947/0/"
alt="shopify analytics ecommerce tracking"></a></div></noscript>
<!-- End of StatCounter Code for Default Guide -->
<!-- /main -->
</body>
</html>
