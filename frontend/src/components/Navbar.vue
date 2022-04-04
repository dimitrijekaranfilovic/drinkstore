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
    <!--mozda ovo stavi u neki dropdown, koji ce imati tekst options ako nije ulogovan, a username ako jeste-->

    <v-menu offset-y v-model="menu">
      <template v-slot:activator="{ on, attrs }">
        <v-btn color="primary" depressed v-bind="attrs" v-on="on">
          {{ dropdownText }}
          <v-icon v-if="menu">
            mdi-chevron-up-circle-outline
            <!-- mdi-arrow-up-drop-circle-outline -->
          </v-icon>
          <v-icon v-else>
            mdi-chevron-down-circle-outline
            <!-- mdi-arrow-down-drop-circle-outline -->
          </v-icon>
        </v-btn>
      </template>
      <v-list>
        <v-list-item-group>
          <v-list-item @click="redirect({ name: 'Login' }, '/login')">
            <v-icon color="primary"> mdi-clipboard-account-outline </v-icon>
            <v-list-item-content>
              <v-list-item-title>Login</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item @click="redirect({ name: 'Register' }, '/register')">
            <v-icon color="primary"> mdi-account-plus </v-icon>
            <v-list-item-content>
              <v-list-item-title>Register</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-divider></v-divider>
          <v-list-item>
            <v-icon color="black"> mdi-logout </v-icon>
            <v-list-item-content>
              <v-list-item-title>Log out</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-menu>
  </v-app-bar>
</template>

<script>
export default {
  name: "Navbar",
  data: () => {
    return {
      menu: false,
    };
  },
  methods: {
    redirect(routeObject, disabledPath) {
      if (this.$route.path !== disabledPath) this.$router.push(routeObject);
    },
  },
  computed: {
    dropdownText() {
      return "options";
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
