"use client";
import { Admin, Resource, ListGuesser, EditGuesser } from "react-admin";
import { authProvider, dataProvider } from "../providers";
import { UserCreate, UserList, UserEdit } from "./user";
import { LoginPage } from "./login";

const AdminApp = () => (
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
);

export default AdminApp;
