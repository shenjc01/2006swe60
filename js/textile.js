function storeTextileQuantity() {
    const quantity = document.getElementById('textileQuantity').value;


    if (quantity >= 10) {
     
        sessionStorage.setItem("textileQuantity", quantity);
        return true; 
    } 
    else
    {
        return false;
    }
}

function redirect() {
   
    const isValidQuantity = storeTextileQuantity();

    if (isValidQuantity) {
        window.location.href="/textiles2";
    }
    else
    {
        sessionStorage.setItem("textile","clothesbin");
        window.location.href="/map";
    }
}


