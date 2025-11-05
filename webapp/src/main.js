import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

import axios from 'axios'
import Buefy from 'buefy'

import 'buefy/dist/buefy.css';
import 'material-design-icons-iconfont/dist/material-design-icons.css';
import '@mdi/font/css/materialdesignicons.css';

const app = createApp(App)
app.config.globalProperties.$axios = axios
axios.interceptors.request.use((config) => {
    config.headers['Authorization'] = "test";
  return config
})
import { createUmi } from '@metaplex-foundation/umi-bundle-defaults'
import { mplTokenMetadata } from '@metaplex-foundation/mpl-token-metadata'

const umi = createUmi('https://mainnet.helius-rpc.com/?api-key=193c93fe-5752-487d-9749-8130b7b13c47',
        { "apiKey": "193c93fe-5752-487d-9749-8130b7b13c47" })
    .use(mplTokenMetadata())
/*Vue.http.interceptors.push(
    {
        request: function (request) {
            request.headers['Authorization'] = "test";
            return request
        },
        response: function (response) {
            return response;
        }
    }
);*/
const vuetify = createVuetify({
    components,
    directives,
    theme: {
        defaultTheme: 'dark'
    },
    iconfont: 'mdi',
})
app.provide("umi", umi);
app.use(vuetify)
    .use(mplTokenMetadata)
    .use(umi)
    .use(Buefy)
    .mount("#app");
