<template>
  <v-app-bar app color="primary" dark>
    <div
      class="d-flex align-center home-icon"
      @click="redirect({ name: 'Home' }, '/')"
    >
      <v-icon> mdi-bottle-wine </v-icon>

      <span
        class="shrink mt-1 hidden-sm-and-down app-name"
        contain
        min-width="100"
      >
        drinkstore
      </span>
    </div>
    <v-spacer />
    <div>
      <v-btn @click="redirect({ name: 'Home' }, '/')" color="primary" depressed>
        <v-icon> mdi-home </v-icon>
        Home
      </v-btn>
      <v-btn
        @click="redirect({ name: 'MostSold' }, '/most-sold')"
        color="primary"
        depressed
      >
        <v-icon> mdi-chart-areaspline </v-icon>
        Most sold
      </v-btn>
    </div>
    <v-spacer />
    <!--dodaj v-if-->
    <v-btn color="primary" depressed>
      Log Out
      <v-icon> mdi-logout </v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
export default {
  name: "Navbar",
  methods: {
    redirect(routeObject, disabledPath) {
      if (this.$route.path !== disabledPath) this.$router.push(routeObject);
    },
  },
};
</script>

<style>
.home-icon {
  cursor: pointer;
}
.app-name {
  font-family: "Montserrat", sans-serif;
}
</style>
