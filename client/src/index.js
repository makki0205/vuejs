import Vue from 'vue'
import Buefy from 'buefy'
import VueRouter from 'vue-router'
import Index from './components/index/index.vue'
import Signin from './components/signin/signin.vue'
import Notfound from './components/notfound/notfound.vue'
import Signup from './components/signup/signup.vue'

import 'buefy/lib/buefy.css'

Vue.use(VueRouter)
Vue.use(Buefy)

const routes = [
    { path: "/", component: Index },
    { path: "/signin", component: Signin },
    { path: "/signup", component: Signup },
    { path: "*", component: Notfound },
]

const router = new VueRouter({  mode: 'history', routes })


new Vue({
    router
}).$mount("#app")