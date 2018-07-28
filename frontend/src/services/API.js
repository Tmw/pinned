const BASE_URL = "/api";

const API = {
  FetchChallanges: () => fetch(`${BASE_URL}/challanges`).then(r => r.json()),

  CheckAnswers: answers => {
    // do some async magic and return score
  }
};

export default API;
