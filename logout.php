<?php
	session_start();
	session_destroy();
	unset($_SESSION['username']);
	echo "<script type='text/JavaScript'>
		 alert('You have successfully logged out');
		 window.location.href = 'index.php';
		 </script>";
?>