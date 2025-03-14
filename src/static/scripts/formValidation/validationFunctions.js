// PERSONAL INFO
function validateName(value){
    errorDiv = document.getElementById("name-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isAlphanumeric(errorDiv, "", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 50, value)
}

function validateLastName(value){
    errorDiv = document.getElementById("last-name-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isAlphanumeric(errorDiv, "", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 50, value)
}

function validateDNI(value){
    errorDiv = document.getElementById("dni-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (!isNumeric(errorDiv, ".", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 8, value)
}

function validateBirthday(value){
    errorDiv = document.getElementById("birthday-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    return isValidDate(errorDiv, value, 1940)
}

function validateGender(value){
    errorDiv = document.getElementById("gender-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    genders = ["Masculino", "Femenino", "Otro"]
    return isAValidOption(errorDiv, genders, value)
}

function validateMaritalStatus(value){
    errorDiv = document.getElementById("marital-status-error")
    maritalStatus = ["Soltero", "Casado", "Separado", "Viudo", "Separado", "Divorciado"]
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    return isAValidOption(errorDiv, maritalStatus, value)
}


// ADDRESS INFO
function validateAddress(value){
    errorDiv = document.getElementById("address-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isAlphanumeric(errorDiv, ",.º", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 100, value)
}

function validateDistrict(value){
    errorDiv = document.getElementById("district-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (!isAlphanumeric(errorDiv, ",.º", value)){
        return false
    }
    return isNotLongerThan(errorDiv, 50, value)
}


function validatePostalCode(value){
    errorDiv = document.getElementById("postal-code-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    if (!isNumeric(errorDiv, "", value)){
        return false
    }
    
    return isNotLongerThan(errorDiv, 4, value)
}



// CONTACT INFO
function validatePhone(value){
    errorDiv = document.getElementById("phone-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isNumeric(errorDiv, "+-", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 30, value)
}

function validateEmail(value){
    errorDiv = document.getElementById("email-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isAlphanumeric(errorDiv, "-@._", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 50, value)
}

// SOCIAL SECURITY INFO
function validateMemberNumber(value){
    errorDiv = document.getElementById("member-number-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isNumeric(errorDiv, "", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 20, value)
}

function validateCUIL(value){
    errorDiv = document.getElementById("cuil-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (!isNumeric(errorDiv, "-", value)){
        return false
    }
    return isNotLongerThan(errorDiv, 20, value)
}

function validateIdEnterprise(value){
    fetch('http://localhost:8085/enterprise/getAllEnterprisesId', {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json'
        }})
      .then(res => res.json())
      .then(res => checkIdEnterprise(res, value))
}

function checkIdEnterprise(res, value){
    errorDiv = document.getElementById('enterprise-error')
    if (res.includes(value)){
        errorDiv.style.display = 'none'
        return true
  } else{
        errorDiv.style.display = 'inline'
        return false
  }
}

function validateCategory(value){
    errorDiv = document.getElementById('category-error')
    categories = ["Nivel 1: Oficial Múltiple", "Nivel 2: Oficial Especializado", "Nivel 3: Oficial General", "Nivel 4: Medio Oficial", "Nivel 5: Ayudante", "Nivel 6: Operario Act. Industrial"]
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    return isAValidOption(errorDiv, categories, value)
}

function validateEntryDate(value){
    errorDiv = document.getElementById("entry-date-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    return isValidDate(errorDiv, value, 1960)
}

// ENTERPRISE INFO
function validateEnterpriseName(value){
    errorDiv = document.getElementById("name-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (!isAlphanumeric(errorDiv, ",.º", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 150, value)
}

function validateEnterpriseNumber(value){
    errorDiv = document.getElementById("enterprise-number-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    if (!isNumeric(errorDiv, "", value)){
        return false
    }
    
    return isNotLongerThan(errorDiv, 10, value)

}


function validateCUIT(value){
    errorDiv = document.getElementById("cuit-error")
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (!isNumeric(errorDiv, "-", value)){
        return false
    }
    return isNotLongerThan(errorDiv, 20, value)
}

// PAYMENT INFO
function validateMonth(value){
    var errorDiv = document.getElementById('month-error')
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    var months = ['01', '02', '03', '04', '05', '06', '07', '08', '09', '10', '11', '12'] 
    if (!months.includes(value)){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "Mes erroneo."
        return false
    } else{
        errorDiv.style.display = 'none'
        return true
    }
}



function validateYear(value){
    errorDiv = document.getElementById('year-error')
    // verifico que no este vacio
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    // verifico que sea un entero
    value = parseInt(value)
    if (!value) {
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "Año erroneo."
        return false
    }
    // si es mayor al año actual o menor a 2000 es un año invalido
    if ((value > new Date().getFullYear()) || (value < 2000)) {
        // si es invalido muestro el error
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "Año erroneo."
        return false
    } else{
        // si es valido no muestro el error
        errorDiv.style.display = 'none'
        return true
    }

}


function validateStatus(value){
    var errorDiv = document.getElementById('status-error')
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (!(value == "PAGO" || value == "IMPAGO")){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "Estado de pago erróneo."
        return false
    } else{
        errorDiv.style.display = 'none'
        return true
    }

}

function validateAmount(value){
    var errorDiv = document.getElementById('amount-error')
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isNumeric(errorDiv, "", value)){
        return false
    }
    return isNotLongerThan(errorDiv, 20, value)
}



function validatePaymentDate(value){
    var errorDiv = document.getElementById('payment-date-error')
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    return isValidDate(errorDiv, value, 2000)
    
    // devuelvo true o false dependiendo de la validacion
}

function validateCommentary(value){
    var errorDiv = document.getElementById('commentary-error')

    return isNotLongerThan(errorDiv, value, 400)
}
