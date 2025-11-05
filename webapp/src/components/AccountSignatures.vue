<script setup>
import { ref } from 'vue'
import axios from 'axios'


const headers = [
    {
        title: 'Signature',
        key: 'signature',
        align: 'start',
        maxWidth: 200,
        nowrap: true
    },
    { 
        title: 'Block',
        key: 'slot',
        align: 'center'
    },
    {
        title: 'Block Time',
        key: 'time',
    }
]

const serverItems = ref([])
const loading = ref(true)

async function loadItems() {
    loading.value = true;      
    const response = await axios.post("api/getAccountSignatures", { "value": null });
    const data = await response.data?.result ?? [];
    for (let i in data) {
        const time = new Date(0);
        time.setUTCSeconds(data[i].blockTime)
        data[i].time = time.toUTCString();
    }
    serverItems.value = data;
    loading.value = false
}

async function expandRow(event, row) {
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

    row.toggleExpand(row.internalItem)
    const response = await axios.post("api/GetSignatureTransaction", { value: row.item.signature })
    const transaction = response?.data.result ?? {};
    row.item.transaction = transaction;
    row.item.show = true;
}

async function loadMore() {
    loading.value = true;
    const response = await axios.post("api/getAccountSignatures", { "value": serverItems.value.slice(-1)[0].signature });
    const data = await response.data?.result ?? [];
    serverItems.value.push(...data)
    loading.value = false
}

</script>

<template>
  <v-container class="pa-md-12">
    <section>
      <h3
        class="d-flex justify-space-between align-center text-subtitle-1 font-weight-bold"
      >
        Last signatures
      </h3>
      <v-data-table-server
        class="bg-transparent"
        :headers="headers"
        :items="serverItems"
        :items-length="totalItems"
        :loading="loading"
        :search="search"
        item-value="signature"
        show-expand
        hideDefaultFooter
        
        @update:options="loadItems"
        @click:row="expandRow"
      >
         <template v-slot:footer>
          <v-toolbar>
            <v-toolbar-title>Expandable Table</v-toolbar-title>
          </v-toolbar>
           
        </template>
        <template v-slot:expanded-row="{ _, item }">
          <v-expand-transition>
            <v-card
                v-show="item.show"
                class="mx-auto bg-secondary"
              >
              <div class="container">
                <div class="columns mt-1 is-align-items-center ">
                    <div class="column">
                        <v-text>Fee {{ item.transaction?.Meta?.Fee }}</v-text>
                    </div>
                    <div class="column">
                        <v-text>Pre balance {{ item.transaction?.Meta?.PreBalances[0] }}</v-text>
                    </div>
                    <div class="column">
                        <v-text>Post balance {{ item.transaction?.Meta?.PostBalances[0] }}</v-text>
                    </div>
                </div>
              </div>
            </v-card>
          </v-expand-transition>
        </template>
      
      </v-data-table-server>
       <button @click="loadMore" class="button is-dark is-info is-fullwidth">
        LOAD MORE
      </button>
    </section>
  </v-container>
</template>

<style scoped>
.v-table__wrapper {
    overflow-x: hidden;
}
</style>