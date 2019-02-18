import axios from "axios";

let backendHost = process.env.VUE_APP_BACK_HOST || "localhost";
let backendPort = process.env.VUE_APP_BACK_PORT || "7000";
let baseURL = `http://${backendHost}:${backendPort}`;

export default axios.create({
  baseURL,
  headers: {
    'Access-Control-Allow-Origin': "*",
    'Access-Control-Allow-Headers':'Content-Type',
    'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
    'Content-Type': 'application/json;charset=UTF-8',
}
});
