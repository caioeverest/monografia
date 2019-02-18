<template>
  <div>
    <b-modal ref="modal" id="deletar" size="lg" title="Deletar sala" hide-footer>
      <b-form @submit="onSubmit" v-if="show">
        <b-form-group id="deletar-sala"
                    label="DeletarSala:"
                    label-for="DeleteSalaInput">
            <b-form-select id="DeleteSalaInput"
                        :options="salas"
                        required
                        v-model="identificador">
            </b-form-select>
        </b-form-group>
        <div>
            <div class="row">
                <div class="col-6">
                    <p>Identificador: {{identificador}}</p>
                </div>
                <div class="col-6">
                    <p>Campus: {{campus}}</p>
                </div>
            </div>
            <div class="row">
                <div class="col-6">
                    <p>Bloco: {{bloco}}</p>
                </div>
                <div class="col-6">
                    <p>Numero: {{numero}}</p>
                </div>
            </div>
        </div>
        <b-button  type="submit" 
                variant="danger"
                v-b-tooltip.hover title="Isso ira deletar tambem os agendamentos vinculado a sala">Deletar</b-button>
      </b-form>
    </b-modal>
  </div>
</template>

<script>
import repoSalas from "./../../../repositorio/repositorioDeSalas.js";

export default {
  name: 'DeletarSala',
  data () {
    return {
      identificador: '',
      campus: '',
      bloco: '',
      numero: '',
      salas: [],
      show: true,
    }
  },
  created () {
    this.buscarSalas();
  },
  watch: {
    identificador: async function (val) {
      const dados = await this.retornarDadosDeSala(val);
      this.campus = dados.campus;
      this.bloco = dados.bloco;
      this.numero = dados.numero;
    },
  },
  methods: {
    async onSubmit (evt) {
      this.$refs.modal.hide()
      evt.preventDefault();
      try {
        await repoSalas.deletarSala(this.identificador);
        // eslint-disable-next-line
        M.toast({html: `Sala deletada!`});
        await setTimeout(() => location.reload(), 3000);
      } catch(error) {
        // eslint-disable-next-line
        M.toast({
          html: `Algo deu errado ${error}`
          });
      }      
    },
    async retornarDadosDeSala (id) {
        const { data } = await repoSalas.retornaUmaSalaPeloId(id);
        return data;
    },
    async buscarSalas () {
        const { data } = await repoSalas.retornaIdentificadorDasSalas();
        this.salas = data;
    }
  }
}
</script>