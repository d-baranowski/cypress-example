"use client"

import React, {useRef, useState} from 'react';
import ExampleForm from "@/components/example-form";
import FormModelTable from "@/components/example-list";
import {FormModel} from "../../../gen/example/v1/form_pb";
import {Dialog, DialogActions, DialogContent, DialogTitle} from "@mui/material";
import Button from "@mui/material/Button";

interface Props {

}

const Page: React.FunctionComponent<Props> = function Page() {
  const [selectedModel, setSelectedModel] = React.useState<FormModel | null>(null); // State to hold the selected model
  const tableRef = useRef<{ refetch: () => void }>(null); // Ref for the table's refetch function
  const [open, setOpen] = useState(false); // State to manage dialog open/close


  // Handler to be called when a row is double-clicked
  const handleRowDoubleClick = (model: FormModel) => {
    setSelectedModel(model); // Open the form with the selected model data
    setOpen(true); // Open the dialog
  };

  // After saving the form, refetch the table rows
  const handleAfterSave = () => {
    setSelectedModel(null); // Close the form
    setOpen(false); // Close the dialog
    tableRef.current?.refetch(); // Refetch the table data after saving
  };

  // Function to close the dialog
  const handleClose = () => {
    setOpen(false); // Close the dialog
  };

  return (
    <div>
      <Button variant="contained" onClick={() => {
        setSelectedModel(null)
        setOpen(true)
      }}>Add Form Model</Button>
      <Dialog open={open} onClose={handleClose} maxWidth="md" fullWidth>
        <DialogTitle>Form Model</DialogTitle>
        <DialogContent>
          {selectedModel && (
            <ExampleForm
              id={selectedModel.id}
              afterSave={handleAfterSave}
              onCancel={handleClose} // Close the dialog on cancel
            />
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="primary">
            Cancel
          </Button>
        </DialogActions>
      </Dialog>
      <FormModelTable onRowDoubleClick={handleRowDoubleClick} ref={tableRef}/>
    </div>
  );
};

export default Page;
