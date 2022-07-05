const axios = require("axios");

export default class BaseService {
  buildConfig() {
    let config = {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("jwt")}`,
      },
    };
    return config;
  }

  get(url) {
    const config = this.buildConfig();
    return axios.get(url, config);
  }

  getWithParams(url, params) {
    const config = this.buildConfig();
    return axios.get(
      url,
      {
        params: params,
      },
      config
    );
  }

  post(url, payload) {
    const config = this.buildConfig();
    return axios.post(url, payload, config);
  }

  put(url, payload) {
    const config = this.buildConfig();
    return axios.put(url, payload, config);
  }

  delete(url) {
    const config = this.buildConfig();
    return axios.delete(url, config);
  }
}
