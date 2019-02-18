<template>
  <div id="Orquestrador">
    <SideBar class="shadow-lg"/>
    <div id="content">
        <nav class="navbar navbar-expand-lg navbar-light bg-light">
                <div class="container-fluid">
                    <button type="button" id="sidebarCollapse" class="navbar-btn">
                        <span></span>
                        <span></span>
                        <span></span>
                    </button>

                    <div class="collapse navbar-collapse relogio">
                        <ul class="nav navbar-nav ml-auto">
                            <li class="nav-item active">
                                <a>{{relogio}}</a>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        <router-view/>
    </div>
  </div>
</template>

<script>
import SideBar from './SideBar'
import moment from 'moment';
import JQuery from 'jquery'
let $ = JQuery

export default {
  name: 'Orquestrador',
  components: {
    SideBar
  },
  data () {
      return {
          relogio: moment().locale('pt-br').format('LLL'),
      }
  },
  created: function() {
      setInterval(() => this.relogio = moment().locale('pt-br').format('LLL'), 1000);
  },
  mounted: function() {
        $(document).ready(function () {
            $('#sidebarCollapse').on('click', function () {
                $('#sidebar').toggleClass('active');
                $(this).toggleClass('active');
            });
        });
    },
}
</script>

<style>
.relogio {
    color: #afacac
}

.navbar {
    padding: 15px 10px;
    background: #fff;
    border: none;
    border-radius: 0;
    margin-bottom: 40px;
    box-shadow: 1px 1px 3px rgba(0, 0, 0, 0.1);
}

.navbar-btn {
    box-shadow: none;
    outline: none !important;
    border: none;
}

#Orquestrador {
  color: #2c3e50;
}

#content {
    width: 100%;
    padding: 20px;
    min-height: 100vh;
    transition: all 0.3s;
}

#sidebarCollapse {
    width: 40px;
    height: 40px;
    background: #f5f5f5;
    cursor: pointer;
}

#sidebarCollapse span {
    width: 80%;
    height: 2px;
    margin: 0 auto;
    display: block;
    background: #555;
    transition: all 0.8s cubic-bezier(0.810, -0.330, 0.345, 1.375);
    transition-delay: 0.2s;
}

#sidebarCollapse span:first-of-type {
    transform: rotate(45deg) translate(2px, 2px);
}
#sidebarCollapse span:nth-of-type(2) {
    opacity: 0;
}
#sidebarCollapse span:last-of-type {
    transform: rotate(-45deg) translate(1px, -1px);
}


#sidebarCollapse.active span {
    transform: none;
    opacity: 1;
    margin: 5px auto;
}
</style>
