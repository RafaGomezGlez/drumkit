import { configureStore } from '@reduxjs/toolkit'
import { drumkitAPI } from '../features/drumkitAPI'

export const store = configureStore({
  reducer: {
    [drumkitAPI.reducerPath]: drumkitAPI.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(drumkitAPI.middleware),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
