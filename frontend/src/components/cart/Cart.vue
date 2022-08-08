<template>
  <!-- <div>CART</div> -->
  <v-container>
    <v-row align="center" justify="center">
      <v-col cols="12">
        <v-flex class="text-center">
          <h1>Cart</h1>
        </v-flex>
      </v-col>
    </v-row>
    <v-row align="center" justify="center">
      <v-col cols="12">
        <v-simple-table v-if="$store.state.cart.length > 0">
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-left">Drink name</th>
                <th class="text-left">Amount</th>
                <th class="text-left">Total price</th>
                <th class="text-left">Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in $store.state.cart" :key="item.id">
                <td>
                  <router-link
                    :to="{ name: 'Drink', params: { id: item.id } }"
                    >{{ item.name }}</router-link
                  >
                </td>
                <td>{{ item.amount }}</td>
                <td>{{ item.amount * item.price }}</td>

                <td>
                  <v-btn color="red" dark @click="removeFromCart(item.id)">
                    <v-icon> mdi-delete </v-icon>
                  </v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
        <v-flex class="text-center" v-else>
          <h2>Cart is empty.</h2>
        </v-flex>
      </v-col>
    </v-row>

    <v-row align="center" justify="center" v-if="$store.state.cart.length > 0">
      <v-col cols="12">
        <v-btn
          :loading="confirmBtnLoading"
          @click="createPurchase()"
          color="primary"
        >
          Confirm order
        </v-btn>
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
import { purchaseService } from "../../services/purchaseService";
import { authService } from "../../services/authService";

import Toast from "../other/Toast.vue";

export default {
  name: "Cart",
  components: { Toast },
  data: () => {
    return {
      snackbar: false,
      text: "",
      confirmBtnLoading: false,
    };
  },
  methods: {
    removeFromCart(itemId) {
      this.$store.dispatch("removeCartItem", itemId);
    },
    createPurchase() {
      this.confirmBtnLoading = true;
      const userId = authService.getJwtField("userId");
      let purchasePayload = { user_id: userId, purchase_items: [] };
      for (const item of this.$store.state.cart) {
        let newItem = {};
        newItem.name = item.name;
        newItem.drink_id = item.id;
        newItem.num_items = item.amount;
        newItem.unit_price = item.price;

        purchasePayload.purchase_items.push(newItem);
      }
      purchaseService
        .createPurchase(purchasePayload)
        .then((_) => {
          this.confirmBtnLoading = false;
          this.text = "Purchase created successfully.";
          this.snackbar = true;
          this.$store.dispatch("emptyCart");
        })
        .catch((_) => {
          this.confirmBtnLoading = false;
          this.text = "An error occurred while creating purchase.";
          this.snackbar = true;
        });
    },
    updateAmount() {},
  },
};
</script>
