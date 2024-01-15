var formdata = {}
var schema_data = []

function updateSchemaData(parentId, subId, bool) {
    const schema = schema_data.find((schema) => schema.id == parentId);
    const subSchema = schema.values.find((subSchema) => subSchema.subId == subId);
    subSchema.checked = bool;
}

const changehandler = (checkbox) => {
    const parentId = checkbox.value.split("-")[0];
    const subId = checkbox.value.split("-")[1];
    if (checkbox.checked) {
        updateSchemaData(parentId, subId, true);
    } else {
        updateSchemaData(parentId, subId, false);   
    }
}

const submitHandler = async (e) => {
    e.preventDefault();
    console.log(schema_data);

    const req_body = [];
    for (const s of schema_data) {
        const filteredValues = s.values.filter(v => v.checked);
        if (filteredValues.length > 0) {
            let newSchema = { ...s, values: filteredValues };
            req_body.push(newSchema);
        }
    }

    const dateDeveloped = document.getElementById("date-developed").value;
    const approvalDate = document.getElementById("approval-date").value;
    const dateLastReviewed = document.getElementById("date-last-reviewed").value;
    const nextReviewDate = document.getElementById("next-review-date").value;
 

    const response = await fetch("/api/swms", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            projectAddress: document.getElementById("projectAddress").value,
            scopeOfWork: document.getElementById("scopeOfWork").value,
            dateDeveloped: dateDeveloped ? new Date(dateDeveloped).toLocaleDateString() : null,
            approvalDate: approvalDate ? new Date(approvalDate).toLocaleDateString() : null,
            dateLastReviewed: dateLastReviewed ? new Date(dateLastReviewed).toLocaleDateString() : null,
            nextReviewDate: nextReviewDate ? new Date(nextReviewDate).toLocaleDateString() : null,
            tableData: req_body
        })
    });
    if (response.ok) {
        //redirect to home page
        window.location.href = "/";
    }
}

document.addEventListener("DOMContentLoaded", async function(event) {
    console.log('is running')
    var response = await fetch("/api/swms/schema");
    if (response.ok) {
        var data = await response.json();
        if (Array.isArray(data)) {
            schema_data = data;
      
        }
    }
});

document.addEventListener("DOMContentLoaded", function(event) {
    document.getElementById("form").addEventListener("submit", submitHandler);
});
