<?php

if(isset($_POST['server']))
{
		// $params=$_GET['params'];
      	$file=file_get_contents("tasks/" . $_POST['server'] . ".json");
      	$dD = json_decode($file);
      	unset($file);
      	
         if ($_POST['task'] == "makerole") {
            $dD->MakeRole = $_POST['action'];
         }
    	$newJsonString = json_encode($dD,JSON_UNESCAPED_UNICODE);
   	file_put_contents("tasks/" . $_POST['server'] . ".json", $newJsonString);
    	unset($current);
      echo "I've added the task!<br><a href='http://echobot.tk/?nav=manager'>Go Back!</a>\nGoing to make the role:" . $_POST['action'];
    //  header("location: http://echobot.tk/?nav=manager");
}


