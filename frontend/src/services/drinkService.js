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
}

export const drinkService = new DrinkService();
