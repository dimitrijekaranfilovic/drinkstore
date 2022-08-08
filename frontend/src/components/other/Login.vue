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
                <v-btn
                  depressed
                  color="primary"
                  type="submit"
                  :loading="loginBtnLoading"
                  :disabled="payload.username === '' || payload.password === ''"
                >
                  Login
                </v-btn>
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
    <toast
      :snackbar="snackbar"
      :text="text"
      @toast-closing="snackbar = false"
    />
  </v-container>
</template>

<script>
import { userService } from "../../services/userService";
import jwtDecode from "jwt-decode";
import Toast from "../other/Toast.vue";
export default {
  components: { Toast },
  data: () => {
    return {
      valuePassword: String,
      snackbar: false,
      text: "",
      loginBtnLoading: false,

      payload: {
        username: "",
        password: "",
      },
    };
  },
  methods: {
    login() {
      this.loginBtnLoading = true;
      userService
        .authenticate(this.payload)
        .then((response) => {
          const token = response.data.jwt;
          this.loginBtnLoading = false;
          const decodedToken = jwtDecode(token);
          const user = {
            username: decodedToken.sub,
            authority: decodedToken.authority,
            token: token,
          };
          this.$store.dispatch("loginUser", user);
          this.$router.push({ name: "Home" });
        })
        .catch((error) => {
          this.loginBtnLoading = false;
          if (error.response) this.text = error.response.data.message;
          else this.text = "An error occurred while logging in.";
          this.snackbar = true;
        });
    },
  },
};
</script>
