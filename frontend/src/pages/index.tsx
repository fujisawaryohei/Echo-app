import { useEffect } from 'react'
import PostCard from '../components/PostCard';
import { useAppDispatch, useAppSelector } from '../hooks';
import { postSelector, requestGetPostsSuccess} from "../stores/post";

const Index = () => {
  const dispatch = useAppDispatch();
  const posts = useAppSelector(postSelector.selectAll)

  useEffect(() => {
    dispatch(requestGetPostsSuccess())
  }, [dispatch])
  
  return (
    <div className='flex flex-wrap justify-between'>
      {posts.map((post) => {
        return (
          <PostCard key={post.id} post={post} />
        )
      })}
    </div>
  )
}
export default Index;