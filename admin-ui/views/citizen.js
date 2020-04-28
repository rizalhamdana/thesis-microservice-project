function showCitizenListView(){
    isAuth()
    token = getCookie("Token")
    citizens = getAllCitizen(token)
    citizens.then(data => {
        if (data instanceof Error){
            document.getElementById("numberCitizen").innerHTML = "-";
            return
        }
        var isListView = true;
        citizenRecordsTableMapping(data, isListView)
    }).catch(error=>{
    console.log(error)
    })
}

function showCitizenForm(){
    isAuth();
    token = getCookie("Token");
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const nik = urlParams.get("nik");
    
    if (nik != null){
        const citizen = getOneCitizen(token, nik);

        citizen.then(data => {
        
        if(data instanceof Error){
            errorMessage = document.getElementById("error-message");
            errorMessage.innerHTML = data;

            hiddenForm = document.getElementById("form-section");
            hiddenForm.style.visibility = 'hidden';
            return 
        }

        hiddenInput = document.getElementById("hidden-nik");
        hiddenInput.value = data.NIK;

        deleteButton = document.getElementById("delete-button");
        deleteButton.setAttribute("onclick", "deleteCitizen()")

        citizenFormMapping(data);

        }).catch(error => {
            errorMessage = document.getElementById("error-message");
            errorMessage.innerHTML = error;

            hiddenForm = document.getElementById("form-section");
            hiddenForm.style.visibility = 'hidden';
            return 
        })
    } else {
        deleteButton = document.getElementById("delete-button");
        deleteButton.style.visibility = 'hidden';

        document.getElementById("modify-button").style.visibility = 'hidden';
        document.getElementById("dissability").removeAttribute("disabled");
        submitButton = document.getElementById("submit-button");
        submitButton.setAttribute("onclick", "saveCitizen()");

    } 
    
}

function saveCitizen(){
    isAuth();
    token = getCookie("Token");
    const birthDate = stringToDate(document.getElementById("input-date-of-birth").value)
    const data = {
        NIK: document.getElementById("input-nik").value,
        name: document.getElementById("input-name").value,
        sex: document.getElementById("input-sex").value,
        family_card_number: document.getElementById("input-family-card").value,
        birth_place: document.getElementById("input-birth-place").value,
        blood_type: document.getElementById("input-blood-type").value,
        religion: document.getElementById("input-religion").value,
        married_status: document.getElementById("input-marriage-status").value,
        occupation: document.getElementById("input-occupation").value,
        dissability: document.getElementById("input-dissability").value,
        current_address: document.getElementById("input-address").value,
        NIK_of_mother: document.getElementById("input-mothers-nik").value,
        NIK_of_father: document.getElementById("input-fathers-nik").value,
        birth_date: birthDate,
    };
    result = insertCitizen(token, data)
    result.then(data => {
        if(data instanceof Error){
            window.location = "../citizens.html"
            return
        }
        window.location = "../citizens.html"
        return
    }).catch(error => {
        window.location = "../citizens.html"
        return
    })
}


function deleteCitizen(){
    isAuth();
    token = getCookie("Token");

    nik = document.getElementById("hidden-nik").value;
    

    const retVal = confirm("Are you sure want to delete this record?")
    if (retVal == true){
        result = deleteOneCitizen(token, nik);

        result.then(data=>{
            console.log(data)
            if(data instanceof Error){
                alert("Could not delete this record, message:"+data);
                return
            }
            alert("Record has successfully deleted");
            window.location = "../citizens.html"

        })
        return 
    } else {
        return
    }   
}

function updateCitizen(){
    isAuth();
    token = getCookie("Token");
    const birthDate = stringToDate(document.getElementById("input-date-of-birth").value)
    console.log(birthDate);
    const data = {
        NIK: document.getElementById("input-nik").value,
        name: document.getElementById("input-name").value,
        sex: document.getElementById("input-sex").value,
        family_card_number: document.getElementById("input-family-card").value,
        birth_place: document.getElementById("input-birth-place").value,
        blood_type: document.getElementById("input-blood-type").value,
        religion: document.getElementById("input-religion").value,
        married_status: document.getElementById("input-marriage-status").value,
        occupation: document.getElementById("input-occupation").value,
        dissability: document.getElementById("input-dissability").value,
        current_address: document.getElementById("input-address").value,
        NIK_of_mother: document.getElementById("input-mothers-nik").value,
        NIK_of_father: document.getElementById("input-fathers-nik").value,
        birth_date: birthDate,
    };
    result = updateOneCitizen(token, data, document.getElementById("input-nik").value)

    result.then(data => {
        if(data instanceof Error){
            alert("Could not update this record, message:"+data);
            window.location = "../citizens.html"
            return
        }
        alert("Record has successfully updated");
        window.location = "../citizens.html"
        return
    }).catch(error => {

        alert("Could not update this record, message:"+ error);
        window.location = "../citizens.html"
        return
    })
}



function modifyUnlock(){
    document.getElementById("dissability").removeAttribute("disabled")
    document.getElementById("input-nik").setAttribute("disabled", true)
}