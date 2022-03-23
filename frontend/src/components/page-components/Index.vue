<template>
  <v-container>
    <!-- <main-page-search-sort-bar
      :searchParams="searchParams"
      @searchEvent="search"
    /> -->
    <!-- <v-row dense align="center">
      <v-col v-for="drink in drinks.slice(0, 3)" :key="drink.id">
        <drink-card :drink="drink" />
      </v-col>
    </v-row> -->
    <!-- <b-row class="row-spacing" v-for="(row, rowI) in numRows" :key="rowI">
        <b-col
          v-bind:lg="cardWidth"
          sm="12"
          v-for="(s, index) in stocksPage.content.slice(
            rowI * numCols,
            rowI * numCols + numCols
          )"
          :key="index"
        >
          <MedicineStockCard
            v-bind:stock="s"
            :image="getMedicineImage(s)"
            :key="s.id"
          />
        </b-col>
      </b-row> -->
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
export default {
  name: "Index",
  components: { DrinkCard },
  data: () => {
    return {
      numRows: 2,
      numCols: 4,
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
