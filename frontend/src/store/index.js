import Vue from "vue";
import Vuex from "vuex";
import createPersistedState from "vuex-persistedstate";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    cart: [],
    isUserLoggedIn: false,
    user: {
      username: undefined,
      authority: undefined,
    },
  },
  mutations: {
    EMPTY_CART(state) {
      state.cart = [];
    },
    LOGIN_USER(state, user) {
      state.isUserLoggedIn = true;
      state.user.username = user.username;
      state.user.authority = user.authority;
      localStorage.setItem("jwt", user.token);
    },
    LOGOUT_USER(state) {
      state.isUserLoggedIn = false;
      state.user.username = undefined;
      state.user.authority = undefined;
      localStorage.removeItem("jwt");
    },
    ADD_TO_CART(state, cartItem) {
      state.cart.push(cartItem);
    },
    REMOVE_FROM_CART(state, cartItemId) {
      state.cart = state.cart.filter((item) => item.id !== cartItemId);
    },
    UPDATE_CART_ITEM(state, updatedCartItem) {
      let itemToUpdate = state.cart.find(
        (item) => item.id === updatedCartItem.id
      );
      itemToUpdate.amount = updatedCartItem.amount;
    },
    ITEM_EXISTS(state, itemId) {
      let result = state.cart.find((item) => item.id === itemId);
      return result !== undefined;
    },
  },
  actions: {
    emptyCart(store) {
      store.commit("EMPTY_CART");
    },
    loginUser(store, user) {
      store.commit("LOGIN_USER", user);
    },
    logoutUser(store) {
      store.commit("LOGOUT_USER");
    },

    cartItemExists(store, itemId) {
      return store.commit("ITEM_EXISTS", itemId);
    },

    //item: id, naziv, jedinicna cijena i kolicina
    addCartItem(store, cartItem) {
      store.commit("ADD_TO_CART", cartItem);
    },
    removeCartItem(store, cartItemId) {
      store.commit("REMOVE_FROM_CART", cartItemId);
    },
    //updatedItem: id i nova kolicina
    updateCartItem(store, updatedCartItem) {
      store.commit("UPDATE_CART_ITEM", updatedCartItem);
    },
  },
  modules: {},
  getters: {
    getCart(state) {
      return state.cart;
    },
  },
  plugins: [createPersistedState()],
});
