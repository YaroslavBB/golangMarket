import * as actionTypes from './actionTypes'
import * as mutationTypes from './mutationTypes'
import moment from 'moment'
import { MutationTree, ActionTree } from 'vuex/types/index'

interface RootState {
    products: Product[]
    productForms: ProductForm[]
    searchProduct: Product[]
}

interface Product {
    productId: Number
    name: string
    dateAdded: Date
}

interface ProductForm {
    form: String
    amount: Number
    price: Number
    dateStart: Date
    dateEnd: Date
}



export const state = (): RootState => ({
    products: [],
    productForms: [],
    searchProduct: [],
})

export const mutations: MutationTree<RootState> = {
    [mutationTypes.SET_ALL_PRODUCT](state, payload) {
        state.products = payload
    },

    [mutationTypes.SET_PRODUCT_FORM](state, payload) {
        state.productForms = payload
    },

    [mutationTypes.ADD_PRODUCT](state, payload) {
        state.products.push(payload)
    },

    [mutationTypes.SEARCH_PRODUCT](state, searchQuery) {
        moment.locale("en")
        let startDate = moment(searchQuery.searchStartDate).format("l")
        let endDate = moment(searchQuery.searchEndDate).format("l")

        state.searchProduct = state.products.filter((product) => {
            let productDate = moment(product.dateAdded).format("l")

            if (moment(productDate).isAfter(startDate) && moment(productDate).isBefore(endDate)) {
                return product.name.toLowerCase().includes(searchQuery.searchName.toLowerCase())
            }
        })
    },
    [mutationTypes.REFRESH](state) {
        state.searchProduct = []
    }
}

export const actions: ActionTree<RootState, RootState> = {
    async [actionTypes.GET_ALL_PRODUCTS]({ commit }) {
        let response = await this.$axios.$get('http://127.0.0.1:8080/products')
        commit(mutationTypes.SET_ALL_PRODUCT, response)
    },

    async [actionTypes.DELETE_BY_ID]({ dispatch }, id) {
        await this.$axios.$delete(`http://127.0.0.1:8080/delete/${id}`)

        dispatch(actionTypes.GET_ALL_PRODUCTS)
    },

    async [actionTypes.ADD_NEW_PRODUCT]({ dispatch }, newProduct) {
        await this.$axios.$post('http://127.0.0.1:8080/product/add', newProduct)

        dispatch(actionTypes.GET_ALL_PRODUCTS)
    },

    async [actionTypes.GET_PRODUCT_FORMS]({ commit }, productId) {
        let response = await this.$axios.$get(`http://127.0.0.1:8080/product/${productId}`)
        commit(mutationTypes.SET_PRODUCT_FORM, response)
    },

    async [actionTypes.ADD_NEW_FORM]({ dispatch }, newProduct) {
        await this.$axios.$post('http://127.0.0.1:8080/product/add', newProduct)

        dispatch(actionTypes.GET_PRODUCT_FORMS, newProduct.productId)
    }
}
