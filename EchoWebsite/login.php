<?php
error_reporting(E_ALL);
$data = "";
if(isset($_POST['username'])) {
	if(file_exists("servers/" . $_POST['username'] . ".json")) {
		$backup = file_get_contents("servers/" . $_POST['username'] . ".json");
		$data = json_decode($backup);
		unset($backup);

		if($_POST['password'] == $data->Password) {
			$_SESSION['username'] = $_POST['username'];
			setcookie("username", $_POST['username'], time()+60*60*24*365, "/");
            setcookie("password", $_POST['password'], time()+60*60*24*365, "/");
			$dat = "<font color='green'>You've logged in to server: " . $_POST['username'] . " You're being redirected to the <u>Server Manager</u> Page!</font>";
			echo '<META http-equiv="refresh" content="1;URL=http://echobot.tk/?nav=manager">';
		} else {
			$dat = "<font color='red'>Password for server " . $_POST['username'] . " Is incorrect.</font>";
		}
	} else {
		$dat = "<font color='red'>You need to type <b>--webmaster</b> in your server. Echo will pm your info!</font>";
	}
}

$username=$_COOKIE['username'];
$password=$_COOKIE['password'];
?>
<p class="msg info">Forget your password or username? type <b>--webinfo</b> in your server.</p>
<center>
<?php echo $dat; ?>
</center>
<form action="" method="POST">
<table>
<tr>
  <th>Login to Server.</th>
  <th>&nbsp;</th>
</tr>
<tr>
<td><b>Username:</b></td><td><input type="text" name="username" value="<?php echo $_COOKIE['username']; ?>"></td></tr>
<tr class="bg">
<td><b>Password:</b></td><td><input type="password" name="password" value="<?php echo $_COOKIE['password']; ?>"></td></tr>
<!--
<tr>
<td>Remember Me:</td><td><input type="checkbox" value="1" name="remember"></td></tr>
<tr class="bg">
-->
<tr>
<td>&nbsp;</td><td><input type="submit" value="Login"></td></tr>
</table>
</form>

      <div class="fix"></div>
      <div class="fix"></div>
      <!-- Text Alignment -->