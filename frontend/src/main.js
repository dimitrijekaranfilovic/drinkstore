import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import VueToast from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

//filters
import "./filters/capitalizeFilter";
import "./filters/removeUnderscoreFilter";
import store from './store'

Vue.use(VueToast);
Vue.config.productionTip = false;

new Vue({
  router,
  vuetify,
  store,
  render: (h) => h(App)
}).$mount("#app");
