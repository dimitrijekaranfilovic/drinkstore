<template>
  <v-container>
    <search-filter-drinks
      :searchParams="searchParams"
      @searchParamsUpdated="search"
    />
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
        ratingFrom: 0,
        ratingTo: 5,
        categories: ["LIQUOR", "WINE", "BEER", "CARBONATED", "NON_CARBONATED"],
      },
      numRows: 1,
      numCols: 5,
      drinks: [
        {
          id: 1,
          name: "Vinjak",
          description: "Opis",
          volume: 0.75,
          volumeMark: "l",
          averageGrade: 4.8,
          category: "LIQUOR",
        },
        {
          id: 2,
          name: "Vranac",
          description: "Opis",
          volume: 0.75,
          volumeMark: "l",
          averageGrade: 4.8,
          category: "WINE",
        },
        {
          id: 3,
          name: "Zajecarsko",
          description: "Opis",
          volume: 0.75,
          volumeMark: "l",
          averageGrade: 4.8,
          category: "BEER",
        },
        {
          id: 4,
          name: "Koka kola",
          description: "Opis",
          volume: 0.75,
          volumeMark: "l",
          averageGrade: 4.8,
          category: "CARBONATED",
        },
        {
          id: 5,
          name: "Multivitamin",
          description: "Opis",
          volume: 0.75,
          volumeMark: "l",
          averageGrade: 4.8,
          category: "NON_CARBONATED",
        },
      ],
      page: 1,
    };
  },
  methods: {
    search() {
      console.log(this.searchParams);
    },
  },
  watch: {
    page(newPage) {
      console.log(newPage);
    },
  },
};
</script>
