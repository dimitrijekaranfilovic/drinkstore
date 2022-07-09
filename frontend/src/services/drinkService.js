import BaseService from "./baseService";

class DrinkService extends BaseService {
  constructor() {
    super();
    this.basePath = process.env.VUE_APP_DRINK_SERVICE_BASE_PATH;
  }

  getDrinks(params) {
    return this.getWithParams(`${this.basePath}/get`, params);
  }

  getSingleDrink(drinkId) {
    return this.get(`${this.basePath}/get/${drinkId}`);
  }

  createDrink(payload) {
    return this.post(`${this.basePath}/create`, payload);
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

  addUserGrade(drinkId, payload) {
    return this.post(`${this.basePath}/${drinkId}/grades`, payload);
  }

  updateUserGrade(drinkId, gradeId, payload) {
    return this.put(`${this.basePath}/${drinkId}/grades/${gradeId}`, payload);
  }

  deleteUserGrade(drinkId, gradeId) {
    return this.delete(`${this.basePath}/${drinkId}/grades/${gradeId}`);
  }

  checkUserGrade(drinkId) {
    return this.get(`${this.basePath}/${drinkId}/grade-for-user`);
  }
}

export const drinkService = new DrinkService();
