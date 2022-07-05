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
        <v-row align="center" justify="center">
          <v-col cols="10">
            <v-rating
              hover
              length="5"
              color="yellow darken-3"
              background-color="yellow darken-3"
              v-model="grade"
              size="36"
            ></v-rating>
          </v-col>
          <v-col cols="2">
            <h5>Grade: {{ grade }}</h5>
          </v-col>
        </v-row>
      </v-card-text>
      <v-card-actions>
        <v-btn @click="$emit('dialog-closing')">Cancel</v-btn>
        <v-btn color="primary" @click="emitEvent()">Confirm</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  props: ["dialog", "dialogTitleText", "gradeToChange"],
  name: "GradingDialog",
  data: () => {
    return {
      grade: 0,
    };
  },
  mounted() {
    this.updateGrade();
  },
  methods: {
    emitEvent() {
      if (this.gradeToChange === undefined)
        this.$emit("grade-created", this.grade);
      else this.$emit("grade-updated", this.grade);
    },
    updateGrade() {
      if (this.gradeToChange !== undefined) this.grade = this.gradeToChange;
    },
  },
  watch: {
    dialog(newDialog) {
      this.updateGrade();
    },
  },
};
</script>
