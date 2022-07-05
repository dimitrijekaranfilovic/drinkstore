import BaseService from "./baseService";

class UserService extends BaseService {
  authenticate(payload) {
    return this.post(
      `${process.env.VUE_APP_USER_SERVICE_BASE_PATH}/authenticate`,
      payload
    );
  }

  register(payload) {
    return this.post(
      `${process.env.VUE_APP_USER_SERVICE_BASE_PATH}/register`,
      payload
    );
  }
}

export const userService = new UserService();
