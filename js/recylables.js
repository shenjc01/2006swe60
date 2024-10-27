function storerecyclingtype() {
    const plasticBottlesSelected = document.getElementById("plasticBottles").checked;
    const cansSelected = document.getElementById("cans").checked;
    const others = document.getElementById("others").checked;
    const anythElseSelected = document.getElementById("glass").checked || 
                          document.getElementById("paper").checked || 
                          document.getElementById("plasticBag").checked ;

    // console.log("Plastic Bottles Selected:", plasticBottlesSelected);
    // console.log("Cans Selected:", cansSelected);
    // console.log("Any Other Selected:", anythElseSelected);

    if ((plasticBottlesSelected || cansSelected) && !anythElseSelected && !others) {
        sessionStorage.setItem("category", "recyclensave");
        // console.log("Category set to: recyclensave");
    } else if (anythElseSelected && !others) {
        sessionStorage.setItem("category", "cashfortrash");
        // console.log("Category set to: cashfortrash");
    }
    else if (others)
    {
        sessionStorage.setItem("category", "recyclingbins");
    }

}


function nextpage()
{
    storerecyclingtype();
    
     const category = sessionStorage.getItem("category");
     console.log("Category:", category);

    window.location.href = '/map';
    // redirect to the map page 
}

