const API_BASE_URL = "http://localhost:8080/api";
function apiUrl(endpoint) {
  return `${API_BASE_URL}${endpoint}`;
}

export { apiUrl as a };
