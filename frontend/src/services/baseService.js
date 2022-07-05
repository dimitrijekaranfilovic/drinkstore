const axios = require("axios");

export default class BaseService {
  get(url) {
    return axios.get(url, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("jwt"),
      },
    });
  }

  getWithParams(url, params) {
    let newParams = { ...params };

    for (let key in params) {
      if (Array.isArray(params[key])) {
        newParams[key] = params[key].join(",");
      }
    }

    return axios.get(url, {
      params: newParams,
      headers: {
        Authorization: "Bearer " + localStorage.getItem("jwt"),
      },
    });
  }

  post(url, payload) {
    return axios.post(url, payload, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("jwt"),
      },
    });
  }

  put(url, payload) {
    return axios.put(url, payload, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("jwt"),
      },
    });
  }

  delete(url) {
    return axios.delete(url, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("jwt"),
      },
    });
  }

  postFormData(url, file, filename) {
    const formData = new FormData();
    formData.append(filename, file);
    return axios.post(url, formData, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("jwt"),
        "Content-Type": "multipart/form-data",
      },
    });
  }
}
