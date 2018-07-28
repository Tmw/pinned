const BASE_URL = "http://localhost:4000/api";

const API = {
  FetchChallanges: () => fetch(`${BASE_URL}/challanges`).then(r => r.json()),

  CheckAnswers: answers => {
    // do some async magic and return score
  }
};

export default API;
