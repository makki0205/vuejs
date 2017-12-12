<template>
    <section>
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

</template>
<script>
    export default {
        name:"signup",
    }
</script>