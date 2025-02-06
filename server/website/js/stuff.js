import { makeRequest } from "./utils.js"

const sep = ", "
makeRequest("select-city", {})
    .then(res => res.json())
    .then(res => console.log("CITY:", res.res.result.join(sep)))
makeRequest("select-method", {})
    .then(res => res.json())
    .then(res => console.log("METHOD:", res.res.result.join(sep)))
makeRequest("select-shop", {})
    .then(res => res.json())
    .then(res => console.log("SHOP:", res.res.result.join(sep)))
makeRequest("select-category", {})
    .then(res => res.json())
    .then(res => console.log("CATEGORY:", res.res.result.join(sep)))
makeRequest("select-item", {})
    .then(res => res.json())
    .then(res => console.log("ITEM:", res.res.result.join(sep)))
