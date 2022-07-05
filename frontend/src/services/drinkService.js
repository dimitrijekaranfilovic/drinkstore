import BaseService from "./baseService";

class DrinkService extends BaseService {
  constructor() {
    super();
    this.basePath = process.env.VUE_APP_DRINK_SERVICE_BASE_PATH;
  }

  getDrinks(params) {
    return this.getWithParams(`${this.basePath}`, params);
  }

  getSingleDrink(drinkId) {
    return this.get(`${this.basePath}/${drinkId}`);
  }

  createDrink(payload) {
    return this.post(`${this.basePath}`, payload);
  }

  createDrinkImage(drinkId, imageFile) {
    return this.postFormData(
      `${this.basePath}/${drinkId}/images`,
      imageFile,
      "image"
    );
  }

  deleteDrink(drinkId) {
    return this.delete(`${this.basePath}/${drinkId}`);
  }

  updateDrink(drinkId, payload) {
    return this.put(`${this.basePath}/${drinkId}`, payload);
  }
}

export const drinkService = new DrinkService();
