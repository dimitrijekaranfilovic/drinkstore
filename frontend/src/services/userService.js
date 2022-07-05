const axios = require("axios");

class UserService {
  authenticate(payload) {
    return axios.post(
      `${process.env.VUE_APP_USER_SERVICE_BASE_PATH}/authenticate`,
      payload
    );
  }
}

export const userService = new UserService();
