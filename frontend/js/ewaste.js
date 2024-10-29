function storecategories() {
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
    const checkboxes = document.querySelectorAll('.checkbox');
    if(!Array.from(checkboxes).some(checkbox => checkbox.checked))
    {
        alert("Select at least one category");
        return;
    }
    storecategories();
    const selectedCategories = JSON.parse(sessionStorage.getItem("SelectedCategories"));
    console.log("Selected Categories:", selectedCategories); 
}
