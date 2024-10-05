function storerecyclingtype() {
    const plasticBottlesSelected = document.getElementById("plasticBottles").checked;
    const cansSelected = document.getElementById("cans").checked;
    const anythElseSelected = document.getElementById("glass").checked || 
                          document.getElementById("paper").checked || 
                          document.getElementById("plasticBag").checked || 
                          document.getElementById("others").checked;

    // console.log("Plastic Bottles Selected:", plasticBottlesSelected);
    // console.log("Cans Selected:", cansSelected);
    // console.log("Any Other Selected:", anythElseSelected);

    if ((plasticBottlesSelected || cansSelected) && !anythElseSelected) {
        sessionStorage.setItem("category", "recyclensave");
        // console.log("Category set to: recyclensave");
    } else if (anythElseSelected) {
        sessionStorage.setItem("category", "cashfortrash");
        // console.log("Category set to: cashfortrash");
    }
}


function nextpage()
{
    storerecyclingtype();
    
     const category = sessionStorage.getItem("category");
     console.log("Category:", category);

    
    // redirect to the map page 
}

