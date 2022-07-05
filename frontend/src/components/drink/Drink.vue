<template>
  <v-container>
    <v-row align="center" justify="center" v-if="drink">
      <v-col cols="2">
        <v-img :src="imagePath" height="400px"></v-img>
      </v-col>
      <v-col cols="10">
        <v-card class="mx-auto">
          <v-card-title>
            {{ drink.name }}

            <v-chip class="ma-2" color="secondary" outlined>
              {{ drink.volume }}{{ drink.volumeLabel }}
            </v-chip>
            <v-chip :color="drinkCategoryChipColor" text-color="white">
              {{ drink.category | capitalize | removeUnderscore }}
            </v-chip>
            <!--TODO: dodaj da se prikazuje ocjena korisnika, ako je ostavio, ako nije, samo neko dugme grade ili tako nesto-->

            <v-col align="right">
              <v-btn
                color="red"
                rounded
                dark
                v-if="itemAlreadyinCart && userAuthority === 'USER'"
                @click="removeFromCart()"
              >
                <v-icon> mdi-delete </v-icon>
                Remove from cart
              </v-btn>
              <v-btn
                color="primary"
                rounded
                dark
                v-else-if="!itemAlreadyinCart && userAuthority === 'USER'"
                @click="addToCart()"
              >
                <v-icon> mdi-cart </v-icon>
                Add to cart
              </v-btn>
              <v-btn
                color="#ea9b09"
                dark
                rounded
                @click="drinkDialog = true"
                v-if="userAuthority === 'ADMIN'"
              >
                <v-icon color="white"> mdi-pencil </v-icon>
                Edit
              </v-btn>
              <v-btn
                color="red"
                dark
                rounded
                @click="confirmDialog = true"
                v-if="userAuthority === 'ADMIN'"
              >
                <v-icon color="white"> mdi-delete-forever </v-icon>
                Delete
              </v-btn>
            </v-col>
          </v-card-title>
          <v-chip class="ma-2" color="red" text-color="white">
            <strong>{{ drink.price }} RSD</strong>
          </v-chip>
          <v-card-actions>
            <span class="avg-grade">Average grade:</span>
            <v-spacer></v-spacer>
            <span class="text--lighten-2 text-caption mr-2"> (1.1) </span>
            <v-rating
              :value="1.2"
              half-increments
              hover
              length="5"
              readonly
              color="yellow darken-3"
              background-color="yellow darken-3"
              size="18"
            ></v-rating>
          </v-card-actions>
          <v-divider />

          <v-card-text>
            {{ drink.description }}
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-divider />
    <v-card>
      <v-card-title> Comments </v-card-title>
      <v-card-text>
        <v-container>
          <v-row v-if="userAuthority === 'USER'">
            <v-col cols="12" md="10">
              <v-text-field label="Enter comment" />
            </v-col>
            <v-col cols="12" md="2">
              <v-btn color="primary" class="send-btn">
                <v-icon color="white"> mdi-send </v-icon>
                Send
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
        <comments :comments="comments" v-if="comments.length > 0" />
      </v-card-text>
    </v-card>
    <drink-dialog
      :dialog="drinkDialog"
      :drinkToEdit="drink"
      @dialog-closing="drinkDialog = false"
      @drink-edited="editDrink($event)"
    />
    <confirm-dialog
      :confirmDialog="confirmDialog"
      :text="dialogDeleteText"
      :title="dialogDeleteTitle"
      @dialog-closing="closeConfirmDialog($event)"
    />
    <toast
      :snackbar="snackbar"
      :text="text"
      @toast-closing="snackbar = false"
    />
  </v-container>
</template>

<script>
import { categories } from "../../util/categories";
import Comments from "../comments/Comments.vue";
import DrinkDialog from "./DrinkDialog.vue";
import { drinkService } from "../../services/drinkService";
import ConfirmDialog from "../other/ConfirmDialog.vue";
import Toast from "../other/Toast.vue";

export default {
  name: "Drink",
  components: {
    Comments,
    DrinkDialog,
    ConfirmDialog,
    Toast,
  },
  data: () => {
    return {
      snackbar: false,
      text: "",
      confirmDialog: false,
      drinkDialog: false,
      drink: null,
      comments: [],
      dialogDeleteTitle: "Drink deletion",
    };
  },
  mounted() {
    drinkService.getSingleDrink(this.$route.params.id).then((response) => {
      this.drink = response.data;
    });
  },
  methods: {
    closeConfirmDialog(event) {
      if (event.answer === "OK") {
        this.deleteDrink();
        this.confirmDialog = false;
      } else this.confirmDialog = false;
    },
    editDrink(drink) {
      let drinkUpdatePayload = { ...drink };
      drinkUpdatePayload.price += "";
      drinkUpdatePayload.volume += "";
      delete drinkUpdatePayload.image;
      drinkService
        .updateDrink(this.drink.id, drinkUpdatePayload)
        .then((_) => {
          if (drink.image !== null) {
            drinkService
              .createDrinkImage(drink.id, drink.image)
              .then((_) => {
                this.$router.go(0);
              })
              .catch((error) => {
                if (error.response) this.text = error.response.data.message;
                else
                  this.text = "An error occurred while uploading drink image.";
                this.snackbar = true;
              });
          } else {
            this.$router.go(0);
          }
        })
        .catch((error) => {
          if (error.response) this.text = error.response.data.message;
          else this.text = "An error occurred while updating drink.";
          this.snackbar = true;
        });
    },
    deleteDrink() {
      drinkService.deleteDrink(this.$route.params.id).then((_) => {
        this.$router.push({ name: "Home" });
      });
    },
    addToCart() {
      let drinkId = this.drink.id;
      let drinkName = this.drink.name;
      let drinkPrice = this.drink.price;
      this.$store.dispatch("addCartItem", {
        id: drinkId,
        name: drinkName,
        price: drinkPrice,
        amount: 1,
      });
    },
    removeFromCart() {
      let drinkId = this.drink.id;
      this.$store.dispatch("removeCartItem", drinkId);
    },
  },
  computed: {
    dialogDeleteText() {
      if (this.drink)
        return `Are you sure you want to delete ${this.drink.name}?`;
      else "Are you sure you want to delete this drink?";
    },
    imagePath() {
      return `${process.env.VUE_APP_DRINK_SERVICE_BASE_PATH}/${this.drink.imagePath}`;
    },
    userAuthority() {
      return this.$store.state.user.authority;
    },
    drinkCategoryChipColor() {
      return categories.find((c) => c.value === this.drink.category).color;
    },
    itemAlreadyinCart() {
      let result =
        this.$store.state.cart.find((item) => item.id === this.drink.id) !==
        undefined;
      return result;
    },
  },
};
</script>

<style scoped>
.avg-grade {
  color: rgba(0, 0, 0, 0.6);
  font-size: 14px;
}

.no-bottom-margin {
  margin-bottom: 5px !important;
}

.send-btn {
  margin-top: 15px;
}
</style>
