<template>
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
          <v-list-item>
            <v-list-item-title>
              <v-icon color="primary">mdi-reply</v-icon>

              Reply
            </v-list-item-title>
          </v-list-item>

          <v-list-item>
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

<style>
.node-tree {
  height: 70px;
}
</style>
