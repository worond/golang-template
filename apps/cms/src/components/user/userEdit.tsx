import {
  Edit,
  SimpleForm,
  TextInput,
  DateInput,
  ReferenceManyField,
  Datagrid,
  TextField,
  DateField,
  EditButton,
  required,
  PasswordInput,
} from "react-admin";

export const UserEdit = () => (
  <Edit>
    <SimpleForm>
      <TextInput disabled label="Id" source="id" />
      <TextInput source="name" validate={required()} />
      <TextInput source="email" validate={required()} />
      <DateInput source="created_at" />
    </SimpleForm>
  </Edit>
);
