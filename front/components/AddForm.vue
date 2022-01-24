<template lang="pug">
form(@submit.prevent='addProduct()')
  .modal-card
    header.modal-card-head
      p.modal-card-title Добавление формы
      button.delete(type='button', @click='$emit("close")')
    .modal-card-body
      b-field.mb-5
        b-input(
          v-model='form',
          placeholder='Укажите форму товара',
          rounded='',
          required=''
        )
      b-field
        b-numberinput(
          v-model='amount',
          placeholder='Введите колличество товара',
          rounded='',
          required=''
        )
      b-field
        b-numberinput(
          v-model='price',
          placeholder='Укажите цену товара',
          rounded='',
          required=''
        )
      b-field
        b-datepicker(
          append-to-body
          v-model='dateStart',
          ref='datepicker',
          expanded='',
          placeholder='Укажите дату начала цен',
          required='',
          locale='ru-RU'
        )
        b-button(
          @click='$refs.datepicker.toggle()',
          icon-left='calendar-today',
          type='is-primary'
        )
      b-field
        b-datepicker(
          append-to-body
          v-model='dateEnd',
          ref='datepicker2',
          expanded='',
          placeholder='Укажите дату конца цен',
          required='',
          locale='ru-RU'
        )
        b-button(
          @click='$refs.datepicker2.toggle()',
          icon-left='calendar-today',
          type='is-primary'
        )
    footer.modal-card-foot
      button.button.is-success.is-outlined(type='submit') Добавить форму
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
    const newProduct = {
      productId: this.productId,
      name: this.name,
      form: this.form,
      amount: Number(this.amount),
      price: Number(this.price),
      dateStart: new Date(Date.parse(this.dateStart)),
      dateEnd: new Date(Date.parse(this.dateEnd)),
    }

    try {
      await this.ADD_NEW_FORM(newProduct)
      ;(this as any).$emit('close')
      ;(this as any).$buefy.toast.open({
        message: 'Успешно добавлено!',
        type: 'is-success',
      })
    } catch (err) {
      ;(this as any).$buefy.toast.open({
        message: err.response.data.error,
        type: 'is-danger',
      })
    }
  }
}
</script>