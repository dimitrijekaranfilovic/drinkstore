<template>
  <v-card
    class="mx-auto my-12 drink-card"
    width="374"
    max-width="374"
    @click="redirectToDrink"
    elevation="4"
  >
    <v-img
      height="250"
      src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
    ></v-img>
    <v-card-title
      >{{ drink.name }}

      <v-chip class="ma-2" color="secondary" outlined>
        {{ drink.volume }}{{ drink.volumeMark }}
      </v-chip>
    </v-card-title>

    <v-card-actions class="pa-4">
      <span class="avg-grade">Average rating:</span>
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
      <v-btn color="primary">Add to cart</v-btn>
      <v-btn color="primary lighten-2" text @click="redirectToDrink">
        View drink
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: "DrinkCard",
  props: ["drink"],
  methods: {
    redirectToDrink() {
      console.log(`REDIRECT TO DRINK ${this.drink.id}`);
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
</style>
