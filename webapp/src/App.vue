<script setup>
import { render, h } from 'vue';
import Hello from './components/Hello.vue'
import WalletConnect from './components/WalletConnect.vue'
import TokenList from './components/TokenList.vue';

const back = { apiName: ""};
function onWalletConnected(modal) {
  debugger;
  render(TokenList, document.getElementById("main-grid"));
}
</script>

<template>
  <div class="container">
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
      <TokenList />
    </main>
    <footer>
      <p>Here's some contact info</p>
    </footer>
  </div>
 </template>
<style scoped>
header {
  line-height: 1.5;
  border-bottom: 1px solid #0b340c;
}

.logo {
  font-size: 1.3em;
  color: #273849;
  font-family: "Dosis", "Source Sans Pro", "Helvetica Neue", Arial, sans-serif;
  font-weight: 500;
  display: inline-flex;
}

#wallet {
    float: right;
}

@media (min-width: 1024px) {
  header {
    position: fixed;
    width: 100%;
    top: 0;
    /*height: 40px;*/
    padding: 10px 50px;
    z-index: 20;
    background-color: var(--color-background);
    /*place-items: center;
    padding-right: calc(var(--section-gap) / 2);*/
  }

  main {
    margin-top: 45px;
    width: 70%;
    margin-left: 15%;
  }

  .logo {
    margin: 0 2rem 0 0;
    color: #48a478;
  }
/*
  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }*/
}
</style>
