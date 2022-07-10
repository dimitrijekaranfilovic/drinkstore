import BaseService from "./baseService";

class PurchaseService extends BaseService {
  constructor() {
    super();
    this.basePath = process.env.VUE_APP_PURCHASE_SERVICE_BASE_PATH;
  }

  createPurchase(payload) {
    return this.post(`${this.basePath}`, payload);
  }

  getUserHistory(userId) {
    return this.get(`${this.basePath}/history-for-user/${userId}`);
  }

  getMostSold(period) {
    return this.get(`${this.basePath}/most-sold?period=${period}`);
  }
}

export const purchaseService = new PurchaseService();
