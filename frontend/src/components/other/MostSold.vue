<template>
  <v-container>
    <v-row>
      <v-card>
        <v-card-title>Most sold</v-card-title>
        <v-card-text>
          <v-chip-group active-class="primary--text" mandatory v-model="chip">
            <v-chip> Last 24 hours </v-chip>
            <v-chip> Last week </v-chip>
            <v-chip> Last month </v-chip>
            <v-chip> Last year </v-chip>
          </v-chip-group>
        </v-card-text>
      </v-card>
    </v-row>
    <v-row dense v-for="(row, rowI) in numRows" :key="rowI">
      <v-col
        v-for="(drink, drinkIndex) in drinks.slice(
          rowI * numCols,
          rowI * numCols + numCols
        )"
        :key="drinkIndex"
      >
        <drink-card :drink="drink" :key="drink.id" :index="drinkIndex" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import DrinkCard from "../drink/DrinkCard.vue";

import { drinkService } from "../../services/drinkService";

export default {
  components: { DrinkCard },
  name: "MostSold",
  data: () => {
    return {
      page: 1,
      chip: 0,
      numRows: 1,
      numCols: 5,
      drinks: [],
    };
  },
  mounted() {
    drinkService.getUnfilteredDrinks().then((response) => {
      this.drinks = response.data;
    });
  },
  watch: {
    chip(newChip) {
      if (newChip == 0) console.log("day");
      else if (newChip == 1) console.log("week");
      else if (newChip == 2) console.log("month");
      else if (newChip == 3) console.log("year");
    },
  },
};
</script>
