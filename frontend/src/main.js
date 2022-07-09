import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import VueToast from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

//filters
import "./filters/capitalizeFilter";
import "./filters/removeUnderscoreFilter";
import "./filters/formatDateFilter";
import store from "./store";

Vue.use(VueToast);
Vue.config.productionTip = false;

router.beforeEach((to, _, next) => {
  const { authenticated, authorities } = to.meta;
  if (authenticated) {
    let userAuthority = store.state.user.authority;
    if (authorities.some((element) => element === userAuthority)) {
      next();
    } else {
      next({ name: "Login" });
    }
  } else {
    next();
  }
});

new Vue({
  router,
  vuetify,
  store,
  render: (h) => h(App),
}).$mount("#app");
