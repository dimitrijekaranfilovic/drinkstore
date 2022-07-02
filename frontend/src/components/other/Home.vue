<template>
  <v-container>
    <search-filter-drinks
      :searchParams="searchParams"
      @searchParamsUpdated="search"
    />
    <!--dodaj v-if za to je li korisnik ulogovan i je li admin-->
    <v-row>
      <v-col cols="12">
        <v-btn color="primary" depressed @click="drinkDialog = true">
          <v-icon> mdi-bookmark-plus </v-icon>
          Add drink
        </v-btn>
      </v-col>
    </v-row>
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
    <drink-dialog
      @dialog-closing="drinkDialog = false"
      :dialog="drinkDialog"
      @drink-added="addDrink($event)"
    />
  </v-container>
</template>

<script>
import DrinkCard from "../drink/DrinkCard.vue";
import SearchFilterDrinks from "./SearchFilterDrinks.vue";
import DrinkDialog from "../drink/DrinkDialog.vue";

export default {
  name: "Home",
  components: { DrinkCard, SearchFilterDrinks, DrinkDialog },
  data: () => {
    return {
      drinkDialog: false,
      searchParams: {
        query: "",
        sortCriteria: "name",
        sortDirection: "asc",
        ratingFrom: 0,
        ratingTo: 5,
        volumeLabels: ["l", "ml", "fl.oz."],
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
          volumeLabel: "l",
          averageGrade: 4.8,
          category: "LIQUOR",
          price: 3200,
        },
        {
          id: 2,
          name: "Vranac",
          description: "Opis",
          volume: 0.75,
          volumeLabel: "l",
          averageGrade: 4.8,
          category: "WINE",
          price: 3100,
        },
        {
          id: 3,
          name: "Zajecarsko",
          description: "Opis",
          volume: 0.75,
          volumeLabel: "l",
          averageGrade: 4.8,
          category: "BEER",
          price: 300,
        },
        {
          id: 4,
          name: "Koka kola",
          description: "Opis",
          volume: 0.75,
          volumeLabel: "l",
          averageGrade: 4.8,
          category: "CARBONATED",
          price: 3000,
        },
        {
          id: 5,
          name: "Multivitamin",
          description: "Opis",
          volume: 0.75,
          volumeLabel: "l",
          averageGrade: 4.8,
          category: "NON_CARBONATED",
          price: 3000,
        },
      ],
      page: 1,
    };
  },
  methods: {
    search() {
      console.log(this.searchParams);
    },
    closeDialog() {
      this.drinkDialog = false;
    },
    addDrink(drink) {
      console.log("Add drink ", drink);
    },
  },
  watch: {
    page(newPage) {
      console.log(newPage);
    },
  },
};
</script>