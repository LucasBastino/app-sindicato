// PERSONAL INFO
function validateName(value){
    errorDiv = Array.from(document.getElementsByClassName("name-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isAlphabetic(errorDiv, "", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 50, value)
}

function validateLastName(value){
    errorDiv = Array.from(document.getElementsByClassName("last-name-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isAlphabetic(errorDiv, "", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 50, value)
}

function validateDNI(value){
    errorDiv = Array.from(document.getElementsByClassName("dni-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (!isNumeric(errorDiv, ".", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 8, value)
}

function validateBirthday(value){
    errorDiv = Array.from(document.getElementsByClassName("birthday-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    return isValidDate(errorDiv, value, 1940)
}

function validateGender(value){
    errorDiv = Array.from(document.getElementsByClassName("gender-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    genders = ["Masculino", "Femenino", "Otro"]
    return isAValidOption(errorDiv, genders, value)
}

function validateMaritalStatus(value){
    errorDiv = Array.from(document.getElementsByClassName("marital-status-error")).pop()
    maritalStatus = ["Soltero", "Casado", "Separado", "Viudo", "Separado", "Divorciado"]
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    return isAValidOption(errorDiv, maritalStatus, value)
}


// ADDRESS INFO
function validateAddress(value){
    errorDiv = Array.from(document.getElementsByClassName("address-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isAlphanumeric(errorDiv, ",.º", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 100, value)
}

function validateDistrict(value){
    errorDiv = Array.from(document.getElementsByClassName("district-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (!isAlphanumeric(errorDiv, ",.º", value)){
        return false
    }
    return isNotLongerThan(errorDiv, 50, value)
}


function validatePostalCode(value){
    errorDiv = Array.from(document.getElementsByClassName("postal-code-error")).pop()
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
    errorDiv = Array.from(document.getElementsByClassName("phone-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isNumeric(errorDiv, "+-", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 20, value)
}

function validateEmail(value){
    errorDiv = Array.from(document.getElementsByClassName("email-error")).pop()
    if (isNotEmpty(errorDiv, value)){
        if (!isAlphanumeric(errorDiv, "-@._", value)){
          
            return false
        } else if(!value.includes("@")){
            errorDiv.style.display = 'inline'
            errorDiv.innerHTML = "No contiene '@'"
            return false
        }
    }
    return isNotLongerThan(errorDiv, 50, value)
}

// SOCIAL SECURITY INFO
function validateMemberNumber(value){
    errorDiv = Array.from(document.getElementsByClassName("member-number-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isNumeric(errorDiv, "", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 20, value)
}

function validateCUIL(value){
    errorDiv = Array.from(document.getElementsByClassName("cuil-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (!isNumeric(errorDiv, "-", value)){
        return false
    }
    return isNotLongerThan(errorDiv, 20, value)
}

function validateIdEnterprise(value){
    // esto no devuelve true o false, devuelve una promesa que devuelve true o false
        return fetch('http://localhost:8085/enterprise/getAllEnterprisesId', {
            method: 'GET',
            headers: {
              Accept: 'application/json',
              'Content-Type': 'application/json'
            }})
          .then(res => res.json())
          .then(res => {return checkIdEnterprise(res, value)})
      
}

function checkIdEnterprise(res, value){
    value = parseInt(value)
    errorDiv = Array.from(document.getElementsByClassName("enterprise-error")).pop()
    if (res.includes(value)){
        errorDiv.style.display = 'none'
        return true
  } else{
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = 'Empresa no válida'
        return false
  }
}

function validateCategory(value){
    errorDiv = Array.from(document.getElementsByClassName("category-error")).pop()
    categories = ["Nivel 1: Oficial Múltiple", "Nivel 2: Oficial Especializado", "Nivel 3: Oficial General", "Nivel 4: Medio Oficial", "Nivel 5: Ayudante", "Nivel 6: Operario Act. Industrial"]
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    return isAValidOption(errorDiv, categories, value)
}

function validateEntryDate(value){
    errorDiv = Array.from(document.getElementsByClassName("entry-date-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    return isValidDate(errorDiv, value, 1960)
}

// ENTERPRISE INFO
function validateEnterpriseName(value){
    errorDiv = Array.from(document.getElementsByClassName("name-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    if (!isAlphanumeric(errorDiv, ",.º", value)){
        return false
    }

    return isNotLongerThan(errorDiv, 150, value)
}

function validateEnterpriseNumber(value){
    errorDiv = Array.from(document.getElementsByClassName("enterprise-number-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    if (!isNumeric(errorDiv, "", value)){
        return false
    }
    
    return isNotLongerThan(errorDiv, 10, value)

}


function validateCUIT(value){
    errorDiv = Array.from(document.getElementsByClassName("cuit-error")).pop()
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
    errorDiv = Array.from(document.getElementsByClassName("month-error")).pop()
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
    errorDiv = Array.from(document.getElementsByClassName("year-error")).pop()
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
    errorDiv = Array.from(document.getElementsByClassName("status-error")).pop()
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
    errorDiv = Array.from(document.getElementsByClassName("amount-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isNumeric(errorDiv, "", value)){
        return false
    }
    return isNotLongerThan(errorDiv, 20, value)
}



function validatePaymentDate(value){
    errorDiv = Array.from(document.getElementsByClassName("payment-date-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    return isValidDate(errorDiv, value, 2000)
    
    // devuelvo true o false dependiendo de la validacion
}

function validateCommentary(value){
    errorDiv = Array.from(document.getElementsByClassName("commentary-error")).pop()
    return isNotLongerThan(errorDiv, 400, value)
}

function validateRelationship(value){
    errorDiv = Array.from(document.getElementsByClassName("relationship-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }

    if (!isAlphabetic(errorDiv, "", value)){
        return false
    }
    return isNotLongerThan(errorDiv, 20, value)
}