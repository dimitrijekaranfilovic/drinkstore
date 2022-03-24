<template>
  <v-app-bar app color="primary" dark>
    <div
      class="d-flex align-center home-icon"
      @click="$router.push({ name: 'Home' })"
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
  </v-app-bar>
</template>

<script>
export default {
  name: "Navbar",
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
