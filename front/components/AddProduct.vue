<template>
  <form @submit.prevent="addProduct()">
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <p class="modal-card-title">Добавление товара</p>
        <button type="button" class="delete" @click="$emit('close')" />
      </header>
      <div class="modal-card-body">
        <b-field>
          <b-input
            v-model="name"
            placeholder="Введите наименование товара"
            minlength="3"
            rounded
            required
          />
        </b-field>
        <b-field>
          <b-input v-model="form" placeholder="Укажите форму товара" rounded />
        </b-field>
        <b-field>
          <b-numberinput
            v-model="amount"
            placeholder="Введите колличество товара"
            rounded
            type="number"
            required
          />
        </b-field>
        <b-field>
          <b-numberinput
            v-model="price"
            placeholder="Укажите цену товара"
            rounded
            type="number"
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
        <button class="button is-success is-outlined" type="submit">Добавить</button>
      </footer>
    </div>
  </form>
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

  addProduct() {
    let newProduct = {
      name: this.name,
      form: this.form,
      amount: this.amount,
      price: this.price,
      dateStart: new Date(Date.parse(this.dateStart)),
      dateEnd: new Date(Date.parse(this.dateEnd)),
    }
    this.ADD_NEW_PRODUCT(newProduct)
    ;(this as any).$emit('close')
    ;(this as any).$buefy.toast.open({
      message: 'Успешно добавлено!',
      type: 'is-success',
    })
  }
}
</script>