<template>
  <div>
    <li class="node-tree">
      <span class="label">{{ comment.content }}</span>

      <span>
        <v-menu offset-y>
          <template v-slot:activator="{ on, attrs }">
            <v-btn elevation="2" x-small v-bind="attrs" v-on="on">
              <v-icon> mdi-dots-vertical </v-icon>
            </v-btn>
          </template>
          <v-list>
            <v-list-item @click="commentReplyDialog = true">
              <v-list-item-title>
                <v-icon color="primary">mdi-reply</v-icon>

                Reply
              </v-list-item-title>
            </v-list-item>

            <v-list-item @click="commentReportDialog = true">
              <v-list-item-title>
                <v-icon color="red"> mdi-alert-octagon </v-icon>
                Report
              </v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </span>

      <ul v-if="comment.children && comment.children.length">
        <comment
          v-for="child in comment.children"
          :comment="child"
          :key="child.id"
        />
      </ul>
    </li>
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
import CommentReplyDialog from "../dialogs/CommentReplyDialog.vue";
import CommentReportDialog from "../dialogs/CommentReportDialog.vue";

export default {
  name: "Comment",
  props: ["comment"],
  components: {
    CommentReplyDialog,
    CommentReportDialog,
  },
  data: () => {
    return {
      commentReplyDialog: false,
      commentReportDialog: false,
    };
  },
  methods: {
    reply(e) {
      console.log(`reply to comment ${this.comment.id} with ${e}`);
      this.commentReplyDialog = false;
    },
    report(e) {
      console.log(`report comment ${this.comment.id} for ${e}`);
    },
  },
};
</script>

<style>
.node-tree {
  height: 70px;
}
</style>
