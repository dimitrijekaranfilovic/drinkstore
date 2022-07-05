import BaseService from "./baseService";

class DrinkService extends BaseService {
  getDrinks(params) {
    return this.getWithParams(
      `${process.env.VUE_APP_DRINK_SERVICE_BASE_PATH}`,
      params
    );
  }

  getSingleDrink(drinkId) {
    return this.get(
      `${process.env.VUE_APP_DRINK_SERVICE_BASE_PATH}/${drinkId}`
    );
  }

  createDrink(payload) {
    return this.post(`${process.env.VUE_APP_DRINK_SERVICE_BASE_PATH}`, payload);
  }

  createDrinkImage(drinkId, imageFile) {
    return this.postFormData(
      `${process.env.VUE_APP_DRINK_SERVICE_BASE_PATH}/${drinkId}/images`,
      imageFile,
      "image"
    );
  }

  deleteDrink(drinkId) {
    return this.delete(
      `${process.env.VUE_APP_DRINK_SERVICE_BASE_PATH}/${drinkId}`
    );
  }
}

export const drinkService = new DrinkService();
