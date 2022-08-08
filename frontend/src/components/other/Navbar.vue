<template>
  <v-app-bar app color="primary" dark>
    <div
      class="d-flex align-center home-icon"
      @click="redirect({ name: 'Home' })"
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
      <v-btn @click="redirect({ name: 'Home' })" color="primary" depressed>
        <v-icon> mdi-home </v-icon>
        Home
      </v-btn>
      <v-btn @click="redirect({ name: 'MostSold' })" color="primary" depressed>
        <v-icon> mdi-chart-areaspline </v-icon>
        Most sold
      </v-btn>
    </div>
    <v-spacer />

    <v-menu offset-y v-model="menu">
      <template v-slot:activator="{ on, attrs }">
        <v-btn color="primary" depressed v-bind="attrs" v-on="on">
          {{ dropdownText }}
          <v-icon v-if="menu"> mdi-chevron-up-circle-outline </v-icon>
          <v-icon v-else> mdi-chevron-down-circle-outline </v-icon>
        </v-btn>
      </template>
      <v-list>
        <v-list-item-group>
          <v-list-item
            @click="redirect({ name: 'Login' })"
            v-if="userAuthority === undefined"
          >
            <v-icon color="primary"> mdi-clipboard-account-outline </v-icon>
            <v-list-item-content>
              <v-list-item-title>Login</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item
            @click="redirect({ name: 'Register' })"
            v-if="userAuthority === undefined"
          >
            <v-icon color="primary"> mdi-account-plus </v-icon>
            <v-list-item-content>
              <v-list-item-title>Register</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item
            @click="redirect({ name: 'ReportedComments' })"
            v-if="userAuthority === 'ADMIN'"
          >
            <v-icon color="primary">mdi-alert</v-icon>
            <v-list-item-content>
              <v-list-item-title> Reported comments </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item
            @click="redirect({ name: 'Cart' })"
            v-if="userAuthority === 'USER'"
          >
            <v-icon color="primary">mdi-cart</v-icon>
            <v-list-item-content>
              <v-list-item-title>
                Cart
                <v-badge
                  :content="$store.state.cart.length"
                  :value="$store.state.cart.length"
                >
                </v-badge>
              </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item
            v-if="userAuthority === 'USER'"
            @click="redirect({ name: 'PurchaseHistory' })"
          >
            <v-icon color="primary">mdi-clipboard-text-clock</v-icon>
            <v-list-item-content>
              <v-list-item-title> Purchase history </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-divider v-if="userAuthority !== undefined"></v-divider>
          <v-list-item v-if="userAuthority !== undefined" @click="logout">
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
    redirect(routeObject) {
      this.$router.push(routeObject).catch(() => {});
    },
    logout() {
      this.$store.dispatch("logoutUser");
      this.$router.push({ name: "Home" }).catch(() => {});
    },
  },
  computed: {
    dropdownText() {
      if (!this.$store.state.isUserLoggedIn) return "options";
      else return this.$store.state.user.username;
    },
    userAuthority() {
      return this.$store.state.user.authority;
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
