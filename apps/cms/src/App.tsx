import { Admin, Resource } from 'react-admin'
import { authProvider, dataProvider } from './providers'
import { LoginPage } from './components/login'
import { UserCreate, UserEdit, UserList } from './components/user'

export const App = () => (
  <Admin
    authProvider={authProvider}
    dataProvider={dataProvider}
    loginPage={LoginPage}
  >
    <Resource
      name="users"
      list={UserList}
      edit={UserEdit}
      create={UserCreate}
      recordRepresentation="name"
    />
  </Admin>
)
