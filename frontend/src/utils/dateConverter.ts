export const date = (str: string) => {
  const date = new Date(str)
  const year = date.getFullYear()
  const month = date.getMonth()
  const day = date.getDay()

  return `${year}年${month}月${day}日`
}

export const timeAndHour = (str: string) => {
  const date = new Date(str)
  const hour = date.getHours()
  const time = date.getMinutes()

  return `${hour}時 ${time}分`
}