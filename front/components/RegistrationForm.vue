<template lang="pug">
section.section
  .columns.is-centered
    .column.is-6
      form(@submit.prevent='register()')
        b-field(label='Имя пользователя')
          b-input(
            v-model='user.username',
            placeholder='Введите имя пользователя',
            required='',
            icon='account',
            minlength='3',
            maxlength='15'
          )
        b-field(label='Введите пароль')
          b-input(
            v-model='user.password',
            placeholder='Введите пароль',
            password-reveal='',
            required='',
            minlength='3',
            maxlength='15',
            type='password',
            icon='lock-outline',
            validation-message='Минимальное количество символов 3. Пароль не должен содержать пробелов',
            pattern='[^\' \']+'
          )
        b-field(label='Введите пароль еще раз')
          b-input(
            v-model='repeatPassword',
            placeholder='Введите пароль еще раз',
            password-reveal='',
            required='',
            type='password',
            icon='lock-outline'
          )
        b-field
          div(v-if='this.comparePassword()')
          .is-size-7.has-text-danger(v-else='')
            | Пароли не совпадают!
        .columns
          b-field.column.is-3
            button.button.is-primary.is-outlined.is-centered(type='submit')
              | Зарегистрироваться
          b-field.column.is-1
            b-button(type='is-ghost is-centered', @click='login') Логин
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { Action } from 'vuex-class'
import * as actionTypes from '../store/actionTypes'
@Component
export default class RegistrationForm extends Vue {
  @Action(actionTypes.REGISTRATION) REGISTRATION
  user = {
    username: null,
    password: null,
  }

  repeatPassword: string = null

  comparePassword() {
    if (this.user.password === this.repeatPassword) {
      return true
    }
    return false
  }

  login() {
    ;(this as any).$router.push('/sign-in')
  }

  async register() {
    try {
      await this.REGISTRATION(this.user)
      ;(this as any).$buefy.toast.open({
        message: 'Регистрация прошла успешно',
        type: 'is-primary',
      })
      ;(this as any).$router.push('/')
    } catch (err) {
      ;(this as any).$buefy.toast.open({
        message: err.response.data.error,
        type: 'is-danger',
      })
    }
  }
}
</script>