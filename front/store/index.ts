
import * as actionTypes from './actionTypes'
import * as mutationTypes from './mutationTypes'
import moment from 'moment'
import { MutationTree, ActionTree } from 'vuex/types/index'

interface RootState {
    products: Product[]
    productForms: ProductForm[]
    searchProduct: Product[]
}

interface AutorisationUserPayload {
    username: string
    password: string
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

interface ProductAddProductPayload {
    productId: Number
    name: string
    form: string
    amount: Number
    price: Number
    dateStart: Date
    dateEnd: Date
}

interface SearchQuery {
    searchStartDate: Date
    searchEndDate: Date
    searchName: string
}

export const state = (): RootState => ({
    products: [],
    productForms: [],
    searchProduct: [],
})

export const mutations: MutationTree<RootState> = {
    [mutationTypes.SET_ALL_PRODUCT](state, payload: Product[]) {
        state.products = payload
    },

    [mutationTypes.SET_PRODUCT_FORM](state, payload: ProductForm[]) {
        state.productForms = payload
    },

    [mutationTypes.SEARCH_PRODUCT](state, searchQuery: SearchQuery) {
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
    },
}

export const actions: ActionTree<RootState, RootState> = {
    async [actionTypes.GET_ALL_PRODUCTS]({ commit }) {
        try {
            let response = await this.$axios.$get('api/products')
            commit(mutationTypes.SET_ALL_PRODUCT, response)
        } catch (err) {
            throw err
        }
    },

    async [actionTypes.DELETE_BY_ID]({ dispatch }, id: Number) {
        await this.$axios.$delete(`api/delete/${id}`)

        dispatch(actionTypes.GET_ALL_PRODUCTS)
    },

    async [actionTypes.ADD_NEW_PRODUCT]({ dispatch }, newProduct: ProductAddProductPayload) {
        try {
            await this.$axios.$post('api/product/add', newProduct)
    
            dispatch(actionTypes.GET_ALL_PRODUCTS)
        } catch (err) {
            throw err
        }
    },

    async [actionTypes.GET_PRODUCT_FORMS]({ commit, dispatch }, productId: Number) {
        let response: Product = await this.$axios.$get(`api/product/${productId}`)
        commit(mutationTypes.SET_PRODUCT_FORM, response)

    },

    async [actionTypes.ADD_NEW_FORM]({ dispatch }, newProduct: ProductAddProductPayload) {
        try {
            await this.$axios.$post('api/product/add', newProduct)
    
            dispatch(actionTypes.GET_PRODUCT_FORMS, newProduct.productId)
        } catch (err) {
            throw err
        }
    },

    async [actionTypes.REGISTRATION]({ commit }, user: AutorisationUserPayload) {
        try {
            await this.$axios.$post('api/sign-up/regisration', user)
        } catch (err: any) {
            throw err
        }
    },

    async [actionTypes.LOGIN]({ commit }, user: AutorisationUserPayload) {
        try {
            await this.$axios.$post('api/sign-in/login', user)
        } catch (err: any) {
            throw err
        }
    }
}
