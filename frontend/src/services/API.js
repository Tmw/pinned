const BASE_URL = "/api";

const API = {
  FetchChallanges: () => fetch(`${BASE_URL}/challanges`).then(r => r.json())
};

export default API;
