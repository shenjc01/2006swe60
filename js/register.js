
async function onSubmit(event){
	event.preventDefault();
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
	const confirmpassword = document.getElementById("confirmpassword").value;
	const email = document.getElementById("email").value;
	if (password != confirmpassword) {
		alert("Passwords do not match!");
		return;
	}
	console.log("Starting handshake");
    const aesKey = await startup();
    console.log("Completed handshake");
    const sessionID = sessionStorage.getItem("sessionID");
    const {ciphertext, iv} = await encrypt(aesKey, password);
    const payload = {
        sessionid: sessionID,
        username: username, // Send the username as a string
        ciphertext: ciphertext, // Convert ciphertext to base64
        iv: iv, // Convert IV to base64
		email: email
    };
    let response = await fetch('/registerProcess', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' }, // JSON content type
        body: JSON.stringify(payload) // Convert payload to JSON
    });
    if (response.ok) {
        const result = await response.json();
        if (result.success) {
            alert(result.message);
            window.location.href = '/login';
        } else {
            alert(result.message);
        }
    } else {
        const result = await response.json();
		alert(result.message);
    }
}