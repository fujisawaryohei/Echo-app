import { Post } from '../types/index'

type IProps = {
  post: Post
}

const PostCard = (props: IProps) => {
  return (
    <div key={props.post.id} className="h-1/5 w-1/3 m-1">
      <img className='h-40 w-full' src="/blog-image.jpeg"></img>
      <p>{props.post.title}</p>
    </div>
  )
}

export default PostCard;