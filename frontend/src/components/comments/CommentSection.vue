<template>
  <v-card>
    <v-card-title> Comments </v-card-title>
    <v-card-text>
      <v-container v-if="userAuthority === 'USER'">
        <v-row>
          <v-col cols="12" md="10">
            <v-text-field label="Enter comment" v-model="commentContent" />
          </v-col>
          <v-col cols="12" md="2">
            <v-btn color="primary" class="send-btn" @click="emitAddComment()">
              <v-icon color="white"> mdi-send </v-icon>
              Send
            </v-btn>
          </v-col>
        </v-row>
      </v-container>
      <comments
        :comments="comments"
        v-if="comments.length > 0"
        :drinkId="drinkId"
      />
    </v-card-text>
  </v-card>
</template>

<script>
import Comments from "./Comments.vue";
export default {
  props: ["comments", "drinkId"],
  components: { Comments },
  name: "CommentSection",
  data: () => {
    return {
      commentContent: "",
    };
  },
  methods: {
    emitAddComment() {
      this.$emit("comment-added", this.commentContent);
      this.commentContent = "";
    },
  },
  computed: {
    userAuthority() {
      return this.$store.state.user.authority;
    },
  },
};
</script>
