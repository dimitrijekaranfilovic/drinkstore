<template>
  <v-container>
    <v-row class="search-bar">
      <v-col cols="12" md="8">
        <!--ovo gleda i naziv i opis i mozda da ih na bekendu rangiram po relevantnosti-->
        <v-text-field
          label="Enter keyword"
          v-model="searchParams.query"
          clearable
        ></v-text-field>
      </v-col>
      <v-col cols="12" md="4">
        <v-btn
          color="primary"
          class="search-btn"
          @click="$emit('searchParamsUpdated')"
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
              <v-row>
                <v-col cols="12">
                  <v-select
                    v-model="searchParams.volumeLabels"
                    :items="volumeLabels"
                    label="Volume label"
                    multiple
                    chips
                    deletable-chips
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6">
                  <v-select
                    label="Sort by"
                    :items="sortOptions"
                    item-text="text"
                    item-value="value"
                    v-model="searchParams.sortCriteria"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-select
                    label="Direction"
                    :items="sortDirections"
                    item-text="text"
                    item-value="value"
                    v-model="searchParams.sortDirection"
                    :append-outer-icon="sortIcon"
                  />
                </v-col>
              </v-row>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-btn @click="dialog = false">Cancel</v-btn>
            <v-btn color="primary" @click="$emit('searchParamsUpdated')"
              >Apply</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script>
import { categories } from "@/util/categories";
import { volumeLabels } from "@/util/volumeLabels";
export default {
  name: "SearchFilterDrinks",
  props: ["searchParams"],
  data: () => {
    return {
      dialog: false,
      categories: categories,
      volumeLabels: volumeLabels,
      sortOptions: [
        {
          text: "Name",
          value: "name",
        },
        {
          text: "Average grade",
          value: "average_grade",
        },
        {
          text: "Category",
          value: "category",
        },
        {
          text: "Price",
          value: "price",
        },
        {
          text: "Description",
          value: "description",
        },
      ],
      sortDirections: [
        {
          text: "ascending",
          value: "asc",
        },
        {
          text: "descending",
          value: "desc",
        },
      ],
    };
  },
  methods: {},
  computed: {
    sortIcon() {
      if (this.searchParams.sortDirection === "desc")
        return "mdi-sort-alphabetical-descending";
      else return "mdi-sort-alphabetical-ascending";
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
