<template>
  <v-container>
    <v-row align="center">
      <v-col cols="12">
        <v-flex class="text-center">
          <h1>Purchase history</h1>
        </v-flex>
      </v-col>
    </v-row>
    <v-row v-if="purchases.length > 0">
      <v-row align="center" v-for="purchase in purchases" :key="purchase.id">
        <h3>
          Purchase made at:
          {{ purchase.purchase_items[0].purchase_datetime | formatDate }}
        </h3>
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-left">Drink</th>
                <th class="text-left">Amount</th>
                <th class="text-left">Total price</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in purchase.purchase_items" :key="item.id">
                <td>
                  <router-link
                    :to="{ name: 'Drink', params: { id: item.drink_id } }"
                    >View drink</router-link
                  >
                </td>
                <td>
                  {{ item.num_items }}
                </td>
                <td>
                  {{ item.num_items * item.unit_price }}
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
        <v-divider />
      </v-row>
    </v-row>
    <v-row v-else>
      <v-col cols="12">
        <v-flex class="text-center">
          <h2>Purchase history is empty.</h2>
        </v-flex>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { authService } from "../../services/authService";
import { purchaseService } from "../../services/purchaseService";
export default {
  name: "PurchaseHistory",
  data: () => {
    return {
      purchases: [],
    };
  },
  mounted() {
    let userId = authService.getJwtField("userId");
    purchaseService.getUserHistory(userId).then((response) => {
      this.purchases = response.data;
      console.log(this.purchases);
    });
  },
};
</script>
