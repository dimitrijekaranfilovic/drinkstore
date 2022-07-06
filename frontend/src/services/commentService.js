import BaseService from "./baseService";

class CommentService extends BaseService {
  constructor() {
    super();
    this.basePath = process.env.VUE_APP_COMMENT_SERVICE_BASE_PATH;
  }

  getCommentsForDrink(drinkId) {
    return this.get(`${this.basePath}/for-drink/${drinkId}`);
  }

  addComment(payload) {
    return this.post(`${this.basePath}`, payload);
  }

  addReport(payload) {
    return this.post(`${this.basePath}/reports`, payload);
  }

  getReports() {
    return this.get(`${this.basePath}/reports`);
  }
}

export const commentService = new CommentService();
