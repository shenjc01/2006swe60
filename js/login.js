
async function onSubmit(){
    const username = document.getElementById("usernameBox").value;
    const password = document.getElementById("passwordBox").value;
    console.log("Starting handshake");
    const aesKey = await startup();
    console.log("Completed handshake");
    const sessionID = sessionStorage.getItem("sessionID");
    const {ciphertext, iv} = await encrypt(aesKey, password)
    document.getElementById("debug").innerText = `Encrypted Password = ${ciphertext}`;
    // Convert ciphertext and iv to base64 strings for transport
    const payload = {
        sessionid: sessionID,
        username: username, // Send the username as a string
        ciphertext: ciphertext, // Convert ciphertext to base64
        iv: iv // Convert IV to base64
    };
    let response = await fetch('/loginattempt', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' }, // JSON content type
        body: JSON.stringify(payload) // Convert payload to JSON
    });
    if (response.ok) {
        document.getElementById("debug").innerText = "Data sent successfully";
    } else {
        document.getElementById("debug").innerText = await response.text();
    }
}