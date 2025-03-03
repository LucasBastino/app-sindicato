function validateEnterpriseName(value){
    errorDiv = document.getElementById("name-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (value.length > 150){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "El nombre tiene mas de 150 caracteres."
        return false
    } else {
        errorDiv.style.display = 'none'
        return true
    }
}

function validateEnterpriseNumber(value){
    errorDiv = document.getElementById("enterprise-number-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    if (!isNumber(value)){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "El nombre posee algún caracter invalido."
        return false
    }
    if (value.length > 150){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "El numero tiene mas de 10 caracteres."
        return false
    } else {
        errorDiv.style.display = 'none'
        return true
    }

}

