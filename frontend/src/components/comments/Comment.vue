<template>
  <li class="node-tree">
    <span class="label">{{ comment.content }}</span>
    <span>
      <v-btn small rounded fab elevation="0">
        <v-icon> mdi-reply </v-icon>
      </v-btn>
    </span>
    <ul v-if="comment.children && comment.children.length">
      <comment
        v-for="(child, i) in comment.children"
        :comment="child"
        :key="i"
      />
    </ul>
  </li>
</template>

<script>
export default {
  name: "Comment",
  props: ["comment"],
};
</script>

<style scoped>
.node-tree {
  height: 70px;
}
</style>
