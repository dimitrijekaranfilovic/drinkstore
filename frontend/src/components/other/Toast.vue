<template>
  <v-snackbar v-model="snackbar" :timeout="5000">
    {{ text }}
    <template v-slot:action="{ attrs }">
      <v-btn
        color="primary"
        text
        v-bind="attrs"
        @click="$emit('toast-closing')"
      >
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script>
export default {
  props: ["snackbar", "text"],
  name: "Toast",
};
</script>
