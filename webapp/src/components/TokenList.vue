<script setup>
import { ref, onMounted, getCurrentInstance } from 'vue'
import axios from 'axios'
import TokenDescriptionList from '../utils/TokensUtils.js'


const headers = [
  {
    title: 'Mint',
    key: 'Mint',
  },
  {
    title: 'Symbol',
    key: 'Symbol',
  },
  {
    title: 'UIAmount',
    key: 'UIAmount',
  },
  {
    title: 'USD Amount',
    key: 'UsdAmount',
  }
]

const itemsPerPage = ref(5)
const serverItems = ref([])
const loading = ref(true)
const totalItems = ref(0)
const umi = getCurrentInstance().appContext.provides.umi;

async function loadItems() {
  loading.value = true;      
  const response = await axios.post("api/getTokens", arguments?.length > 0 ? arguments[0] : null);
  const data = await response.data?.result ?? [];

  /*const transactionsResponse = await axios.post("api/getTokensTransactions", arguments?.length > 0 ? arguments[0] : null);
  const transactionsData = await transactionsResponse.data?.result ?? [];
  console.log(transactionsData);*/

  for (let i in data) {
    const tokenInfo = await TokenDescriptionList.getTokenInfoByMintAsync(umi, data[i].Mint, "5C5oKvXWWA9T2GtG5rHsp4xQoF4G93TqwbWJ91Po3Ric");
    data[i].Symbol = tokenInfo?.symbol;
    data[i].UIAmount = (data[i].Amount / (tokenInfo.decimals === 0 ? 1 : Math.pow(10, tokenInfo.decimals))).toFixed(4).toString();
    data[i].UsdAmount = 0;
    data[i].image = tokenInfo.image;

  }
  serverItems.value = data;
  totalItems.value = data.length;
  loading.value = false
}

function handleClick(event, row) {
 /*if (row.item.show) {
    row.item.show = false;
     setTimeout(() => {
       row.toggleExpand(row.internalItem)
    }, 100);
    return;
  }

  row.toggleExpand(row.internalItem)
  row.item.loading = true;
  setTimeout(() => {
    row.item.show = true;
  }, 100);
  setTimeout(() => {
    row.item.loading = false;
  }, 1200);*/
}

onMounted(async function () {
  
})

async function GetAssets() {
  const response = await axios.post("/api/GetAssetByMint", { Value: "jtojtomepa8beP8AuQc6eXt5FriJwfFMwQx2v2f9mCL" })
  var data = response.data?.result;
  console.log(data);
}

GetAssets();

</script>

<template>
  <v-container class="pa-md-12">
    <section>
      <h3
        class="d-flex justify-space-between align-center text-subtitle-1 font-weight-bold"
      >
        Your assets

        <!-- <v-btn
          class="text-none"
          color="primary"
          prepend-icon="mdi-plus"
          rounded="lg"
          slim
          text="Create promotion"
          variant="flat"
        /> -->
      </h3>

      <!-- <div class="text-body-2 text-medium-emphasis mb-4 w-100 w-md-75">
        A list of all the promotions that are currently running.
      </div> -->

      <v-data-table-server
        v-model:items-per-page="itemsPerPage"
        class="bg-transparent"
        :headers="headers"
        :items="serverItems"
        :items-length="totalItems"
        :loading="loading"
        :search="search"
        item-value="Mint"
        show-expand
        hideDefaultFooter
        @update:options="loadItems"
        @click:row="handleClick"
      >

      <template v-slot:item.Symbol="{ item }">
          <div class="container is-inline-flex is-align-items-center">
            <v-img
              :src="`${item.image}`"
              height="24"
              width="24"
              cover
            ></v-img>
            <v-text>{{ item.Symbol }}</v-text>
          </div>
      </template>
      <!-- <template v-slot:body="{items, columns }">
        <tr @click="handleClick(this)" v-for="(item, index) in items" :key="index">
          <td v-for="(column, index)  in columns" :key="index">
            <span>
              {{ item[column.value] }}
            </span>
          </td>
          <v-progress-circular v-show="index === 0"
                color="primary"
                indeterminate
          ></v-progress-circular>
        </tr>
      </template> -->
        <!-- <template v-slot:top>
        <v-toolbar name="tr-fade" flat>
          <v-toolbar-title>Expandable Table</v-toolbar-title>
        </v-toolbar>
      </template> -->
        <!-- <template v-slot:body="props">
            <table>
              <v-data-table-row class="row" @click="handleClick" v-for="(item, index) in props.items" :key="index">
                <v-data-table-column>{{item.Mint}}</v-data-table-column>
                <v-data-table-column>{{item.Symbol}}</v-data-table-column>
                <v-data-table-column>{{item.UIAmount}}</v-data-table-column>
                <v-data-table-column>{{item.UsdAmount}}</v-data-table-column>
              </v-data-table-row>
          </table>
        </template>  -->
         <template v-slot:footer>
          <v-toolbar>
            <v-toolbar-title>Expandable Table</v-toolbar-title>
          </v-toolbar>
           
        </template>
        <template v-slot:expanded-row="{ columns, item }">
          <v-expand-transition>
            <v-card
                v-show="item.show"
                class="mx-auto bg-secondary"
                height="35px"
              >
               <p>dfsgdfgd</p>
            </v-card>
          </v-expand-transition>
          <!-- <v-progress-circular v-show="item.loading"
                color="primary"
                indeterminate
          ></v-progress-circular> -->
        </template>
      
      </v-data-table-server>
       <button class="button is-dark is-info is-fullwidth">
        LOAD MORE
      </button>
    </section>
  </v-container>
</template>
