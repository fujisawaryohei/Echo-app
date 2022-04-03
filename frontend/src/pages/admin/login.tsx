import { useState } from 'react'
import { Box, TextField } from '@mui/material'

const Login = () => {
  const [email, setEmail] = useState<String>();
  const [password, setPassword] = useState<String>();

  return (
    <Box
      component="form"
      sx={{
        '& .MuiTextField-root': { m: 1, width: 300 },
        textAlign: 'center'
      }}
      noValidate
      autoComplete="off"
    >
      <TextField
        id="filled-email-input"
        label="Email"
        type="email"
        autoComplete="current-email"
        variant="outlined"
      />
      <TextField
        id="filled-password-input"
        label="Password"
        type="password"
        autoComplete="current-password"
        variant="outlined"
      />
    </Box>
  )
}


export default Login;
