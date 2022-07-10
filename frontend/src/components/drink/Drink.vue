<template>
  <v-container>
    <v-row align="center" justify="center" v-if="drink">
      <v-col cols="2">
        <v-img :src="imagePath" height="400px"></v-img>
      </v-col>
      <v-col cols="10">
        <v-card class="mx-auto">
          <v-card-title>
            {{ drink.name }}

            <v-chip class="ma-2" color="secondary" outlined>
              {{ drink.volume }}{{ drink.volumeLabel }}
            </v-chip>
            <v-chip :color="drinkCategoryChipColor" text-color="white">
              {{ drink.category | capitalize | removeUnderscore }}
            </v-chip>

            <v-col align="right">
              <v-btn
                color="red"
                rounded
                dark
                v-if="itemAlreadyinCart && userAuthority === 'USER'"
                @click="removeFromCart()"
              >
                <v-icon> mdi-delete </v-icon>
                Remove from cart
              </v-btn>
              <v-btn
                color="primary"
                rounded
                dark
                v-else-if="!itemAlreadyinCart && userAuthority === 'USER'"
                @click="addToCart()"
              >
                <v-icon> mdi-cart </v-icon>
                Add to cart
              </v-btn>
              <v-btn
                color="#ea9b09"
                dark
                rounded
                @click="drinkDialog = true"
                v-if="userAuthority === 'ADMIN'"
              >
                <v-icon color="white"> mdi-pencil </v-icon>
                Edit
              </v-btn>
              <v-btn
                color="red"
                dark
                rounded
                @click="confirmDialog = true"
                v-if="userAuthority === 'ADMIN'"
              >
                <v-icon color="white"> mdi-delete-forever </v-icon>
                Delete
              </v-btn>
            </v-col>
          </v-card-title>
          <v-chip class="ma-2" color="red" text-color="white">
            <strong>{{ drink.price }} RSD</strong>
          </v-chip>
          <v-card-actions v-if="userAuthority === 'USER'">
            <span>Your grade:</span>
            <v-spacer />

            <span
              v-if="!userGrade.gradeExists"
              class="text--lighten-2 text-caption mr-2"
            >
              Not graded
            </span>
            <v-btn v-if="!userGrade.gradeExists" @click="grade()">Grade</v-btn>
            <span
              class="text--lighten-2 text-caption mr-2"
              v-if="userGrade.gradeExists"
            >
              {{ userGrade.gradeValue }}
            </span>
            <v-rating
              v-if="userGrade.gradeExists"
              :value="userGrade.gradeValue"
              half-increments
              hover
              length="5"
              readonly
              color="yellow darken-3"
              background-color="yellow darken-3"
              size="18"
            ></v-rating>
            <v-btn
              color="orange"
              dark
              v-if="userGrade.gradeExists"
              @click="changeGrade()"
              >Change grade</v-btn
            >
            <v-btn
              dark
              color="red"
              v-if="userGrade.gradeExists"
              @click="removeGrade()"
              >Remove grade</v-btn
            >
          </v-card-actions>
          <v-card-actions>
            <span class="avg-grade">Average grade:</span>
            <v-spacer />
            <span class="text--lighten-2 text-caption mr-2">
              {{ drink.averageGrade }}
            </span>
            <v-rating
              :value="drink.averageGrade"
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
      </v-col>
    </v-row>

    <v-divider />
    <comment-section
      :comments="comments"
      @comment-added="addComment($event)"
      :drinkId="$route.params.id"
    />
    <drink-dialog
      :dialog="drinkDialog"
      :drinkToEdit="drink"
      @dialog-closing="drinkDialog = false"
      @drink-edited="editDrink($event)"
    />
    <confirm-dialog
      :confirmDialog="confirmDialog"
      :text="dialogDeleteText"
      :title="dialogDeleteTitle"
      @dialog-closing="closeConfirmDialog($event)"
    />
    <toast
      :snackbar="snackbar"
      :text="text"
      @toast-closing="snackbar = false"
    />
    <grading-dialog
      :dialog="gradingDialog"
      :dialogTitleText="gradingDialogTitleText"
      :gradeToChange="userGrade.gradeValue"
      @dialog-closing="gradingDialog = false"
      @grade-created="createGrade($event)"
      @grade-updated="updateGrade($event)"
    />
  </v-container>
</template>

<script>
import { categories } from "../../util/categories";
import DrinkDialog from "./DrinkDialog.vue";
import { drinkService } from "../../services/drinkService";
import ConfirmDialog from "../other/ConfirmDialog.vue";
import Toast from "../other/Toast.vue";
import GradingDialog from "../other/GradingDialog.vue";
import CommentSection from "../comments/CommentSection.vue";
import { commentService } from "../../services/commentService";

export default {
  name: "Drink",
  components: {
    CommentSection,
    DrinkDialog,
    ConfirmDialog,
    Toast,
    GradingDialog,
  },
  data: () => {
    return {
      userGrade: {
        gradeExists: false,
        gradeValue: undefined,
        gradeId: -1,
      },
      gradingDialog: false,
      gradingDialogTitleText: "Grade",
      snackbar: false,
      text: "",
      confirmDialog: false,
      drinkDialog: false,
      drink: null,
      comments: [],
      dialogDeleteTitle: "Drink deletion",
    };
  },
  mounted() {
    drinkService.getSingleDrink(this.$route.params.id).then((response) => {
      this.drink = response.data;
    });
    drinkService.checkUserGrade(this.$route.params.id).then((response) => {
      this.userGrade = response.data;
      if (this.userGrade.gradeValue === -1)
        this.userGrade.gradeValue = undefined;
    });
    commentService
      .getCommentsForDrink(this.$route.params.id)
      .then((response) => {
        this.comments = response.data;
      });
  },
  methods: {
    addComment(commentContent) {
      const payload = {
        content: commentContent,
        user: this.$store.state.user.username,
        drinkId: this.drink.id,
      };
      commentService
        .addComment(payload)
        .then((response) => {
          this.comments.push({ ...response.data });
        })
        .catch((error) => {
          if (error.response) this.text = error.response.data.message;
          else this.text = "An error occured while commenting.";
          this.snackbar = true;
        });
    },
    removeGrade() {
      drinkService
        .deleteUserGrade(this.$route.params.id, this.userGrade.gradeId)
        .then((_) => {
          this.userGrade.gradeExists = false;
          this.userGrade.gradeValue = undefined;
        });
    },
    createGrade(grade) {
      drinkService
        .addUserGrade(this.$route.params.id, { grade: grade })
        .then((response) => {
          this.gradingDialog = false;
          this.userGrade.gradeExists = true;
          this.userGrade.gradeValue = grade;
          this.userGrade.gradeId = response.data.id;
        })
        .catch((error) => {
          if (error.response) this.text = error.response.data.message;
          else this.text = "An error occured while grading drink.";
          this.snackbar = true;
        });
    },
    updateGrade(grade) {
      drinkService
        .updateUserGrade(this.$route.params.id, this.userGrade.gradeId, {
          grade: grade,
        })
        .then((response) => {
          this.gradingDialog = false;
          this.userGrade.gradeExists = true;
          this.userGrade.gradeValue = grade;
          this.userGrade.gradeId = response.data.id;
        });
    },

    grade() {
      this.gradingDialogTitleText = "Grade";
      this.gradingDialog = true;
      this.gradeToChange = 0;
    },
    changeGrade() {
      this.dialogDeleteTitle = "Change grade";
      this.gradingDialog = true;
      this.gradeToChange = this.userGrade.gradeValue;
    },
    closeConfirmDialog(event) {
      if (event.answer === "OK") {
        this.deleteDrink();
        this.confirmDialog = false;
      } else this.confirmDialog = false;
    },
    editDrink(drink) {
      let drinkUpdatePayload = { ...drink };
      drinkUpdatePayload.price += "";
      drinkUpdatePayload.volume += "";
      delete drinkUpdatePayload.image;
      drinkService
        .updateDrink(this.drink.id, drinkUpdatePayload)
        .then((_) => {
          if (drink.image !== null) {
            drinkService
              .createDrinkImage(drink.id, drink.image)
              .then((_) => {
                this.$router.go(0);
              })
              .catch((error) => {
                if (error.response) this.text = error.response.data.message;
                else
                  this.text = "An error occurred while uploading drink image.";
                this.snackbar = true;
              });
          } else {
            this.$router.go(0);
          }
        })
        .catch((error) => {
          if (error.response) this.text = error.response.data.message;
          else this.text = "An error occurred while updating drink.";
          this.snackbar = true;
        });
    },
    deleteDrink() {
      drinkService.deleteDrink(this.$route.params.id).then((_) => {
        this.$router.push({ name: "Home" });
      });
    },
    addToCart() {
      this.$store.dispatch("addCartItem", {
        id: this.drink.id,
        name: this.drink.name,
        price: this.drink.price,
        amount: 1,
      });
    },
    removeFromCart() {
      this.$store.dispatch("removeCartItem", this.drink.id);
    },
  },
  computed: {
    dialogDeleteText() {
      if (this.drink)
        return `Are you sure you want to delete ${this.drink.name}?`;
      else "Are you sure you want to delete this drink?";
    },
    imagePath() {
      return `${process.env.VUE_APP_DRINK_SERVICE_BASE_PATH}/${this.drink.imagePath}`;
    },
    userAuthority() {
      return this.$store.state.user.authority;
    },
    drinkCategoryChipColor() {
      return categories.find((c) => c.value === this.drink.category).color;
    },
    itemAlreadyinCart() {
      let result =
        this.$store.state.cart.find((item) => item.id === this.drink.id) !==
        undefined;
      return result;
    },
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
