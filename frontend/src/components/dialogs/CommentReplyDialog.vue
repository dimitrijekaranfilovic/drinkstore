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
          Reply to comment
          <v-icon>mdi-reply</v-icon>
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
              <v-text-field label="Enter reply" v-model="replyText" />
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-btn @click="closeDialog()">Cancel</v-btn>
        <v-btn color="primary" @click="reply()">Reply</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "CommentReplyDialog",
  props: ["dialog"],
  data: () => {
    return {
      replyText: "",
    };
  },
  methods: {
    closeDialog() {
      this.$emit("dialog-closing");
      this.replyText = "";
    },
    reply() {
      this.$emit("replied", this.replyText);
      this.replyText = "";
    },
  },
};
</script>
