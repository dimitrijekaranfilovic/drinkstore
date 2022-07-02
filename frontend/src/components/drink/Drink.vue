<template>
  <v-container>
    <v-card class="mx-auto" v-if="drink">
      <v-img
        src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
        height="200px"
      ></v-img>
      <v-card-title>
        {{ drink.name }}

        <v-chip class="ma-2" color="secondary" outlined>
          {{ drink.volume }}{{ drink.volumeLabel }}
        </v-chip>
        <v-chip :color="drinkCategoryChipColor" text-color="white">
          {{ drink.category | capitalize | removeUnderscore }}
        </v-chip>
        <!-- TODO: za ovo dodaj v-if-->
        <!-- TODO: mozda dodaj da admin moze da oznaci da vise nije na stanju-->

        <v-col align="right">
          <v-btn
            color="red"
            rounded
            dark
            v-if="itemAlreadyinCart"
            @click="removeFromCart()"
          >
            <v-icon> mdi-delete </v-icon>
            Remove from cart
          </v-btn>
          <v-btn color="primary" rounded dark v-else @click="addToCart()">
            <v-icon> mdi-cart </v-icon>
            Add to cart
          </v-btn>
          <v-btn color="#ea9b09" dark rounded @click="drinkDialog = true">
            <v-icon color="white"> mdi-pencil </v-icon>
            Edit
          </v-btn>
          <v-btn color="red" dark rounded @click="deleteDrink()">
            <v-icon color="white"> mdi-delete-forever </v-icon>
            Delete
          </v-btn>
        </v-col>
      </v-card-title>
      <v-chip class="ma-2" color="red" text-color="white">
        <strong>{{ drink.price }} RSD</strong>
      </v-chip>
      <v-card-actions>
        <span class="avg-grade">Rating:</span>
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
    <v-divider />
    <v-card>
      <v-card-title> Comments </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
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
  </v-container>
</template>

<script>
import { categories } from "../../util/categories";
import Comments from "../comments/Comments.vue";
import DrinkDialog from "./DrinkDialog.vue";

export default {
  name: "Drink",
  components: {
    Comments,
    DrinkDialog,
  },
  data: () => {
    //TODO: ovaj drink je samo placeholder
    return {
      drinkDialog: false,
      drink: {
        id: 1,
        name: "Vinjak",
        description: "Opis",
        volume: 0.75,
        volumeLabel: "l",
        averageGrade: 4.8,
        category: "BEER",
        price: 2000,
      },
      comments: [
        {
          id: 1,
          user: "user 1",
          content: "AAA",
          children: [
            {
              id: 2,
              user: "user 2",
              content: "BBB",
              children: [
                {
                  id: 3,
                  user: "user 3",

                  content: "CCC",
                },
              ],
            },
            {
              id: 4,
              user: "user 4",
              content: "DDD",
            },
          ],
        },
      ],
    };
  },
  methods: {
    editDrink(drink) {
      console.log("Edit drink ", drink);
    },
    deleteDrink() {
      console.log("delete drink");
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
  mounted() {
    //TODO: ovdje dobavi dato pice na osnovu id-ija
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
