"use client";
import {
  Admin,
  Resource,
  ListGuesser,
  EditGuesser,
  radiantLightTheme,
} from "react-admin";
import { authProvider, dataProvider } from "../providers";

const AdminApp = () => (
  <Admin
    theme={radiantLightTheme}
    authProvider={authProvider}
    dataProvider={dataProvider}
  >
    <Resource
      name="users"
      list={ListGuesser}
      edit={EditGuesser}
      recordRepresentation="name"
    />
    <Resource
      name="posts"
      list={ListGuesser}
      edit={EditGuesser}
      recordRepresentation="title"
    />
    <Resource name="comments" list={ListGuesser} edit={EditGuesser} />
  </Admin>
);

export default AdminApp;
