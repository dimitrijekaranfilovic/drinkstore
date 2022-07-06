<template>
  <v-container>
    <span>
      <span class="comment-user">{{ comment.user }}</span>
      <v-btn
        rounded
        elevation="0"
        small
        @click="emitReport()"
        v-if="userAuthority === 'USER' && username !== comment.user"
      >
        <v-icon color="red">mdi-alert-octagon</v-icon>
        Report
      </v-btn>
    </span>
    <p>
      {{ comment.postedAt | formatDate }}
    </p>
    <p class="comment-content">
      {{ comment.content }}
    </p>

    <v-divider />
  </v-container>
</template>

<script>
import CommentReportDialog from "./CommentReportDialog.vue";
import { eventBus } from "../../util/eventBus";

export default {
  name: "Comment",
  props: ["comment"],
  components: {
    CommentReportDialog,
  },
  methods: {
    emitReport() {
      eventBus.$emit("report-btn-clicked", this.comment.id);
    },
  },
  computed: {
    userAuthority() {
      return this.$store.state.user.authority;
    },
    username() {
      return this.$store.state.user.username;
    },
  },
};
</script>

<style scoped>
.comment-content {
  font-size: 22px;
  color: black;
}

.comment-user {
  font-size: 1.5rem;
  letter-spacing: 0.0071428571em;
  color: rgba(0, 0, 0, 0.6);
}
</style>
