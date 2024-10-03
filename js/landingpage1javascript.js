function storerecyclingtype() {
    
    const plasticBottlesSelected = document.getElementById("plasticBottles").checked;
    const cansSelected = document.getElementById("cans").checked;
    const anythElseSelected = document.getElementById("glass").checked || 
                          document.getElementById("paper").checked || 
                          document.getElementById("plasticBag").checked || 
                          document.getElementById("others").checked;

    if ((plasticBottlesSelected || cansSelected) && !anythElseSelected) {
        sessionStorage.setItem("category", "recyclensave");
    } else if (anythElseSelected) {
        sessionStorage.setItem("category", "cashfortrash");
    }
}


function nextpage()
{
    storerecyclingtype();
    
     const category = sessionStorage.getItem("category");
     console.log("Category:", category);

    
    // redirect to the map page 
}

