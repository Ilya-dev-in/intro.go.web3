<script setup>
import { createAppKit } from '@reown/appkit'
import { SolanaAdapter } from '@reown/appkit-adapter-solana'

import { solana } from '@reown/appkit/networks'
import { mainnet, arbitrum, sepolia } from '@reown/appkit/networks'

import { SolflareWalletAdapter, PhantomWalletAdapter } from '@solana/wallet-adapter-wallets'

const solanaWeb3JsAdapter = new SolanaAdapter({
  wallets: [new PhantomWalletAdapter(), new SolflareWalletAdapter()]
})
const projectId = 'YOUR_PROJECT_ID';

const metadata = {
  name: 'AppKit',
  description: 'AppKit Example',
  url: 'http://localhost:5173/',
  icons: ['https://avatars.githubusercontent.com/u/179229932']
}

const modal = createAppKit({
  adapters: [solanaWeb3JsAdapter],
  networks: [mainnet, arbitrum, sepolia, solana],
  metadata,
  projectId,
  features: {
    analytics: true,
  }
});

const emit = defineEmits(['onWalletConnected'])
modal.subscribeEvents(function () {
   console.log(modal.getEvent().data.event)
})
modal.subscribeState(function () {
  emit("onWalletConnected", modal.getAddress());
})


</script>

<template>
  <div>
    <w3m-button />
    <!-- <w3m-network-button /> -->
  </div>
</template>