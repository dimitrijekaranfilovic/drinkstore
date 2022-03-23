import Vue from "vue";

Vue.filter("removeUnderscore", (param) => {
  return param.replaceAll("_", " ");
});
