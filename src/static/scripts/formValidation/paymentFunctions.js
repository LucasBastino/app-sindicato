
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

    if (!isNumber(errorDiv, value)){
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






