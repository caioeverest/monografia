import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const Agendamentos = () => import(`@/components/paginas/Agendamentos.vue`);
const Salas = () => import(`@/components/paginas/Salas.vue`);

export default new VueRouter({
  routes: [
    {
      name: `agendamentos`,
      path: `/`,
      component: Agendamentos,
    },
    {
      name: `salas`,
      path: `/salas`,
      component: Salas,
    },
  ],
  mode: `history`,
});