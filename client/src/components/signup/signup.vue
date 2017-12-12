<template>
    <section>
        <button class="button is-primary is-medium" @click="cardModal()">
            Launch card modal (Component)
        </button>
    </section>
</template>

<script>
    import ModalForm from './signupModal.vue'

    export default {
        methods: {
            cardModal() {
                this.$modal.open({
                    parent: this,
                    component: ModalForm,
                    hasModalCard: true
                })
            }
        }
    }
</script>
