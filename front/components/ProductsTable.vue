<template lang="pug">
section.section
  SearchProduct
  b-field
    b-button(type='is-success', outlined='', @click='openAddProductModal') Добавить товар
  .column.is-full
    b-table(:data='this.allProducts()', fullwidth='')
      b-table-column(
        field='productId',
        label='ID',
        v-slot='props',
        sortable='',
        centered=''
      )
        | {{ props.row.productId }}
      b-table-column(
        field='name',
        label='Наименование продукта',
        v-slot='props',
        sortable='',
        centered=''
      )
        | {{ props.row.name }}
      b-table-column(
        field='dateAdded',
        label='Дата добавления',
        v-slot='props',
        sortable='',
        centered=''
      )
        | {{ formatDate(props.row.dateAdded) }}
      b-table-column(label='', v-slot='props', centered='')
        b-button(
          type='is-primary',
          outlined='',
          @click='openProductFormTableModal(props.row.productId, props.row.name)'
        ) Детали
      b-table-column(label='', v-slot='props', centered='')
        b-button(
          type='is-danger',
          icon-right='delete',
          @click='deleted(props.row.productId)'
        )
</template>
<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { Action, State } from 'vuex-class'
import moment, { locale } from 'moment'
import * as actionTypes from '../store/actionTypes'
import ProductsFormTable from './ProductFormTable.vue'
import AddProduct from './AddProduct.vue'
import SearchProduct from './SearchProduct.vue'
@Component({
  components: { ProductsFormTable, AddProduct, SearchProduct },
})
export default class ProductTable extends Vue {
  @Action(actionTypes.GET_ALL_PRODUCTS) GET_ALL_PRODUCTS
  @Action(actionTypes.DELETE_BY_ID) DELETE_BY_ID
  @State('products') products
  @State('searchProduct') searchProduct

  showProductFormTableModal = false
  showAddProductModal = false
  productId = false

  openProductFormTableModal(productId: number, name: string) {
    ;(this as any).$buefy.modal.open({
      parent: this,
      component: ProductsFormTable,
      trapFocus: true,
      props: {
        productId,
        name,
      },
    })
  }

  openAddProductModal() {
    ;(this as any).$buefy.modal.open({
      parent: this,
      component: AddProduct,
      canCancel: false,
      fullScreen: false,
    })
  }

  formatDate(date: Date) {
    locale('ru')
    return moment(date).format('L')
  }

  async deleted(id: Number) {
    await this.DELETE_BY_ID(id)
    ;(this as any).$buefy.toast.open({
      message: 'Успешно удалено!',
      type: 'is-danger',
    })
  }

  allProducts() {
    if (this.searchProduct.length === 0) {
      return this.products
    } else {
      return this.searchProduct
    }
  }

  async created() {
    try {
      await this.GET_ALL_PRODUCTS()
    } catch (error) {
      ;(this as any).$router.push('/sign-in')
    }
  }
}
</script>