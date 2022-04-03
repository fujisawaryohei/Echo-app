import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit';
import PostReducer from './stores/post';

export const store = configureStore({
  reducer: {
    post: PostReducer
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
