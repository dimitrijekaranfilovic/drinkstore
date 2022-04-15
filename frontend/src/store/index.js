import Vue from "vue";
import Vuex from "vuex";
import createPersistedState from "vuex-persistedstate";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    cart: [],
  },
  mutations: {
    ADD_TO_CART(cartItem) {
      state.cart.push(cartItem);
    },
    REMOVE_FROM_CART(cartItemId) {
      state.cart = state.cart.filter((item) => item.id !== cartItemId);
    },
    UPDATE_CART_ITEM(updatedCartItem) {
      let itemToUpdate = state.cart.find(
        (item) => item.id === updatedCartItem.id
      );
      itemToUpdate.amount = updatedCartItem.amount;
    },
  },
  actions: {
    //item: id i kolicina
    addToCart(store, cartItem) {
      store.commit("ADD_TO_CART", cartItem);
    },
    removeFromCart(store, cartItemId) {
      store.commit("REMOVE_FROM_CART", cartItemId);
    },
    //updatedItem: id i nova kolicina
    updateCartItem(store, updatedCartItem) {
      store.commit("UPDATE_CART_ITEM", updatedCartItem);
    },
  },
  modules: {},
  getters: {
    getCart() {
      return state.cart;
    },
  },
  plugins: [createPersistedState()],
});
