<template>
  <div>
    <li class="node-tree">
      <span class="label"
        ><strong>{{ comment.user }}</strong
        >: {{ comment.content }}</span
      >

      <span>
        <v-menu offset-y>
          <template v-slot:activator="{ on, attrs }">
            <v-btn elevation="0" x-small v-bind="attrs" v-on="on">
              <v-icon> mdi-dots-vertical </v-icon>
            </v-btn>
          </template>
          <v-list>
            <v-list-item @click="emitReply()">
              <v-list-item-title>
                <v-icon color="primary">mdi-reply</v-icon>

                Reply
              </v-list-item-title>
            </v-list-item>

            <v-list-item @click="emitReport()">
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
  </div>
</template>

<script>
import CommentReplyDialog from "./CommentReplyDialog.vue";
import CommentReportDialog from "./CommentReportDialog.vue";
import { eventBus } from "../../util/eventBus";

export default {
  name: "Comment",
  props: ["comment"],
  components: {
    CommentReplyDialog,
    CommentReportDialog,
  },
  methods: {
    emitReply() {
      eventBus.$emit("reply-btn-clicked", this.comment.id);
    },
    emitReport() {
      eventBus.$emit("report-btn-clicked", this.comment.id);
    },
  },
};
</script>

<style>
.node-tree {
  height: 70px;
}
</style>
