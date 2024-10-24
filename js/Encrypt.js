function str2ab(str) {
    const buf = new ArrayBuffer(str.length);
    const bufView = new Uint8Array(buf);
    for (let i = 0, strLen = str.length; i < strLen; i++) {
        bufView[i] = str.charCodeAt(i);
    }
    return buf;
}
function _arrayBufferToBase64( buffer ) {
    let binary = '';
    const bytes = new Uint8Array(buffer);
    const len = bytes.byteLength;
    for (let i = 0; i < len; i++) {
        binary += String.fromCharCode( bytes[ i ] );
    }
    return window.btoa( binary );
}

async function generateAESKey() {
    return crypto.subtle.generateKey(
        {
            name: "AES-GCM",
            length: 256
        },
        true, // extractable to allow encryption
        ["encrypt"] // restrict the usage of the key to encryption
    );
}

// Convert PEM public key string to a CryptoKey object
function importRSAPublicKey(pem) {
    // Remove the PEM header/footer and decode base64
    document.getElementById("debug").innerText = pem.toString();
    const pemHeader = "-----BEGIN PUBLIC KEY-----";
    const pemFooter = "-----END PUBLIC KEY-----";
    const pemContents = pem.substring(
        pemHeader.length,
        pem.length - pemFooter.length - 1,
    );
    document.getElementById("debug").innerText = pemContents;
    const binaryDerString = window.atob(pemContents);
    // convert from a binary string to an ArrayBuffer
    const binaryDer = str2ab(binaryDerString);
    return crypto.subtle.importKey(
        "spki",
        binaryDer,
        {
            name: "RSA-OAEP",
            hash: "SHA-256",
        },
        false, // non-extractable key
        ["encrypt"] // restrict usage to encryption
    );
}

async function startup(){
    let sessionID;
    while(!(sessionID = sessionStorage.getItem("sessionID"))){
        sessionStorage.setItem(
            "sessionID",
            btoa(String.fromCharCode(...crypto.getRandomValues(new Uint8Array(32)))).
            replace('+','')
        );
        console.log("Generated new SessionID")
    }
    console.log("Received SessionID");
    // Step 1: Generate AES Key
    const aesKey = await generateAESKey();
    console.log("AES key Generated");
    // Step 2: Fetch the RSA Public Key from the server
    const uri = `${window.location.href.slice(0, location.href.lastIndexOf("/"))}/getkey?sessionID=${sessionID}`;
    console.log(`Fetching RSA key from ${uri}`);
    let response = await fetch(uri);
    const pemKey = await response.text();
    console.log(`Successful Fetch\nKey: ${_arrayBufferToBase64(pemKey)}`);
    sessionStorage.setItem("pubkey",pemKey);
    const rsaPublicKey = await importRSAPublicKey(pemKey).
    catch(err=>console.log(err));
    if(!rsaPublicKey) return null;
    console.log("Successful Integration");
    // Step 3: Encrypt the AES Key with the RSA Public Key
    // Export the AES key to a raw format (ArrayBuffer)
    const exported = await window.crypto.subtle.exportKey("raw", aesKey);
    const exportedKeyBuffer = new Uint8Array(exported);
    console.log(`AES key extraction successful.\n[${exportedKeyBuffer}]`);
    // Encrypt the AES key with the RSA public key
    const aespromise = crypto.subtle.encrypt(
        {
            name: "RSA-OAEP"
        },
        rsaPublicKey,
        exported,
    );
    console.log("Encrypting AES key with RSA key");
    const encryptedAESKey = await aespromise.
    catch(err=>console.log(err));
    console.log(`AES key Encrypted: ${encryptedAESKey}`);
    // Step 4: Send the Encrypted AES Key to the server
    response = await fetch(`/sendkey?sessionID=${sessionID}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/octet-stream' },
        body: encryptedAESKey
    });
    console.log("AES key Sent");
    console.log(await response.text());
    console.log(`Encrypted AES Key ${_arrayBufferToBase64(encryptedAESKey)}`);
    // <--Add confirmation code here-->
    return aesKey;
}

async function encrypt(aesKey, plaintext){
    console.log("Password encryption attempted");
    const params = new URLSearchParams({
        key: sessionStorage.getItem("pubkey")
    }).toString();
    let response = await fetch(`/checkRSA?${params.toString()}`).catch(
        err=>document.getElementById("debug").innerText = err
    )
    const check = await response.text().catch(
        err=>document.getElementById("debug").innerText = err
    )
    if(check==="false") aesKey = await startup();

    let iv =  crypto.getRandomValues(new Uint8Array(12)); // Generate a unique IV
    const encoder = new TextEncoder();
    const encodedPlaintext = encoder.encode(plaintext); // Convert the plaintext to a Uint8Array
    console.log("Plaintext encoded successfully");
    // Encrypt the data
    let ciphertext = await crypto.subtle.encrypt(
        {
            name: "AES-GCM",
            iv: iv, // The generated IV
            // optional: additionalData: encoder.encode("metadata"), // Optional AAD
            // optional: tagLength: 128, // Optional, default is 128 bits
        },
        aesKey, // CryptoKey object for AES
        encodedPlaintext // The data to be encrypted (Uint8Array)
    ).catch(err=>console.log(err));
    if(ciphertext) console.log("Plaintext encrypted successfully");
    else console.log("Plaintext encryption failed.")
    ciphertext = _arrayBufferToBase64(ciphertext);
    iv = _arrayBufferToBase64(iv);
    return {ciphertext, iv};
}