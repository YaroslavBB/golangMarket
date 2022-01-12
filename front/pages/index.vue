<template>
  <section class="section">
    <ProductsTable></ProductsTable>
  </section>
</template>
<script type="ts">
import ProductsTable from '../components/ProductsTable.vue'
import { Component, Vue } from 'nuxt-property-decorator'
@Component({
  components: { ProductsTable }
})
export default class Index{}
</script>
