<template>
  <v-container>
    <v-card class="mx-auto">
      <v-img
        src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
        height="200px"
      ></v-img>
      <v-card-title>
        {{ drink.name }}

        <v-chip class="ma-2" color="secondary" outlined>
          {{ drink.volume }}{{ drink.volumeMark }}
        </v-chip>
        <v-chip :color="drinkCategoryChipColor" text-color="white">
          {{ drink.category | capitalize | removeUnderscore }}
        </v-chip>
      </v-card-title>
      <v-container>
        <v-row>
          <v-col>
            <v-btn color="#ea9b09" dark rounded>
              <v-icon color="white"> mdi-pencil </v-icon>
              Edit
            </v-btn>
          </v-col>
        </v-row>
      </v-container>
      <v-card-actions>
        <span class="avg-grade">Rating:</span>
        <v-spacer></v-spacer>
        <span class="text--lighten-2 text-caption mr-2"> (1.1) </span>
        <v-rating
          :value="1.2"
          half-increments
          hover
          length="5"
          readonly
          color="yellow darken-3"
          background-color="yellow darken-3"
          size="18"
        ></v-rating>
      </v-card-actions>
      <v-divider />

      <v-card-text>
        {{ drink.description }}
      </v-card-text>
    </v-card>
    <v-divider />
    <v-card>
      <v-card-title> Comments </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12" md="10">
              <v-text-field label="Enter comment" />
            </v-col>
            <v-col cols="12" md="2">
              <v-btn color="primary" class="send-btn">
                <v-icon color="white"> mdi-send </v-icon>
                Send
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
        <comments :comments="comments" />
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
import { categories } from "../../util/categories";
import Comments from "../comments/Comments.vue";

export default {
  name: "Drink",
  components: {
    Comments,
  },
  data: () => {
    //ovaj drink je samo placeholder
    return {
      drink: {
        id: 1,
        name: "Vinjak",
        description: "Opis",
        volume: 0.75,
        volumeMark: "l",
        averageGrade: 4.8,
        category: "BEER",
      },
      comments: [
        {
          id: 1,
          content: "AAA",
          children: [
            {
              id: 2,
              content: "BBB",
              children: [
                {
                  id: 3,
                  content: "CCC",
                },
              ],
            },
            {
              id: 4,
              content: "DDD",
            },
          ],
        },
      ],
    };
  },
  computed: {
    drinkCategoryChipColor() {
      return categories.find((c) => c.value === this.drink.category).color;
    },
  },
  mounted() {
    //TODO: ovdje dobavi dato pice na osnovu id-ija
  },
};
</script>

<style scoped>
.avg-grade {
  color: rgba(0, 0, 0, 0.6);
  font-size: 14px;
}

.no-bottom-margin {
  margin-bottom: 5px !important;
}

.send-btn {
  margin-top: 15px;
}
</style>
