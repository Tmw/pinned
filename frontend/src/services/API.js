const BASE_URL = "/api";

const API = {
  FetchChallenges: () => fetch(`${BASE_URL}/challenges`).then(r => r.json())
};

export default API;
