<template>
    <b-nav vertical id="sidebar">
        <div class="sidebar-header">
            <img alt="Logo" class="d-inline-block align-top" src="https://www.uva.br/sites/default/files/logos/logo-footer.png">
        </div>
        <b-navbar-nav class="mx-auto">
            <div class="row">
                <b-btn v-b-modal.criar variant="outline-light" class="crud-item">Criar <criar-icon /></b-btn>
                <b-btn v-b-modal.deletar variant="outline-light" class="crud-item">Apagar <apagar-icon /></b-btn>
            </div>
        </b-navbar-nav>
        <b-nav-item to="/">Agendamentos</b-nav-item>
        <b-nav-item to="/salas">Salas</b-nav-item>
    </b-nav>
</template>

<script>
import CriarIcon from "vue-material-design-icons/Plus.vue"
import ApagarIcon from "vue-material-design-icons/DeleteForever.vue"

export default {
  name: 'Sidebar',
  components: {
      CriarIcon,
      ApagarIcon
  }
}
</script>

<style scoped>
#sidebar {
    min-width: 250px;
    max-width: 250px;
    background: #194B75;
    color: rgb(48, 48, 48);
    transition: all 0.6s cubic-bezier(0.945, 0.020, 0.270, 0.665);
    transform-origin: bottom left;
}

.crud-item {
    margin: 0.5rem;
}

a {
    color: #fefefe;
}

a:hover {
    color: #D6588E;
}

a:focus {
    color: #87A5B9;
    text-decoration: none;
    transition: all 0.3s;
}

#sidebar.active {
    margin-left: -250px;
    transform: rotateY(100deg);
}

#sidebar .sidebar-header {
    padding: 20px;
    background: #194B75;
}

#sidebar ul.components {
    padding: 20px 0;
    border-bottom: 1px solid #47748b;
}

#sidebar ul p {
    color: #fff;
    padding: 10px;
}

#sidebar ul li a {
    padding: 10px;
    font-size: 1.1em;
    display: block;
}
#sidebar ul li a:hover {
    color: #7386D5;
    background: #fff;
}

#sidebar ul li.active > a, a[aria-expanded="true"] {
    color: #fff;
    background: #6d7fcc;
}


a[data-toggle="collapse"] {
    position: relative;
}

.dropdown-toggle::after {
    display: block;
    position: absolute;
    top: 50%;
    right: 20px;
    transform: translateY(-50%);
}

ul ul a {
    font-size: 0.9em !important;
    padding-left: 30px !important;
    background: #6d7fcc;
}

ul.CTAs {
    padding: 20px;
}

ul.CTAs a {
    text-align: center;
    font-size: 0.9em !important;
    display: block;
    border-radius: 5px;
    margin-bottom: 5px;
}

a.download {
    background: #fff;
    color: #7386D5;
}

a.article, a.article:hover {
    background: #6d7fcc !important;
    color: #fff !important;
}

@media (max-width: 768px) {
    #sidebar {
        margin-left: -250px;
        transform: rotateY(90deg);
    }
    #sidebar.active {
        margin-left: 0;
        transform: none;
    }
    #sidebarCollapse span:first-of-type,
    #sidebarCollapse span:nth-of-type(2),
    #sidebarCollapse span:last-of-type {
        transform: none;
        opacity: 1;
        margin: 5px auto;
    }
    #sidebarCollapse.active span {
        margin: 0 auto;
    }
    #sidebarCollapse.active span:first-of-type {
        transform: rotate(45deg) translate(2px, 2px);
    }
    #sidebarCollapse.active span:nth-of-type(2) {
        opacity: 0;
    }
    #sidebarCollapse.active span:last-of-type {
        transform: rotate(-45deg) translate(1px, -1px);
    }

}
</style>