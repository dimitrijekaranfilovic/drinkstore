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
    <toast :snackbar="snackbar" :text="text" @toastClosing="snackbar = false" />
  </v-container>
</template>

<script>
import DrinkCard from "../drink/DrinkCard.vue";
import SearchFilterDrinks from "./SearchFilterDrinks.vue";
import DrinkDialog from "../drink/DrinkDialog.vue";
import { drinkService } from "../../services/drinkService";
import Toast from "../other/Toast.vue";

export default {
  name: "Home",
  components: { DrinkCard, SearchFilterDrinks, DrinkDialog, Toast },
  data: () => {
    return {
      snackbar: false,
      text: "",
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
      let drinkPayload = { ...drink };
      delete drinkPayload.image;
      drinkService
        .createDrink(drinkPayload)
        .then((response) => {
          let drinkId = response.data.id;
          drinkService
            .createDrinkImage(drinkId, drink.image)
            .then((_) => {
              this.text = "Drink successfully created.";
              this.snackbar = true;
            })
            .catch((error) => {
              if (error.response) this.text = error.response.data.message;
              else this.text = "An error occurred while creating drink.";
              this.snackbar = true;
            });
        })
        .catch((error) => {
          if (error.response) this.text = error.response.data.message;
          else this.text = "An error occurred while uploading image for drink.";
          this.snackbar = true;
        });
    },
  },
  watch: {
    page(newPage) {
      this.search(newPage);
    },
  },
};
</script>
