document.onkeydown = function(e){
    if (e.key === 'Enter' ){
        console.log("se apreto enter")
        e.preventDefault()
    }
}

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

function blankSearchInput(){
    console.log("blankl ")
    document.getElementById("search-input").value = ""
}

function showConfirmButton(model){
    document.getElementById(`btn-${model}-confirm`).style.display = 'inline'
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


function showEnterpriseButton(){
    document.getElementById('enterprise-button').style.display = 'inline'
}

function hideEnterpriseButton(){
    document.getElementById('enterprise-button').style.display = 'none'
}

function hideAddPaymentButton(){
    document.getElementById('add-payment-btn').style.display = 'none'
}

function showAddPaymentBtn(){
    document.getElementById('add-payment-btn').style.display = 'inline'
}

function showXEnterprise(){
    document.getElementById('x-enterprise').style.display = 'inline'
}


function searchEnterpriseAgain(){
    document.getElementById('x-enterprise').style.display = 'none'
    document.getElementById('name-enterprise-input').style.display = 'none'
    document.getElementById('enterprise-search-box').style.display = 'inline'
}

function selectEnterprise(IdEnterprise, EnterpriseName){
    document.getElementById("id-enterprise-input").value = IdEnterprise
    document.getElementById("name-enterprise-input").value = EnterpriseName
    document.getElementById('enterprise-search-box').style.display = 'none'
    document.getElementById('optionsTable').innerHTML = ''
    document.getElementById('name-enterprise-input').style.display = 'inline'
    document.getElementById('x-enterprise').style.display = 'inline'
}


// function makeBackup(){
//     console.log("asdadss")
// 	spawn('cmd.exe', ['/c',`mysqldump -h ndk.h.filess.io -P 3307 --no-tablespaces -u sindicatoDB_settingcry -pdatabasefilessio sindicatoDB_settingcry > C:\Users\cuent\desktop\mydb.sql`]);
//   }


// import { spawn } from 'child_process';
// function makeBackup(){
//     execSync = require('child_process').execSync
//     // console.log("asddas")
//     // spawn('cmd.exe', ['/c',`mysqldump -h ndk.h.filess.io -P 3307 --no-tablespaces -u sindicatoDB_settingcry -pdatabasefilessio sindicatoDB_settingcry > C:\Users\cuent\desktop\mydb.sql`]);
//     const output = execSync('ls', { encoding: 'utf-8' });  // the default is 'buffer'
// console.log('Output was:\n', output);
// }