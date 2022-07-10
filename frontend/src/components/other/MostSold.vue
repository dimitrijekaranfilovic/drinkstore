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
import { purchaseService } from "../../services/purchaseService";

export default {
  components: { DrinkCard },
  name: "MostSold",
  data: () => {
    return {
      page: 1,
      chip: 3,
      numRows: 1,
      numCols: 5,
      drinks: [],
      period: "year",
    };
  },
  mounted() {
    this.search(this.period);
  },
  methods: {
    search(period) {
      drinkService.getUnfilteredDrinks().then((response) => {
        let ds = response.data;
        //this.drinks = response.data;
        for (const drink of ds) {
          drink.numberOfSales = 0;
        }
        purchaseService.getMostSold(period).then((response) => {
          const mostSold = response.data;
          for (const item of mostSold) {
            const drinkId = item.drink_id;
            const sold = item.sold_items;
            ds.find((drink) => drink.id === drinkId)["numberOfSales"] = sold;
          }

          //console.log(this.drinks);
          ds.sort((d1, d2) => d1.numberOfSales > d2.numberOfSales);
          this.drinks = [...ds];
        });
      });
    },
  },
  watch: {
    chip(newChip) {
      if (newChip == 0) this.period = "day";
      else if (newChip == 1) this.period = "week";
      else if (newChip == 2) this.period = "month";
      else if (newChip == 3) this.period = "year";

      this.search(this.period);
    },
  },
};
</script>
