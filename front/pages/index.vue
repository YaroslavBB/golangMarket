<template lang="pug">
  section.section
    ProductsTable
</template>
<script type="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import ProductsTable from '../components/ProductsTable.vue'
@Component({
  components: { ProductsTable },
})
export default class extends Vue {}
</script>
