<template>
  <div class="columns">
    <b-field class="column is-3">
      <b-input
        v-model="searchQuery.searchName"
        placeholder="Наименование товара..."
        icon="magnify"
        icon-clickable
      >
      </b-input>
    </b-field>

    <b-field class="column is-3">
      <b-datepicker
        v-model="searchQuery.searchStartDate"
        ref="datepicker1"
        expanded
        placeholder="С '28.12.2021'..."
      >
      </b-datepicker>
      <b-button
        @click="$refs.datepicker1.toggle()"
        icon-left="calendar-today"
        type="is-primary"
      />
    </b-field>

    <b-field class="column is-3">
      <b-datepicker
        v-model="searchQuery.searchEndDate"
        ref="datepicker"
        expanded
        placeholder="По '5.01.2022'..."
      >
      </b-datepicker>
      <b-button
        @click="$refs.datepicker.toggle()"
        icon-left="calendar-today"
        type="is-primary"
      />
    </b-field>

    <b-field class="column">
      <b-button
        icon-left="magnify"
        type="search is-info is-light"
        @click="search"
        >Поиск</b-button
      >
    </b-field>

    <b-field class="column is-2">
      <b-button icon-left="refresh" type="is-warning is-light" @click="refresh"
        >Очистить</b-button
      >
    </b-field>
  </div>
</template>

<script lang="ts">
import * as mutationType from '../store/mutationTypes'
import { State, Mutation } from 'vuex-class'
import { Component, Vue } from 'nuxt-property-decorator'
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
    if (this.searchProduct.length == 0) {
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
  refresh() {
    this.REFRESH()
  }
}
</script>