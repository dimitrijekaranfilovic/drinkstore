<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-col cols="12">
        <v-flex class="text-center">
          <h1>Register</h1>
        </v-flex>
      </v-col>
    </v-row>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="12" md="8" lg="6" xl="4">
        <v-stepper v-model="currentStep">
          <v-stepper-header>
            <v-stepper-step :complete="currentStep > 1" step="1">
              Credentials
            </v-stepper-step>

            <v-divider />

            <v-stepper-step :complete="currentStep > 2" step="2">
              Address
            </v-stepper-step>
            <v-divider />

            <v-stepper-step :complete="currentStep > 3" :step="3">
              Finish
            </v-stepper-step>
          </v-stepper-header>

          <v-stepper-items>
            <v-stepper-content step="1">
              <v-card>
                <v-card-title> Credentials </v-card-title>
                <v-card-text>
                  <v-form
                    v-model="validFirstStep"
                    @submit.prevent="moveForward"
                  >
                    <v-text-field
                      :rules="firstNameRules"
                      v-model="payload.firstName"
                      label="First name"
                      :counter="40"
                      maxlength="40"
                    />
                    <v-text-field
                      :rules="lastNameRules"
                      v-model="payload.lastName"
                      label="Last name"
                      :counter="40"
                      maxlength="40"
                    />
                    <v-text-field
                      v-model="payload.username"
                      :rules="usernameRules"
                      :counter="30"
                      maxlength="30"
                      label="Username"
                      required
                    />
                    <v-text-field
                      v-model="payload.password"
                      :rules="passwordRules"
                      label="Password"
                      :type="valuePassword ? 'password' : 'text'"
                      @click:append="() => (valuePassword = !valuePassword)"
                      :append-icon="valuePassword ? 'mdi-eye-off' : 'mdi-eye'"
                      hint="Remember your password, because if you forget it, you cannot recover it!"
                      persistent-hint
                      required
                    />
                    <v-text-field
                      v-model="confirmedPassword"
                      :rules="confirmPasswordRules"
                      label="Confirm password"
                      :type="valueConfirmPassword ? 'password' : 'text'"
                      @click:append="
                        () => (valueConfirmPassword = !valueConfirmPassword)
                      "
                      :append-icon="
                        valueConfirmPassword ? 'mdi-eye-off' : 'mdi-eye'
                      "
                      required
                    />

                    <v-row>
                      <v-col md="12" class="text-right">
                        <v-btn
                          depressed
                          color="primary"
                          type="submit"
                          :disabled="!validFirstStep"
                        >
                          Next
                        </v-btn>
                      </v-col>
                    </v-row>
                    <p>
                      Already have an account? Log in
                      <router-link :to="{ name: 'Login' }">here.</router-link>
                    </p>
                  </v-form>
                </v-card-text>
              </v-card>
            </v-stepper-content>

            <v-stepper-content step="2">
              <v-card>
                <v-card-title> Address for delivery </v-card-title>
                <v-card-text>
                  <v-form
                    @submit.prevent="moveForward"
                    v-model="validSecondStep"
                  >
                    <v-row>
                      <v-col cols="8">
                        <v-text-field
                          label="Street"
                          v-model="payload.address.street"
                          :rules="stringRules"
                        />
                      </v-col>
                      <v-col cols="4">
                        <v-text-field
                          label="Number"
                          v-model="payload.address.number"
                          :rules="stringRules"
                        />
                      </v-col>
                    </v-row>

                    <v-row>
                      <v-col cols="8">
                        <v-text-field
                          label="City"
                          v-model="payload.address.city"
                          :rules="stringRules"
                        />
                      </v-col>
                      <v-col cols="4">
                        <v-text-field
                          label="Zip code"
                          v-model="payload.address.zipCode"
                          :rules="stringRules"
                        />
                      </v-col>
                    </v-row>
                    <v-row>
                      <v-col cols="12">
                        <v-text-field
                          label="Country"
                          v-model="payload.address.country"
                          :rules="stringRules"
                        />
                      </v-col>
                    </v-row>

                    <v-row>
                      <v-col cols="8" sm="8" md="8" lg="8" xl="8">
                        <v-btn depressed @click="moveBackwards">Back</v-btn>
                      </v-col>

                      <v-col
                        cols="4"
                        sm="4"
                        md="4"
                        lg="4"
                        xl="4"
                        class="text-right"
                      >
                        <v-btn
                          depressed
                          color="primary"
                          type="submit"
                          class="next-btn"
                          :disabled="!validSecondStep"
                          >Next</v-btn
                        >
                      </v-col>
                    </v-row>
                  </v-form>
                </v-card-text>
              </v-card>
            </v-stepper-content>

            <v-stepper-content step="3">
              <v-card elevation="2" outlined>
                <v-card-title>Finishing steps</v-card-title>
                <v-card-text>
                  Please, make sure that you've entered all your information
                  correctly.
                  <br />
                  If you are not sure, simply go back and check. When you're
                  ready, simply click the 'Register' button.
                </v-card-text>
              </v-card>
              <v-row class="button-row">
                <v-col cols="6" sm="6" md="6" lg="6" xl="6">
                  <v-btn depressed @click="moveBackwards">Back</v-btn>
                </v-col>
                <v-col cols="6" sm="6" md="6" lg="6" xl="6" class="text-right">
                  <v-btn
                    depressed
                    color="primary"
                    @click="register"
                    :disabled="!validFirstStep || !validSecondStep"
                    >Register</v-btn
                  >
                </v-col>
              </v-row>
            </v-stepper-content>
          </v-stepper-items>
        </v-stepper>
      </v-col>
    </v-row>
    <toast :snackbar="snackbar" :text="text" @toastClosing="snackbar = false" />
  </v-container>
</template>

<script>
import { passwordRules } from "../../util/passwordRules";
import { usernameRules } from "../../util/usernameRules";
import { firstNameRules } from "../../util/firstNameRules";
import { lastNameRules } from "../../util/lastNameRules";
import { userService } from "../../services/userService";
import Toast from "../other/Toast.vue";

export default {
  name: "Register",
  components: { Toast },
  data() {
    return {
      snackbar: false,
      text: "",
      currentStep: 1,
      usernameRules: usernameRules,
      passwordRules: passwordRules,
      firstNameRules: firstNameRules,
      lastNameRules: lastNameRules,
      stringRules: [(s) => !!s || "This field is required!"],
      validFirstStep: false,
      validSecondStep: false,
      valuePassword: String,
      valueConfirmPassword: String,
      confirmedPassword: "",
      payload: {
        username: "",
        password: "",
        firstName: "",
        lastName: "",
        address: {
          street: "",
          number: "",
          city: "",
          zipCode: "",
          country: "",
        },
      },
    };
  },
  computed: {
    confirmPasswordRules() {
      return [
        ...passwordRules,
        (pass) => this.payload.password === pass || "Passwords don't match.",
      ];
    },
  },

  methods: {
    moveForward() {
      this.currentStep += 1;
    },
    moveBackwards() {
      this.currentStep -= 1;
    },
    register() {
      userService
        .register(this.payload)
        .then((_) => {
          this.$router.push({ name: "Login" });
        })
        .catch((error) => {
          if (error.response) this.text = error.response.data.message;
          else this.text = "An error occurred while registering user.";
          this.snackbar = true;
        });
      //console.log(this.payload);
    },
  },
};
</script>

<style scoped>
.container {
  margin-top: 2%;
}
.next-btn {
  margin-right: 1%;
}
.button-row {
  margin-top: 1%;
}
</style>
