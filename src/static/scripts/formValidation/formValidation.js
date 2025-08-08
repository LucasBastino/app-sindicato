async function validateMember(action){
    
    var validationFields = {
        name: validateName(getInputValue('name')),
        lastName: validateLastName(getInputValue('last-name')),
        dni: validateDNI(getInputValue('dni')),
        birthday: validateBirthday(getInputValue('birthday')),
        gender: validateGender(getInputValue('gender')),
        maritalStatus: validateMaritalStatus(getInputValue('marital-status')),
        phone: validatePhone(getInputValue('phone')),
        email: validateEmail(getInputValue('email')),
        address: validateAddress(getInputValue('address')),
        postalCode: validatePostalCode(getInputValue('postal-code')),
        district: validateDistrict(getInputValue('district')),
        memberNumber: validateMemberNumber(getInputValue('member-number')),
        affiliated: validateAffiliated(getInputValue('affiliated')),
        cuil: validateCUIL(getInputValue('cuil')),
        idEnterprise: validateIdEnterprise(getInputValue('id-enterprise')),
        category: validateCategory(getInputValue('category')),
        entryDate: validateEntryDate(getInputValue('entry-date')),
        observations: validateObservations(getInputValue('observations')),
    }
    
    if (checkValidationFields(validationFields)){
        postForm("member", action) 
    }

}

function validatePayment(action){
    
    var validationFields = {
        month: validateMonth(getInputValue('month')),
        year: validateYear(getInputValue('year')),
        status: validateStatus(getInputValue('status')),
        amount: validateAmount(getInputValue('amount')),
        paymentDate: validatePaymentDate(getInputValue('payment-date')),
        observations: validateObservations(getInputValue('observations')),
    }
    
    if (checkValidationFields(validationFields)){
        postForm("payment", action)
    }
}


async function validateEnterprise(action){
    
    var validationFields = {
        name: validateEnterpriseName(getInputValue('name')),
        number: await validateEnterpriseNumber(getInputValue('enterprise-number')),
        address: validateAddress(getInputValue('address')),
        cuit: validateCUIT(getInputValue('cuit')),
        district: validateDistrict(getInputValue('district')),
        postalCode: validatePostalCode(getInputValue('postal-code')),
        phone: validatePhone(getInputValue('phone')),
        contact: validateContact(getInputValue('contact')),
        observations: validateObservations(getInputValue('observations')),
    }
    if (checkValidationFields(validationFields)){
        postForm("enterprise", action)
    }
}


function validateParent(action){
    
    var validationFields = {
        name: validateName(getInputValue('name')),
        lastName: validateLastName(getInputValue('last-name')),
        relationship: validateRelationship(getInputValue('rel')),
        birthday: validateBirthday(getInputValue('birthday')),
        gender: validateGender(getInputValue('gender')),
        cuil: validateCUIL(getInputValue('cuil')),
    }
    
    if (checkValidationFields(validationFields)){
        postForm("parent", action)
    }
}


function checkValidationFields(validation){
    for (let field in validation){
        if (!validation[field]) {
            return false
        }
    }
    return true
}

function postForm(model, action){
    document.getElementById(`submit-${action}-${model}-btn`).click()
}
