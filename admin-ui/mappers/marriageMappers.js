function marriageRecordsTableMapping(data, isListView = false){
    var limit = data.length >= 5 ? 5 : data.length;
    for(var i=0; i<limit; i++){
        var tableRow = document.createElement("tr");

        var certificateNumber = document.createElement("th");
        certificateNumber.innerHTML = data[i].married_certificate_number;

        var husbandName = document.createElement("td");
        husbandName.innerHTML = data[i].husband_name;

        var wifeName = document.createElement("td");
        wifeName.innerHTML = data[i].wife_name;

        tableRow.appendChild(certificateNumber);
        tableRow.appendChild(husbandName);
        tableRow.appendChild(wifeName);

        if(isListView){
            const verifiedStatus = document.createElement("td");
            const verifiedValue = document.createElement("span");
            const verifiedClass = data[i].verified_status ? "ni ni-check-bold" : "ni ni-fat-remove"
            verifiedValue.setAttribute("class", verifiedClass);

            verifiedStatus.appendChild(verifiedValue);
            tableRow.appendChild(verifiedStatus);

            const action = document.createElement("td");
            const anchorDetails = document.createElement("a");
            var url = "forms/marriageForm.html?married_certificate_number="+data[i].married_certificate_number;
            anchorDetails.setAttribute("href", url);
            anchorDetails.setAttribute("class", "ni ni-badge");

            action.appendChild(anchorDetails);
            tableRow.appendChild(action);
        }

        document.getElementById("marriageRecords").appendChild(tableRow);
    }
}

function marriageFormMapping(data){
    const fullNameField = document.getElementById("certificate-number");
    fullNameField.innerHTML = "Married Certificate Number: " + data.married_certificate_number;

    const husbandNIK = document.getElementById("input-husbands-nik");
    husbandNIK.value = data.husband_nik;

    const husbandName = document.getElementById("input-husbands-name");
    husbandName.value = data.husband_name;

    const wifeNIK = document.getElementById("input-wifes-nik");
    wifeNIK.value = data.wife_nik;

    const wifeName = document.getElementById("input-wifes-name");
    wifeName.value = data.wife_name;

    const courtName = document.getElementById("input-court-name");
    courtName.value = data.court_name;
    
    const courtDecisionNumber = document.getElementById("input-court-decision-number");
    courtDecisionNumber.value = data.court_decision_number;

    const isVerifiedLogo = document.getElementById("is-verified");
    const verifiedClass = data.verified_status ? "ni ni-check-bold" : "ni ni-fat-remove";
    isVerifiedLogo.setAttribute("class", verifiedClass)

    const regisNumber = document.getElementById("input-regis-number");
    regisNumber.value = data.regis_number;
    

    // if (data.verified_status){
    //     document.getElementById("verify-button").style.visibility = "hidden";
    // }
}