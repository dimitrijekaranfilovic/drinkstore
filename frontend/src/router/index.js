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
  {
    path: "/drink/:id",
    name: "Drink",
    component: () => import("../views/DrinkView.vue"),
  },
  {
    path: "/cart",
    name: "Cart",
    component: () => import("../views/CartView.vue"),
  },
  {
    path: "/reported-comments",
    name: "ReportedComments",
    component: () => import("../views/ReportedCommentsView.vue"),
  },
];
const router = new VueRouter({
  routes,
  mode: "history",
});

export default router;
