import repo from "./repositorio";

const endpoint = "/agenda";
export default {
    retornaAgenda() {
        return repo.get(`${endpoint}`);
    },
    
    retornaUmaAgendamentoPeloId(id) {
        return repo.get(`/agendamento/${id}`);
    },
    
    criaAgendamento(payload) {
        return repo.post(`/agendar`, payload);
    },

    deletar(id) {
        return repo.delete(`/agendamento/${id}`);
    },
};