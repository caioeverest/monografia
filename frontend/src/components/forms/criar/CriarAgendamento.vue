<template>
  <div>
    <b-modal ref="modal" id="criar" size="lg" title="Criar novo agendamento" hide-footer>
      <b-form @submit="onSubmit" @reset="onReset" v-if="show">
        <b-form-row>
          <div class="col 2">
            <b-form-group id="professor"
                        label="Professor solicitante:"
                        label-for="professorInput">
              <b-form-input id="professorInput"
                            type="text"
                            v-model="form.professor"
                            required
                            placeholder="Fulano da Silva">
              </b-form-input>
            </b-form-group>
          </div>
          <div class="col 2">
            <b-form-group id="salas"
                      label="Salas:"
                      label-for="salasInput">
              <b-form-select id="salasInput"
                            :options="salas"
                            required
                            v-model="form.sala">
              </b-form-select>
            </b-form-group>
          </div>
        </b-form-row>
        <b-form-group id="inicio"
                    label="Data e hora do inicio da aula:"
                    label-for="inicioInput">
            <date-picker v-model="form.inicio" :config="options" required/>
        </b-form-group>
        <b-form-group id="fim"
                    label="Data e hora do fim da aula:"
                    label-for="fimInput">
            <date-picker v-model="form.fim" :config="options" required/>
        </b-form-group>
        <b-button type="submit" variant="primary">Criar!</b-button>
        <b-button type="reset" variant="danger">Limpar</b-button>
      </b-form>
    </b-modal>
  </div>
</template>

<script>
import repoAgendamentos from "./../../../repositorio/repositorioDeAgedamentos.js";
import repoSalas from "./../../../repositorio/repositorioDeSalas.js";
import datePicker from 'vue-bootstrap-datetimepicker';
import moment from 'moment';

export default {
  name: 'CriarAgendamento',
  components: {
    datePicker,
  },
  data () {
    return {
      form: {
        professor: '',
        sala: '',
        inicio: moment(),
        fim: moment()
      },
      date: moment(),
        options: {
          locale: "pt-br",
          useCurrent: true,
      },
      salas: [],
      show: true,
    }
  },
  computed: {
    inicio() {
      return this.form.inicio;
    },
    fim() {
      return this.form.fim;
    }
  },
  watch: {
    inicio: function (val) {
      this.salasCaller(val, this.form.fim);
    },
    fim: function (val) {
      this.salasCaller(this.form.inicio, val);
    },
  },
  created () {
    // eslint-disable-next-line
    $.extend(true, $.fn.datetimepicker.defaults, {
      icons: {
        time: 'far fa-clock',
        date: 'far fa-calendar',
        up: 'fas fa-arrow-up',
        down: 'fas fa-arrow-down',
        previous: 'fas fa-chevron-left',
        next: 'fas fa-chevron-right',
        today: 'fas fa-calendar-check',
        clear: 'far fa-trash-alt',
        close: 'far fa-times-circle'
      }
    });
  },
  methods: {
    async salasCaller (inicio, fim) {
      const { data } = await repoSalas.retornaIdentificadorDeSalasDisponiveis(inicio, fim);
      this.salas = data;
    },
    async onSubmit (evt) {
      this.$refs.modal.hide()
      evt.preventDefault();
      try {
        await repoAgendamentos.criaAgendamento(this.form);
        // eslint-disable-next-line
        M.toast({html: `Criado com sucesso`});
        await setTimeout(() => location.reload(), 3000);
      } catch(error) {
        // eslint-disable-next-line
        M.toast({
          html: `Algo deu errado na criação do agendamento`
          });
      }
    },
    onReset (evt) {
      evt.preventDefault();
      
      this.form.professor = '';
      this.form.sala = '';
      this.form.inicio = null;
      this.form.fim = null;
      
      this.show = false;
      this.$nextTick(() => { this.show = true });
    }
  },
}
</script>