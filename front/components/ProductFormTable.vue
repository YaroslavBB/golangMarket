<template lang="pug">
.modal-card(style='width: auto')
  header.modal-card-head
    p.modal-card-title Детальнее
    button.delete(type='button', @click='$emit("close")')
  .modal-card-body
    b-table(:data='productForms || []', fullwidth='')
      b-table-column(field='form', label='Форма', v-slot='props', centered='')
        | {{ props.row.form }}
      b-table-column(
        field='amount',
        label='Количество',
        v-slot='props',
        centered=''
      )
        | {{ props.row.amount }}
      b-table-column(field='price', label='Цена', v-slot='props', centered='')
        | {{ props.row.price }}
      b-table-column(
        field='dateStart',
        label='Дата старта цены',
        v-slot='props',
        centered=''
      )
        | {{ formatDate(props.row.dateStart) }}
      b-table-column(
        field='dateEnd',
        label='Дата окончания цены',
        v-slot='props',
        centered=''
      )
        | {{ formatDate(props.row.dateEnd) }}
  footer.modal-card-foot
    b-button(type='is-success', @click='openAddFormModal', outlined='')
      | Добавить форму
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { Action, State } from 'vuex-class'
import moment, { locale } from 'moment'
import * as actionTypes from '../store/actionTypes'
import AddForm from './AddForm.vue'
@Component({
  components: { AddForm },
})
export default class extends Vue {
  @Action(actionTypes.GET_PRODUCT_FORMS) GET_PRODUCT_FORMS!: any
  @State('productForms') productForms!: any

  @Prop()
  name!: string

  @Prop()
  productId!: number

  formatDate(date: Date) {
    locale('ru')
    return moment(date).format('L')
  }

  openAddFormModal() {
    ;(this as any).$buefy.modal.open({
      parent: this,
      component: AddForm,
      hasModalCard: true,
      trapFocus: true,
      customClass: "custom-class",
      props: {
        name: this.name,
        productId: this.productId,
      },
    })
  }

  async created() {
    await this.GET_PRODUCT_FORMS(this.productId)
  }
}
</script>

