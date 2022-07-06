<template>
  <v-container>
    <comment v-for="comment in comments" :comment="comment" :key="comment.id" />
    <comment-report-dialog
      :dialog="commentReportDialog"
      @dialog-closing="commentReportDialog = false"
      @reported="report($event)"
    />
  </v-container>
</template>

<script>
import Comment from "./Comment.vue";
import CommentReportDialog from "./CommentReportDialog.vue";
import { eventBus } from "../../util/eventBus";

export default {
  name: "Comments",
  props: ["comments"],
  data: () => {
    return {
      commentId: undefined,
      commentReportDialog: false,
    };
  },
  components: {
    Comment,
    CommentReportDialog,
  },
  methods: {
    report(reportReason) {
      console.log(`report comment ${this.commentId} for ${reportReason}`);
      this.commentReportDialog = false;
    },
  },
  mounted() {
    eventBus.$on("reply-btn-clicked", (e) => {
      this.commentId = e;
      this.commentReplyDialog = true;
    });
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
