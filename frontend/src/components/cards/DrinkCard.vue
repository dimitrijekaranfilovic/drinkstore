<template>
  <v-card
    class="mx-auto my-12 drink-card"
    width="300"
    max-width="300"
    elevation="4"
  >
    <v-img
      height="275"
      src="https://upload.wikimedia.org/wikipedia/commons/f/f9/Rubinov_vinjak.jpg"
    ></v-img>
    <v-card-title
      ><span v-if="index !== undefined">{{ index + 1 }}. </span>
      {{ drink.name }}
      <!--TODO: dodaj cijenu-->
      <v-chip class="ma-2" color="secondary" outlined>
        {{ drink.volume }}{{ drink.volumeLabel }}
      </v-chip>
      <v-chip :color="drinkCategoryChipColor" text-color="white">
        {{ drink.category | capitalize | removeUnderscore }}
      </v-chip>
    </v-card-title>

    <v-chip
      v-if="index !== undefined"
      color="green"
      text-color="white"
      class="ma-2"
    >
      Sales: <strong> {{ drink.numberOfSales }}</strong>
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
    <v-card-actions>
      <!--TODO: dodaj v-if za admina-->
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn color="primary" dark v-bind="attrs" v-on="on" rounded>
            <v-icon> mdi-cart </v-icon>
          </v-btn>
        </template>
        <span>Add to cart</span>
      </v-tooltip>

      <v-btn color="primary lighten-2" text @click="redirectToDrink">
        Details
      </v-btn>
      <!--TODO: dodaj v-if-->
    </v-card-actions>
  </v-card>
</template>

<script>
import { categories } from "../../util/categories";

export default {
  name: "DrinkCard",
  props: ["drink", "index"],
  methods: {
    redirectToDrink() {
      this.$router.push({ name: "Drink", params: { id: this.drink.id } });
    },
  },
  computed: {
    drinkCategoryChipColor() {
      return categories.find((c) => c.value === this.drink.category).color;
    },
  },
};
</script>

<style scoped>
.drink-card:hover {
  box-shadow: #0000005e 0 12px 26px !important;
  transform: translate(0px, -10px);
}
.drink-card {
  transition: 0.5s ease all;
  border-radius: 1.5em !important;
}

.avg-grade {
  color: rgba(0, 0, 0, 0.6);
  font-size: 14px;
}

.no-bottom-margin {
  margin-bottom: 5px !important;
}
</style>
