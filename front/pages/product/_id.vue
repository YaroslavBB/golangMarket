<template>
  <section class="section">
    <div>
      <div class="columns">
        <b-field class="column is-2">
          <b-button type="is-success" @click="openAddFormModal" outlined>
            Добавить форму
          </b-button>
        </b-field>
        <b-field class="column is-2">
          <b-button type="is-primary" outlined @click="$router.push('/')"
            >Обратно</b-button
          >
        </b-field>
      </div>

      <b-table :data="productForms || []" fullwidth>
        <b-table-column field="form" label="Форма" v-slot="props" centered>
          {{ props.row.form }}
        </b-table-column>

        <b-table-column
          field="amount"
          label="Количество"
          v-slot="props"
          centered
        >
          {{ props.row.amount }}
        </b-table-column>

        <b-table-column field="price" label="Цена" v-slot="props" centered>
          {{ props.row.price }}
        </b-table-column>

        <b-table-column
          field="dateStart"
          label="Дата старта цены"
          v-slot="props"
          centered
        >
          {{ formatDate(props.row.dateStart) }}
        </b-table-column>

        <b-table-column
          field="dateEnd"
          label="Дата окончания цены"
          v-slot="props"
          centered
        >
          {{ formatDate(props.row.dateEnd) }}
        </b-table-column>
      </b-table>
    </div>
  </section>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { Action, State } from 'vuex-class'
import moment from 'moment'
import AddForm from '../../components/AddForm.vue'
import * as actionTypes from '../../store/actionTypes'
@Component({
  components: { AddForm },
})
export default class Product extends Vue {
  @Action(actionTypes.GET_PRODUCT_FORMS) GET_PRODUCT_FORMS
  @State('productForms') productForms

  id: string
  name: string

  formatDate(date: Date) {
    return moment(date).format('L')
  }

  openAddFormModal() {
    ;(this as any).$buefy.modal.open({
      parent: this,
      component: AddForm,
      hasModalCard: true,
      trapFocus: true,
      props: {
        name: this.name,
        productId: this.id,
      },
    })
  }
  async created() {
    console.log(this.name)
    this.id = (this as any).$route.params.id
    this.name = (this as any).$route.params.name
    this.GET_PRODUCT_FORMS(this.id)
  }
}
</script>