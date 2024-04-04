import { List, Datagrid, TextField, DateField } from "react-admin";
import { userFilter } from "./userFilter";

export const UserList = () => (
  <List filters={userFilter}>
    <Datagrid>
      <TextField source="id" />
      <TextField source="email" />
      <TextField source="name" />
      <DateField source="created_at" />
    </Datagrid>
  </List>
);
