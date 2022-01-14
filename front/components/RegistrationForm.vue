<template>
  <section class="section">
    <div class="columns is-centered">
      <div class="column is-6">
        <form @submit.prevent="register()">
          <b-field label="Имя пользователя">
            <b-input
              v-model="user.username"
              placeholder="Введите имя пользователя"
              required
              icon="account"
              minlength="3"
              maxlength="15"
            ></b-input>
          </b-field>

          <b-field label="Введите пароль">
            <b-input
              v-model="user.password"
              placeholder="Введите пароль"
              password-reveal
              required
              minlength="3"
              maxlength="15"
              type="password"
              icon="lock-outline"
            ></b-input>
          </b-field>

          <b-field label="Введите пароль еще раз">
            <b-input
              v-model="repeatPassword"
              placeholder="Введите пароль еще раз"
              password-reveal
              required
              type="password"
              icon="lock-outline"
            ></b-input>
          </b-field>
          <b-field>
            <div v-if="this.comparePassword()"></div>
            <div class="is-size-7 has-text-danger" v-else>
              Пароли не совпадают!
            </div>
          </b-field>
          <b-field>
            <button class="button is-primary" type="submit">
              Зарегистрироваться
            </button>
          </b-field>
        </form>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { Action, State } from 'vuex-class'
import * as actionTypes from '../store/actionTypes'
@Component
export default class RegistrationForm extends Vue {
  @Action(actionTypes.REGISTRATION) REGISTRATION
  @State('message') message
  user = {
    username: null,
    password: null,
  }
  repeatPassword: string = null

  comparePassword() {
    if (this.user.password == this.repeatPassword) {
      return true
    }
    return false
  }

  async register() {
    if (this.comparePassword()) {
      await this.REGISTRATION(this.user)
      if (this.message == '') {
        ;(this as any).$buefy.toast.open({
          message: 'Регистрация прошла успешно!',
          type: 'is-primary',
        })
      } else {
        ;(this as any).$buefy.toast.open({
          message: this.message,
          type: 'is-danger',
        })
      }
    }
  }
}
</script>