import React, {forwardRef, useImperativeHandle, useRef} from 'react';
import {Box, Button} from '@mui/material';
import {useQuery} from "@connectrpc/connect-query";
import {listFormModels} from "../../gen/example/v1/form-FormModelService_connectquery";
import {ConnectError} from "@connectrpc/connect";
import {MaterialReactTable} from "material-react-table";
import {FormModel} from "../../gen/example/v1/form_pb";

interface Props {
  onRowDoubleClick(row: FormModel): void
}

// Main Component
const FormModelTable = forwardRef<{ refetch: () => void }, Props>(({ onRowDoubleClick }, ref) => {
  const {data, isLoading, refetch, error} = useQuery(listFormModels)

  // Expose the refetch function via the ref
  useImperativeHandle(ref, () => ({
    refetch: () => refetch(),
  }));

  if (error instanceof ConnectError) {
    return <div>Error: {error.message}</div>;
  }

  // Define columns for MaterialReactTable
  const columns = [
    {
      accessorKey: 'id', // Access the id field
      header: 'ID',
      Cell: (props: any) => <Button onClick={() => onRowDoubleClick(props.row.original)}>Edit - {String(props.cell.getValue())}</Button>,
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

  return (
    <Box>
      {/* Button to refetch data manually */}
      <Button onClick={() => refetch()} variant="contained" sx={{mb: 2}}>
        Refetch Data
      </Button>

      {/* Render the Material React Table */}
      <MaterialReactTable
        columns={columns}
        enablePagination={false}
        data={data?.formModels || []} // Pass form models from response
        // Handle double click on rows
        muiTableBodyRowProps={({ row }) => ({
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
