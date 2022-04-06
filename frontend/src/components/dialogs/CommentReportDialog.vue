<template>
  <v-dialog
    width="50%"
    v-model="dialog"
    hide-overlay
    transition="dialog-bottom-transition"
  >
    <v-card>
      <v-toolbar dark color="primary">
        <v-toolbar-title>
          Report comment
          <v-icon>mdi-alert-octagon</v-icon>
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon dark @click="$emit('dialog-closing')">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-toolbar>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-chip-group
                active-class="primary--text"
                v-model="chip"
                class="chip-group"
                column
              >
                <v-chip v-for="(reason, i) in reportReasons" :key="i">
                  {{ reason | capitalize | removeUnderscore }}
                </v-chip>
              </v-chip-group>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>

      <v-card-actions>
        <v-btn @click="closeDialog()">Cancel</v-btn>
        <v-btn color="primary" @click="report()" :disabled="chip === undefined"
          >Report</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { reportReasons } from "../../util/reportReasons";
export default {
  name: "CommentReportDialog",
  props: ["dialog"],
  data: () => {
    return {
      chip: undefined,
      reportReasons: reportReasons,
    };
  },
  computed: {
    reason() {
      return this.reportReasons.at(this.chip);
    },
  },
  methods: {
    closeDialog() {
      this.$emit("dialog-closing");
      this.chip = undefined;
    },
    report() {
      this.$emit("reported", this.reason);
      this.chip = undefined;
    },
  },
};
</script>

<style scoped>
/* .chip-group {
  max-width: 500px;
  min-width: 20px;
  overflow-x: auto;
} */
</style>
