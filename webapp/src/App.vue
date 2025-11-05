<script setup>
import { render, h, ref } from 'vue';
import WalletConnect from './components/WalletConnect.vue'
import TokenList from './components/TokenList.vue';
import AccountSignatures from './components/AccountSignatures.vue';

const back = { apiName: ""};
function onWalletConnected(address) {
  console.log(address)
  //render(TokenList, document.getElementById("main-grid"));
}

const tab = ref(0);
</script>

<template>
  <div class="has-text-grey-light is-align-items-center is-justify-content-flex-start">
    <header>
        <div class="logo">
          <img alt="Vue logo" class="logo" src="/favicon.ico" width="30" height="30" />
          <span v-bind:id="back.apiName">Welcome to Web3 {{ back.apiName }}</span>
        </div>
        <div id="wallet">
          <WalletConnect @onWalletConnected="onWalletConnected($event)" />
        </div>
    </header>
    <main>
      <div class="container">
        <div class="columns mt-1 is-align-items-center ">
          <div class="column overview-box ml-5">First column</div>
          <div class="column overview-box ml-5">Second column</div>
          <div class="column overview-box ml-5">Third column</div>
        </div>
        <v-card>
          <v-tabs
            v-model="tab"
            align-tabs="center"
            bg-color="light-blue-darken-4"
            stacked>
            <v-tab value="0">
            <v-icon icon="mdi-hand-coin-outline"></v-icon>
              Tokens
            </v-tab>
            <v-tab value="1">
              <v-icon icon="mdi-transcribe"></v-icon>
              Signatures
            </v-tab>
             <v-tab value="2">
              <v-icon icon="mdi-progress-check"></v-icon>
              Transactions
            </v-tab>

          </v-tabs>
          <v-tabs-window v-model="tab">
            <v-tabs-window-item :key="0" :value="0">
              <v-card>
                <TokenList />
              </v-card>
            </v-tabs-window-item>

            <v-tabs-window-item :key="1" :value="1">
              <v-card>
                <AccountSignatures />
              </v-card>
            </v-tabs-window-item>

            <v-tabs-window-item :key="2" :value="2">
              <v-card>
                <v-card-text>tab 3</v-card-text>
              </v-card>
            </v-tabs-window-item>

          </v-tabs-window>
        </v-card>
      </div>
    </main>
    <footer>
      <p>Here's some contact info</p>
    </footer>
  </div>
 </template>
<style scoped>
  body {
    display: block;
  }

  header {
      position: sticky;
      width: 100%;
      padding: 10px 50px;
      top: 0;
      z-index: 99;
      background-color: #181818 ;
      max-height: 55px;
      line-height: 1.5;
      border-bottom: 1px solid #0b340c;
    }

    main {
      width: 100%;
      margin-top: 40px;
    }

    .logo {
      font-size: 1.3em;
      color: #273849;
      font-family: "Dosis", "Source Sans Pro", "Helvetica Neue", Arial, sans-serif;
      font-weight: 500;
      display: inline-flex;
      margin: 0 2rem 0 0;
        color: #48a478;
    }

    #wallet {
        float: right;
    }
  
</style>
