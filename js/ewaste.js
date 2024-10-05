function storerecyclingtype() {
    const selectedcategories = []; 

    
    if (document.getElementById("ICT equipment").checked) {
        selectedcategories.push(document.getElementById("ICT equipment").value);
    }
    if (document.getElementById("Batteries").checked) {
        selectedcategories.push(document.getElementById("Batteries").value);
    }
    if (document.getElementById("Lamps").checked) {
        selectedcategories.push(document.getElementById("Lamps").value);
    }
    if (document.getElementById("ConsumerProducts").checked) {
        selectedcategories.push(document.getElementById("ConsumerProducts").value);
    }
    if (document.getElementById("others").checked) {
        selectedcategories.push(document.getElementById("others").value);
    }

 
    sessionStorage.setItem("SelectedCategories", JSON.stringify(selectedcategories));
  
}

function nextpage() {
    storerecyclingtype();
    window.location.href = `${window.location.href.slice(0, location.href.lastIndexOf("/"))}/map`
    const selectedCategories = JSON.parse(sessionStorage.getItem("SelectedCategories"));
    console.log("Selected Categories:", selectedCategories); 
}
