import {Map} from 'immutable'

// --begin-actions-types--
export const LOGIN = 'LOGIN'
export const LOGIN_SUCCESS = 'LOGIN_SUCCESS'
export const LOGIN_FAILURE = 'LOGIN_FAILURE'
export const SIGNUP = 'SIGNUP'
export const SIGNUP_SUCCESS = 'SIGNUP_SUCCESS'
export const SIGNUP_FAILURE = 'SIGNUP_FAILURE'
// --end-actions-types--

// --begin-actions-creators--
export const login = (data) => ({
  type: LOGIN,
  payload: {data}
})

export const loginSuccess = (response) => ({
  type: LOGIN_SUCCESS,
  payload: {response}
})

export const loginFailure = (error) => ({
  type: LOGIN_FAILURE,
  payload: {error}
})
export const signup = (data) => ({
  type: SIGNUP,
  payload: {data}
})

export const signupSuccess = (response) => ({
  type: SIGNUP_SUCCESS,
  payload: {response}
})

export const signupFailure = (error) => ({
  type: SIGNUP_FAILURE,
  payload: {error}
})
// --end-actions-creators--

// --begin-initial-state--
const INITIAL_STATE = Map({
  loginFetching: false,
  loginFetched: false,
  loginError: null,
  loginData: null,
  loginResponse: null,

  signupFetching: false,
  signupFetched: false,
  signupError: null,
  signupData: null,
  signupResponse: null,
})
// --end-initial-state--

// --begin-reducer--
export default function reducer (state = INITIAL_STATE, { type, payload }) {
  switch (type) {
    case LOGIN: {
      const {data} = payload
      return {
        ...state,
        loginData: data,
        loginFetching: true,
      }
    }
    
    case LOGIN_SUCCESS: {
      const {response} = payload
      return {
        ...state,
        loginResponse: response,
        loginFetching: false,
        loginFetched: true,
      }
    }

    case LOGIN_FAILURE: {
      const {error} = payload
      return {
        ...state,
        loginError: error,
        loginFetching: false,
        loginFetched: true,
      }
    }
    case SIGNUP: {
      const {data} = payload
      return {
        ...state,
        signupData: data,
        signupFetching: true,
      }
    }
    
    case SIGNUP_SUCCESS: {
      const {response} = payload
      return {
        ...state,
        signupResponse: response,
        signupFetching: false,
        signupFetched: true,
      }
    }

    case SIGNUP_FAILURE: {
      const {error} = payload
      return {
        ...state,
        signupError: error,
        signupFetching: false,
        signupFetched: true,
      }
    }
    
    default:
      return state
  }
}
// --end-reducer--
