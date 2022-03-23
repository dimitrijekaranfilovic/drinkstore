<template>
  <v-container>
    <search-filter-drinks :searchParams="searchParams" />
    <v-row dense v-for="(row, rowI) in numRows" :key="rowI">
      <v-col
        v-for="(drink, drinkIndex) in drinks.slice(
          rowI * numCols,
          rowI * numCols + numCols
        )"
        :key="drinkIndex"
      >
        <drink-card :drink="drink" :key="drink.id" />
      </v-col>
    </v-row>
    <div class="text-center">
      <!--TODO:  ovo vidi da se malo ispravi, za sad je samo placeholder -->
      <v-pagination v-model="page" :length="6"></v-pagination>
    </div>
  </v-container>
</template>

<script>
import DrinkCard from "../cards/DrinkCard.vue";
import SearchFilterDrinks from "../SearchFilterDrinks.vue";
export default {
  name: "Index",
  components: { DrinkCard, SearchFilterDrinks },
  data: () => {
    return {
      searchParams: {
        query: "",
        sortCriteria: "name",
        sortDirection: "asc",
      },
      numRows: 2,
      numCols: 5,
      drinks: [
        {
          id: 1,
          name: "Vinjak",
          description: "Opis",
          volume: 0.75,
          volumeMark: "l",
          averageGrade: 4.8,
        },
        {
          id: 2,
          name: "Vinjak",
          description: "Opis",
          volume: 0.75,
          volumeMark: "l",
          averageGrade: 4.8,
        },
        {
          id: 3,
          name: "Vinjak",
          description: "Opis",
          volume: 0.75,
          volumeMark: "l",
          averageGrade: 4.8,
        },
        {
          id: 4,
          name: "Vinjak",
          description: "Opis",
          volume: 0.75,
          volumeMark: "l",
          averageGrade: 4.8,
        },
        {
          id: 5,
          name: "Vinjak",
          description: "Opis",
          volume: 0.75,
          volumeMark: "l",
          averageGrade: 4.8,
        },
      ],
      page: 1,
    };
  },
};
</script>
