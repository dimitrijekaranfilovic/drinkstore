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
      <v-col cols="12" v-if="reports.length > 0">
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
              <tr v-for="report in reports" :key="report.id">
                <td>{{ report.commentContent }}</td>
                <td>
                  {{ report.reportReason | removeUnderscore | capitalize }}
                </td>
                <td>{{ report.postedBy }}</td>
                <td>{{ report.reportedBy }}</td>
                <td>{{ report.reportedAt | formatDate }}</td>

                <td>
                  <router-link
                    :to="{ name: 'Drink', params: { id: report.drinkId } }"
                    >View drink</router-link
                  >
                </td>
                <td>
                  <v-btn color="primary" dark @click="deleteReport(report.id)"
                    >Delete report</v-btn
                  >
                </td>
                <td>
                  <v-btn
                    color="red"
                    dark
                    @click="deleteComment(report.commentId)"
                    >Delete comment</v-btn
                  >
                </td>
                <td>
                  <v-btn color="purple" dark @click="banUser(report)"
                    >Ban user</v-btn
                  >
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
      <v-col v-else cols="12">
        <v-flex class="text-center">
          <h2>No reports.</h2>
        </v-flex>
      </v-col>
    </v-row>
    <toast
      :snackbar="snackbar"
      :text="text"
      @toast-closing="snackbar = false"
    />
  </v-container>
</template>

<script>
import { commentService } from "../../services/commentService";
import { userService } from "../../services/userService";
import Toast from "../other/Toast.vue";
export default {
  name: "ReportedComments",
  components: {
    Toast,
  },
  data: () => {
    return {
      snackbar: false,
      text: "",
      reports: [],
    };
  },
  mounted() {
    commentService.getReports().then((response) => {
      this.reports = response.data;
    });
  },
  methods: {
    deleteReport(reportId) {
      commentService
        .deleteReport(reportId)
        .then((_) => {
          this.text = "Report successfully deleted.";
          this.snackbar = true;
          this.reports = this.reports.filter((item) => item.id !== reportId);
        })
        .catch((error) => {
          if (error.response) this.text = error.response.data.message;
          else this.text = "An error occurred while deleting report.";
          this.snackbar = true;
        });
    },
    deleteComment(commentId) {
      commentService
        .deleteComment(commentId)
        .then((_) => {
          this.text = "Comment successfully deleted.";
          this.snackbar = true;
          this.reports = this.reports.filter(
            (item) => item.commentId !== commentId
          );
        })
        .catch((error) => {
          if (error.response) this.text = error.response.data.message;
          else this.text = "An error occurred while deleting comment.";
          this.snackbar = true;
        });
    },
    banUser(report) {
      const username = report.postedBy;
      userService
        .banUser(username)
        .then((_) => {
          this.text = `User ${username} successfully banned`;
          this.snackbar = true;
          this.deleteComment(report.commentId);
        })
        .catch((error) => {
          if (error.response) this.text = error.response.data.message;
          else this.text = `An error occurred while banning user ${username}.`;
          this.snackbar = true;
        });
    },
  },
};
</script>
