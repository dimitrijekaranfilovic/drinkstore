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
  </div>
</template>

<script>
import Comment from "./Comment.vue";
export default {
  name: "Comments",
  props: ["comments"],
  components: {
    Comment,
  },
};
</script>

<style scoped>
.tree-list ul {
  padding-left: 16px;
  margin: 6px 0;
}

.tree {
  max-height: 500px;
  overflow-y: auto;
}
</style>
