import axios from "axios";

const API = axios.create({
    headers: {
       'Content-Type': 'multipart/form-data',
    },
    timeout: 2000
})

export default API