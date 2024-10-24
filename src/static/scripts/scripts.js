function enableInputs(){
    inputs = Array.from(document.getElementsByTagName("input"))
    inputs.forEach(input => {
        input.disabled = false
    }); 
}

function enableSelects(){
    selects = Array.from(document.getElementsByTagName("select"))
    selects.forEach(select => {
        select.disabled = false
    }); 
}


function showConfirmButton(){
    button = document.querySelector(".btn-submit")
    button.style.display = 'inline'
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

function showMemberSearchInput(){
    document.getElementById('member-search-input').style.display = 'inline'
}

function hideMemberSearchInput(){
    input = document.getElementById('member-search-input').style.display = 'none'
}

function showEnterpriseSearchInput(){
    document.getElementById('enterprise-search-input').style.display = 'inline'
}

function hideEnterpriseSearchInput(){
    document.getElementById('enterprise-search-input').style.display = 'none'
}

function showEnterpriseMemberSearchNav(){
    document.getElementById('nav-enterprise-member-search').style.display = 'inline'
}


function showParentSearchInput(){
    document.getElementById('parent-search-input').style.display = 'inline'
}

function hideParentSearchInput(){
    document.getElementById('parent-search-input').style.display = 'none'
}

function showAddParentButton(){
    document.getElementById('add-parent-button').style.display = 'inline'
}

function hideAddParentButton(){
    document.getElementById('add-parent-button').style.display = 'none'
}

function showEnterpriseSelect(){
    document.getElementById('enterprise-select').style.display = 'inline'
}

function hideEnterpriseSelect(){
    document.getElementById('enterprise-select').style.display = 'none'
}

function showEnterpriseButton(){
    document.getElementById('enterprise-button').style.display = 'inline'
}

function hideEnterpriseButton(){
    document.getElementById('enterprise-button').style.display = 'none'
}