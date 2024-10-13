<script setup>
import { createAppKit } from '@reown/appkit'
import { SolanaAdapter } from '@reown/appkit-adapter-solana'

import { solana, solanaTestnet, solanaDevnet } from '@reown/appkit/networks'
import { mainnet, arbitrum, sepolia } from '@reown/appkit/networks'

import { SolflareWalletAdapter, PhantomWalletAdapter } from '@solana/wallet-adapter-wallets'

const solanaWeb3JsAdapter = new SolanaAdapter({
  wallets: [new PhantomWalletAdapter(), new SolflareWalletAdapter()]
})
const projectId = '0eb6b5fca832a86275177e9fa462eaca'

const metadata = {
  name: 'AppKit',
  description: 'AppKit Example',
  url: 'http://localhost:5173/', // origin must match your domain & subdomain
  icons: ['https://avatars.githubusercontent.com/u/179229932']
}

const modal = createAppKit({
  adapters: [solanaWeb3JsAdapter],
  networks: [mainnet, arbitrum, sepolia, solana, solanaTestnet, solanaDevnet],
  metadata,
  projectId,
  features: {
    analytics: true,
  }
});

modal.subscribeEvents(function () {
  //this.$emit("onWalletConnected", modal);
})

</script>

<template>
  <w3m-button />
  <w3m-network-button />
  <w3m-connect-button />
  <w3m-account-button />

  <button @click="modal.open()">Open Connect Modal</button>
  <button @click="modal.open({ view: 'Networks' })">Open Network Modal</button>
  <button @click="setThemeMode(themeMode === 'dark' ? 'light' : 'dark')">Toggle Theme Mode</button>
  <button @click="this.$emit('onWalletConnected',modal)">Toggle Theme Mode</button>
  <pre>{{ JSON.stringify(state, null, 2) }}</pre>
  <pre>{{ JSON.stringify({ themeMode, themeVariables }, null, 2) }}</pre>
  <!-- <pre>{{ JSON.stringify(events, null, 2) }}</pre> -->
</template>