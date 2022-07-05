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
    <div class="text-center">
      <!--TODO:  ovo vidi da se malo ispravi, za sad je samo placeholder -->
      <v-pagination v-model="page" :length="6"></v-pagination>
    </div>
  </v-container>
</template>

<script>
import DrinkCard from "../drink/DrinkCard.vue";

export default {
  components: { DrinkCard },
  name: "MostSold",
  data: () => {
    return {
      page: 1,
      chip: 0,
      numRows: 1,
      numCols: 5,
      drinks: [
        {
          id: 1,
          name: "Vinjak",
          description: "Opis",
          volume: 0.75,
          volumeLabel: "l",
          averageGrade: 4.8,
          category: "LIQUOR",
          numberOfSales: 100,
          price: 3000,
        },
        {
          id: 2,
          name: "Vranac",
          description: "Opis",
          volume: 0.75,
          volumeLabel: "l",
          averageGrade: 4.8,
          category: "WINE",
          numberOfSales: 99,
          price: 3000,
        },
        {
          id: 3,
          name: "Zajecarsko",
          description: "Opis",
          volume: 0.75,
          volumeLabel: "l",
          averageGrade: 4.8,
          category: "BEER",
          numberOfSales: 98,
          price: 3000,
        },
      ],
    };
  },
};
</script>
