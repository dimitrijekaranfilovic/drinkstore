<template>
  <v-container>
    <comment v-for="comment in comments" :comment="comment" :key="comment.id" />
    <comment-report-dialog
      :dialog="commentReportDialog"
      @dialog-closing="commentReportDialog = false"
      @reported="report($event)"
    />
    <toast
      :snackbar="snackbar"
      :text="text"
      @toast-closing="snackbar = false"
    />
  </v-container>
</template>

<script>
import Comment from "./Comment.vue";
import CommentReportDialog from "./CommentReportDialog.vue";
import Toast from "../other/Toast.vue";
import { eventBus } from "../../util/eventBus";
import { commentService } from "../../services/commentService";

export default {
  name: "Comments",
  props: ["comments", "drinkId"],
  data: () => {
    return {
      snackbar: false,
      text: "",
      commentId: undefined,
      commentReportDialog: false,
    };
  },
  components: {
    Toast,
    Comment,
    CommentReportDialog,
  },
  methods: {
    report(reportReason) {
      const commentToReport = this.comments.find(
        (item) => item.id === this.commentId
      );

      const reportPayload = {
        commentContent: commentToReport.content,
        commentId: commentToReport.id,
        postedBy: commentToReport.user,
        drinkId: this.drinkId,
        reportReason: reportReason,
        reportedBy: this.$store.state.user.username,
      };

      this.commentReportDialog = false;
      commentService
        .addReport(reportPayload)
        .then((_) => {
          this.text = "Comment reported successfully.";
          this.snackbar = true;
        })
        .catch((error) => {
          if (error.response) this.text = error.response.data.message;
          else this.text = "An error occurred while reporting comment.";
          this.snackbar = true;
        });
    },
  },
  mounted() {
    eventBus.$on("report-btn-clicked", (e) => {
      this.commentId = e;
      this.commentReportDialog = true;
    });
  },
};
</script>

<style>
.tree-list ul {
  padding-left: 16px;
  margin: 6px 0;
}

.tree {
  max-height: 250px;
  height: 200px;
  overflow-y: auto;
}
</style>
