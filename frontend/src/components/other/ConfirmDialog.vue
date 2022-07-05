<template>
  <v-dialog v-model="confirmDialog" persistent max-width="290">
    <v-card>
      <v-toolbar flat dark color="primary">
        <v-toolbar-title>{{ title }}</v-toolbar-title>
        <v-spacer />
        <v-btn icon dark @click="$emit('dialog-closing', { answer: 'CANCEL' })">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-toolbar>

      <v-card-text>
        {{ text }}
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn text @click="$emit('dialog-closing', { answer: 'CANCEL' })">
          Cancel
        </v-btn>
        <v-btn
          color="primary"
          text
          @click="$emit('dialog-closing', { answer: 'OK' })"
        >
          Ok
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "ConfirmDialog",
  props: ["confirmDialog", "title", "text"],
};
</script>
