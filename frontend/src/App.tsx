import { useQuery, useMutation } from '@tanstack/react-query';
import { FC, useCallback, useEffect, useState } from 'react';
import { listUsers, createUser, deleteUser } from './gen/proto/users/v1/users-UserService_connectquery';
import { SortDirection, User, ListUsersRequest } from './gen/proto/users/v1/users_pb';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { TransportProvider } from '@bufbuild/connect-query';
import { createConnectTransport } from '@bufbuild/connect-web';
import { DataGrid, GridColDef, GridRenderCellParams, GridSortModel, GridToolbar } from '@mui/x-data-grid';

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
import { PartialMessage } from '@bufbuild/protobuf';

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


const UserList = () => {
  const [queryOptions, setQueryOptions] = useState<PartialMessage<ListUsersRequest>>({
    offset: 0,
    pageSize: 5,
    sorting: undefined,
    query: {
      text: ""
    }
  });

  const { data, isLoading } = useQuery(
    // listUsers.useQuery({ sorting: { field: 'created_at', direction: SortDirection.ASC } })
    listUsers.useQuery(queryOptions)
  );
  const { mutate } = useMutation({
  if (users.length === 0) return <Typography>No users in db yet...</Typography>;

  const deleteUserMutation  = useMutation<DeleteUserResponse, Error, DeleteUserRequest>({
    ...deleteUser.useMutation(),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['users.v1.UserService', 'ListUsers'],
      });
    },
  });

  // Some API clients return undefined while loading
  // Following lines are here to prevent `rowCountState` from being undefined during the loading
  const [rowCountState, setRowCountState] = useState(data?.total || 0);
  useEffect(() => {
    setRowCountState((prevRowCountState) => (data?.total !== undefined ? data?.total : prevRowCountState));
  }, [data?.total, setRowCountState]);

  const handleSortModelChange = useCallback(
    (sortModel: GridSortModel) => {
      if (sortModel.length > 0) {
        setQueryOptions({
          ...queryOptions,
          sorting: { field: sortModel[0].field, direction: sortModel[0].sort?.toUpperCase() as unknown as SortDirection },
        });
      } else {
        setQueryOptions({ ...queryOptions, sorting: undefined });
      }
    },
    [queryOptions]
  );

  console.log(queryOptions);

  const columns: GridColDef[] = [
    { field: 'id', headerName: 'ID'},
    { field: 'name', headerName: 'Name', flex: 1 },
    {
      field: 'delete_action',
      headerName: 'Delete',
      sortable: false,
      renderCell: (params: GridRenderCellParams<User>) => (
        <IconButton
          aria-label="delete"
          size="small"
          onClick={() => {
            mutate({ userId: params.row.id });
          }}
        >
          <DeleteIcon fontSize="inherit" />
        </IconButton>
      ),
    },
  ];


  return (
    <div style={{ width: 1000, marginTop: 30 }}>
      <TextField
      sx={{ marginBottom: "10px" }}
        margin="dense"
        id="search"
        label="Search"
        variant="standard"
        value={queryOptions.query?.text}
        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
          setQueryOptions({...queryOptions, query: {...queryOptions.query, text:event.target.value}});
        }}
      />
      <DataGrid
        columns={columns}
        rows={data?.users || []}
        rowCount={Number(rowCountState)}
        loading={isLoading}
        pageSizeOptions={[5, 10]}
        paginationModel={{ page: queryOptions.offset as number, pageSize: queryOptions.pageSize as number }}
        paginationMode="server"
        onPaginationModelChange={(newPagination) =>
          setQueryOptions({ ...queryOptions, pageSize: newPagination.pageSize, offset: newPagination.page })
        }
        sortingMode="server"
        onSortModelChange={handleSortModelChange}
        disableColumnFilter
        disableColumnSelector
        disableDensitySelector
        autoHeight
      />
    </div>
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

  return (
    <div>
      <AddUser />
      <UserList />
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

