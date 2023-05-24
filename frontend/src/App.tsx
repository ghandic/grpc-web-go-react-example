import { useQuery, useMutation } from '@tanstack/react-query';
import { FC, useState } from 'react';
import { listUsers, createUser, deleteUser } from './gen/proto/users/v1/users-UserService_connectquery';
import { ListUsersResponse, User, CreateUserResponse, CreateUserRequest, DeleteUserRequest, DeleteUserResponse } from './gen/proto/users/v1/users_pb';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { TransportProvider } from '@bufbuild/connect-query';
import { createConnectTransport } from '@bufbuild/connect-web';
import {
  Alert,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Typography,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  TextField,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';

const AddUser = () => {
  const [open, setOpen] = useState(false);
  const [name, setName] = useState<string | undefined>(undefined);
  const [error, setError] = useState<string | undefined>(undefined);
  const [showError, setShowError] = useState(false);

  const handleOpen = () => setOpen(true);

  const handleClose = () => {
    setName(undefined);
    setError(undefined);
    setOpen(false);
  }

  const createUserMutation = useMutation<CreateUserResponse, Error, CreateUserRequest>({
    ...createUser.useMutation(),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['users.v1.UserService', 'ListUsers'],
      });
      handleClose()
    },
    onError: (error: Error) => {
      error ? setError(error.message) : null;
      setShowError(true);
      window.setTimeout(() => {
        setShowError(false);
      }, 5000);
    }
  });

  const handleCreateUser = () => {
    if(!name) {
      return
    }

    const payload = {
      name: name
    }
    createUserMutation.mutate(payload);
  }

  return (
    <div>
      <Button variant="outlined" onClick={handleOpen}>
        Add User
      </Button>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Add User</DialogTitle>
        <DialogContent>
          <DialogContentText>To add a user, please enter the name here.</DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            id="name"
            label="Name"
            fullWidth
            variant="standard"
            value={name}
            onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
              setName(event.target.value);
            }}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleCreateUser}>
            Add User
          </Button>
        </DialogActions>
      </Dialog>
      {showError && (
        <div className="position-fixed bottom-0 p-3 z-index-2">
          <Alert severity="error">
            {error ? `${error}` : 'Error adding user.'}
          </Alert>
        </div>
      )}
    </div>
  );
}

const UserList = ({ users }: {users: User[]}) => {

  if (users.length === 0) return <Typography>No users in db yet...</Typography>;

  const deleteUserMutation  = useMutation<DeleteUserResponse, Error, DeleteUserRequest>({
    ...deleteUser.useMutation(),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['users.v1.UserService', 'ListUsers'],
      });
    },
  });

  const handleDeleteUser = (userId) => {
    const payload = {
      userId: userId
    }
    deleteUserMutation.mutate(payload);
  }
  
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>ID</TableCell>
            <TableCell>Name</TableCell>
            <TableCell align="right">Delete</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {users.map((user) => (
            <TableRow key={user.id} sx={{ '&:last-child td, &:last-child th': { border: 0 } }}>
              <TableCell>{user.id}</TableCell>
              <TableCell>{user.name}</TableCell>
              <TableCell align="right">
                <IconButton aria-label="delete" size="small" onClick={()=>{handleDeleteUser(user.id)}}>
                  <DeleteIcon fontSize="inherit" />
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

const UserManagement: FC = () => {
  const { data: users, isSuccess } = useQuery<ListUsersResponse, Error>(listUsers.useQuery({}));
  if (!isSuccess) return <div>Fail...</div>

  return (
    <div>
      <AddUser />
      <UserList users={users.users} />
    </div>
  );
};

const queryClient = new QueryClient();

const App: FC = () => {

  /*TODO Fetch from env variables */
  const transport = createConnectTransport({
    baseUrl: 'http://127.0.0.1:8080',
  });

  return (
    <TransportProvider transport={transport}>
      <QueryClientProvider client={queryClient}>
        <UserManagement />
        <ReactQueryDevtools initialIsOpen={true} />
      </QueryClientProvider>
    </TransportProvider>
  );
}

export default App;

