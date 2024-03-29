<template>
  <v-dialog
    width="50%"
    v-model="dialog"
    hide-overlay
    transition="dialog-bottom-transition"
    persistent
  >
    <v-card>
      <v-toolbar dark color="primary">
        <v-toolbar-title> {{ dialogTitleText }} </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon dark @click="$emit('dialog-closing')">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-toolbar>

      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field label="Drink name" v-model="drink.name" />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                label="Drink description"
                v-model="drink.description"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                label="Volume"
                type="number"
                step="0.1"
                min="0"
                v-model="drink.volume"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                label="Volume label"
                :items="volumeLabels"
                v-model="drink.volumeLabel"
              />
            </v-col>
          </v-row>
          <v-row align="center" justify="center">
            <v-col cols="12">
              <v-text-field
                label="Price"
                v-model="drink.price"
                type="number"
                append-icon="mdi-cash-multiple"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-file-input
                accept="image/*"
                show-size
                counter
                chips
                label="Image"
                prepend-icon="mdi-camera"
                v-model="drink.image"
              ></v-file-input>
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                label="Category"
                :items="categories"
                v-model="drink.category"
              />
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-btn @click="$emit('dialog-closing')">Cancel</v-btn>
        <v-btn
          color="primary"
          @click="emitConfirmEvent()"
          :disabled="btnDisabled"
          >{{ confirmButtonText }}</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { volumeLabels } from "../../util/volumeLabels";
import { categories } from "../../util/categories";

export default {
  name: "DrinkDialog",
  props: ["dialog", "drinkToEdit"],
  data: () => {
    return {
      volumeLabels: volumeLabels,
      categories: categories,
      drink: {
        name: "",
        description: "",
        volume: 0,
        volumeLabel: "",
        image: null,
        category: "",
        price: 0,
      },
    };
  },
  mounted() {
    if (this.drinkToEdit) {
      this.updateDrink();
    }
  },
  methods: {
    updateDrink() {
      this.drink = { ...this.drinkToEdit };
      this.drink.image = null;
      delete this.drink.averageGrade;
      delete this.drink.imagePath;
    },
    emitConfirmEvent() {
      let eventName;
      if (this.drinkToEdit) {
        eventName = "drink-edited";
      } else {
        eventName = "drink-added";
      }

      this.$emit(eventName, { ...this.drink });
      this.drink = {
        name: "",
        description: "",
        volume: 0,
        volumeLabel: "",
        image: null,
        category: "",
        price: 0,
      };
    },
  },
  computed: {
    btnDisabled() {
      return (
        this.drink.name === "" ||
        this.drink.description === "" ||
        this.drink.volume == 0 ||
        this.drink.volumeLabel === "" ||
        (this.drink.image === null && this.drinkToEdit === undefined) ||
        this.drink.category === "" ||
        this.drink.price == 0
      );
    },
    confirmButtonText() {
      return this.drinkToEdit === undefined ? "Add" : "Save changes";
    },
    dialogTitleText() {
      return this.drinkToEdit === undefined ? "Add drink" : "Edit drink";
    },
  },
  watch: {
    dialog(newDialog) {
      if (newDialog === true) this.updateDrink();
    },
  },
};
</script>
