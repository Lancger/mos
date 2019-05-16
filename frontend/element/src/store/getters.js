const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  nick_name: state => state.user.nick_name,
  perms: state => state.user.perms,
}
export default getters
