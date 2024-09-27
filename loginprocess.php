<?php
$dbhost="localhost";
$dbuser="root";
$dbpass="";
$dbname="recyclo";

session_start();

//create new connection
$connectToDatabase = mysqli_connect($dbhost, $dbuser, $dbpass , $dbname);

if($_SERVER["REQUEST_METHOD"] == "POST") {
      // username and password sent from form 
      
      $username = mysqli_real_escape_string($connectToDatabase,$_POST['username']);
      $password = mysqli_real_escape_string($connectToDatabase,$_POST['password']); 
      $sql = "SELECT * FROM account WHERE username = '$username' and password = '$password'";
      $result = mysqli_query($connectToDatabase,$sql);
      
      $count = mysqli_num_rows($result);
      
      // If result matched $myusername and $mypassword, table row must be 1 row
		
      if($count == 1) {
		 if (isset($_POST['rememberme'])) {
		 setcookie("username", $username, time() + 3600, '/');
		 setcookie("password", $password, time() + 3600, '/');
		 $_SESSION['username'] = $username ;
		 echo "<script type='text/JavaScript'>
		 alert('Login Success');
		 window.location.href = 'index.php';
		 </script>";
		 } else { 
			 setcookie("username",$username, time() - 3600, '/');
			 setcookie("password",$password, time() - 3600, '/');
			 $_SESSION['username'] = $username ;
		     echo "<script type='text/JavaScript'>
		     alert('Login Success');
		     window.location.href = 'index.php';
		     </script>";
		 }
      } else {
         echo "<script type='text/JavaScript'>
		 alert('Wrong Username or Password');
		 window.location.href = 'login.php';
		 </script>";
      }
}

?>
