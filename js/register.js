
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
    const payload = {
        username: username,
        password: password,
		email: email,
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
            document.getElementById("debug").innerText = result.message;
        }
    } else {
        document.getElementById("debug").innerText = await response.text();
    }
}