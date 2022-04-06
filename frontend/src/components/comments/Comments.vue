<template>
  <div class="tree">
    <ul class="tree-list">
      <!-- <v-for comment in comments>
        <comment :comment="comment" :key="i" />
      </v-for> -->
      <comment
        v-for="comment in comments"
        :comment="comment"
        :key="comment.id"
      />
    </ul>
    <comment-reply-dialog
      :dialog="commentReplyDialog"
      @dialog-closing="commentReplyDialog = false"
      @replied="reply($event)"
    />
    <comment-report-dialog
      :dialog="commentReportDialog"
      @dialog-closing="commentReportDialog = false"
      @reported="report($event)"
    />
  </div>
</template>

<script>
import Comment from "./Comment.vue";
import CommentReplyDialog from "../dialogs/CommentReplyDialog.vue";
import CommentReportDialog from "../dialogs/CommentReportDialog.vue";
import { eventBus } from "../../util/eventBus";

export default {
  name: "Comments",
  props: ["comments"],
  data: () => {
    return {
      commentId: undefined,
      commentReplyDialog: false,
      commentReportDialog: false,
    };
  },
  components: {
    Comment,
    CommentReplyDialog,
    CommentReportDialog,
  },
  methods: {
    reply(replyText) {
      console.log(`reply to comment ${this.commentId} with ${replyText}`);
      this.commentReplyDialog = false;
    },
    report(reportReason) {
      console.log(`report comment ${this.commentId} for ${reportReason}`);
      this.commentReportDialog = false;
    },
  },
  mounted() {
    eventBus.$on("reply-btn-clicked", (e) => {
      this.commentId = e;
      this.commentReplyDialog = true;
      //console.log(`reply to ${e}`);
    });
    eventBus.$on("report-btn-clicked", (e) => {
      this.commentId = e;
      this.commentReportDialog = true;
      //console.log(`report ${e}`);
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
