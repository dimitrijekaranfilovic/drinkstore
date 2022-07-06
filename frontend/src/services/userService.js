import BaseService from "./baseService";

class UserService extends BaseService {
  constructor() {
    super();
    this.basePath = process.env.VUE_APP_USER_SERVICE_BASE_PATH;
  }
  authenticate(payload) {
    return this.post(`${this.basePath}/authenticate`, payload);
  }

  register(payload) {
    return this.post(`${this.basePath}/register`, payload);
  }

  banUser(username) {
    return this.get(`${this.basePath}/ban/${username}`);
  }
}

export const userService = new UserService();
