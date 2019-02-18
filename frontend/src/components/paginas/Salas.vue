<template>
  <div class="d-flex justify-content-center tipografia-tab position-relative">
    <b-alert 
            v-if="!isLoading && salas.length === 0" 
            show
            variant="info">
            Não existem salas registradas
        </b-alert>
    <b-table 
      v-if="!isLoading" 
      responsive
      hover 
      :fields="fields"
      :items="salas">
    </b-table>
    <div v-if="isLoading"  class="spinner-grow text-success" role="status">
      <span class="sr-only">Carregando...</span>
    </div>
    <CriarSala />
    <DeletarSala />
  </div>
</template>

<script>
import repo from "./../../repositorio/repositorioDeSalas.js";
import CriarSala from "./../forms/criar/CriarSala";
import DeletarSala from "./../forms/deletar/DeletarSala";

export default {
  name: 'Salas',
  components: {
    CriarSala,
    DeletarSala
  },
  data() {
    return {
      isLoading: false,
      fields: [ 'identificador', 'bloco', 'campus', 'numero'],
      salas: [],
    };
  },
  created () {
    this.fetch();
    //setInterval(this.fetch, 10000)
  },
  destroyed () {
    clearInterval(this.fetch);
  },
  methods: {
    async fetch () {
      this.isLoading = true;
      const { data } = await repo.retornaSalas();
      this.isLoading = false;
      this.salas = data;
    }
  },
  computed: {
    computedSalas () {
      return this.salas.slice(0, 10);
    }
  }
}
</script>

<style>
</style>