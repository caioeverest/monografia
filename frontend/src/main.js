import Vue from 'vue';
import BootstrapVue from 'bootstrap-vue';
import App from './components/App.vue';
import router from './rota';

Vue.use(BootstrapVue);

window.$ = window.jQuery = require('jquery');
require('popper.js');
require('bootstrap');

import 'materialize-css/dist/js/materialize';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import "vue-material-design-icons/styles.css";
import 'pc-bootstrap4-datetimepicker/build/css/bootstrap-datetimepicker.css';
import '@fortawesome/fontawesome-free/css/fontawesome.css';
import '@fortawesome/fontawesome-free/css/regular.css';
import '@fortawesome/fontawesome-free/css/solid.css';
import './assets/style.css'

new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
