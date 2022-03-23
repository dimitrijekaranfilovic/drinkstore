<template>
  <v-container>
    <v-row class="search-bar">
      <v-col cols="8">
        <v-text-field
          label="Enter keyword"
          v-model="searchParams.query"
        ></v-text-field>
      </v-col>
      <v-col cols="2">
        <v-btn color="primary" class="search-btn" @click="emitSearchEvent()"
          >Search</v-btn
        >
      </v-col>
      <v-col cols="2" align="right">
        <v-btn @click="drawer = !drawer" class="search-btn">
          <v-icon> mdi-filter </v-icon>
          Filters
        </v-btn>
      </v-col>
      <!-- <v-spacer /> -->
      <!-- <v-col cols="6" md="3">
      <v-select
        label="Sort by"
        :items="sortOptions"
        v-model="searchParams.sortCriteria"
        item-text="name"
        item-value="value"
        @change="emitSearchEvent()"
      ></v-select>
    </v-col>
    <v-col cols="6" md="1">
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            elevation="1"
            class="direction-btn"
            v-bind="attrs"
            v-on="on"
            hint="Sorting direction"
            @click="changeSortingDirection()"
          >
            <v-icon>
              {{
                searchParams.sortDirection === "asc"
                  ? "mdi-sort-numeric-ascending"
                  : "mdi-sort-numeric-descending"
              }}
            </v-icon>
          </v-btn>
        </template>
        <span>Sorting direction</span>
      </v-tooltip>
    </v-col> -->
    </v-row>
    <v-row dense>
      <v-navigation-drawer
        v-model="drawer"
        absolute
        temporary
        width="400"
        right
      >
        <v-card elevation="4">
          <v-card-title>
            Filters
            <v-spacer></v-spacer>

            <v-btn :ripple="false" elevation="0" @click="drawer = false">
              <v-icon> mdi-close </v-icon>
            </v-btn>
          </v-card-title>
          <v-card-text> </v-card-text>
        </v-card>
      </v-navigation-drawer>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "SearchFilterDrinks",
  props: ["searchParams"],
  data: () => {
    return {
      drawer: false,
      sortOptions: [
        {
          name: "Name",
          value: "name",
        },
        {
          name: "Date created",
          value: "dateCreated",
        },
      ],
    };
  },
  methods: {
    emitSearchEvent() {
      this.$emit("searchParamsUpdated");
    },
    changeSortingDirection() {
      this.searchParams.sortDirection =
        this.searchParams.sortDirection === "asc" ? "desc" : "asc";
      this.emitSearchEvent();
    },
  },
  watch: {
    page: (newPage) => {
      // console.log(newPage);
      // TODO: ovdje pozovi metodu za dobavljanje kolekcija
    },
  },
};
</script>

<style scoped>
.search-bar {
  margin-top: 5px;
}
.search-btn,
.direction-btn {
  margin-top: 15px;
}
</style>
