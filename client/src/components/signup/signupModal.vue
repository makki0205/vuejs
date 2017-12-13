<template>
    <div class="modal-card">
        <header class="modal-card-head">
            <p class="modal-card-title">SignUp</p>
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

        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="$parent.close()">Close</button>
            <button class="button is-primary" @click="signUp()" >SignUp</button>
        </footer>
    </div>
</template>

<script>
    import user from '../../services/user'
    import validationUtils from '../../lib/ValidationUtils'

    export default {
        name:"login-modal",
        data(){
            return {
                email: "",
                password: ""
            }
        },
        methods:{
            signUp(){
                if (validationUtils.isBlank(this.email)) {
                    validationUtils.pushMessage('メールアドレスガニュウリョクサレテイマセン', this)
                    return
                }
                if (validationUtils.isBlank(this.password)) {
                    validationUtils.pushMessage('パスワードガニュウリョクサレテイマセン', this)
                    return
                }
                user.create(this.email, this.password)
                    .then((response) => {
                        if (response.status === 200) {
                            this.$toast.open({
                                message: '登録が完了しました',
                                type: 'is-success'
                            })
                        }
                    })
                    .catch((error) => {
                        if (error.response.status === 400) {
                            this.$toast.open({
                                duration: 5000,
                                message: `既に使用されているメールアドレスです`,
                                type: 'is-danger'
                            })
                        }
                    })
            },
        }
    }
</script>

<style scoped>
    .modal-card {
        width: auto;
    }
</style>