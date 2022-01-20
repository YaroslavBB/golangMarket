<template lang="pug">
section.section
  .columns.is-centered
    .column.is-6
      form(@submit.prevent='login()')
        b-field(label='Имя пользователя')
          b-input(
            v-model='user.username',
            placeholder='Введите имя пользователя',
            required='',
            icon='account',
            minlength='3',
            maxlength='15',
            validation-message='Минимальное количество символов 3. Имя не может содержать пробелов',
            pattern='[^\' \']+'
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
            icon='lock-outline'
          )
        .columns
          b-field.column.is-1
            button.button.is-primary.is-outlined.is-centered(type='submit')
              | Войти
          b-field.column.is-1
            b-button(type='is-ghost is-centered', @click='registration') Регистрация
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { Action } from 'vuex-class'
import * as actionTypes from '../store/actionTypes'
@Component
export default class LoginForm extends Vue {
  @Action(actionTypes.LOGIN) LOGIN
  user = {
    username: null,
    password: null,
  }

  registration() {
    ;(this as any).$router.push('/sign-up/registration')
  }

  async login() {
    try {
      await this.LOGIN(this.user)
      ;(this as any).$buefy.toast.open({
        message: 'Добро пожаловать!',
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