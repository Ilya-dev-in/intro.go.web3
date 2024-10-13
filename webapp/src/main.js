import './assets/main.css'
import { registerPlugins } from '@/plugins'

import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)
registerPlugins(app);
createApp(App).mount('#app')
