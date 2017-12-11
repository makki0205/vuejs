import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from './components/index/index.vue'
import Signin from './components/signin/signin.vue'
import Notfound from './components/notfound/notfound.vue'
 Vue.use(VueRouter)

const routes = [
    { path: "/", component: Index },
    { path: "/signin", component: Signin },
    { path: "*", component: Notfound },
]

const router = new VueRouter({  mode: 'history', routes })


new Vue({
    router
}).$mount("#app")