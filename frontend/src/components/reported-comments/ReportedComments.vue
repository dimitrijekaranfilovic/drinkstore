<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-col cols="12">
        <v-flex class="text-center">
          <h1>Reported comments</h1>
        </v-flex>
      </v-col>
    </v-row>
    <v-row align="center" justify="center">
      <v-col cols="12">
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-left">Comment content</th>
                <th class="text-left">Reported for</th>
                <th class="text-left">Posted by</th>
                <th class="text-left">Reported by</th>
                <th class="text-left">Reported at</th>

                <th class="text-left">Drink</th>
                <!--id pica na kojem je ostavljen komentar-->
                <th class="text-left">Delete report</th>
                <th class="text-left">Delete comment</th>
                <th class="text-left">Ban user</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="comment in reports" :key="comment.id">
                <td>{{ comment.commentContent }}</td>
                <td>
                  {{ comment.reportReason | removeUnderscore | capitalize }}
                </td>
                <td>{{ comment.postedBy }}</td>
                <td>{{ comment.reportedBy }}</td>
                <td>{{ comment.reportedAt | formatDate }}</td>

                <td>
                  <router-link
                    :to="{ name: 'Drink', params: { id: comment.drinkId } }"
                    >View drink</router-link
                  >
                </td>
                <td>
                  <v-btn color="primary" dark>Delete report</v-btn>
                </td>
                <td>
                  <v-btn color="red" dark>Delete comment</v-btn>
                </td>
                <td>
                  <v-btn color="purple" dark>Ban user</v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { commentService } from "../../services/commentService";
export default {
  name: "ReportedComments",
  data: () => {
    return {
      reports: [],
    };
  },
  mounted() {
    commentService.getReports().then((response) => {
      this.reports = response.data;
    });
  },
};
</script>
