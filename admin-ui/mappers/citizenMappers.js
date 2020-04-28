function citizenRecordsTableMapping(data, isListView = false){
    var limit = data.length >= 5 ? 5 : data.length;
    for(var i=0; i<limit; i++){
        var tableRow = document.createElement("tr");

        var nik = document.createElement("th");
        nik.innerHTML = data[i].NIK;

        var name = document.createElement("td")
        name.innerHTML = data[i].name;

        var sex = document.createElement("td");
        sex.innerHTML = data[i].sex;

        var address = document.createElement("td");
        address.innerHTML = data[i].current_address;

        tableRow.appendChild(nik);
        tableRow.appendChild(name);
        tableRow.appendChild(sex);
        tableRow.appendChild(address)

        if(isListView){
            addActionButtonsToRow(tableRow, data[i].NIK);
        }

        document.getElementById("citizenRecords").appendChild(tableRow)

    }
}

function addActionButtonsToRow(tableRow,nik){
    var actions = document.createElement("td");

    var modifyButton = document.createElement("a");
    var url = "forms/citizenForm.html?nik="+nik;
    modifyButton.setAttribute("href", url);
    modifyButton.innerHTML = "MODIFY";

    actions.appendChild(modifyButton);
    tableRow.appendChild(actions);
}

function citizenFormMapping(data){
    const fullNameField = document.getElementById("input-name");
    fullNameField.value = data.name;

    const nikField = document.getElementById("input-nik");
    nikField.value = data.NIK;

    const birthPlaceField = document.getElementById("input-birth-place");
    birthPlaceField.value = data.birth_place;

    const occupationField = document.getElementById("input-occupation");
    occupationField.value = data.occupation

    const dateOfBirthField = document.getElementById("input-date-of-birth");
    dateOfBirthField.value = dateConverter(data.birth_date);

    const dissabilityField = document.getElementById("input-dissability");
    dissabilityField.value = data.dissability;

    const familyCardField = document.getElementById("input-family-card");
    familyCardField.value = data.family_card_number;

    const addressField = document.getElementById("input-address");
    addressField.value = data.current_address;

    const fathersNIK = document.getElementById("input-fathers-nik");
    fathersNIK.value = data.NIK_of_father;  

    const mothersNIK = document.getElementById("input-mothers-nik");
    mothersNIK.value = data.NIK_of_mother;
    
    const sexOptions = document.getElementById("input-sex").children;
    for(i=0;i<sexOptions.length;i++){
        console.log();
        if(data.sex === sexOptions[i].value){
            sexOptions[i].setAttribute("selected", "true");
        }
    }

    const marriageOptions = document.getElementById("input-marriage-status").children;
    for(i=0;i<marriageOptions.length;i++){
        if(data.married_status == marriageOptions[i].value){
            marriageOptions[i].setAttribute("selected", "true");
        }
    }

    const bloodTypeOptions = document.getElementById("input-blood-type").children;
    for(i=0;i<bloodTypeOptions.length;i++){
        if(data.blood_type == bloodTypeOptions[i].value){
            bloodTypeOptions[i].setAttribute("selected", "true")
        }
    }

    const religionOptions = document.getElementById("input-religion").children;
    for(i=0;i<religionOptions.length;i++){
        if(data.religion == religionOptions[i].value){
            religionOptions[i].setAttribute("selected", "true")
        }
    }
}

function dateConverter(stringDate){
    const [day, month, year] = stringDate.split("-")
    date =  new Date(year, month - 1, day);
    return [date.getFullYear(), month, day].join("-");
}

function stringToDate(dateInput){
    const [year, month, day] = dateInput.split("-")
    date =  new Date(year, month - 1, day);
    return [day, month, date.getFullYear()].join("-");
}