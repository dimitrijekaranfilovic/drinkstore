import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("../views/HomeView.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/LoginView.vue"),
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../views/RegisterView.vue"),
  },
  {
    path: "/most-sold",
    name: "MostSold",
    component: () => import("../views/MostSoldView.vue"),
  },
];
//TODO: dodati redirekt sa Home stranice na login ili index u zavisnosti od toga da li je korisnik ulogovan
const router = new VueRouter({
  routes,
  mode: "history",
});

export default router;
