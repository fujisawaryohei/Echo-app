import { createEntityAdapter, createSelector, createSlice, PayloadAction } from '@reduxjs/toolkit';
import { AppThunk, RootState, store } from '../store';
import { getPosts } from '../apis/blog.api';
import { Post } from '../types/index';

type IPostState = {
  selectPost: Post | null,
  errorMsg?: string
}

const postAdapter = createEntityAdapter<Post>({
  selectId: (post) => post.id,
})

const slice = createSlice({
  name: 'post',
  // { ids: [], entities: [], selectPost: null }
  initialState: postAdapter.getInitialState<IPostState>({
    selectPost: null,
    errorMsg: '',
  }),
  reducers: {
    getPostsSuccess: (state, action: PayloadAction<Post[]>) => {
      postAdapter.setAll(state, action.payload)
    },
    getPostsFailure: (state, action: PayloadAction<Error>) => {
      state.errorMsg = action.payload.message
    },
    setSelectPost: (state, action: PayloadAction<Post>) => {
      state.selectPost = action.payload
    }
  }
})

export default slice.reducer;

export const {
  getPostsSuccess,
  setSelectPost,
  getPostsFailure,
} = slice.actions;


export const requestGetPostsSuccess = (): AppThunk => {
  return async (dispatch, getState) => {
    await getPosts('/posts')
      .then(res => dispatch(getPostsSuccess(res)))
      .catch(error => dispatch(getPostsFailure(error)))
  }
}

export const postSelector = postAdapter.getSelectors((state: RootState) => state.post )