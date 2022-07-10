import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const roles = { admin: "ADMIN", user: "USER" };

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("../views/HomeView.vue"),
    meta: {
      authenticated: false,
      authorities: [],
    },
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/LoginView.vue"),
    meta: {
      authenticated: false,
      authorities: [],
    },
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../views/RegisterView.vue"),
    meta: {
      authenticated: false,
      authorities: [],
    },
  },
  {
    path: "/most-sold",
    name: "MostSold",
    component: () => import("../views/MostSoldView.vue"),
    meta: {
      authenticated: false,
      authorities: [],
    },
  },
  {
    path: "/drink/:id",
    name: "Drink",
    component: () => import("../views/DrinkView.vue"),
    meta: {
      authenticated: false,
      authorities: [],
    },
  },
  {
    path: "/cart",
    name: "Cart",
    component: () => import("../views/CartView.vue"),
    meta: {
      authenticated: true,
      authorities: [roles.user],
    },
  },
  {
    path: "/reported-comments",
    name: "ReportedComments",
    component: () => import("../views/ReportedCommentsView.vue"),
    meta: {
      authenticated: true,
      authorities: [roles.admin],
    },
  },
  {
    path: "/purchase-history",
    name: "PurchaseHistory",
    component: () => import("../views/PurchaseHistoryView.vue"),
    meta: {
      authenticated: true,
      authorities: [roles.user],
    },
  },
];
const router = new VueRouter({
  routes,
  mode: "history",
});

export default router;
