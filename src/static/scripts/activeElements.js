function enableInputs(){
    inputs = Array.from(document.getElementsByTagName("input"))
    inputs.forEach(input => {
        input.disabled = false
    }); 
}

function showConfirmButton(){
    button = document.querySelector(".btn-submit")
    button.style.display = 'inline'
}

function disableBrowserInput(){
    browserInput = document.getElementById('browser-input')
    browserInput.disabled = true
}

function enableBrowserInput(){
    browserInput = document.getElementById('browser-input')
    browserInput.disabled = false
}