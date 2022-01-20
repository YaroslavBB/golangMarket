<template lang="pug">
.columns
  b-field.column.is-3
    b-input(
      v-model='searchQuery.searchName',
      placeholder='Наименование товара...',
      icon='magnify',
      icon-clickable=''
    )
  b-field.column.is-3
    b-datepicker(
      v-model='searchQuery.searchStartDate',
      locale='ru-RU',
      ref='datepicker1',
      expanded='',
      placeholder='С \'28.12.2021\'...'
    )
    b-button(
      @click='$refs.datepicker1.toggle()',
      icon-left='calendar-today',
      type='is-primary'
    )
  b-field.column.is-3
    b-datepicker(
      v-model='searchQuery.searchEndDate',
      locale='ru-RU',
      ref='datepicker',
      expanded='',
      placeholder='По \'5.01.2022\'...'
    )
    b-button(
      @click='$refs.datepicker.toggle()',
      icon-left='calendar-today',
      type='is-primary'
    )
  b-field.column
    b-button(
      icon-left='magnify',
      type='search is-info is-light',
      @click='search'
    ) Поиск
  b-field.column.is-2
    b-button(icon-left='refresh', type='is-warning is-light', @click='refresh') Очистить
</template>

<script lang="ts">
import { State, Mutation } from 'vuex-class'
import { Component, Vue } from 'nuxt-property-decorator'
import moment, { locale } from 'moment'
import * as mutationType from '../store/mutationTypes'
@Component
export default class SearchProduct extends Vue {
  @State('searchProduct') searchProduct
  @Mutation(mutationType.SEARCH_PRODUCT) SEARCH_PRODUCT
  @Mutation(mutationType.REFRESH) REFRESH

  date = new Date()
  searchQuery = {
    searchStartDate: new Date(this.date.getFullYear(), this.date.getMonth(), 1),
    searchEndDate: new Date(
      this.date.getFullYear(),
      this.date.getMonth() + 1,
      0
    ),
    searchName: '',
  }

  search() {
    this.SEARCH_PRODUCT(this.searchQuery)
    if (this.searchProduct.length === 0) {
      ;(this as any).$buefy.toast.open({
        message: 'Ничего не найдено!',
      })
    } else {
      ;(this as any).$buefy.toast.open({
        message: 'Вот что мы нашли!',
        type: 'is-success',
      })
    }
  }

  formatDate(date: Date) {
    locale('ru')
    return moment(date).format('L')
  }

  refresh() {
    this.REFRESH()
  }
}
</script>