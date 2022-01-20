<template lang="pug">
  section.section
    LoginForm
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import LoginForm from '../../components/LoginForm.vue'
@Component({
  components: { LoginForm },
})
export default class Login extends Vue {}
</script>

<style>
</style>