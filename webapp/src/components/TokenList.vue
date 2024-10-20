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
  },
  {
    title: '',
    key: 'actions',
  },
]

const itemsPerPage = ref(5)
const serverItems = ref([])
const loading = ref(true)
const totalItems = ref(0)
const umi = getCurrentInstance().appContext.provides.umi;

async function loadItems() {
  loading.value = true;      
  const response = await axios.post("api/getTokens", arguments?.length > 0 ? arguments[0] : null);
  const data = await response.data
  for (let i in data) {
    const tokenInfo = await TokenDescriptionList.getTokenInfoByMintAsync(umi, data[i].Mint, "5C5oKvXWWA9T2GtG5rHsp4xQoF4G93TqwbWJ91Po3Ric");
    data[i].Symbol = tokenInfo?.symbol;
    data[i].UIAmount = data[i].Amount / (tokenInfo.decimals === 0 ? 1 : Math.pow(10, tokenInfo.decimals));
    data[i].UsdAmount = 0;

  }
  serverItems.value = data;
  totalItems.value = data.length;
  loading.value = false
}

function handleClick(event, row) {
  if (row.item.show) {
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
  }, 1200);
}

onMounted(function () {
})

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
        @update:options="loadItems"
        @click:row="handleClick"
      >
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
          <tbody name="fade" is="transition-group">
            <template >
            <tr class="row" v-for="(item, index) in props.items" :key="index">
              <td>{{item.title}}</td>
              <td>{{item.description}}</td>
              <td>{{item.status}}</td>
              <td>{{item.start}}</td>
              <td>{{item.end}}</td>
            </tr>
          </template>
        </tbody>
        </template> -->
        <template  v-slot:expanded-row="{ columns, item }">
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
    </section>
  </v-container>
</template>
