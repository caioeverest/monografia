import repo from "./repositorio";

export default {
  retornaSalas() {
    return repo.get(`/salas`);
  },

  retornaIdentificadorDeSalasDisponiveis(inicio, fim) {
    return repo.get(`/salas/disponiveis`, { 
      params: {
        inicio,
        fim
      }
    });
  },

  retornaIdentificadorDasSalas() {
    return repo.get(`/salas/identificadores`);
  },

  retornaUmaSalaPeloId(id) {
    return repo.get(`sala/identificador/${id}`);
  },

  criaSala(sala) {
    return repo.post(`/sala`, sala);
  },

  deletarSala(id) {
    return repo.delete(`/sala/${id}`);
  }
};