<template>
  <div>
    <b-modal ref="modal" id="deletar" size="lg" title="Deletar agendamento" hide-footer>
      <b-form @submit="onSubmit" v-if="show">
        <b-form-group id="deletar-agendamento"
                    label="DeletarAgendamento"
                    label-for="DeleteAgendamentoInput">
            <b-form-select id="DeleteAgendamentoInput"
                        :options="agendamentos"
                        required
                        v-model="id">
            </b-form-select>
        </b-form-group>
        <div>
            <div class="col-6">
                <p>Id: {{id}}</p>
            </div>
            <div class="row">
                <div class="col-6">
                    <p>Professor: {{professor}}</p>
                </div>
                <div class="col-6">
                    <p>Sala: {{sala}}</p>
                </div>
            </div>
            <div class="row">
                <div class="col-6">
                    <p>Inicio: {{inicio}}</p>
                </div>
                <div class="col-6">
                    <p>Fim: {{fim}}</p>
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
import repo from "./../../../repositorio/repositorioDeAgedamentos.js";

export default {
  name: 'DeletarSala',
  data () {
    return {
      id: '',
      professor: '',
      sala: '',
      inicio: '',
      fim: '',
      agendamentos: [],
      show: true,
    }
  },
  created () {
    this.buscarAgenda();
  },
  watch: {
    id: async function (val) {
      const { data } = await repo.retornaUmaAgendamentoPeloId(val);
      this.professor = data.professor;
      this.sala = data.sala;
      this.inicio = data.inicio;
      this.fim = data.fim;
    },
  },
  methods: {
    async onSubmit (evt) {
      this.$refs.modal.hide()
      evt.preventDefault();
      try {
        await repo.deletar(this.id);
        // eslint-disable-next-line
        M.toast({html: `Agendamento deletada!`});
        await setTimeout(() => location.reload(), 3000);
      } catch(error) {
        // eslint-disable-next-line
        M.toast({
          html: `Algo deu errado ${error}`
          });
      }      
    },
    async buscarAgenda () {
        const { data } = await repo.retornaAgenda();
        data.forEach(agendamento => {
          this.agendamentos.push(agendamento.id);
        });
    }
  }
}
</script>