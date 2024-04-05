import {
  Create,
  SimpleForm,
  ReferenceInput,
  TextInput,
  useRedirect,
  useNotify,
} from "react-admin";

export const UserCreate = () => {
  const redirect = useRedirect();
  const notify = useNotify();

  const onSuccess = (data: any) => {
    notify(`Changes saved`);
    redirect(`/users/${data.id}`);
  };

  return (
    <Create mutationOptions={{ onSuccess }}>
      <SimpleForm mode="onSubmit" reValidateMode="onChange">
        <TextInput label="Name" source="name" />
        <TextInput label="Email" source="email" />
        <TextInput label="Password" source="password" />
      </SimpleForm>
    </Create>
  );
};
