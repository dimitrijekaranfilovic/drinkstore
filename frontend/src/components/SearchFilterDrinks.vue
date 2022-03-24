<template>
  <v-container>
    <v-row class="search-bar">
      <v-col cols="12" md="8">
        <!--ovo gleda i naziv i opis i mozda da ih na bekendu rangiram po relevantnosti-->
        <v-text-field
          label="Enter keyword"
          v-model="searchParams.query"
        ></v-text-field>
      </v-col>
      <v-col cols="12" md="4">
        <v-btn color="primary" class="search-btn" @click="emitSearchEvent()"
          >Search</v-btn
        >
        <v-btn @click="dialog = !dialog" class="search-btn">
          <v-icon> mdi-filter </v-icon>
          Filters
        </v-btn>
      </v-col>
    </v-row>
    <div class="dialogs">
      <v-dialog
        width="50%"
        v-model="dialog"
        hide-overlay
        transition="dialog-bottom-transition"
      >
        <v-card>
          <v-toolbar dark color="primary">
            <v-toolbar-title>
              Filters
              <v-icon> mdi-filter </v-icon>
            </v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn icon dark @click="dialog = false">
              <v-icon>mdi-close</v-icon>
            </v-btn>
          </v-toolbar>

          <v-card-text>
            <v-container>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    label="From grade"
                    type="number"
                    min="0"
                    max="5"
                    step="0.1"
                    v-model="searchParams.gradeFrom"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field
                    label="To grade"
                    type="number"
                    min="0"
                    max="5"
                    step="0.1"
                    hide-details
                    v-model="searchParams.gradeTo"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-select
                    v-model="searchParams.categories"
                    :items="categories"
                    item-text="text"
                    item-value="value"
                    label="Category"
                    multiple
                    chips
                    deletable-chips
                  />
                </v-col>
              </v-row>
            </v-container>
          </v-card-text>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script>
//TODO: mozda dodaj neke napredne filtere za poredjenje zapremine
//TODO: dodaj ono za najprodavanije, ali to mozda na zasebnu stranicu sa drugacijim karticama
export default {
  name: "SearchFilterDrinks",
  props: ["searchParams"],
  data: () => {
    return {
      dialog: false,
      categories: [
        {
          text: "Liquor",
          value: "LIQUOR",
        },
        {
          text: "Beer",
          value: "BEER",
        },
        {
          text: "Wine",
          value: "WINE",
        },
        {
          text: "Carbonated",
          value: "CARBONATED",
        },
        {
          text: "Non carbonated",
          value: "NON_CARBONATED",
        },
      ],
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
