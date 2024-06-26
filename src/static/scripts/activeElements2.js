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

function showAddMemberBtn(){
    document.getElementById('add-member-btn').style.display = 'inline'
}

function showAddEnterpriseBtn(){
    document.getElementById('add-enterprise-btn').style.display = 'inline'
}

function hideAddMemberBtn(){
    document.getElementById('add-member-btn').style.display = 'none'
}

function hideAddEnterpriseBtn(){
    document.getElementById('add-enterprise-btn').style.display = 'none'
}