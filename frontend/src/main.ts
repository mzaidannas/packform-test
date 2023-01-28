import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { VueQueryPlugin } from 'vue-query';
import VueAwesomePaginate from 'vue-awesome-paginate';

import App from './App.vue';
import router from './router';

import './assets/main.css';
import 'mosha-vue-toastify/dist/style.css';
import 'vue-awesome-paginate/dist/style.css';

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(VueQueryPlugin);
app.use(VueAwesomePaginate);

app.mount('#app');
