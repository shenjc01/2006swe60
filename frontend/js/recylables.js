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

/*document.addEventListener('DOMContentLoaded', () => {
    fetch('sc2006 software eng webdev/html/navbar.html')
        .then(res => res.text())
        .then(data => {
            document.getElementById('navbar').innerHTML = data;
        });
});*/

function nextpage()
{
    const checkboxes = document.querySelectorAll('.checkbox');
    if(!Array.from(checkboxes).some(checkbox => checkbox.checked))
    {
        alert("Select at least one category");
        return;
    }
    storerecyclingtype();
    
     const category = sessionStorage.getItem("category");
     console.log("Category:", category);

    
    // redirect to the map page 
}

