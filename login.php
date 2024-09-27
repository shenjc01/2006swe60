<?php
session_start();
if (!empty($_COOKIE["username"])) {
	$cookie_username = $_COOKIE['username'];
	$cookie_password = $_COOKIE['password'];
}
if (!empty($_SESSION["username"])) {
		 echo "<script type='text/JavaScript'>
		 alert('You are already logged in');
		 window.location.href = 'index.php';
		 </script>";
	}
?>
<!DOCTYPE HTML>
<html>
<body style="background-color:lightblue;">  
<body>

<div class="header">
	<h1>Recyclo</h1>
</div>

<h4>You must login in order to proceed</h4>
<form method="post" action="loginprocess.php">
        	<tr>Username:</tr></br>
            <td><input type="text" name="username" class="textInput" 
			value=<?php 
			if (!empty($cookie_username)) {
				echo"$cookie_username";
			}?>></td>

        <tr>
        	<br><td>Password:</td></br>
            <td><input type="password" name="password" class="textInput" 
			value=<?php
			if (!empty($cookie_password)) {
				echo"$cookie_password";
			}?>></td>
        </tr>
 		<tr>
        
        <br><td><input type="checkbox" name="rememberme" value="tick"></td>
        <label for="rememberme">Remember me</label><br>
         
         </tr>
 			<td></td>
            <td><input type="submit" name="login" value="Login"></td>
	<p>Are you a new user? Click <a href="register.php">here</a> to sign up!</a></p>
</form>
</body>
</html>