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

function validateAffiliated(value){
    errorDiv = Array.from(document.getElementsByClassName("affiliated-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    if (value == "true" && Array.from(document.getElementsByName("id-enterprise")).pop().value == "1"){
        return false
    }
    if (value == "false" && Array.from(document.getElementsByName("id-enterprise")).pop().value != "1"){
        return false
    }
    return isBoolean(value)
}

function isBoolean(value){
    if (value == true || value == false || value == "true" || value =="false"){
        errorDiv.style.display = 'none'
        return true
    } else {
        errorDiv.innerHTML = 'El campo no contiene un valor válido'
        errorDiv.style.display = 'inline'
        return false
    }
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

function validateContact(value){
    errorDiv = Array.from(document.getElementsByClassName("contact-error")).pop()
    return isNotLongerThan(errorDiv, 200, value)
}

async function validateEnterpriseNumber(value){
    errorDiv = Array.from(document.getElementsByClassName("enterprise-number-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    if (!isNumeric(errorDiv, "", value)){
        return false
    }
    
    if (!isNotLongerThan(errorDiv, 10, value)){
        return false
    }
    return await checkEnterpriseNumber(value)
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

function validateIdEnterprise(value){
    console.log("entro a validate")
    console.log(value)
    console.log(typeof value)
    console.log(Array.from(document.getElementsByName("affiliated")).pop().value)
    console.log(typeof Array.from(document.getElementsByName("affiliated")).pop().value)
    errorDiv = Array.from(document.getElementsByClassName("enterprise-error")).pop()
    if (!isNotEmpty(errorDiv, value)){
        return false
    }
    
    if (!isNumeric(errorDiv, "", value)){
        return false
    }
    
    if (!isNotLongerThan(errorDiv, 10, value)){
        return false
    }

    if (value == "0"){
        errorDiv.innerHTML = "Empresa no válida."
        errorDiv.style.display = 'inline'
        return false
    }

    if (value == "1" && Array.from(document.getElementsByName("affiliated")).pop().value == "true"){
        errorDiv.innerHTML = "Si está afiliado debe pertenecer a una empresa."
        errorDiv.style.display = 'inline'
        return false
    }
    if (value != "1" && Array.from(document.getElementsByName("affiliated")).pop().value == "false"){
        errorDiv.innerHTML = "Si no está afiliado no puede pertenecer a una empresa."
        errorDiv.style.display = 'inline'
        return false
    }

    return true
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
    return isBoolean(value)

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
    // en este caso, puede estar vacio, por eso el true
    if (!isNotEmpty(errorDiv, value)){
        errorDiv.style.display = 'none'
        return true
    }
    
    return isValidDate(errorDiv, value, 2000)
    
    // devuelvo true o false dependiendo de la validacion
}

function validateObservations(value){
    errorDiv = Array.from(document.getElementsByClassName("observations-error")).pop()
    return isNotLongerThan(errorDiv, 1000, value)
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

// other functions

// function getEnterprisesId(){ 
//     // esto no devuelve true o false, devuelve una promesa que devuelve true o false
//         return fetch('http://localhost:8080/enterprise/getAllEnterprisesId', {
//             method: 'GET',
//             headers: {
//               Accept: 'application/json',
//               'Content-Type': 'application/json'
//             }})
//           .then(res => res.json())      
// }

// async function checkIdEnterprise(errorDiv, value){
//     try {
//         enterprisesId = await getEnterprisesId()
//         value = parseInt(value)
//         if (enterprisesId.includes(value)){
//             errorDiv.style.display = 'none'
//             return true
//         } else{
//             errorDiv.style.display = 'inline'
//             errorDiv.innerHTML = 'Empresa no válida'
//             return false
//   }
//     } catch (error) {
//         console.log(error)
//     }
   
// }


//  function getEnterpriseNumbers(){
//     // esto no devuelve true o false, devuelve una promesa que devuelve true o false
//     return fetch('http://192.168.100.2:8080/enterprise/getAllEnterprisesNumber', {
//         method: 'GET',
//         headers: {
//             Accept: 'application/json',
//             'Content-Type': 'application/json'
//         }})
//         .then(res => res.json())
// }

// async function checkEnterpriseNumber(value){
//     try {
//         let enterprisesNumbers = await getEnterpriseNumbers()
//         // para poder editar sin cambiar el number enterprise, sino te tira que ya existe
//         console.log(enterprisesNumbers)
//         oldValue = Array.from(document.getElementsByName("old-enterprise-number")).pop().value
//         if (oldValue == value) {
//             errorDiv.style.display = 'none'
//             return true
//         } else if (enterprisesNumbers.includes(value)){
//             errorDiv.innerHTML = 'Numero de empresa ya registrado'
//             errorDiv.style.display = 'inline'
//             return false
//         } else{
//             errorDiv.style.display = 'none'
//             return true
//     }
//         } catch (error) {
//             console.log(err)
//         }
    
// }