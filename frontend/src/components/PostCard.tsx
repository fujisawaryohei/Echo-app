import { Post } from '../types'
import { Grid, Card, CardMedia, CardContent, Typography, CardActions, Button } from '@mui/material'

type IProps = {
  post: Post
}

const PostCard = (props: IProps) => {
  return (
    <Grid item xs={4}>
      <Card>
        <CardMedia
          component="img"
          alt="post-card"
          height="300"
          image="/blog-image.jpeg"
        />
        <CardContent sx={ { backgroundColor: '#FFFFFF' } } >
          <Typography gutterBottom variant="h6" component="div">
            { props.post.title }
          </Typography>
        </CardContent>
        <Typography noWrap={ true } sx={ { marginLeft: '15px' } }>
          { props.post.body }
        </Typography>
        <CardActions>
          <Button size="small">詳細ページへ</Button>
        </CardActions>
      </Card>
    </Grid>
  )
}

export default PostCard;