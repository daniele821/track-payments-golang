export async function makeRequest(typeField, dataField) {
    return fetch("", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ "type": typeField, "data": dataField }),
    }).then(res => res.json())
}
