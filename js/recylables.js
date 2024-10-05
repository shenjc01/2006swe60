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
        sessionStorage.setItem("category", "RecycleNSave");
        // console.log("Category set to: RecycleNSave");
    } else if (anythElseSelected) {
        sessionStorage.setItem("category", "CashForTrash");
        // console.log("Category set to: CashForTrash");
    }
}


function nextpage()
{
    storerecyclingtype();
    window.location.href = `${window.location.href.slice(0, location.href.lastIndexOf("/"))}/map`
     const category = sessionStorage.getItem("category");
     console.log("Category:", category);

    
    // redirect to the map page 
}

