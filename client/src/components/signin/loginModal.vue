<template>
    <div class="modal-card">
        <header class="modal-card-head">
            <p class="modal-card-title">Login</p>
        </header>
        <section class="modal-card-body">
            <b-field label="Email">
                <b-input
                        type="email"
                        v-model="email"
                        placeholder="Your email"
                        required>
                </b-input>
            </b-field>

            <b-field label="Password">
                <b-input
                        type="password"
                        v-model="password"
                        password-reveal
                        placeholder="Your password"
                        required>
                </b-input>
            </b-field>

            <b-checkbox>Remember me</b-checkbox>
        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="$parent.close()">Close</button>
            <button class="button is-primary" @click="signUp()" >Login</button>
        </footer>
    </div>
</template>

<script>
    export default {
        name:"login-modal",
        // props: ['email', 'password'],
        data(){
            return {
                email: "",
                password: ""
            }
        },
        methods:{
            signUp(){
                console.log(this.email)
                console.log(this.password)
            }
        }
    }
</script>

<style scoped>
    .modal-card {
        width: auto;
    }
</style>