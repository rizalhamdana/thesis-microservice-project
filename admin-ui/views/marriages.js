function showMarriagesListView(){
    isAuth();
    var token = getCookie("Token");
    marriages = getAllMarriage(token);
    marriages.then(data => {
    if (data instanceof Error){
        document.getElementById("numberMarriage").innerHTML = "-"
        return
    }
    
    marriageRecordsTableMapping(data, true);
   });
}
function showMarriageForm(){
    isAuth();
    token = getCookie("Token");
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const certificateNumber = urlParams.get("married_certificate_number");
    

    fetchResult = getOneMarriage(token, certificateNumber)
    fetchResult.then(marriage => {
        if (marriage instanceof Error){
            errorMessage = document.getElementById("error-message");
            errorMessage.innerHTML = data;

            hiddenForm = document.getElementById("form-section");
            hiddenForm.style.visibility = 'hidden';
            return 
        }

        marriageFormMapping(marriage);

    })
}
function verifyMarriage(){
    isAuth();
    token = getCookie("Token");
    const regisNumber = document.getElementById("input-regis-number").value;
    
    const verify = verifyOneMarriage(token, regisNumber)
    verify.then(data => {
        if(data instanceof Error){
            alert("Failed Verifying Record: "+ data)
        }
        alert("Verifying record success")
        window.location  = "../marriages.html"

    })
}