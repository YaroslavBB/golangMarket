<template lang="pug">
  section.section
    RegistrationForm
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import RegistrationForm from '../../components/RegistrationForm.vue'
@Component({
  components: { RegistrationForm },
})
export default class Logun extends Vue {}
</script>