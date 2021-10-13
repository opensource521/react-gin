import { configureStore } from '@reduxjs/toolkit'
import alkaneReducer from './modules/alkaneSlice'

const store = configureStore({
  reducer: {
    alkane: alkaneReducer,
  },
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch

export default store
