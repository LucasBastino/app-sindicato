

function validateForm(model, action){
    switch (model){
        case 'member': validateMember(action)
        case 'enterprise': validateEnterprise(action)
        case 'parent': validateParent(action)
        case 'payment': validatePayment(action)
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


function validatePayment(action){
    
    var validationFields = {
        month: validateMonth(getInputValue('month')),
        year: validateYear(getInputValue('year')),
        status: validateStatus(getInputValue('status')),
        amount: validateAmount(getInputValue('amount')),
        paymentDate: validatePaymentDate(getInputValue('payment-date')),
        commentary: validateCommentary(getInputValue('commentary')),
    }
    
    if (checkValidationFields(validationFields)){
        postForm("payment", action)
    }
}


function validateEnterprise(action){
    
    var validationFields = {
        name: validateEnterpriseName(getInputValue('name')),
        number: validateEnterpriseNumber(getInputValue('enterprise-number')),
        address: validateEnterpriseAddress(getInputValue('address'))
        /* year: validateYear(getInputValue('year')),
        status: validateStatus(getInputValue('status')),
        amount: validateAmount(getInputValue('amount')),
        paymentDate: validatePaymentDate(getInputValue('payment-date')),
        commentary: validateCommentary(getInputValue('commentary')), */
    }
    
    if (checkValidationFields(validationFields)){
        // postForm("enterprise", action)
        console.log("todos verdaderos")
    } else{
         console.log("hay alguno falso")
    }
}