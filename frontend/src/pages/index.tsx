import { useEffect } from 'react'
import { useAppDispatch, useAppSelector } from '../hooks';
import { postSelector, requestGetPostsSuccess} from "../stores/post";
import PostCard from '../components/PostCard'
import { Grid } from '@mui/material'

const Index = () => {
  const dispatch = useAppDispatch();
  const posts = useAppSelector(postSelector.selectAll)

  useEffect(() => {
    dispatch(requestGetPostsSuccess())
  }, [dispatch])
  
  return (
    <Grid container spacing={2}>
      {posts.map((post) => {
        return <PostCard post={post} key={post.id.toString()} />
      })}
    </Grid>
  )
}
export default Index;