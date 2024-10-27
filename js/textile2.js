


function nextpage()
{
    const dropoffSelected  = document.getElementById("dropoff").checked;
    const pickupSelected = document.getElementById("pickup").checked;

    console.log("dropoffSelected:", dropoffSelected);
    console.log("pickupSelected:", pickupSelected);
    
    if (dropoffSelected)
    {
        sessionStorage.setItem("textile","refashdropoff");
    
        window.location.href="/map";
    }
    if (pickupSelected)
    {
        window.location.href = "https://refash.sg/"; 
    }
    
}

