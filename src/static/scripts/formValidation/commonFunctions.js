function getInputValue(name){
    return Array.from(document.getElementsByName(name)).pop().value
}


function isNotEmpty(errorDiv, string){
    if (string == ""){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "El campo esta vacío."
        return false
    } else{
        errorDiv.style.display = 'none'
        return true
    }
}


function isAlphanumeric(errorDiv, allowedCharacter, value){
    validCharacters = " abcdefghijklmnñopqrstuvwxyzáéíóúüÃ0123456789"
	validCharacters += allowedCharacter
    return isValidCharacter(errorDiv, validCharacters, value)
}

function isAlphabetic(errorDiv, allowedCharacter, value){
    validCharacters = " abcdefghijklmnñopqrstuvwxyzáéíóúüÃ"
	validCharacters += allowedCharacter
    return isValidCharacter(errorDiv, validCharacters, value)
}

function isNumeric(errorDiv, allowedCharacter, value){
    validCharacters = " 0123456789"
	validCharacters += allowedCharacter
    return isValidCharacter(errorDiv, validCharacters, value)
}

function isValidCharacter(errorDiv, validCharacters, value){
    value = value.toLowerCase()
    for (let i=0; i<value.length; i++){
        if (validCharacters.includes(value[i])){
            continue
        } else {
            errorDiv.style.display = 'inline'
            errorDiv.innerHTML = "Caracter inválido."
            return false
        }
    }
    return true
}

function isNotLongerThan(errorDiv, limit, value){
    if (value.length>limit){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = `El campo contiene más de ${limit} caracteres.`
        return false
    } else {
        errorDiv.style.display = 'none'
        return true
    }
}

function isAValidOption(errorDiv, options, value){
    if (options.includes(value)){
        errorDiv.style.display = 'none'
        return true
    } else {
        errorDiv.style.display = 'inline'
        return false
    }
}

function isValidDate(errorDiv, paymentDate, maxYear){
    // verifico que la longitud sea igual a 10
    if (paymentDate.length != 10){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "Formato de fecha erroneo."
        return false
    }
    
    // verifico que solo contenga numeros o "/"
    if (!isNumeric(errorDiv, "/", paymentDate)){
        return false
    }
    // cambiarlo por is alphanumeric

    // separo el string en partes aislando el dia, mes y año, ignorando las "/"
    var dateParts = paymentDate.split("/")
    var day = parseInt(dateParts[0])
    var month = parseInt(dateParts[1])
    var year = parseInt(dateParts[2])

    // verifico que no sea una fecha posterior a la actual
   
    // primero creo la fecha actual
    var actualDate = new Date(Date.now())

    // creo en formato fecha la del input, el segundo parametro es el index del mes, por ejemplo enero es el 0
    var inputDate = new Date(year, month-1, day)

    // ahora los comparo
    if (actualDate.getTime() < inputDate.getTime()){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "Fecha inválida"
        return false
    }

    // verifico que el año no sea menor al maximo indicado
    if (year < maxYear) {
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "Año inválido."
        return false
    }

    // verifico que sea un mes entre 1 y 12
    if ((month < 1) || (month > 12)){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "Mes inválido."
        return false
    }

    // verifico que sea un dia entre 1 y 31
    if ((day < 1) || (day > 31)){
        errorDiv.style.display = 'inline'
        errorDiv.innerHTML = "Día inválido."
        return false
    }

    switch (month){
        case 2:
            // en febrero, si no es bisiesto y es mayor a 28 es una fecha invalida
            if ((year % 4 != 0) && (day > 28)) {   
                errorDiv.style.display = 'inline'
                errorDiv.innerHTML = "Fecha inválida."
                return false
            }
            // en febrero, si es bisiesto y es mayor a 29 es una fecha invalida
            if ((year % 4 == 0) && (day > 29)) {
                errorDiv.style.display = 'inline'
                errorDiv.innerHTML = "Fecha inválida."
                return false                    
            }      
        case 4, 6, 9, 11: if (day > 30) {
            errorDiv.style.display = 'inline'
            errorDiv.innerHTML = "Fecha inválida."
			return false
		}
	}

    // si no hay ningun error devolvemos true
    errorDiv.style.display = 'none'
    return true
}

