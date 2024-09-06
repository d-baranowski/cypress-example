import React, {forwardRef, useImperativeHandle} from 'react';
import {Box, Button} from '@mui/material';
import {useMutation, useQuery} from "@connectrpc/connect-query";
import {deleteFormModel, listFormModels} from "../../gen/example/v1/form-FormModelService_connectquery";
import {ConnectError} from "@connectrpc/connect";
import {MaterialReactTable, MRT_ActionMenuItem} from "material-react-table";
import {FormModel} from "../../gen/example/v1/form_pb";
import {Delete} from "@mui/icons-material";
import toast from "react-hot-toast";

interface Props {
  onRowDoubleClick(row: FormModel): void
}

// Main Component
const FormModelTable = forwardRef<{ refetch: () => void }, Props>(({onRowDoubleClick}, ref) => {
  const {data, isLoading, refetch, error} = useQuery(listFormModels)

  const deleteMutation = useMutation(deleteFormModel);

  // Expose the refetch function via the ref
  useImperativeHandle(ref, () => ({
    refetch: () => {
      console.log("refetch in table via ref called");
      window.setTimeout(() => {
        toast.promise(refetch(), {
          loading: 'Loading',
          success: 'Data refetched',
          error: 'Error when refetching',
        });
      }, 300)

    },
  }), [refetch]);

  if (error instanceof ConnectError) {
    return <div>Error: {error.message}</div>;
  }

  // Define columns for MaterialReactTable
  const columns = [
    {
      accessorKey: 'id', // Access the id field
      header: 'ID',
      Cell: (props: any) => <Button onClick={() => onRowDoubleClick(props.row.original)}>Edit
        - {String(props.cell.getValue())}</Button>,
    },
    {
      accessorKey: 'name', // Access the name field
      header: 'Name',
    },
    {
      accessorKey: 'age', // Access the age field
      header: 'Age',
    },
    {
      accessorKey: 'city', // Access the city field
      header: 'City',
    },
    {
      accessorKey: 'gender', // Access the gender field
      header: 'Gender',
    },
    {
      accessorKey: 'satisfactionLevel', // Access the satisfaction level field
      header: 'Satisfaction Level',
    },
    {
      accessorKey: 'termsAccepted', // Access termsAccepted
      header: 'Terms Accepted',
      Cell: (props: any) => (props.cell.getValue() ? 'Yes' : 'No'),
    },
  ];

  // Handle deleting a row
  const handleDelete = async (id: bigint) => {
    await deleteMutation.mutateAsync({id: id});
    refetch(); // Refetch data after deletion
  };


  return (
    <Box data-testid={"example-list"}>
      {/* Button to refetch data manually */}
      <Button onClick={() => refetch()} variant="contained" sx={{mb: 2}}>
        Refetch Data
      </Button>
      <MaterialReactTable
        columns={columns}
        enablePagination={false}
        data={data?.formModels || []} // Pass form models from response
        enableCellActions={true}
        renderCellActionMenuItems={({closeMenu, cell, row, table}) => [
          //array required
          <MRT_ActionMenuItem //or just use the normal MUI MenuItem
            icon={<Delete/>}
            key={1}
            label="Delete"
            onClick={() => {
              handleDelete(row.original.id as unknown as bigint)
              closeMenu(); //close the menu after the action is performed
            }}
            table={table}
          />,
        ]}

        // Handle double click on rows
        muiTableBodyRowProps={({row}) => ({
          onDoubleClick: () => onRowDoubleClick(row.original),
        })}
        state={{
          isLoading,
        }}
      />
    </Box>
  );
});

export default FormModelTable;
