<template lang="pug">
form(@submit.prevent='addProduct()')
  .modal-card(style='width: auto')
    header.modal-card-head
      p.modal-card-title Добавление товара
      button.delete(type='button', @click='$emit("close")')
    .modal-card-body
      b-field
        b-input(
          v-model='name',
          placeholder='Введите наименование товара',
          minlength='3',
          rounded='',
          required=''
        )
      b-field
        b-input(v-model='form', placeholder='Укажите форму товара', rounded='')
      b-field
        b-numberinput(
          v-model='amount',
          placeholder='Введите колличество товара',
          rounded='',
          type='number',
          required=''
        )
      b-field
        b-numberinput(
          v-model='price',
          placeholder='Укажите цену товара',
          rounded='',
          type='number',
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
      button.button.is-success.is-outlined(type='submit') Добавить
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { Action } from 'vuex-class'
import * as actionTypes from '../store/actionTypes'

@Component
export default class AddProduct extends Vue {
  @Action(actionTypes.ADD_NEW_PRODUCT) ADD_NEW_PRODUCT

  name = null
  form = null
  amount = null
  price = null
  dateStart = null
  dateEnd = null

  async addProduct() {
    const newProduct = {
      name: this.name,
      form: this.form,
      amount: this.amount,
      price: this.price,
      dateStart: new Date(Date.parse(this.dateStart)),
      dateEnd: new Date(Date.parse(this.dateEnd)),
    }
    try {
      await this.ADD_NEW_PRODUCT(newProduct)
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