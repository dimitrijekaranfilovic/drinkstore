<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-col cols="12">
        <v-flex class="text-center">
          <h1>Login</h1>
        </v-flex>
      </v-col>
    </v-row>
    <form @submit.prevent="login">
      <v-row align="center" justify="center">
        <v-col cols="12" sm="8" md="8" lg="6" xl="4">
          <v-card elevation="4">
            <v-row align="center" justify="center">
              <v-col cols="8" sm="8" md="8" lg="10" xl="10">
                <v-text-field
                  label="Username"
                  required
                  v-model="payload.username"
                >
                </v-text-field>
              </v-col>
            </v-row>
            <v-row align="center" justify="center">
              <v-col cols="8" sm="8" md="8" lg="10" xl="10">
                <v-text-field
                  label="Password"
                  :type="valuePassword ? 'password' : 'text'"
                  @click:append="() => (valuePassword = !valuePassword)"
                  :append-icon="valuePassword ? 'mdi-eye-off' : 'mdi-eye'"
                  required
                  v-model="payload.password"
                >
                </v-text-field>
              </v-col>
            </v-row>

            <v-row align="center" justify="center">
              <v-col cols="8" sm="8" md="8" lg="10" xl="10">
                <v-btn depressed color="primary" type="submit"> Login </v-btn>
                <!-- <v-checkbox label="Remember me" v-model="rememberMe"></v-checkbox> -->
                <p>
                  Don't have an account? Register
                  <router-link :to="{ name: 'Register' }">here.</router-link>
                </p>
              </v-col>
            </v-row>
          </v-card>
        </v-col>
      </v-row>
    </form>
  </v-container>
</template>

<script>
export default {
  data: () => {
    return {
      valuePassword: String,
      payload: {
        rememberMe: false,
        username: "",
        password: "",
      },
    };
  },
  methods: {
    login() {
      console.log("LOGIN");
    },
  },
};
</script>
