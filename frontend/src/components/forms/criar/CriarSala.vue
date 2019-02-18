<template>
  <div>
    <b-modal ref="modal" id="criar" size="lg" title="Criar nova sala" hide-footer>
      <b-form @submit="onSubmit" @reset="onReset" v-if="show">
        <b-form-row>
          <div class="col 2">
            <b-form-group id="campus"
                      label="Campus:"
                      label-for="campusInput">
              <b-form-select id="campusInput"
                            :options="todosOsCampus"
                            required
                            v-model="form.campus">
              </b-form-select>
            </b-form-group>
          </div>
          <div class="col 2">
            <b-form-group id="bloco"
                      label="Bloco:"
                      label-for="blocoInput">
              <b-form-select id="blocoInput"
                            :options="blocos"
                            required
                            v-model="form.bloco">
              </b-form-select>
            </b-form-group>
          </div>
        </b-form-row>
        <b-form-row>
          <div class="col 2">
            <b-form-group id="identificador"
                        label="Identificador:"
                        label-for="identificadorInput">
              {{form.identificador}}
            </b-form-group>
          </div>
          <div class="col 2">
            <b-form-group id="numero"
                        label="Numero da sala:"
                        label-for="numeroInput">
              <b-form-input id="numeroInput"
                            type="text"
                            v-model="form.numero"
                            required
                            placeholder="123">
              </b-form-input>
            </b-form-group>
          </div>
        </b-form-row>
        <b-button type="submit" variant="primary">Criar!</b-button>
        <b-button type="reset" variant="danger">Limpar</b-button>
      </b-form>
    </b-modal>
  </div>
</template>

<script>
import repoSalas from "./../../../repositorio/repositorioDeSalas.js";

export default {
  name: 'CriarSala',
  data () {
    return {
      form: {
        identificador: '',
        numero: '',
        bloco: [],
        campus: []
      },
      blocos: [
        { text: 'Selecione um bloco', value: null },
        'A', 'B', 'C', 'D'
      ],
      todosOsCampus: [
        { text: 'Selecione um campus', value: null },
        'TIJUCA', 'CABO-FRIO', 'BARRA', 'CENTRO'
      ],
      show: true,
    }
  },
  computed: {
    campus() {
      return this.form.campus;
    },
    bloco() {
      return this.form.bloco;
    },
    numero() {
      return this.form.numero;
    },
    identificador() {
      return this.form.identificador;
    },
    retornaChave(camp) {
      return this.retornaChaveCampus(camp)
    }
  },
  watch: {
    campus: function (val) {
      this.form.identificador = val.replace(/[aeiou]/gi, '').slice(0,2) + this.form.bloco + this.form.numero;
    },
    bloco: function (val) {
      this.form.identificador = this.form.campus.replace(/[aeiou]/gi, '').slice(0,2) + val + this.form.numero;
    },
    numero: function (val) {
      this.form.identificador = this.form.campus.replace(/[aeiou]/gi, '').slice(0,2) + this.form.bloco + val;
    }
  },
  methods: {
    async onSubmit (evt) {
      this.$refs.modal.hide()
      evt.preventDefault();
      try {
        await repoSalas.criaSala(this.form);
        // eslint-disable-next-line
        M.toast({html: `Criado com sucesso`});
        await setTimeout(() => location.reload(), 3000);
      } catch(error) {
        // eslint-disable-next-line
        M.toast({
          html: `Algo deu errado na criação da sala`
          });
      }
    },
    onReset (evt) {
      evt.preventDefault();
      this.form.numero = '';
      this.form.bloco = null;
      this.form.campus = null;
      this.form.identificador = '';
      this.show = false;
      this.$nextTick(() => { this.show = true });
    },
    // format (value, event) {
    //   return value.toUpperCase();
    // },
  }
}
</script>