<template>
<form @submit.prevent="addProduct()">
  <div class="modal-card">
    <header class="modal-card-head">
      <p class="modal-card-title">Добавление формы</p>
      <button type="button" class="delete" @click="$emit('close')" />
    </header>
    <div class="modal-card-body">
      <b-field class="mb-5">
        <b-input v-model="form" placeholder="Укажите форму товара" rounded required/>
        
      </b-field>

      <b-field>
        <b-numberinput
          v-model="amount"
          placeholder="Введите колличество товара"
          rounded
          required
        />
      </b-field>
      <b-field>
        <b-numberinput
          v-model="price"
          placeholder="Укажите цену товара"
          rounded
          required
        />
      </b-field>
      <b-field>
        <b-datepicker
          v-model="dateStart"
          ref="datepicker"
          expanded
          placeholder="Укажите дату начала цен"
          required
        >
        </b-datepicker>
        <b-button
          @click="$refs.datepicker.toggle()"
          icon-left="calendar-today"
          type="is-primary"
        />
      </b-field>
      <b-field>
        <b-datepicker
          v-model="dateEnd"
          ref="datepicker2"
          expanded
          placeholder="Укажите дату конца цен"
          required
        >
        </b-datepicker>
        <b-button
          @click="$refs.datepicker2.toggle()"
          icon-left="calendar-today"
          type="is-primary"
        />
      </b-field>
    </div>
    <footer class="modal-card-foot">
      <button class="button is-success is-outlined" type="submit">
        Добавить форму
      </button>
    </footer>
  </div>
</form>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { Action } from 'vuex-class'
import * as actionTypes from '../store/actionTypes'
@Component
export default class AddForm extends Vue {
  @Action(actionTypes.ADD_NEW_FORM) ADD_NEW_FORM

  @Prop()
  name: string
  @Prop()
  productId: number

  form = null
  amount = null
  price = null
  dateStart = null
  dateEnd = null

  async addProduct() {
    let newProduct = {
      productId: this.productId,
      name: this.name,
      form: this.form,
      amount: Number(this.amount),
      price: Number(this.price),
      dateStart: new Date(Date.parse(this.dateStart)),
      dateEnd: new Date(Date.parse(this.dateEnd)),
    }

    this.ADD_NEW_FORM(newProduct)
    ;(this as any).$emit('close')
    ;(this as any).$buefy.toast.open({
      message: 'Успешно добавлено!',
      type: 'is-success',
    })
  }
}
</script>