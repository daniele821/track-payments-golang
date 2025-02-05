import { makeRequest } from "./utils.js"

makeRequest("select-city", {})
    .then(res => console.log(res))
