<?php
if($_GET['do'] != "") {
	unlink("tasks/" . $_GET['do'] . ".json");
}