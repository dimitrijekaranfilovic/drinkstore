<template>
  <v-container>
    <search-filter-drinks
      :searchParams="searchParams"
      @searchParamsUpdated="search"
    />
    <v-row v-if="userAuthority === 'ADMIN'">
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
      <v-pagination v-model="page" :length="totalPages"></v-pagination>
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
import { drinkService } from "../../services/drinkService";

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
        gradeFrom: 0,
        gradeTo: 5,
        volumeLabels: ["l", "ml", "fl.oz."],
        categories: ["LIQUOR", "WINE", "BEER", "CARBONATED", "NON_CARBONATED"],
      },
      numRows: 1,
      numCols: 5,
      drinks: [],
      page: 1,
      size: 4,
      totalPages: 1,
    };
  },
  mounted() {
    this.search(this.page);
  },
  computed: {
    userAuthority() {
      return this.$store.state.user.authority;
    },
  },
  methods: {
    search(pageNumber) {
      const searchPayload = {
        ...this.searchParams,
        page: pageNumber - 1,
        size: this.size,
      };

      drinkService.getDrinks(searchPayload).then((response) => {
        const pageObj = response.data;
        this.totalPages = pageObj.totalPages;
        this.drinks = pageObj.drinks;
      });
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
      this.search(newPage);
    },
  },
};
</script>
