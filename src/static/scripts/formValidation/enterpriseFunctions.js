function validateEnterpriseName(value){
    errorDiv = document.getElementById("name-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    return isNotLongerThan(errorDiv, 150, value)
}

function validateEnterpriseNumber(value){
    errorDiv = document.getElementById("enterprise-number-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    if (!isNumber(errorDiv, value)){
        return false
    }
    
    return isNotLongerThan(errorDiv, 10, value)

}

function validateEnterpriseAddress(value){
    errorDiv = document.getElementById("address-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isAlphanumeric(errorDiv, ",.", value)){
        return false
    }
    return isNotLongerThan(errorDiv, 100, value)
}