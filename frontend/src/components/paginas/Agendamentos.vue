<template>
    <div class="d-flex justify-content-center tipografia-tab">
        <b-alert 
            v-if="!isLoading && agedamentos.length === 0" 
            show
            variant="info">
            Não existem Agendamentos
        </b-alert>
        <b-table 
        v-if="!isLoading && agedamentos.length > 0" 
        responsive
        hover 
        :fields="fields"
        :items="agedamentos">
        </b-table>
        <div v-if="isLoading"  class="spinner-grow text-warning" role="status">
            <span class="sr-only">Carregando...</span>
        </div>
        <CriarAgendamento/>
        <DeletarAgendamento/>
    </div>
</template>

<script>
import repo from "./../../repositorio/repositorioDeAgedamentos.js";
import CriarAgendamento from "./../forms/criar/CriarAgendamento";
import DeletarAgendamento from "./../forms/deletar/DeletarAgendamento";

export default {
  name: 'Agedamentos',
  components: {
    CriarAgendamento,
    DeletarAgendamento,
  },
  data() {
    return {
      isLoading: false,
      fields: [ 'id', 'inicio', 'fim', 'professor', 'sala', 'status' ],
      agedamentos: [],
    };
  },
  created () {
    this.fetch()
    //setInterval(this.fetch, 10000)
  },
  destroyed () {
    clearInterval(this.fetch)
  },
  methods: {
    async fetch () {
      this.isLoading = true
      const { data } = await repo.retornaAgenda()
      this.isLoading = false
      this.agedamentos = data;
    }
  },
  computed: {
    computedSalas () {
      return this.agedamentos.slice(0, 10)
    }
  }
}
</script>

<style>
</style>